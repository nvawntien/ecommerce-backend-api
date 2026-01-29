package implements

import (
	"context"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"

	"github.com/jmoiron/sqlx"
)

type reviewRepositoryImpl struct {
	db *sqlx.DB
}

func NewReviewRepository() repository.ReviewRepository {
	return &reviewRepositoryImpl{
		db: global.Pdbx,
	}
}

func (r *reviewRepositoryImpl) CreateReview(ctx context.Context, review *models.Review) error {
	query := `INSERT INTO reviews (product_id, user_id, rating, comment, created_at) 
		VALUES (:product_id, :user_id, :rating, :comment, :created_at)
	`

	_, err := r.db.NamedExecContext(ctx, query, review)
	return err
}

func (r *reviewRepositoryImpl) GetReviewsByProductID(ctx context.Context, productID string) ([]models.Review, error) {
	var reviews []models.Review
	query := `
		SELECT 
			r.id, r.product_id, r.user_id, r.rating, r.comment, r.created_at,
			u.full_name  
		FROM reviews r
		JOIN users u ON r.user_id = u.user_id
		WHERE product_id = $1
	`

	err := r.db.SelectContext(ctx, &reviews, query, productID)
	return reviews, err
}