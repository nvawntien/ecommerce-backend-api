package wire

import (
	controller "go-ecommerce-backend-api/internal/controller"
	repositoryImpl "go-ecommerce-backend-api/internal/repository/implements"
	serviceImpl "go-ecommerce-backend-api/internal/services/implements"
)

func InitReviewRouterHandler() *controller.ReviewController {
	reviewRepo := repositoryImpl.NewReviewRepository()
	reviewSvc := serviceImpl.NewReviewService(reviewRepo)
	reviewCtrl := controller.NewReviewController(reviewSvc)

	return reviewCtrl
}
