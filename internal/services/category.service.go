package services

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
)

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]models.Category, error)
}
