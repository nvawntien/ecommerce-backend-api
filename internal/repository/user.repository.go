package repository

import (
	"context"
	"go-ecommerce-backend-api/internal/models"
)

type UserRepository interface {
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, user *models.User) error
	ActiveUserByEmail(ctx context.Context, email string) error
	GetNameByEmail(ctx context.Context, email string) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserIDByEmail(ctx context.Context, email string) (string, error)
	UpdatePasswordByUserID(ctx context.Context, userID string, newPassword string) error
}
