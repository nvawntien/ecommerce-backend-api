package services

import (
	"context"
	"go-ecommerce-backend-api/pkg/request"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req request.CreateProductRequest) error
}
