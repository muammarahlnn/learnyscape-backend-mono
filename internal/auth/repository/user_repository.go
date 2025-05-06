package repository

import (
	"context"
	"database/sql"
	"errors"
	"learnyscape-backend-mono/internal/auth/entity"
	"learnyscape-backend-mono/internal/data"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
}

type userRepositoryImpl struct {
	db data.DBTX
}

func NewUserRepository(db data.DBTX) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `
	SELECT
		u.id,
		u.username,
		u.email,
		u.hash_password,
		u.full_name,
		u.profile_pic_url,
		r.name
	FROM
		users u
	JOIN
		roles r
	ON
		u.role_id = r.id
		AND r.deleted_at IS NULL
	WHERE
		u.username = $1
		AND u.deleted_at IS NULL
	`

	var user entity.User
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.HashPassword,
		&user.FullName,
		&user.ProfilePicURL,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
