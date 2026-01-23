package implements

import (
	"context"
	"database/sql"
	"fmt"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/pkg/database"
	"go-ecommerce-backend-api/pkg/request"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type productServiceImpl struct {
	productRepo repository.ProductRepository
	transactor  database.Transactor
}

func NewProductService(productRepo repository.ProductRepository, transactor database.Transactor) services.ProductService {
	return &productServiceImpl{
		productRepo: productRepo,
		transactor:  transactor,
	}
}

func (p *productServiceImpl) CreateProduct(ctx context.Context, req request.CreateProductRequest) error {
	return p.transactor.WithTransaction(ctx, func(txCtx context.Context) error {
		product := &models.Product{
			ID:          uuid.NewString(),
			CategoryID:  req.CategoryID,
			Name:        req.Name,
			Slug:        slug.Make(req.Name),
			Description: req.Description,
			Brand:       req.Brand,
			BasePrice:   req.BasePrice,
		}

		if err := p.productRepo.CreateProduct(txCtx, product); err != nil {
			return fmt.Errorf("faild to create product: %w", err)
		}

		for _, variants := range req.Variants {
			variant := &models.ProductVariant{
				ID:            uuid.NewString(),
				ProductID:     product.ID,
				SKU:           variants.SKU,
				VariantName:   variants.VariantName,
				PriceModifier: variants.PriceModifier,
				StockQuantity: variants.StockQuantity,
				ImageURL:      variants.ImageURL,
			}

			if err := p.productRepo.CreateProductVariant(txCtx, variant); err != nil {
				return fmt.Errorf("failde to create product variant: %w", err)
			}
		}
		return nil
	})
}

func (p *productServiceImpl) GetProduct(ctx context.Context, productID string) (*models.Product, error) {
	product, err := p.productRepo.GetProductByID(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	variants, err := p.productRepo.GetVariantsByProductID(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product variants: %w", err)
	}

	product.Variants = variants
	return product, nil
}

func (p *productServiceImpl) GetListProducts(ctx context.Context, filter request.ProductListRequest) (*models.ProductListData, error) {
	products, total, err := p.productRepo.GetListProducts(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get list products: %w", err)
	}

	totalPages := (total + filter.Limit - 1) / filter.Limit

	result := &models.ProductListData{
		Data:       products,
		Total:      total,
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalPages: totalPages,
	}

	return result, nil
}

func (p *productServiceImpl) UpdateProduct(ctx context.Context, productID string, req request.UpdateProductRequest) error {
	currentProduct, err := p.productRepo.GetProductByID(ctx, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("product not found")
		}
		return fmt.Errorf("failed to get current product: %w", err)
	}

	if req.CategoryID != nil {
		currentProduct.CategoryID = *req.CategoryID
	}

	if req.Name != nil {
		currentProduct.Name = *req.Name
		currentProduct.Slug = slug.Make(*req.Name)
	}

	if req.Description != nil {
		currentProduct.Description = *req.Description
	}

	if req.Brand != nil {
		currentProduct.Brand = *req.Brand
	}

	if req.BasePrice != nil {
		currentProduct.BasePrice = *req.BasePrice
	}

	if err := p.productRepo.UpdateProduct(ctx, currentProduct); err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}
