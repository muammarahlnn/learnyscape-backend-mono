package entity

import "time"

type User struct {
	ID            int64
	Username      string
	Email         string
	HashPassword  string
	FullName      string
	ProfilePicURL *string
	Role          string
	IsVerified    bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CreateUserParams struct {
	Username     string
	Email        string
	HashPassword string
	FullName     string
	RoleID       int64
}

type SearchUserParams struct {
	Query string
	Limit int64
	Page  int64
}

type UpdateUserParams struct {
	ID       int64
	Username string
	Email    string
	FullName string
}

type ChangePasswordParams struct {
	UserID          int64
	NewHashPassword string
}
