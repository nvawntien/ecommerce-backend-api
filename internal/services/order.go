package services

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/pkg/request"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req request.CreateOrderRequest) (*models.Order, error)
}
