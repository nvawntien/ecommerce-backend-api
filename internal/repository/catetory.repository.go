package repository

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
)

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) ([]models.Category, error)
}
