package service

import (
	"context"
	"errors"
	"fmt"
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/domain/auth/constant"
	"learnyscape-backend-mono/internal/domain/auth/dto"
	"learnyscape-backend-mono/internal/domain/auth/entity"
	"learnyscape-backend-mono/internal/domain/auth/httperror"
	"learnyscape-backend-mono/internal/domain/auth/repository"
	tokenutil "learnyscape-backend-mono/internal/domain/auth/util/token"
	redisx "learnyscape-backend-mono/internal/shared/redis"
	"learnyscape-backend-mono/pkg/mq"
	encryptutil "learnyscape-backend-mono/pkg/util/encrypt"
	jwtutil "learnyscape-backend-mono/pkg/util/jwt"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuthService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Refresh(ctx context.Context, req *dto.RefreshRequest) (*dto.LoginResponse, error)
	Verify(ctx context.Context, req *dto.VerificationRequest) (*dto.VerificationResponse, error)
}

type authServiceImpl struct {
	dataStore                 repository.AuthDataStore
	hasher                    encryptutil.Hasher
	jwt                       jwtutil.JWTUtil
	redis                     redisx.RedisClient
	config                    *config.AuthConfig
	sendVerificationPublisher mq.AMQPPublisher
	accountVerifiedPublisher  mq.AMQPPublisher
}

func NewAuthService(
	datastore repository.AuthDataStore,
	hasher encryptutil.Hasher,
	jwt jwtutil.JWTUtil,
	redis redisx.RedisClient,
	config *config.AuthConfig,
	sendVerificationPublisher mq.AMQPPublisher,
	accountVerifiedPublisher mq.AMQPPublisher,
) AuthService {
	return &authServiceImpl{
		dataStore:                 datastore,
		hasher:                    hasher,
		jwt:                       jwt,
		redis:                     redis,
		config:                    config,
		sendVerificationPublisher: sendVerificationPublisher,
		accountVerifiedPublisher:  accountVerifiedPublisher,
	}
}

func (s *authServiceImpl) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.dataStore.UserRepository().FindByIdentifier(ctx, req.Identifier)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, httperror.NewInvalidCredentialError()
	}

	if ok := s.hasher.Check(req.Password, user.HashPassword); !ok {
		return nil, httperror.NewInvalidCredentialError()
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

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authServiceImpl) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	var res *dto.RegisterResponse
	err := s.dataStore.WithinTx(ctx, func(ds repository.AuthDataStore) error {
		userRepo := ds.UserRepository()
		verificationRepo := ds.VerificationRepository()

		user, err := userRepo.FindByIdentifier(ctx, req.Username)
		if err != nil {
			return err
		}
		if user != nil {
			return httperror.NewUserAlreadyExistsError()
		}

		user, err = userRepo.FindByIdentifier(ctx, req.Email)
		if err != nil {
			return err
		}
		if user != nil {
			return httperror.NewUserAlreadyExistsError()
		}

		hashedPassword, err := s.hasher.Hash(req.Password)
		if err != nil {
			return err
		}

		user, err = userRepo.Create(ctx, &entity.CreateUserParams{
			Username:     req.Username,
			Email:        req.Email,
			HashPassword: hashedPassword,
			FullName:     req.FullName,
			RoleID:       req.RoleID,
		})
		if err != nil {
			return err
		}

		params := &entity.CreateVerificationsParams{
			UserID:   user.ID,
			Token:    tokenutil.GenerateOTPCode(),
			ExpireAt: time.Now().Add(time.Duration(s.config.AccountVerificationTokenDuration) * time.Minute),
		}
		verification, err := verificationRepo.Create(ctx, params)
		if err != nil {
			return err
		}

		if err := s.sendVerificationPublisher.Publish(ctx, &dto.SendVerificationEvent{
			Email: user.Email,
			Name:  user.FullName,
			Token: verification.Token,
		}); err != nil {
			return err
		}

		res = dto.ToRegisterResponse(user)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authServiceImpl) Refresh(ctx context.Context, req *dto.RefreshRequest) (*dto.LoginResponse, error) {
	claims, err := s.jwt.ParseRefresh(req.RefreshToken)
	if err != nil {
		return nil, httperror.NewInvalidRefreshTokenError()
	}

	if claims.UserID == 0 {
		return nil, httperror.NewInvalidRefreshTokenError()
	}

	if claims.Role == "" {
		return nil, httperror.NewInvalidRefreshTokenError()
	}

	refreshTokenKey := fmt.Sprintf(constant.RefreshTokenKey, claims.ID)
	userID, err := s.redis.Get(ctx, refreshTokenKey)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, httperror.NewInvalidRefreshTokenError()
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

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authServiceImpl) Verify(ctx context.Context, req *dto.VerificationRequest) (*dto.VerificationResponse, error) {
	var res *dto.VerificationResponse
	err := s.dataStore.WithinTx(ctx, func(ds repository.AuthDataStore) error {
		userRepo := ds.UserRepository()
		verificationRepo := ds.VerificationRepository()

		user, err := userRepo.FindByIdentifier(ctx, req.Email)
		if err != nil {
			return err
		}
		if user == nil {
			return httperror.NewUserNotFoundError()
		}
		if user.IsVerified {
			return httperror.NewUserAlreadyVerifiedError()
		}

		verification, err := verificationRepo.FindByUserID(ctx, user.ID)
		if err != nil {
			return err
		}

		if verification.ExpireAt.Before(time.Now()) {
			return httperror.NewVerificationTokenExpiredError()
		}
		if verification.Token != req.Token {
			return httperror.NewInvalidVerificationTokenError()
		}

		user.IsVerified = true
		if err := userRepo.VerifyByUserID(ctx, user.ID); err != nil {
			return err
		}
		if err := verificationRepo.DeleteByID(ctx, verification.ID); err != nil {
			return err
		}
		if err := s.accountVerifiedPublisher.Publish(ctx, &dto.AccountVerifiedEvent{
			Email: user.Email,
			Name:  user.FullName,
		}); err != nil {
			return err
		}

		res = dto.ToVerificationResponse(user)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
