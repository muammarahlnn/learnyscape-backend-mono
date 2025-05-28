package repository

import (
	"context"
	"learnyscape-backend-mono/internal/domain/shared/repository"
	"learnyscape-backend-mono/internal/shared/datastore"
)

type AdminDataStore interface {
	datastore.DataStore
	WithinTx(ctx context.Context, fn func(AdminDataStore) error) error
	RoleRepository() RoleRepository
	UserRepository() repository.UserRepository
	VerificationRepository() repository.VerificationRepository
}

type adminDataStore struct {
	datastore.DataStore
}

func NewAdminDataStore(ds datastore.DataStore) AdminDataStore {
	return &adminDataStore{
		DataStore: ds,
	}
}

func (ds *adminDataStore) WithinTx(ctx context.Context, fn func(AdminDataStore) error) error {
	return datastore.WithinTx(ctx, ds.DataStore, NewAdminDataStore, fn)
}

func (ds *adminDataStore) RoleRepository() RoleRepository {
	return NewRoleRepository(ds.DB())
}

func (ds *adminDataStore) UserRepository() repository.UserRepository {
	return repository.NewUserRepository(ds.DB())
}

func (ds *adminDataStore) VerificationRepository() repository.VerificationRepository {
	return repository.NewVerificationRepository(ds.DB())
}
