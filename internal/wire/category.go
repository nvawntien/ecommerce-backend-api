package wire

import (
	controller "go-ecommerce-backend-api/internal/controller"
	repoIml "go-ecommerce-backend-api/internal/repository/implements"
	serviceIml "go-ecommerce-backend-api/internal/services/implements"
)

func InitCategoryRouterHandler() *controller.CategoryController {
	cateRepo := repoIml.NewCategoryRepository()
	cateSvc := serviceIml.NewCategoryService(cateRepo)
	cateCtrl := controller.NewCategoryController(cateSvc)
	return cateCtrl
}
