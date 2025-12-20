package implements

import (
	"context"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"

	"github.com/jmoiron/sqlx"
)

type userRepositoryImpl struct {
	pdb *sqlx.DB
}

func NewUserRepository() repository.UserRepository {
	return &userRepositoryImpl{
		pdb: global.Pdbx,
	}
}

func (ur *userRepositoryImpl) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 from users WHERE email = $1)`
	err := ur.pdb.GetContext(ctx, &exists, query, email)
	return exists, err
}

func (ur *userRepositoryImpl) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (user_id, full_name, email, password, is_active, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := ur.pdb.ExecContext(ctx, query,
		user.UserID,
		user.Fullname,
		user.Email,
		user.Password,
		user.IsActive,
		user.CreateAt,
		user.UpdateAt,
	)

	return err
}

func (ur *userRepositoryImpl) ActiveUserByEmail(ctx context.Context, email string) error {
	query := `UPDATE users SET is_active = true, updated_at = NOW() WHERE email = $1`
	_, err := ur.pdb.ExecContext(ctx, query, email)
	return err
}

func (ur *userRepositoryImpl) GetNameByEmail(ctx context.Context, email string) (string, error) {
	var fullname string
	query := `SELECT full_name FROM users WHERE email = $1`
	err := ur.pdb.GetContext(ctx, &fullname, query, email)
	return fullname, err
}