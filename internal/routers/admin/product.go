package admin

import (
	"go-ecommerce-backend-api/internal/middlewares"
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(r *gin.RouterGroup) {
	productCtrl := wire.InitProductRouterHandler()

	product := r.Group("/products")
	{  
		product.POST("/create", middlewares.IsAuthenticated(), middlewares.IsAdmin(), productCtrl.CreateProduct)
		product.PUT("/update/:id", middlewares.IsAuthenticated(), middlewares.IsAdmin(), productCtrl.UpdateProduct)
		product.DELETE("/:id", middlewares.IsAuthenticated(), middlewares.IsAdmin(), productCtrl.DeleteProduct)
	}
}
