package wire

import (
	controller "go-ecommerce-backend-api/internal/controller/user"
	repoimpl "go-ecommerce-backend-api/internal/repository/implements"
	svcimpl "go-ecommerce-backend-api/internal/services/implements"
)

func InitUserRouterHandler() *controller.AuthController {
	userRepo := repoimpl.NewUserRepository()
	authSvc := svcimpl.NewAuthService(userRepo)
	authCtrl := controller.NewAuthController(authSvc)
	return authCtrl
}
