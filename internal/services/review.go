package services

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/pkg/request"
)

type ReviewService interface {
	CreateReview(ctx context.Context, req request.CreateReviewRequest, userID string) error

	GetProductReviews(ctx context.Context, productID string) ([]models.Review, error)
}
