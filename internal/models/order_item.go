package models

type OrderItem struct {
	ID              int     `json:"id" db:"id"`
	OrderID         string  `json:"order_id" db:"order_id"`
	VariantID       int     `json:"variant_id" db:"variant_id"`
	Quantity        int     `json:"quantity" db:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase" db:"price_at_purchase"`
}
