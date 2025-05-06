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
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
