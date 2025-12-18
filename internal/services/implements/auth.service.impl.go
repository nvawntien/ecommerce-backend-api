package implements

import (
	"go-ecommerce-backend-api/internal/repository"
	"go-ecommerce-backend-api/internal/services"
)

type authServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) services.AuthService {
	return &authServiceImpl{
		userRepo: userRepo,
	}
}
