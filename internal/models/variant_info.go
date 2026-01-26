package models

type VariantInfo struct {
	ID int `db:"id"`
	StockQuantity int `db:"stock_quantity"` 
	PriceModifier float64 `db:"price_modifier"`
	BasePrice float64 `db:"base_price"`
	ProductName string `db:"product_name"`
	VariantName string `db:"variant_name"`
}