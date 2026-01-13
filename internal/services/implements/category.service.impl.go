package implements

import (
	"context"
	"fmt"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/pkg/request"
	"github.com/gosimple/slug"
)

type categoryServiceImpl struct {
	cateRepo repository.CategoryRepository
}

func NewCategoryService(cateRepo repository.CategoryRepository) services.CategoryService {
	return &categoryServiceImpl{
		cateRepo: cateRepo,
	}
}

func (cs *categoryServiceImpl) GetAllCategories(ctx context.Context) ([]*models.Category, error) { 
    categories, err := cs.cateRepo.GetAllCategories(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to get categories: %w", err)
    }

    nodes := make(map[int]*models.Category)
    for i := range categories {
        categories[i].Children = []*models.Category{} 
        nodes[categories[i].ID] = &categories[i]
    }

    var tree []*models.Category 

    for i := range categories {
        cat := &categories[i]
        if cat.ParentID == nil {
            tree = append(tree, cat)
        } else {
            if parent, ok := nodes[*cat.ParentID]; ok {
                parent.Children = append(parent.Children, cat)
            }
        }
    }

    return tree, nil
}

func (cs *categoryServiceImpl) CreateCategory(ctx context.Context, req request.CreateCategoryRequest) error {
	if req.ParentID != nil {
		_, err := cs.cateRepo.GetCategoryByID(ctx, *req.ParentID)
		if err != nil {
			return fmt.Errorf("parent category not found: %w", err)
		}
	}

	category := models.Category{
		ParentID: req.ParentID,
		Name: req.Name,	
		Slug: slug.Make(req.Name),
	}

	if err := cs.cateRepo.CreateCategory(ctx, category); err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}
	
	return nil
}
