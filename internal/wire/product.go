package wire

import (
	controller "go-ecommerce-backend-api/internal/controller"
	repoIml "go-ecommerce-backend-api/internal/repository/implements"
	serviceIml "go-ecommerce-backend-api/internal/services/implements"
	"go-ecommerce-backend-api/pkg/database"
)

func InitProductRouterHandler() *controller.ProductController {
	productRepo := repoIml.NewProductRepository()
	transactor := database.NewTransactor()
	productSvc := serviceIml.NewProductService(productRepo, transactor)
	productCtrl := controller.NewProductController(productSvc)
	return productCtrl
}
