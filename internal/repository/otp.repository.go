package repository

import (
	"context"
	"time"
)

type OTPRepository interface {
	SetOTP(ctx context.Context, email string, code string, expiration time.Duration) error
	GetOTP(ctx context.Context, email string) (string, error)
	DeleteOTP(ctx context.Context, email string) error
}
