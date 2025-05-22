package repository

import (
	"context"
	"learnyscape-backend-mono/internal/shared/datastore"
)

type AuthDataStore interface {
	datastore.DataStore
	WithinTx(ctx context.Context, fn func(AuthDataStore) error) error
	UserRepository() UserRepository
	VerificationRepository() VerificationRepository
}

type authDataStore struct {
	datastore.DataStore
}

func NewAuthDataStore(ds datastore.DataStore) AuthDataStore {
	return &authDataStore{
		DataStore: ds,
	}
}

func (ds *authDataStore) WithinTx(ctx context.Context, fn func(AuthDataStore) error) error {
	return datastore.WithinTx(ctx, ds.DataStore, NewAuthDataStore, fn)
}

func (ds *authDataStore) UserRepository() UserRepository {
	return NewUserRepository(ds.DB())
}

func (ds *authDataStore) VerificationRepository() VerificationRepository {
	return NewVerificationRepository(ds.DB())
}
