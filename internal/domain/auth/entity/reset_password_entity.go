package entity

import "time"

type ResetPasswordToken struct {
	ID          int64
	UserID      int64
	Token       string
	TokenExpiry time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateResetPasswordTokenParams struct {
	UserID      int64
	Token       string
	TokenExpiry time.Time
}
