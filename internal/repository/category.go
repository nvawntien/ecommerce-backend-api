package repository

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
)

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) ([]models.Category, error)
	GetCategoryByID(ctx context.Context, id int) (*models.Category, error)
	CreateCategory(ctx context.Context, category models.Category) error
	UpdateCategory(ctx context.Context, categoryID int, updatedData models.Category) error
	DeleteCategory(ctx context.Context, categoryID int) error
}
