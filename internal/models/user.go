package models

import (
	"time"
)

type User struct {
	UserID    string    `json:"user_id" db:"user_id"`
	Fullname  string    `json:"full_name" db:"full_name" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required,email"`
	Password  string    `json:"-" db:"password" binding:"required,min=8"`
	Role      string    `json:"role" db:"role"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
