package repository

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/pkg/request"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	
	CreateProductVariant(ctx context.Context, variant *models.ProductVariant) error
	
	GetProductByID(ctx context.Context, productID string) (*models.Product, error)
	
	GetVariantsByProductID(ctx context.Context, productID string) ([]models.ProductVariant, error)
	
	GetListProducts(ctx context.Context, filter request.ProductListRequest) ([]models.Product, int, error)
	
	UpdateProduct(ctx context.Context, product *models.Product) error 
	
	DeleteProductVariantsByID(ctx context.Context, productID string) error
	
	DeleteProductByID(ctx context.Context, productID string) error
}
