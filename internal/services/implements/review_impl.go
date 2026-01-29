package implements

import (
	"context"
	"fmt"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/pkg/request"
)

type reviewServiceImpl struct {
	reviewRepo repository.ReviewRepository
}

func NewReviewService(reviewRepo repository.ReviewRepository) services.ReviewService {
	return &reviewServiceImpl{
		reviewRepo: reviewRepo,
	}
}

func (r *reviewServiceImpl) CreateReview(ctx context.Context, req request.CreateReviewRequest, userID string) error {
	review := &models.Review{
		ProductID: req.ProductID,
		UserID:    userID,
		Rating:    req.Rating,
		Comment:   req.Comment,
	}

	return r.reviewRepo.CreateReview(ctx, review)
}

func (r *reviewServiceImpl) GetProductReviews(ctx context.Context, productID string) ([]models.Review, error) {
	reviews, err := r.reviewRepo.GetReviewsByProductID(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("faild to get product reviews:with error: %v", err)
	}
	return reviews, nil
}