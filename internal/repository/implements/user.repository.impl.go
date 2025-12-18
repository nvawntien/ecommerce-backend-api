package implements

import (
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/repository"

	"github.com/jmoiron/sqlx"
)

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository() repository.UserRepository {
	return &userRepositoryImpl{
		db: global.Pdbx,
	}
}
