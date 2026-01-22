package controller

import (
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/pkg/request"
	"go-ecommerce-backend-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController struct {
	productSvc services.ProductService
}

func NewProductController(productSvc services.ProductService) *ProductController {
	return &ProductController{
		productSvc: productSvc,
	}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var req request.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "Dữ liệu không hợp lệ")
		return
	}

	if err := pc.productSvc.CreateProduct(c.Request.Context(), req); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Tạo sản phẩm thành công", nil)
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	// products/:id
	productID := c.Param("id")
	
	if _, err := uuid.Parse(productID); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "ID sản phẩm không hợp lệ")
		return
	}

	product, err := pc.productSvc.GetProduct(c.Request.Context(), productID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Lấy sản phẩm thành công", product)	
}