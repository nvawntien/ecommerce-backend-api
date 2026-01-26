package models

import "time"

type Order struct {
	ID              string      `json:"id" db:"id"`
	UserID          string      `json:"user_id" db:"user_id"`
	TotalAmount     float64     `json:"total_amount" db:"total_amount"`
	Status          string      `json:"status" db:"status"`
	ShippingAddress string      `json:"shipping_address" db:"shipping_address"`
	PaymentMethod   string      `json:"payment_method" db:"payment_method"`
	CreatedAt       time.Time   `json:"created_at" db:"created_at"`
	Items           []OrderItem `json:"order_item"`
}
