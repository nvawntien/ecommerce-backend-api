package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Init Logger complete", zap.String("Init", "Success"))
	m := global.Config.PostgreSQL
	fmt.Println("Loading PostgreSQL with config:", m.Username, m.Host, m.DBName)
	InitPostgres()
	r := InitRouter()
	return r
}
