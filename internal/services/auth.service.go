package services

import (
	"context"
	"go-ecommerce-backend-api/pkg/request"
)

type AuthService interface {
	Register(ctx context.Context, req request.RegisterRequest) error
	VerifyOTP(ctx context.Context, req request.VerifyOTPRequest) error
	ResendOTP(ctx context.Context, req request.ResendOTPRequest) error
}
