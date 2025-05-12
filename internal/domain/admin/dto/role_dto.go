package dto

import "learnyscape-backend-mono/internal/domain/admin/entity"

type RoleResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func ToRoleResponse(role *entity.Role) *RoleResponse {
	return &RoleResponse{
		ID:   role.ID,
		Name: role.Name,
	}
}

func ToRoleResponses(roles []*entity.Role) []*RoleResponse {
	res := make([]*RoleResponse, len(roles))

	for i, role := range roles {
		res[i] = ToRoleResponse(role)
	}

	return res
}
