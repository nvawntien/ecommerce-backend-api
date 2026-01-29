package user

import (
	"go-ecommerce-backend-api/internal/middlewares"
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type ReviewRouter struct {
}

func (rt *ReviewRouter) InitReviewRouter(r *gin.RouterGroup) {
	reviewCtrl := wire.InitReviewRouterHandler()

	review := r.Group("/reviews")
	{
		review.GET("/product/:id", reviewCtrl.GetProductReviews)
		review.POST("/", middlewares.IsAuthenticated(), reviewCtrl.CreateReview)
	}
}
