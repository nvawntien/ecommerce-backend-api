package admin

import (
	"go-ecommerce-backend-api/internal/middlewares"
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
}

func (cr *CategoryRouter) InitCategoryRouter(r *gin.RouterGroup) {
	cateCtrl := wire.InitCategoryRouterHandler()

	category := r.Group("/categories")
	{
		category.POST("/",  middlewares.IsAuthenticated(), middlewares.IsAdmin(), cateCtrl.CreateCategory)
		category.PUT("/:id", middlewares.IsAuthenticated(), middlewares.IsAdmin(), cateCtrl.UpdateCategory)
		category.DELETE("/:id", middlewares.IsAuthenticated(), middlewares.IsAdmin(), cateCtrl.DeleteCategory)
	}
}
