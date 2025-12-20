package controller

import (
	"fmt"
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/pkg/errors"
	"go-ecommerce-backend-api/pkg/request"
	"go-ecommerce-backend-api/pkg/response"
	"net/http"

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

func (ac *AuthController) Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "")
		return
	}

	if err := ac.authSvc.Register(c.Request.Context(), req); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Đăng ký tài khoản thành công. Vui lòng kiểm tra email để xác thực tài khoản.", nil)
}

func (ac *AuthController) VerifyOTP(c *gin.Context) {
	var req request.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "")
		return
	}

	if err := ac.authSvc.VerifyOTP(c.Request.Context(), req); err != nil {
		if err == errors.ErrUserNotFound {
			response.Error(c, http.StatusNotFound, response.CodeUnauthorized, "Người dùng không tồn tại.")
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Xác thực OTP thành công. Tài khoản của bạn đã được kích hoạt.", nil)
}

func (ac *AuthController) ResendOTP(c *gin.Context) {
	var req request.ResendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "")
		return
	}

	if err := ac.authSvc.ResendOTP(c.Request.Context(), req); err != nil {
		if err == errors.ErrUserNotFound {
			response.Error(c, http.StatusNotFound, response.CodeUnauthorized, "Người dùng không tồn tại.")
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Gửi lại mã OTP thành công. Vui lòng kiểm tra email của bạn.", nil)
}
