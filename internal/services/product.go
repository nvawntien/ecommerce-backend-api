package services

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/pkg/request"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req request.CreateProductRequest) error
	
	GetProduct(ctx context.Context, productID string) (*models.Product, error)
	
	GetListProducts(ctx context.Context, filter request.ProductListRequest) (*models.ProductListData, error)
	
	UpdateProduct(ctx context.Context, productID string, req request.UpdateProductRequest) error
	
	DeleteProduct(ctx context.Context, productID string) error
}
