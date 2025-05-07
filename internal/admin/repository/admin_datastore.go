package repository

import "learnyscape-backend-mono/internal/data"

type AdminDataStore interface {
	data.DataStore
	RoleRepository() RoleRepository
}

type adminDataStore struct {
	data.DataStore
}

func NewAdminDataStore(ds data.DataStore) AdminDataStore {
	return &adminDataStore{
		DataStore: ds,
	}
}

func (ds *adminDataStore) RoleRepository() RoleRepository {
	return NewRoleRepository(ds.DB())
}
