package controller

import (
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/pkg/request"
	"go-ecommerce-backend-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderController struct {
	orderSvc services.OrderService
}

func NewOrderController(orderSvc services.OrderService) *OrderController {
	return &OrderController{
		orderSvc: orderSvc,
	}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var req request.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "Dữ liệu không hợp lệ")
		return
	}

	order, err := oc.orderSvc.CreateOrder(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}
	
	response.Success(c, "Create order successfully", order)
}

func (oc *OrderController) GetOrderDetail(c *gin.Context) {
	orderID := c.Param("id")
	if _, err := uuid.Parse(orderID); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "orderID is invalid")
		return
	}

	order, err := oc.orderSvc.GetOrderDetail(c.Request.Context(), orderID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "get order detail successfuly", order)
}