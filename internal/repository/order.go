package repository

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
)

type OrderRepository interface {
	GetVariantInfo(ctx context.Context, variantID int) (*models.VariantInfo, error)

	CreateOrder(ctx context.Context, order *models.Order) error

	CreateOrderItem(ctx context.Context, item *models.OrderItem) error 

	DecreaseStock(ctx context.Context, variantID int, quantity int) error
}
