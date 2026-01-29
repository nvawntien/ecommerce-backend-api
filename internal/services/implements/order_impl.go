package implements

import (
	"context"
	"fmt"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/pkg/database"
	"go-ecommerce-backend-api/pkg/request"
	"time"

	"github.com/google/uuid"
)

type orderServiceImpl struct {
	orderRepo  repository.OrderRepository
	transactor database.Transactor
}

func NewOrderService(orderRepo repository.OrderRepository, transactor database.Transactor) services.OrderService {
	return &orderServiceImpl{
		orderRepo:  orderRepo,
		transactor: transactor,
	}
}

func (os *orderServiceImpl) CreateOrder(ctx context.Context, req request.CreateOrderRequest) (*models.Order, error) {
	var totalAmount float64
	var itemsToSave []models.OrderItem

	for _, itemReq := range req.Items {
		vInfo, err := os.orderRepo.GetVariantInfo(ctx, itemReq.VariantID)
		if err != nil {
			return nil, fmt.Errorf("not found product variant: %w", err)
		}

		if vInfo.StockQuantity < itemReq.Quantity {
			return nil, fmt.Errorf("product '%s - %s' is out of stock (remaining: %d)", vInfo.ProductName, vInfo.VariantName, vInfo.StockQuantity)
		}

		finalPrice := vInfo.BasePrice + vInfo.PriceModifier
		if finalPrice < 0 {
			finalPrice = 0
		}

		totalAmount += finalPrice * float64(itemReq.Quantity)

		itemsToSave = append(itemsToSave, models.OrderItem{
			VariantID:       itemReq.VariantID,
			Quantity:        itemReq.Quantity,
			PriceAtPurchase: finalPrice,
		})
	}

	order := &models.Order{
		ID:              uuid.NewString(),
		UserID:          req.UserID,
		TotalAmount:     totalAmount,
		Status:          "pending",
		ShippingAddress: req.ShippingAddress,
		PaymentMethod:   req.PaymentMethod,
		CreatedAt:       time.Now(),
	}

	err := os.transactor.WithTransaction(ctx, func (txCtx context.Context) error {
		if err := os.orderRepo.CreateOrder(txCtx, order); err != nil {
			return fmt.Errorf("faild to create order:%w", err)
		}

		for _, item := range itemsToSave {
			item.OrderID = order.ID
			if err := os.orderRepo.CreateOrderItem(txCtx, &item); err != nil {
				return fmt.Errorf("faild to create order item: %w", err)
			}

			if err := os.orderRepo.DecreaseStock(txCtx, item.VariantID, item.Quantity); err != nil {
				return fmt.Errorf("faild to decrease stock quantity: %w", err)
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	
	return order, nil
} 

func (os *orderServiceImpl) GetOrderDetail(ctx context.Context, orderID string) (*models.OrderDetail, error) {
	order, err := os.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}

	items, err := os.orderRepo.GetOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order items: %w", err)
	}

	orderDetail := &models.OrderDetail{
		Order: *order,
		Items: items,
	}

	return orderDetail, nil
}