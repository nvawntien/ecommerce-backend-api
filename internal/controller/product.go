package controller

import (
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/pkg/request"
	"go-ecommerce-backend-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
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

func (pc *ProductController) GetListProducts(c *gin.Context) {
	var req request.ProductListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "Dữ liệu không hợp lệ")
		return
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}

	if req.Limit > 100 {
		req.Limit = 100
	}

	if req.Page < 1 {
		req.Page = 1
	}

	products, err := pc.productSvc.GetListProducts(c.Request.Context(), req)
	if err != nil {
		//global.Logger.Error("err is here: ", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Lấy danh sách sản phẩm thành công", products)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	if _, err := uuid.Parse(productID); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "ID sản phẩm không hợp lệ")
		return
	}

	var req request.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logger.Error("Bind JSON error: ", zap.Error(err))
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "Dữ liệu không hợp lệ")
		return
	}

	if err := pc.productSvc.UpdateProduct(c.Request.Context(), productID, req); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Cập nhật sản phẩm thành công", nil)
}
