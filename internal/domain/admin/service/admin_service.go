package service

import (
	"context"
	"learnyscape-backend-mono/internal/config"
	. "learnyscape-backend-mono/internal/domain/admin/dto"
	"learnyscape-backend-mono/internal/domain/admin/httperror"
	"learnyscape-backend-mono/internal/domain/admin/repository"
	. "learnyscape-backend-mono/internal/domain/shared/dto"
	"learnyscape-backend-mono/internal/domain/shared/entity"
	tokenutil "learnyscape-backend-mono/internal/domain/shared/util/token"
	"learnyscape-backend-mono/pkg/mq"
	encryptutil "learnyscape-backend-mono/pkg/util/encrypt"
	"time"
)

type AdminService interface {
	GetRoles(ctx context.Context) ([]*RoleResponse, error)
	CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error)
	GetAllUsers(ctx context.Context) ([]*UserResponse, error)
}

type adminServiceimpl struct {
	config                    *config.AdminConfig
	dataStore                 repository.AdminDataStore
	hasher                    encryptutil.Hasher
	sendVerificationPublisher mq.AMQPPublisher
}

func NewAdminService(
	config *config.AdminConfig,
	dataSstore repository.AdminDataStore,
	hasher encryptutil.Hasher,
	sendVerificationPublisher mq.AMQPPublisher,
) AdminService {
	return &adminServiceimpl{
		config:                    config,
		dataStore:                 dataSstore,
		hasher:                    hasher,
		sendVerificationPublisher: sendVerificationPublisher,
	}
}

func (s *adminServiceimpl) GetRoles(ctx context.Context) ([]*RoleResponse, error) {
	roleRepo := s.dataStore.RoleRepository()

	roles, err := roleRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return ToRoleResponses(roles), nil
}

func (s *adminServiceimpl) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
	var res *UserResponse
	err := s.dataStore.WithinTx(ctx, func(ds repository.AdminDataStore) error {
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

		verification, err := verificationRepo.Create(ctx, &entity.CreateVerificationParams{
			UserID:   user.ID,
			Token:    tokenutil.GenerateOTPCode(),
			ExpireAt: time.Now().Add(time.Duration(s.config.AccountVerificationTokenDuration) * time.Minute),
		})
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

		res = ToUserResponse(user)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *adminServiceimpl) GetAllUsers(ctx context.Context) ([]*UserResponse, error) {
	users, err := s.dataStore.UserRepository().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return ToUserResponses(users), nil
}
