package repository

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	CreateProductVariant(ctx context.Context, variant *models.ProductVariant) error
}
