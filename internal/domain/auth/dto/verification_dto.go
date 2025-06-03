package dto

import (
	"learnyscape-backend-mono/internal/domain/auth/constant"
	"learnyscape-backend-mono/internal/domain/shared/entity"
)

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

type AccountVerifiedEvent struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (e AccountVerifiedEvent) Key() string {
	return constant.AccountVerifiedKey
}

type ForgotPasswordEvent struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (e ForgotPasswordEvent) Key() string {
	return constant.ForgotPasswordKey
}
