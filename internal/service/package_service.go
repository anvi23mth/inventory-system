package service

import (
	"context"
	"errors" // Required for validation errors

	"github.com/anvi23mth/inventory-system/internal/model"
	"github.com/anvi23mth/inventory-system/internal/repository"
)

// ProductService must start with a Capital 'P' to be exported
type ProductService struct {
	Repo repository.ProductRepository
}

// NewProductService initializes the service
func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{Repo: r}
}

func (s *ProductService) CreateProduct(ctx context.Context, p model.Product) (model.Product, error) {
	// --- WEEK 3 VALIDATION LOGIC ---
	if p.Price < 0 {
		return model.Product{}, errors.New("price cannot be negative")
	}
	if p.Name == "" {
		return model.Product{}, errors.New("product name is required")
	}
	// -------------------------------

	err := s.Repo.Create(ctx, p)
	return p, err
}
func (s *ProductService) ListProducts(ctx context.Context) ([]model.Product, error) {
	return s.Repo.GetAll(ctx)
}

func (s *ProductService) GetProductByID(ctx context.Context, id string) (model.Product, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id string, p model.Product) (model.Product, error) {
	err := s.Repo.Update(ctx, id, p)
	return p, err
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}
