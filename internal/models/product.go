package models

type Product struct {
	ID          string           `json:"id" db:"id"`
	CategoryID  int              `json:"category_id" db:"category_id"`
	Name        string           `json:"name" db:"name"`
	Slug        string           `json:"slug" db:"slug"`
	Description string           `json:"description" db:"description"`
	Brand       string           `json:"brand" db:"brand"`
	BasePrice   float64          `json:"base_price" db:"base_price"`
	CreatedAt   string           `json:"created_at" db:"created_at"`
	UpdatedAt   string           `json:"updated_at" db:"updated_at"`
	Variants    []ProductVariant `json:"variants"`
}
