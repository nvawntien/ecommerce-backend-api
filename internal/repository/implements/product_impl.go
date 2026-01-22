package implements

import (
	"context"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"
	"go-ecommerce-backend-api/pkg/database"

	"github.com/jmoiron/sqlx"
)

type productRepositoryImpl struct {
	db *sqlx.DB
}

func NewProductRepository() repository.ProductRepository {
	return &productRepositoryImpl{
		db: global.Pdbx,
	}
}

func (pr *productRepositoryImpl) CreateProduct(ctx context.Context, product *models.Product) error {
	executor := database.GetExecutor(ctx, pr.db)
	query := `INSERT INTO products (id, category_id, name, slug, description, brand, base_price)
		VALUES (:id, :category_id, :name, :slug, :description, :brand, :base_price)`
	_, err := executor.NamedExecContext(ctx, query, product)
	return err
}

func (pr *productRepositoryImpl) CreateProductVariant(ctx context.Context, variant *models.ProductVariant) error {
	executor := database.GetExecutor(ctx, pr.db)
	query := `INSERT INTO product_variants (product_id, sku, variant_name, price_modifier, stock_quantity, image_url)
		VALUES (:product_id, :sku, :variant_name, :price_modifier, :stock_quantity, :image_url)`
	_, err := executor.NamedExecContext(ctx, query, variant)
	return err
}

func (pr *productRepositoryImpl) GetProductByID(ctx context.Context, productID string) (*models.Product, error) {
	product := &models.Product{}
	query := `SELECT * FROM products WHERE id = $1`
	err := pr.db.GetContext(ctx, product, query, productID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productRepositoryImpl) GetVariantsByProductID(ctx context.Context, productID string) ([]models.ProductVariant, error) {
	variants := []models.ProductVariant{}
	query := `SELECT * FROM product_variants WHERE product_id = $1`
	err := p.db.SelectContext(ctx, &variants, query, productID)
	if err != nil {
		return nil, err
	}
	return variants, nil
}