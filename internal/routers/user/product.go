package user

import (
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
}

func (p *ProductRouter) InitProductRouter(r *gin.RouterGroup) {
	productCtrl := wire.InitProductRouterHandler()

	product := r.Group("/products")
	{
		product.GET("/:id", productCtrl.GetProduct)
		product.GET("/", productCtrl.GetListProducts)
	}
}
