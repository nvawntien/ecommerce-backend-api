package models

type ProductVariant struct {
	ID            string  `json:"id" db:"id"`
	ProductID     string  `json:"product_id" db:"product_id"`
	SKU           string  `json:"sku" db:"sku"`
	VariantName   string  `json:"variant_name" db:"variant_name"`
	PriceModifier float64 `json:"price_modifier" db:"price_modifier"`
	StockQuantity int     `json:"stock_quantity" db:"stock_quantity"`
	ImageURL      string  `json:"image_url" db:"image_url"`
}
