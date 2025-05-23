package repository

import (
	"context"
	"learnyscape-backend-mono/internal/domain/auth/entity"
	"learnyscape-backend-mono/internal/shared/datastore"
)

type VerificationRepository interface {
	Create(ctx context.Context, params *entity.CreateVerificationsParams) (*entity.Verification, error)
	FindByUserID(ctx context.Context, userId int64) (*entity.Verification, error)
	DeleteByUserID(ctx context.Context, userId int64) error
}

type verificationRepositoryImpl struct {
	db datastore.DBTX
}

func NewVerificationRepository(db datastore.DBTX) VerificationRepository {
	return &verificationRepositoryImpl{
		db: db,
	}
}

func (r *verificationRepositoryImpl) Create(ctx context.Context, params *entity.CreateVerificationsParams) (*entity.Verification, error) {
	query := `
	INSERT INTO
		user_verifications (
			user_id,
			token,
			expire_at
		)
	VALUES
		($1, $2, $3)
	RETURNING
		id,
		user_id,
		token,
		expire_at,
		created_at,
		updated_at
	`

	var verification entity.Verification
	if err := r.db.QueryRowContext(
		ctx,
		query,
		params.UserID,
		params.Token,
		params.ExpireAt,
	).Scan(
		&verification.ID,
		&verification.UserID,
		&verification.Token,
		&verification.ExpireAt,
		&verification.CreatedAt,
		&verification.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &verification, nil
}

func (r *verificationRepositoryImpl) FindByUserID(ctx context.Context, userId int64) (*entity.Verification, error) {
	query := `
	SELECT
		id,
		user_id,
		token,
		expire_at,
		created_at,
		updated_at
	FROM
		user_verifications
	WHERE
		user_id = $1
		AND deleted_at IS NULL
	`

	var verification entity.Verification
	if err := r.db.QueryRowContext(ctx, query, userId).Scan(
		&verification.ID,
		&verification.UserID,
		&verification.Token,
		&verification.ExpireAt,
		&verification.CreatedAt,
		&verification.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &verification, nil
}

func (r *verificationRepositoryImpl) DeleteByUserID(ctx context.Context, userId int64) error {
	query := `
	UPDATE
		user_verifications
	SET
		deleted_at = NOW()
	WHERE
		user_id = $1
		AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, query, userId); err != nil {
		return err
	}

	return nil
}
