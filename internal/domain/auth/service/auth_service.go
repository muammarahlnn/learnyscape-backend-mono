package service

import (
	"context"
	"errors"
	"fmt"
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/domain/auth/constant"
	. "learnyscape-backend-mono/internal/domain/auth/dto"
	. "learnyscape-backend-mono/internal/domain/auth/entity"
	. "learnyscape-backend-mono/internal/domain/auth/httperror"
	"learnyscape-backend-mono/internal/domain/auth/repository"
	. "learnyscape-backend-mono/internal/domain/shared/dto"
	. "learnyscape-backend-mono/internal/domain/shared/entity"
	. "learnyscape-backend-mono/internal/domain/shared/httperror"
	tokenutil "learnyscape-backend-mono/internal/domain/shared/util/token"
	redisx "learnyscape-backend-mono/internal/shared/redis"
	"learnyscape-backend-mono/pkg/mq"
	encryptutil "learnyscape-backend-mono/pkg/util/encrypt"
	jwtutil "learnyscape-backend-mono/pkg/util/jwt"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuthService interface {
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	Refresh(ctx context.Context, req *RefreshRequest) (*LoginResponse, error)
	Verify(ctx context.Context, req *VerificationRequest) (*VerificationResponse, error)
	ResendVerification(ctx context.Context, req *ResendVerificationRequest) error
	ForgotPassword(ctx context.Context, req *ForgotPasswordRequest) error
}

type authServiceImpl struct {
	config                    *config.AuthConfig
	dataStore                 repository.AuthDataStore
	redis                     redisx.RedisClient
	hasher                    encryptutil.Hasher
	jwt                       jwtutil.JWTUtil
	sendVerificationPublisher mq.AMQPPublisher
	accountVerifiedPublisher  mq.AMQPPublisher
	forgotPasswordPublisher   mq.AMQPPublisher
}

func NewAuthService(
	config *config.AuthConfig,
	datastore repository.AuthDataStore,
	redis redisx.RedisClient,
	hasher encryptutil.Hasher,
	jwt jwtutil.JWTUtil,
	sendVerificationPublisher mq.AMQPPublisher,
	accountVerifiedPublisher mq.AMQPPublisher,
	forgotPasswordPublisher mq.AMQPPublisher,
) AuthService {
	return &authServiceImpl{
		config:                    config,
		dataStore:                 datastore,
		redis:                     redis,
		hasher:                    hasher,
		jwt:                       jwt,
		sendVerificationPublisher: sendVerificationPublisher,
		accountVerifiedPublisher:  accountVerifiedPublisher,
		forgotPasswordPublisher:   forgotPasswordPublisher,
	}
}

func (s *authServiceImpl) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	user, err := s.dataStore.UserRepository().FindByIdentifier(ctx, req.Identifier)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, NewInvalidCredentialError()
	}
	if req.IsEmail() && !user.IsVerified {
		return nil, NewEmailNotVerifiedError()
	}

	if ok := s.hasher.Check(req.Password, user.HashPassword); !ok {
		return nil, NewInvalidCredentialError()
	}

	jwtPayload := &jwtutil.JWTPayload{
		UserID: user.ID,
		Role:   user.Role,
	}

	accessToken, err := s.jwt.SignAccess(jwtPayload)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwt.SignRefresh(jwtPayload)
	if err != nil {
		return nil, err
	}

	refreshClaims, err := s.jwt.ParseRefresh(refreshToken)
	if err != nil {
		return nil, err
	}
	if err := s.redis.Set(
		ctx,
		fmt.Sprintf(constant.RefreshTokenKey, refreshClaims.ID),
		user.ID,
		time.Duration(s.config.RefreshTokenDuration)*time.Minute,
	); err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authServiceImpl) Refresh(ctx context.Context, req *RefreshRequest) (*LoginResponse, error) {
	claims, err := s.jwt.ParseRefresh(req.RefreshToken)
	if err != nil {
		return nil, NewInvalidRefreshTokenError()
	}

	if claims.UserID == 0 {
		return nil, NewInvalidRefreshTokenError()
	}

	if claims.Role == "" {
		return nil, NewInvalidRefreshTokenError()
	}

	refreshTokenKey := fmt.Sprintf(constant.RefreshTokenKey, claims.ID)
	userID, err := s.redis.Get(ctx, refreshTokenKey)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, NewInvalidRefreshTokenError()
		}
		return nil, err
	}

	if err := s.redis.Delete(ctx, refreshTokenKey); err != nil {
		return nil, err
	}

	payload := &jwtutil.JWTPayload{
		UserID: claims.UserID,
		Role:   claims.Role,
	}
	accessToken, err := s.jwt.SignAccess(payload)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwt.SignRefresh(payload)
	if err != nil {
		return nil, err
	}

	newClaims, err := s.jwt.ParseRefresh(refreshToken)
	if err != nil {
		return nil, err
	}

	if err := s.redis.Set(
		ctx,
		fmt.Sprintf(constant.RefreshTokenKey, newClaims.ID),
		userID,
		time.Duration(s.config.RefreshTokenDuration)*time.Minute,
	); err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authServiceImpl) Verify(ctx context.Context, req *VerificationRequest) (*VerificationResponse, error) {
	var res *VerificationResponse
	err := s.dataStore.WithinTx(ctx, func(ds repository.AuthDataStore) error {
		userRepo := ds.UserRepository()
		verificationRepo := ds.VerificationRepository()

		user, err := userRepo.FindByIdentifier(ctx, req.Email)
		if err != nil {
			return err
		}
		if user == nil {
			return NewUserNotFoundError()
		}
		if user.IsVerified {
			return NewUserAlreadyVerifiedError()
		}

		verification, err := verificationRepo.FindByUserID(ctx, user.ID)
		if err != nil {
			return err
		}

		if verification.ExpireAt.Before(time.Now()) {
			return NewVerificationTokenExpiredError()
		}
		if verification.Token != req.Token {
			return NewInvalidVerificationTokenError()
		}

		user.IsVerified = true
		if err := userRepo.VerifyByUserID(ctx, user.ID); err != nil {
			return err
		}
		if err := verificationRepo.DeleteByUserID(ctx, verification.UserID); err != nil {
			return err
		}
		if err := s.accountVerifiedPublisher.Publish(ctx, &AccountVerifiedEvent{
			Email: user.Email,
			Name:  user.FullName,
		}); err != nil {
			return err
		}

		res = ToVerificationResponse(user)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authServiceImpl) ResendVerification(ctx context.Context, req *ResendVerificationRequest) error {
	err := s.dataStore.WithinTx(ctx, func(ds repository.AuthDataStore) error {
		userRepo := ds.UserRepository()
		verificationRepo := ds.VerificationRepository()

		user, err := userRepo.FindByIdentifier(ctx, req.Email)
		if err != nil {
			return err
		}
		if user == nil {
			return NewUserNotFoundError()
		}
		if user.IsVerified {
			return NewUserAlreadyVerifiedError()
		}

		if token, err := s.redis.Get(
			ctx,
			fmt.Sprintf(constant.VerificationTokenKey, user.ID),
		); token != "" && err != redis.Nil {
			return NewVerificationCooldownError()
		}

		params := &CreateVerificationParams{
			UserID:   user.ID,
			Token:    tokenutil.GenerateOTPCode(),
			ExpireAt: time.Now().Add(time.Duration(s.config.AccountVerificationTokenDuration) * time.Minute),
		}
		if err := s.redis.Set(
			ctx,
			fmt.Sprintf(constant.VerificationTokenKey, params.UserID),
			params.Token,
			time.Duration(s.config.AccountVerificationTokenCooldownDuration)*time.Minute,
		); err != nil {
			return err
		}

		if err := verificationRepo.DeleteByUserID(ctx, user.ID); err != nil {
			return err
		}
		verification, err := verificationRepo.Create(ctx, params)
		if err != nil {
			return err
		}

		if err := s.sendVerificationPublisher.Publish(ctx, &SendVerificationEvent{
			Email: user.Email,
			Name:  user.FullName,
			Token: verification.Token,
		}); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *authServiceImpl) ForgotPassword(ctx context.Context, req *ForgotPasswordRequest) error {
	err := s.dataStore.WithinTx(ctx, func(ds repository.AuthDataStore) error {
		userRepo := ds.UserRepository()
		resetPasswordRepo := ds.ResetPasswordRepository()

		user, err := userRepo.FindByEmail(ctx, req.Email)
		if err != nil {
			return err
		}
		if user == nil {
			return nil
		}

		if token, err := s.redis.Get(
			ctx,
			fmt.Sprintf(constant.ResetPasswordTokenKey, user.ID),
		); token != "" && err != redis.Nil {
			return NewForgotPasswordCooldownError()
		}

		token, err := resetPasswordRepo.FindUnexpiredTokenByUserID(ctx, user.ID)
		if err != nil {
			return err
		}
		if token != nil {
			if err := s.redis.Set(
				ctx,
				fmt.Sprintf(constant.ResetPasswordTokenKey, token.UserID),
				token.Token,
				time.Duration(s.config.ResetPasswordTokenCooldownDuration)*time.Minute,
			); err != nil {
				return err
			}

			if err := s.forgotPasswordPublisher.Publish(ctx, &ForgotPasswordEvent{
				Email: user.Email,
				Name:  user.FullName,
				Token: token.Token,
			}); err != nil {
				return err
			}
			return nil
		}

		params := &CreateResetPasswordTokenParams{
			UserID:      user.ID,
			Token:       tokenutil.GenerateOTPCode(),
			TokenExpiry: time.Now().Add(time.Duration(s.config.ResetPasswordTokenDuration) * time.Minute),
		}
		if err := s.redis.Set(
			ctx,
			fmt.Sprintf(constant.ResetPasswordTokenKey, params.UserID),
			params.Token,
			time.Duration(s.config.ResetPasswordTokenCooldownDuration)*time.Minute,
		); err != nil {
			return err
		}

		token, err = resetPasswordRepo.Create(ctx, params)
		if err != nil {
			return err
		}

		if err := s.forgotPasswordPublisher.Publish(ctx, &ForgotPasswordEvent{
			Email: user.Email,
			Name:  user.FullName,
			Token: token.Token,
		}); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
