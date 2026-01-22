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

	//public api
	{
		userGroup := mainGroup.Group("/user")
		userRouter.InitAuthRouter(userGroup)
		userRouter.InitCategoryRouter(userGroup)
	}

	//private api
	{
		adminGroup := mainGroup.Group("/admin")
		adminRouter.InitCategoryRouter(adminGroup)
		adminRouter.InitProductRouter(adminGroup)
	}
	return r
}
