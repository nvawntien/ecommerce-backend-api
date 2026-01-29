package repository

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
)

type ReviewRepository interface {
	CreateReview(ctx context.Context, review *models.Review) error

	GetReviewsByProductID(ctx context.Context, productID string) ([]models.Review, error)
}
