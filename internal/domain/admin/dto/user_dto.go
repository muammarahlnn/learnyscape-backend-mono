package dto

import (
	"learnyscape-backend-mono/internal/domain/shared/entity"
	"time"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
	FullName string `json:"full_name" binding:"required"`
	RoleID   int64  `json:"role_id" binding:"required"`
}

type UserResponse struct {
	ID            int64     `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	FullName      string    `json:"full_name"`
	ProfilePicURL *string   `json:"profile_pic_url"`
	Role          string    `json:"role"`
	IsVerified    bool      `json:"is_verified"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

func ToUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		FullName:      user.FullName,
		ProfilePicURL: user.ProfilePicURL,
		Role:          user.Role,
		IsVerified:    user.IsVerified,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
