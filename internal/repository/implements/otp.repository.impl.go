package implements

import (
	"context"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

type otpRepositoryImpl struct {
	rdb *redis.Client
}

func NewOTPRepository() repository.OTPRepository {
	return &otpRepositoryImpl{
		rdb: global.Rdb,
	}
}

func (or *otpRepositoryImpl) SetOTP(ctx context.Context, email string, code string, expiration time.Duration) error {
	key := "otp:" + email
	return or.rdb.Set(ctx, key, code, expiration).Err()
}


func (or *otpRepositoryImpl) GetOTP(ctx context.Context, email string) (string, error) {
	key := "otp:" + email
	return or.rdb.Get(ctx, key).Result()
}

func (or *otpRepositoryImpl) DeleteOTP(ctx context.Context, email string) error {
	key := "otp:" + email
	return or.rdb.Del(ctx, key).Err()
}

