package repository

import "learnyscape-backend-mono/internal/data"

type AuthDataStore interface {
	data.DataStore
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

func (ds *authDataStore) UserRepository() UserRepository {
	return NewUserRepository(ds.DB())
}
