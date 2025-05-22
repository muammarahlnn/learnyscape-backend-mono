package dto

import "learnyscape-backend-mono/internal/domain/auth/constant"

type SendVerificationEvent struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (e SendVerificationEvent) Key() string {
	return constant.SendVerificationKey
}

type AccountVerifiedEvent struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (e AccountVerifiedEvent) Key() string {
	return constant.AccountVerifiedKey
}
