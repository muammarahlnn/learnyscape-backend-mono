package repository

import (
	"context"
	"learnyscape-backend-mono/internal/data"
)

type AuthDataStore interface {
	data.DataStore
	Atomic(ctx context.Context, fn func(AuthDataStore) error) error
	UserRepository() UserRepository
}

type authDataStore struct {
	data.DataStore
}

func NewAuthDataStore(ds data.DataStore) AuthDataStore {
	return &authDataStore{
		DataStore: ds,
	}
}

func (ds *authDataStore) Atomic(ctx context.Context, fn func(AuthDataStore) error) error {
	return data.Atomic(ctx, ds.DataStore, NewAuthDataStore, fn)
}

func (ds *authDataStore) UserRepository() UserRepository {
	return NewUserRepository(ds.DB())
}
