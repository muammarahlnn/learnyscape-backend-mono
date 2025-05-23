package dto

import (
	"learnyscape-backend-mono/internal/domain/auth/entity"
	"time"
)

type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
	FullName string `json:"full_name" binding:"required"`
	RoleID   int64  `json:"role_id" binding:"required"`
}

type RegisterResponse struct {
	ID            int64     `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	FullName      string    `json:"full_name"`
	ProfilePicURL *string   `json:"profile_pic_url"`
	Role          string    `json:"role"`
	IsVerified    bool      `json:"is_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToRegisterResponse(user *entity.User) *RegisterResponse {
	return &RegisterResponse{
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

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type VerificationRequest struct {
	Email string `json:"email" binding:"required,email"`
	Token string `json:"token" binding:"required"`
}

type VerificationResponse struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}

func ToVerificationResponse(user *entity.User) *VerificationResponse {
	return &VerificationResponse{
		ID:         user.ID,
		Email:      user.Email,
		IsVerified: user.IsVerified,
	}
}

type ResendVerificationRequest struct {
	Email string `json:"email" binding:"required,email"`
}
