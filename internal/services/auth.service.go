package services

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/pkg/request"
)

type AuthService interface {
	Register(ctx context.Context, req request.RegisterRequest) error
	VerifyOTP(ctx context.Context, req request.VerifyOTPRequest) error
	ResendOTP(ctx context.Context, req request.ResendOTPRequest) error
	Login(ctx context.Context, req request.LoginRequest) (*models.User, string, string, error)
	RefreshToken(ctx context.Context, userID string, userRole string) (string, string, error)
	ForgotPassword(ctx context.Context, req request.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req request.ResetPasswordRequest) error
	ChangePassword(ctx context.Context, userID string, req request.ChangePasswordRequest) error
}
