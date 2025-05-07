package service

import (
	"context"
	"learnyscape-backend-mono/internal/admin/dto"
	"learnyscape-backend-mono/internal/admin/repository"
)

type AdminService interface {
	GetAll(ctx context.Context) ([]*dto.RoleResponse, error)
}

type adminServiceimpl struct {
	dataStore repository.AdminDataStore
}

func NewAdminService(ds repository.AdminDataStore) AdminService {
	return &adminServiceimpl{
		dataStore: ds,
	}
}

func (s *adminServiceimpl) GetAll(ctx context.Context) ([]*dto.RoleResponse, error) {
	roleRepo := s.dataStore.RoleRepository()

	roles, err := roleRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToRoleResponses(roles), nil
}
