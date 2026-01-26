package wire

import (
	"go-ecommerce-backend-api/internal/controller"
	repoImpl "go-ecommerce-backend-api/internal/repository/implements"
	svcImpl "go-ecommerce-backend-api/internal/services/implements"
	"go-ecommerce-backend-api/pkg/database"
)	
func InitOrderRouterHandler() *controller.OrderController {
	orderRepo := repoImpl.NewOrderRepository()
	transactor := database.NewTransactor()
	orderSvc :=  svcImpl.NewOrderService(orderRepo, transactor)
	orderCtrl := controller.NewOrderController(orderSvc)
	return orderCtrl 
}
