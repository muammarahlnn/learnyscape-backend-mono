package repository

import (
	"context"
	"learnyscape-backend-mono/internal/data"
)

type UserRepository interface {
	Test(ctx context.Context) (string, error)
}

type userRepositoryImpl struct {
	db data.DBTX
}

func NewUserRepository(db data.DBTX) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) Test(ctx context.Context) (string, error) {
	query := `
	SELECT NOW()
	`

	var now string
	err := r.db.QueryRowContext(ctx, query).Scan(&now)
	if err != nil {
		return "", err
	}

	return now, nil
}
