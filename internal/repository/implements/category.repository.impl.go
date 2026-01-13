package implements

import (
	"context"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"

	"github.com/jmoiron/sqlx"
)

type categoryRepositoryImpl struct {
	pdb *sqlx.DB
}

func NewCategoryRepository() repository.CategoryRepository {
	return &categoryRepositoryImpl{
		pdb: global.Pdbx,
	}
}

func (cr *categoryRepositoryImpl) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	query := `SELECT id, parent_id, name, slug FROM categories ORDER BY id ASC`
	err := cr.pdb.SelectContext(ctx, &categories, query)
	return categories, err
}
