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

func (cr *categoryRepositoryImpl) GetCategoryByID(ctx context.Context, id int) (*models.Category, error) {
	var category models.Category
	query := `SELECT id, parent_id, name, slug FROM categories WHERE id = $1`
	err := cr.pdb.GetContext(ctx, &category, query, id)
	return &category, err
}

func (cr *categoryRepositoryImpl) CreateCategory(ctx context.Context, category models.Category) error {
	query := `INSERT INTO categories (parent_id, name, slug)
		VALUES ($1, $2, $3) RETURNING id`
	err := cr.pdb.QueryRowContext(ctx, query,
		category.ParentID,
		category.Name,
		category.Slug,
	).Scan(&category.ID)
	return err
}