package controller

import (
	"fmt"
	"go-ecommerce-backend-api/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authSvc services.AuthService
}

func NewAuthController(authSvc services.AuthService) *AuthController {
	return &AuthController{
		authSvc: authSvc,
	}
}

func (ac *AuthController) Welcome(c *gin.Context) {
	fmt.Println("Hello World!")
}
