package service

import (
	"context"
	"learnyscape-backend-mono/internal/auth/repository"
)

type AuthService interface {
	Test(ctx context.Context) (string, error)
}

type authServiceImpl struct {
	dataStore repository.AuthDataStore
}

func NewAuthService(ds repository.AuthDataStore) AuthService {
	return &authServiceImpl{
		dataStore: ds,
	}
}

func (s *authServiceImpl) Test(ctx context.Context) (string, error) {
	userRepository := s.dataStore.UserRepository()

	now, err := userRepository.Test(ctx)
	if err != nil {
		return "", err
	}

	return now, nil
}
