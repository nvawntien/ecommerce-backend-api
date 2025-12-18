package user

import (
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (ar *AuthRouter) InitAuthRouter(r *gin.RouterGroup) {
	authCtrl := wire.InitUserRouterHandler()

	authRouter := r.Group("/auth")
	authRouter.GET("/welcome", authCtrl.Welcome)
}
