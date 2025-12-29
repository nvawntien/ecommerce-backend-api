package controller

import (
	"fmt"
	"go-ecommerce-backend-api/global"
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

func (ac *AuthController) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "")
		return
	}

	user, accessToken, refreshToken, err := ac.authSvc.Login(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	c.SetCookie("access_token", accessToken, global.Config.JWT.AccessExpiry, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, global.Config.JWT.RefreshExpiry, "/", "localhost", false, true)
	response.Success(c, "Đăng nhập thành công", user)
}

func (ac *AuthController) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	response.Success(c, "Đăng xuất thành công", nil)
}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	userID := c.GetString("user_id")
	userRole := c.GetString("role")

	if userID == "" || userRole == "" {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "Invalid token claims")
		return
	}

	accessToken, refreshToken, err := ac.authSvc.RefreshToken(c.Request.Context(), userID, userRole)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	c.SetCookie("access_token", accessToken, global.Config.JWT.AccessExpiry, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, global.Config.JWT.RefreshExpiry, "/refresh-token", "localhost", false, true)
	response.Success(c, "Làm mới token thành công", nil)
}

func (ac *AuthController) ForgotPassword(c *gin.Context) {
	var req request.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "")
		return
	}

	if err := ac.authSvc.ForgotPassword(c.Request.Context(), req); err != nil {
		if err == errors.ErrUserNotFound {
			response.Error(c, http.StatusNotFound, response.CodeUnauthorized, "Người dùng không tồn tại.")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Gửi mã đặt lại mật khẩu thành công. Vui lòng kiểm tra email của bạn.", nil)
}

func (ac *AuthController) ResetPassword(c *gin.Context) {
	var req request.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "")
		return
	}

	if err := ac.authSvc.ResetPassword(c.Request.Context(), req); err != nil {
		if err == errors.ErrTokenInvalid {
			response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "Token không hợp lệ.")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, "Đặt lại mật khẩu thành công.", nil)
}
