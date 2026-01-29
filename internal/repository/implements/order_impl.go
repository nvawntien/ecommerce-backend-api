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

func (or *orderRepositoryImpl) GetOrderByID(ctx context.Context, orderID string) (*models.Order, error) {
	executor := database.GetExecutor(ctx, or.db)

	query := `
		SELECT * FROM orders WHERE id = $1
	`

	order := &models.Order{}
	if err := executor.GetContext(ctx, order, query, orderID); err != nil {
		return nil, err
	}

	return order, nil
}

func (or *orderRepositoryImpl) GetOrderItemsByOrderID(ctx context.Context, orderID string) ([]models.OrderItemList, error) {
	executor := database.GetExecutor(ctx, or.db)
	query := `
		SELECT 
			oi.id, oi.variant_id, oi.quantity, oi.price_at_purchase,
			p.name AS product_name,
			pv.variant_name, pv.sku, pv.image_url
		FROM order_items oi
		JOIN product_variants pv ON oi.variant_id = pv.id
		JOIN products p ON pv.product_id = p.id
		WHERE oi.order_id = $1
	`
	var items []models.OrderItemList
	if err := executor.SelectContext(ctx, &items, query, orderID); err != nil {
		return nil, err
	}

	return items, nil
}

func (or *orderRepositoryImpl) GetOrderByUserID(ctx context.Context, userID string) ([]models.Order, error) {
	executor := database.GetExecutor(ctx, or.db)
	query := `
		SELECT * FROM orders WHERE user_id = $1 ORDER BY created_at DESC
	`

	var orders []models.Order
	if err := executor.SelectContext(ctx, &orders, query, userID); err != nil {
		return nil, err
	}

	return orders, nil
}
