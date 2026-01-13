package wire

import (
	controller "go-ecommerce-backend-api/internal/controller"
	repoimpl "go-ecommerce-backend-api/internal/repository/implements"
	svcimpl "go-ecommerce-backend-api/internal/services/implements"
)

func InitUserRouterHandler() *controller.AuthController {
	userRepo := repoimpl.NewUserRepository()
	otpRepo := repoimpl.NewOTPRepository()
	authSvc := svcimpl.NewAuthService(userRepo, otpRepo)
	authCtrl := controller.NewAuthController(authSvc)
	return authCtrl
}
