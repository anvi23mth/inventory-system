package service

import (
	"context"
	"errors"

	"github.com/anvi23mth/inventory-system/internal/model"
)

// ProductCategoryRepository defines the interface for the database layer [cite: 197]
type ProductCategoryRepository interface {
	Create(ctx context.Context, c model.ProductCategory) error
	GetAll(ctx context.Context) ([]model.ProductCategory, error)
	GetByID(ctx context.Context, id string) (model.ProductCategory, error)
	Update(ctx context.Context, id string, c model.ProductCategory) error
	Delete(ctx context.Context, id string) error
} // ProductCategoryService handles business rules for categories [cite: 200, 223]
type ProductCategoryService struct {
	Repo ProductCategoryRepository
}

// NewProductCategoryService initializes the service with a repository [cite: 212]
func NewProductCategoryService(r ProductCategoryRepository) *ProductCategoryService {
	return &ProductCategoryService{Repo: r}
}

// CreateCategory implements business validation for Week 5 [cite: 251]
func (s *ProductCategoryService) CreateCategory(ctx context.Context, c model.ProductCategory) error {
	// Validation: Title is required [cite: 219, 252]
	if c.Title == "" {
		return errors.New("category title is required")
	}
	return s.Repo.Create(ctx, c)
}

func (s *ProductCategoryService) ListCategories(ctx context.Context) ([]model.ProductCategory, error) {
	return s.Repo.GetAll(ctx)
}
func (s *ProductCategoryService) UpdateCategory(ctx context.Context, id string, c model.ProductCategory) error {
	if c.Title == "" {
		return errors.New("category title is required")
	}
	return s.Repo.Update(ctx, id, c)
}

// GetByID retrieves a single category by its ID
func (s *ProductCategoryService) GetByID(ctx context.Context, id string) (model.ProductCategory, error) {
	category, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return model.ProductCategory{}, err // This is the line that must turn GREEN
	}
	return category, nil // FIX: Return the actual category found
}

// DeleteCategory removes a category by its ID
func (s *ProductCategoryService) DeleteCategory(ctx context.Context, id string) error {
	err := s.Repo.Delete(ctx, id)
	return err // This captures the error for 100% coverage
}
