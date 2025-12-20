package models

import (
	"time"
)

type User struct {
	UserID   string    `json:"user_id" db:"user_id"`
	Fullname string    `json:"fullname" db:"fullname" binding:"required"`
	Email    string    `json:"email" db:"email" binding:"required,email"`
	Password string    `json:"-" db:"password" binding:"required,min=8"`
	Role     string    `json:"role" db:"role"`
	IsActive bool      `json:"is_active" db:"is_active"`
	CreateAt time.Time `json:"create_at" db:"create_at"`
	UpdateAt time.Time `json:"update_at" db:"update_at"`
}
