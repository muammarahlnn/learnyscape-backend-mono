package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"learnyscape-backend-mono/internal/domain/shared/entity"
	"learnyscape-backend-mono/internal/shared/datastore"
	pageutil "learnyscape-backend-mono/pkg/util/page"
)

type UserRepository interface {
	FindByIdentifier(ctx context.Context, identifier string) (*entity.User, error)
	Create(ctx context.Context, params *entity.CreateUserParams) (*entity.User, error)
	VerifyByUserID(ctx context.Context, userID int64) error
	Search(ctx context.Context, params *entity.SearchUserParams) ([]*entity.User, int64, error)
	Update(ctx context.Context, params *entity.UpdateUserParams) (*entity.User, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
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

	role, err := r.getRole(ctx, roleID)
	if err != nil {
		return nil, err
	}
	user.Role = role

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

func (r *userRepositoryImpl) Search(ctx context.Context, params *entity.SearchUserParams) ([]*entity.User, int64, error) {
	query := `
	SELECT
		u.id,
		u.username,
		u.email,
		u.full_name,
		u.profile_pic_url,
		u.is_verified,
		r.name,
		COUNT(*) OVER(PARTITION BY 1)
	FROM
		users u
	JOIN
		roles r ON u.role_id = r.id
		AND r.deleted_at IS NULL
	WHERE
		u.deleted_at IS NULL
		AND ($1 = '' OR u.username ILIKE $1)
		AND ($1 = '' OR u.email ILIKE $1)
		AND ($1 = '' OR u.full_name ILIKE $1)
	LIMIT $2 OFFSET $3
	`

	search := fmt.Sprintf("%%%s%%", params.Query)
	offset := pageutil.Offset(params.Page, params.Limit)
	rows, err := r.db.QueryContext(ctx, query, search, params.Limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var (
		users []*entity.User
		total int64
	)
	for rows.Next() {
		var user entity.User

		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.FullName,
			&user.ProfilePicURL,
			&user.IsVerified,
			&user.Role,
			&total,
		); err != nil {
			return nil, 0, err
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepositoryImpl) Update(ctx context.Context, params *entity.UpdateUserParams) (*entity.User, error) {
	query := `
	UPDATE
		users
	SET
		username = $1,
		email = $2,
		full_name = $3,
		updated_at = NOW()
	WHERE
		id = $4
		AND deleted_at IS NULL
	RETURNING
		id,
		username,
		email,
		full_name,
		profile_pic_url,
		is_verified,
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
		params.FullName,
		params.ID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FullName,
		&user.ProfilePicURL,
		&user.IsVerified,
		&roleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	role, err := r.getRole(ctx, roleID)
	if err != nil {
		return nil, err
	}
	user.Role = role

	return &user, nil
}

func (r *userRepositoryImpl) getRole(ctx context.Context, roleID int64) (string, error) {
	query := `
	SELECT
		name
	FROM
		roles
	WHERE
		id = $1
		AND deleted_at IS NULL
	`

	var role string
	if err := r.db.QueryRowContext(ctx, query, roleID).Scan(&role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return role, nil
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id int64) (*entity.User, error) {
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
		u.id = $1
		AND u.deleted_at IS NULL
	`

	var user entity.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
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
