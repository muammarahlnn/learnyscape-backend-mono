package repository

import (
	"context"
	"learnyscape-backend-mono/internal/domain/admin/entity"
	"learnyscape-backend-mono/internal/shared/datastore"
)

type RoleRepository interface {
	GetAll(ctx context.Context) ([]*entity.Role, error)
}

type roleRepositoryImpl struct {
	db datastore.DBTX
}

func NewRoleRepository(db datastore.DBTX) RoleRepository {
	return &roleRepositoryImpl{
		db: db,
	}
}

func (r *roleRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Role, error) {
	query := `
	SELECT
		id,
		name,
		created_at,
		updated_at
	FROM
		roles
	WHERE
		deleted_at IS NULL
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*entity.Role
	for rows.Next() {
		var role entity.Role

		if err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.CreatedAt,
			&role.UpdatedAt,
		); err != nil {
			return nil, err
		}

		roles = append(roles, &role)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
