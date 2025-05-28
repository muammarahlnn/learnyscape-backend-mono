package repository

import (
	"context"
	"database/sql"
	"errors"
	"learnyscape-backend-mono/internal/domain/shared/entity"
	"learnyscape-backend-mono/internal/shared/datastore"
)

type UserRepository interface {
	FindByIdentifier(ctx context.Context, identifier string) (*entity.User, error)
	Create(ctx context.Context, params *entity.CreateUserParams) (*entity.User, error)
	VerifyByUserID(ctx context.Context, userID int64) error
}

type userRepositoryImpl struct {
	db datastore.DBTX
}

func NewUserRepository(db datastore.DBTX) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) FindByIdentifier(ctx context.Context, identifier string) (*entity.User, error) {
	query := `
	SELECT
		u.id,
		u.username,
		u.email,
		u.hash_password,
		u.full_name,
		u.profile_pic_url,
		u.is_verified,
		r.name
	FROM
		users u
	JOIN
		roles r
	ON
		u.role_id = r.id
		AND r.deleted_at IS NULL
	WHERE
		(u.username = $1 OR u.email = $1)
		AND u.deleted_at IS NULL
	`

	var user entity.User
	err := r.db.QueryRowContext(ctx, query, identifier).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.HashPassword,
		&user.FullName,
		&user.ProfilePicURL,
		&user.IsVerified,
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

func (r *userRepositoryImpl) Create(ctx context.Context, params *entity.CreateUserParams) (*entity.User, error) {
	query := `
	INSERT INTO 
		users (
			username,
			email,
			hash_password,
			full_name,
			role_id
		)
	VALUES
		($1, $2, $3, $4, $5)
	RETURNING
		id,
		username,
		email,
		full_name,
		role_id,
		created_at,
		updated_at
	`

	var (
		user   entity.User
		roleID int64
	)
	if err := r.db.QueryRowContext(
		ctx,
		query,
		params.Username,
		params.Email,
		params.HashPassword,
		params.FullName,
		params.RoleID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FullName,
		&roleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	query = `
	SELECT
		name
	FROM
		roles
	WHERE
		id = $1
		AND deleted_at IS NULL
	`
	if err := r.db.QueryRowContext(ctx, query, roleID).Scan(&user.Role); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) VerifyByUserID(ctx context.Context, userID int64) error {
	query := `
	UPDATE
		users
	SET
		is_verified = true,
		updated_at = NOW()
	WHERE
		id = $1
		AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, query, userID); err != nil {
		return err
	}

	return nil
}
