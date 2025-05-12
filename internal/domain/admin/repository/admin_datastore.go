package repository

import "learnyscape-backend-mono/internal/shared/datastore"

type AdminDataStore interface {
	datastore.DataStore
	RoleRepository() RoleRepository
}

type adminDataStore struct {
	datastore.DataStore
}

func NewAdminDataStore(ds datastore.DataStore) AdminDataStore {
	return &adminDataStore{
		DataStore: ds,
	}
}

func (ds *adminDataStore) RoleRepository() RoleRepository {
	return NewRoleRepository(ds.DB())
}
