package user

import (
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
}

func (cr *CategoryRouter) InitCategoryRouter(r *gin.RouterGroup) {
	cateCtrl := wire.InitCategoryRouterHandler()

	category := r.Group("/categories")
	{
		category.GET("/", cateCtrl.GetAllCategories)
	}
}