package service

import (
	"context"
	"testing"

	"github.com/anvi23mth/inventory-system/internal/model"
)

// 1. Define the Mock to match the Repository Interface exactly
type MockRepository struct{}

func (m *MockRepository) Create(ctx context.Context, p model.Product) error {
	return nil
}

func (m *MockRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	return []model.Product{
		{ID: "test-1", Name: "Mock Product"},
	}, nil
}

func (m *MockRepository) GetByID(ctx context.Context, id string) (model.Product, error) {
	return model.Product{ID: id, Name: "Mock Product"}, nil
}

func (m *MockRepository) Update(ctx context.Context, id string, p model.Product) error {
	return nil
}

func (m *MockRepository) Delete(ctx context.Context, id string) error {
	return nil
}

// 2. The Test Function
func TestListProducts(t *testing.T) {
	// Initialize Mock and Service
	// Note: We cast MockRepository to the expected interface type if necessary
	mockRepo := &MockRepository{}

	// Use your actual NewProductService
	// Note: Ensure repository.ProductRepository is an interface in your repo file
	// This should now work because the Service accepts the Interface
	svc := NewProductService(mockRepo)
	ctx := context.Background()

	// Call the CORRECT method name from your service: ListProducts
	products, err := svc.ListProducts(ctx)

	// 3. Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(products))
	}

	if products[0].Name != "Mock Product" {
		t.Errorf("Expected 'Mock Product', got %s", products[0].Name)
	}
}
func TestCreateProduct_Validation(t *testing.T) {
	mockRepo := &MockRepository{}
	svc := NewProductService(mockRepo)
	ctx := context.Background()

	// Define multiple scenarios (Table-Driven Test)
	tests := []struct {
		name    string
		product model.Product
		wantErr bool
	}{
		{
			name: "Valid Product",
			product: model.Product{
				Name:  "Gaming Mouse",
				Price: 25.00,
			},
			wantErr: false,
		},
		{
			name: "Invalid Price - Negative",
			product: model.Product{
				Name:  "Broken Mouse",
				Price: -5.00,
			},
			wantErr: true,
		},
		{
			name: "Invalid Name - Empty",
			product: model.Product{
				Name:  "",
				Price: 10.00,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.CreateProduct(ctx, tt.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
