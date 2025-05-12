package repository

import (
	"context"
	"learnyscape-backend-mono/internal/shared/datastore"
)

type AuthDataStore interface {
	datastore.DataStore
	Atomic(ctx context.Context, fn func(AuthDataStore) error) error
	UserRepository() UserRepository
}

type authDataStore struct {
	datastore.DataStore
}

func NewAuthDataStore(ds datastore.DataStore) AuthDataStore {
	return &authDataStore{
		DataStore: ds,
	}
}

func (ds *authDataStore) Atomic(ctx context.Context, fn func(AuthDataStore) error) error {
	return datastore.Atomic(ctx, ds.DataStore, NewAuthDataStore, fn)
}

func (ds *authDataStore) UserRepository() UserRepository {
	return NewUserRepository(ds.DB())
}
