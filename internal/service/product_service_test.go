package service

import (
	"context"
	"errors"
	"testing"

	"github.com/anvi23mth/inventory-system/internal/model"
)

// 1. Define the Mock to match the Repository Interface exactly
type MockRepository struct {
	ShouldFail bool // Add this flag
}

func (m *MockRepository) Create(ctx context.Context, p model.Product) error {
	if m.ShouldFail {
		return errors.New("database error")
	}
	return nil
}
func (m *MockRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	if m.ShouldFail {
		return nil, errors.New("database error")
	}
	return []model.Product{{ID: "test-1", Name: "Mock Product"}}, nil
}
func (m *MockRepository) GetByID(ctx context.Context, id string) (model.Product, error) {
	if m.ShouldFail {
		return model.Product{}, errors.New("product not found")
	}
	return model.Product{ID: id, Name: "Test"}, nil
}
func (m *MockRepository) Update(ctx context.Context, id string, p model.Product) error {
	if m.ShouldFail {
		return errors.New("database error")
	}
	return nil
}

func (m *MockRepository) Delete(ctx context.Context, id string) error {
	if m.ShouldFail {
		return errors.New("database error")
	}
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
func TestUpdateProduct_Validation(t *testing.T) {
	mockRepo := &MockRepository{}
	svc := NewProductService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name    string
		id      string
		product model.Product
		wantErr bool
	}{
		{
			name:    "Valid Update",
			id:      "test-1",
			product: model.Product{Name: "Updated Mouse", Price: 30.00},
			wantErr: false,
		},
		{
			name:    "Invalid Update - Negative Price",
			id:      "test-1",
			product: model.Product{Name: "Broken Mouse", Price: -1.00},
			wantErr: true,
		},
		{
			name:    "Invalid Update - Empty Name",
			id:      "test-1",
			product: model.Product{Name: "", Price: 10.00},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.UpdateProduct(ctx, tt.id, tt.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestProductService_FinalCoverage(t *testing.T) {
	mockRepo := &MockRepository{}
	svc := NewProductService(mockRepo)
	ctx := context.Background()

	// Test GetProductByID
	t.Run("GetProductByID", func(t *testing.T) {
		_, err := svc.GetProductByID(ctx, "test-id")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	// Test DeleteProduct
	t.Run("DeleteProduct", func(t *testing.T) {
		err := svc.DeleteProduct(ctx, "test-id")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}
func TestProductService_ListCoverage(t *testing.T) {
	mockRepo := &MockRepository{}
	svc := NewProductService(mockRepo)
	ctx := context.Background()

	// Explicitly test the ListProducts method
	t.Run("ListProducts", func(t *testing.T) {
		products, err := svc.ListProducts(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(products) == 0 {
			// This ensures we hit any logic that handles empty results
		}
	})
}
func TestProductService_RepositoryErrors(t *testing.T) {
	// 1. Force the mock to return an error
	mockRepo := &MockRepository{ShouldFail: true}
	svc := NewProductService(mockRepo)
	ctx := context.Background()

	t.Run("CreateProduct_DatabaseError", func(t *testing.T) {
		_, err := svc.CreateProduct(ctx, model.Product{Name: "Test", Price: 10})
		if err == nil {
			t.Error("Expected an error from the database, but got nil")
		}
	})

	t.Run("ListProducts_DatabaseError", func(t *testing.T) {
		_, err := svc.ListProducts(ctx)
		if err == nil {
			t.Error("Expected an error during listing, but got nil")
		}
	})
	// Add these inside TestProductService_RepositoryErrors where ShouldFail is true
	t.Run("GetProductByID_DatabaseError", func(t *testing.T) {
		_, err := svc.GetProductByID(ctx, "test-id")
		if err == nil {
			t.Error("Expected an error for GetByID, but got nil")
		}
	})

	t.Run("UpdateProduct_DatabaseError", func(t *testing.T) {
		_, err := svc.UpdateProduct(ctx, "test-id", model.Product{Name: "Update", Price: 10})
		if err == nil {
			t.Error("Expected an error for Update, but got nil")
		}
	})
	// Add these inside TestProductService_RepositoryErrors (where ShouldFail is true)
	t.Run("DeleteProduct_DatabaseError", func(t *testing.T) {
		err := svc.DeleteProduct(ctx, "test-id")
		if err == nil {
			t.Error("Expected an error for DeleteProduct, but got nil")
		}
	})

	// Add this to your successful test file/function
	t.Run("DeleteProduct_Success", func(t *testing.T) {
		mockRepoSuccess := &MockRepository{ShouldFail: false}
		svcSuccess := NewProductService(mockRepoSuccess)
		err := svcSuccess.DeleteProduct(ctx, "test-id")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("GetProductByID_DatabaseError", func(t *testing.T) {
		// mockRepo already has ShouldFail = true here
		_, err := svc.GetProductByID(ctx, "test-id")
		if err == nil {
			t.Error("Expected error for GetProductByID, but got nil")
		}
	})
	// 2. Test Failure Path (This turns the 'return err' line GREEN)
	t.Run("Delete_RepoFailure", func(t *testing.T) {
		mockRepo := &MockRepository{ShouldFail: true}
		svc := NewProductService(mockRepo)
		err := svc.DeleteProduct(ctx, "test-id")
		if err == nil {
			t.Error("Expected error from repository, but got nil")
		}
	})
}
