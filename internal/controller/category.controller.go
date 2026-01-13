package controller

import (
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/pkg/request"
	"go-ecommerce-backend-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	cateSvc services.CategoryService
}

func NewCategoryController(cateSvc services.CategoryService) *CategoryController {
	return &CategoryController{
		cateSvc: cateSvc,
	}
}

func (cc *CategoryController) GetAllCategories(c *gin.Context) {
	tree, err := cc.cateSvc.GetAllCategories(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Lấy danh sách danh mục thành công", tree)
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var req request.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "Dữ liệu không hợp lệ")
		return
	}

	if err := cc.cateSvc.CreateCategory(c.Request.Context(), req); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Tạo danh mục thành công", nil)
}
