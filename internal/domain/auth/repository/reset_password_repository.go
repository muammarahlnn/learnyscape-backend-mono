package repository

import (
	"context"
	"database/sql"
	"learnyscape-backend-mono/internal/domain/auth/entity"
	"learnyscape-backend-mono/internal/shared/datastore"
)

type ResetPasswordRepository interface {
	Create(ctx context.Context, params *entity.CreateResetPasswordTokenParams) (*entity.ResetPasswordToken, error)
	FindUnexpiredTokenByUserID(ctx context.Context, userId int64) (*entity.ResetPasswordToken, error)
}

type resetPasswordRepositoryImpl struct {
	db datastore.DBTX
}

func NewResetPasswordRepository(db datastore.DBTX) ResetPasswordRepository {
	return &resetPasswordRepositoryImpl{
		db: db,
	}
}

func (r *resetPasswordRepositoryImpl) Create(ctx context.Context, params *entity.CreateResetPasswordTokenParams) (*entity.ResetPasswordToken, error) {
	query := `
	INSERT INTO
		reset_password_tokens (
			user_id,
			token,
			token_expiry
		)
	VALUES
		($1, $2, $3)
	RETURNING
		id,
		user_id,
		token,
		token_expiry,
		created_at,
		updated_at
	`

	var token entity.ResetPasswordToken
	if err := r.db.QueryRowContext(
		ctx,
		query,
		params.UserID,
		params.Token,
		params.TokenExpiry,
	).Scan(
		&token.ID,
		&token.UserID,
		&token.Token,
		&token.TokenExpiry,
		&token.CreatedAt,
		&token.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *resetPasswordRepositoryImpl) FindUnexpiredTokenByUserID(ctx context.Context, userId int64) (*entity.ResetPasswordToken, error) {
	query := `
	SELECT
		id,
		user_id,
		token,
		token_expiry,
		created_at,
		updated_at
	FROM
		reset_password_tokens
	WHERE
		user_id = $1 
		AND token_expiry > NOW()
		AND deleted_at IS NULL
	`

	var token entity.ResetPasswordToken
	if err := r.db.QueryRowContext(
		ctx,
		query,
		userId,
	).Scan(
		&token.ID,
		&token.UserID,
		&token.Token,
		&token.TokenExpiry,
		&token.CreatedAt,
		&token.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &token, nil
}
