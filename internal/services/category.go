package services

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/pkg/request"
)

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]*models.Category, error)
	CreateCategory(ctx context.Context, req request.CreateCategoryRequest) error
	UpdateCategory(ctx context.Context, categoryID int, req request.UpdateCategoryRequest) error
	DeleteCategory(ctx context.Context, categoryID int) error
}
