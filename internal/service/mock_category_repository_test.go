package service

import (
	"context"
	"errors"

	"github.com/anvi23mth/inventory-system/internal/model"
)

// MockCategoryRepository manually mocks the interface to isolate Service logic
type MockCategoryRepository struct {
	Categories []model.ProductCategory
	ShouldFail bool // Add this toggle
}

func (m *MockCategoryRepository) Create(ctx context.Context, c model.ProductCategory) error {
	if m.ShouldFail {
		return errors.New("db error")
	}
	return nil
}
func (m *MockCategoryRepository) GetAll(ctx context.Context) ([]model.ProductCategory, error) {
	if m.ShouldFail {
		return nil, errors.New("db error")
	}
	return m.Categories, nil
}

func (m *MockCategoryRepository) GetByID(ctx context.Context, id string) (model.ProductCategory, error) {
	if m.ShouldFail {
		return model.ProductCategory{}, errors.New("category not found")
	}
	return model.ProductCategory{ID: id, Title: "Mock Category"}, nil
}
func (m *MockCategoryRepository) Update(ctx context.Context, id string, c model.ProductCategory) error {
	if m.ShouldFail {
		return errors.New("db error")
	}
	return nil
}
func (m *MockCategoryRepository) Delete(ctx context.Context, id string) error {
	if m.ShouldFail {
		return errors.New("db error")
	}
	return nil
}
