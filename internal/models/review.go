package models

import "time"

type Review struct {
	ID        int       `json:"id" db:"id"`
	ProductID string    `json:"product_id" db:"product_id"`
	UserID    string    `json:"-" db:"user_id"`
	Fullname  string    `json:"full_name" db:"full_name"`
	Rating    int       `json:"rating" db:"rating"`
	Comment   string    `json:"comment" db:"comment"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
