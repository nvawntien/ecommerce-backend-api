package models

import (
	"time"
)

type User struct {
	UserID   string    `json:"user_id" db:"user_id"`
	Email    string    `json:"email" db:"email" binding:"required,email"`
	Username string    `json:"username" db:"username" binding:"required"`
	Password string    `json:"-" db:"password" binding:"required,min=8"`
	Role     string    `json:"role" db:"role"`
	CreateAt time.Time `json:"create_at" db:"create_at"`
	UpdateAt time.Time `json:"update_at" db:"update_at"`
}
