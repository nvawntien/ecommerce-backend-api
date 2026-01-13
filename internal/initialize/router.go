package initialize

import (
	"go-ecommerce-backend-api/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	userRouter := routers.RouterGroupApp.UserRouter
	adminRouter := routers.RouterGroupApp.AdminRouter
	mainGroup := r.Group("/api/v1")

	{
		userRouter.InitAuthRouter(mainGroup)
		userRouter.InitCategoryRouter(mainGroup)
	}

	{
		adminRouter.InitCategoryRouter(mainGroup)
	}
	return r
}
