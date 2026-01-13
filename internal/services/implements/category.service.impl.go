package implements

import (
	"context"
	"fmt"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"
	"go-ecommerce-backend-api/internal/services"
)

type categoryServiceImpl struct {
	cateRepo repository.CategoryRepository
}

func NewCategoryService(cateRepo repository.CategoryRepository) services.CategoryService {
	return &categoryServiceImpl{
		cateRepo: cateRepo,
	}
}

func (cs *categoryServiceImpl) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	categories, err := cs.cateRepo.GetAllCategories(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}

	nodes := make(map[int]*models.Category)

	for i := range categories {
		categories[i].Children = []models.Category{}
		nodes[categories[i].ID] = &categories[i]
	}

	var tree []models.Category

	for _, child := range nodes {
		if child.ParentID == nil {
			tree = append(tree, *child)
		} else {
			if parent, ok := nodes[*child.ParentID]; ok {
				parent.Children = append(parent.Children, *child)
			}
		}
	}

	return tree, nil
}
