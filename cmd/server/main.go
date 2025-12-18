package main

import (
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/initialize"
)

func main() {
	r := initialize.Run()
	r.Run(":" + global.Config.Server.Port)
}
