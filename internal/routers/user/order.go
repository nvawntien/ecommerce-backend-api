package user

import (
	"go-ecommerce-backend-api/internal/middlewares"
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

func (or *OrderRouter) InitOrderRouter(r *gin.RouterGroup) {
	orderCtrl := wire.InitOrderRouterHandler()

	order := r.Group("/orders")

	{
		order.POST("/", middlewares.IsAuthenticated(), orderCtrl.CreateOrder)
		order.GET("/:id", middlewares.IsAuthenticated(), orderCtrl.GetOrderDetail)
		order.GET("/mine", middlewares.IsAuthenticated(), orderCtrl.GetMyOrders)
	}
}
