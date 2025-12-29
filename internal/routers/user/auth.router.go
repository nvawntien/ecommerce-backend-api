package user

import (
	"go-ecommerce-backend-api/internal/middlewares"
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (ar *AuthRouter) InitAuthRouter(r *gin.RouterGroup) {
	authCtrl := wire.InitUserRouterHandler()

	authRouter := r.Group("/auth")

	authRouter.GET("/welcome", authCtrl.Welcome)

	authRouter.POST("/register", authCtrl.Register)

	authRouter.POST("/verify-otp", authCtrl.VerifyOTP)

	authRouter.POST("/resend-otp", authCtrl.ResendOTP)

	authRouter.POST("/login", authCtrl.Login)

	authRouter.POST("/logout", authCtrl.Logout)

	authRouter.POST("/refresh-token", middlewares.IsExpiredRefreshToken(), authCtrl.RefreshToken)

	authRouter.POST("/forgot-password", authCtrl.ForgotPassword)

	authRouter.POST("/reset-password", authCtrl.ResetPassword)
}
