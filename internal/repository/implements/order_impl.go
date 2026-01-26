package implements

import (
	"context"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"
	"go-ecommerce-backend-api/pkg/database"

	"github.com/jmoiron/sqlx"
)

type orderRepositoryImpl struct {
	db *sqlx.DB
}

func NewOrderRepository() repository.OrderRepository {
	return &orderRepositoryImpl{
		db: global.Pdbx,
	}
}

func (or *orderRepositoryImpl) GetVariantInfo(ctx context.Context, variantID int) (*models.VariantInfo, error) {
	executor := database.GetExecutor(ctx, or.db)

	query := `
		SELECT pv.id, pv.stock_quantity, pv.price_modifier, pv.variant_name, p.base_price, p.name as product_name
		FROM product_variants pv 
		JOIN products p ON p.id = pv.product_id
		WHERE pv.id = $1
	`

	info := &models.VariantInfo{}

	if err := executor.GetContext(ctx, info, query, variantID); err != nil {
		return nil, err
	}

	return info, nil
}

func (or *orderRepositoryImpl) CreateOrder(ctx context.Context, order *models.Order) error {
	executor := database.GetExecutor(ctx, or.db)
	query := `
		INSERT INTO orders (id, user_id, total_amount, status, shipping_address, payment_method, created_at)
		VALUES (:id, :user_id, :total_amount, :status, :shipping_address, :payment_method, :created_at)
	`
	_, err := executor.NamedExecContext(ctx, query, order)
	return err
}

func (or *orderRepositoryImpl) CreateOrderItem(ctx context.Context, item *models.OrderItem) error {
	executor := database.GetExecutor(ctx, or.db)
	query := `
		INSERT INTO order_items (order_id, variant_id, quantity, price_at_purchase)
		VALUES (:order_id, :variant_id, :quantity, :price_at_purchase)
	`

	_, err := executor.NamedExecContext(ctx, query, item)
	return err
}

func (or *orderRepositoryImpl) DecreaseStock(ctx context.Context, variantID int, quantity int) error {
	executor := database.GetExecutor(ctx, or.db)
	query := `
		UPDATE product_variants SET stock_quantity = stock_quantity - $1 WHERE id = $2	
	`
	_, err := executor.ExecContext(ctx, query, quantity, variantID)
	return err
}