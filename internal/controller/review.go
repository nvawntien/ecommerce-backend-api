package controller

import (
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/pkg/request"
	"go-ecommerce-backend-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReviewController struct {
	reviewSvc services.ReviewService
}

func NewReviewController(reviewSvc services.ReviewService) *ReviewController {
	return &ReviewController{
		reviewSvc: reviewSvc,
	}
}

func (rc *ReviewController) CreateReview(c *gin.Context) {
	var req request.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "Dữ liệu không hợp lệ")
		return
	}

	userID := c.GetString("user_id")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "Invalid token claims")
		return
	}

	err := rc.reviewSvc.CreateReview(c.Request.Context(), req, userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Create review successfully", nil)
}

func (rc *ReviewController) GetProductReviews(c *gin.Context) {
	productID := c.Param("id")
	if _, err := uuid.Parse(productID); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "productID is invalid")
		return
	}

	reviews, err := rc.reviewSvc.GetProductReviews(c.Request.Context(), productID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Get product reviews successfully", reviews)
}
