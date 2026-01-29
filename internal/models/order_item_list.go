package models

type OrderItemList struct {
	ID              int     `json:"id" db:"id"`
	VariantID       int     `json:"variant_id" db:"variant_id"`
	ProductName     string  `json:"product_name" db:"product_name"`
	VariantName     string  `json:"variant_name" db:"variant_name"`
	SKU             string  `json:"sku" db:"sku"`
	Quantity        int     `json:"quantity" db:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase" db:"price_at_purchase"`
	ImageURL        string  `json:"image_url" db:"image_url"`
}
