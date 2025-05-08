package service

import (
	"context"
	"learnyscape-backend-mono/internal/auth/dto"
	"learnyscape-backend-mono/internal/auth/entity"
	"learnyscape-backend-mono/internal/auth/httperror"
	"learnyscape-backend-mono/internal/auth/repository"
	encryptutil "learnyscape-backend-mono/pkg/util/encrypt"
	jwtutil "learnyscape-backend-mono/pkg/util/jwt"
)

type AuthService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
}

type authServiceImpl struct {
	dataStore repository.AuthDataStore
	hasher    encryptutil.Hasher
	jwt       jwtutil.JWTUtil
}

func NewAuthService(
	ds repository.AuthDataStore,
	hasher encryptutil.Hasher,
	jwt jwtutil.JWTUtil,
) AuthService {
	return &authServiceImpl{
		dataStore: ds,
		hasher:    hasher,
		jwt:       jwt,
	}
}

func (s *authServiceImpl) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	res := &dto.LoginResponse{}
	err := s.dataStore.Atomic(ctx, func(ds repository.AuthDataStore) error {
		userRepo := ds.UserRepository()

		user, err := userRepo.FindByIdentifier(ctx, req.Identifier)
		if err != nil {
			return err
		}
		if user == nil {
			return httperror.NewInvalidCredentialError()
		}

		if ok := s.hasher.Check(req.Password, user.HashPassword); !ok {
			return httperror.NewInvalidCredentialError()
		}

		token, err := s.jwt.Sign(&jwtutil.JWTPayload{UserID: user.ID, Role: user.Role})
		if err != nil {
			return err
		}

		res.AccessToken = token
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authServiceImpl) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	var res *dto.RegisterResponse
	err := s.dataStore.Atomic(ctx, func(ds repository.AuthDataStore) error {
		userRepo := ds.UserRepository()

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

		params := &entity.CreateUserParams{
			Username:     req.Username,
			Email:        req.Email,
			HashPassword: hashedPassword,
			FullName:     req.FullName,
			RoleID:       req.RoleID,
		}
		user, err = userRepo.Create(ctx, params)
		if err != nil {
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
