package implements

import (
	"context"
	"database/sql"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"
	"go-ecommerce-backend-api/pkg/database"
	"go-ecommerce-backend-api/pkg/request"

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

func (pr *productRepositoryImpl) GetListProducts(ctx context.Context, filter request.ProductListRequest) ([]models.Product, int, error) {
	query := `SELECT * FROM products WHERE 1=1`
	countQuery := `SELECT count(*) FROM products WHERE 1=1`
	args := make(map[string]interface{})

	if filter.Keyword != "" {
		condition := ` AND name ILIKE: keyword`
		query += condition
		countQuery += condition
		args["keyword"] = "%" + filter.Keyword + "%"
	}

	if filter.CategoryID != 0 {
		condition := ` AND category_id = :category_id`
		query += condition
		countQuery += condition
		args["category_id"] = filter.CategoryID
	}

	var total int
	rowsCount, err := pr.db.NamedQueryContext(ctx, countQuery, args)
	if err != nil {
		return nil, 0, err
	}

	defer rowsCount.Close()

	if rowsCount.Next() {
		if err := rowsCount.Scan(&total); err != nil {
			return nil, 0, err
		}
	}

	offset := (filter.Page - 1) * filter.Limit
	args["limit"] = filter.Limit
	args["offset"] = offset

	query += ` ORDER BY created_at DESC LIMIT :limit OFFSET :offset`

	var products []models.Product
	rows, err := pr.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.StructScan(&product); err != nil {
			return nil, 0, err
		}
		products = append(products, product)
	}

	return products, total, nil
}

func (pr *productRepositoryImpl) UpdateProduct(ctx context.Context, product *models.Product) error {
	query := `UPDATE products SET category_id = :category_id, name = :name, slug = :slug,
		description = :description, brand = :brand, base_price = :base_price, updated_at = :updated_at
		WHERE id = :id`

	result, err := pr.db.NamedExecContext(ctx, query, product)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
