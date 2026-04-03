package service

import (
	"context"
	"testing"

	"github.com/anvi23mth/inventory-system/internal/model"
)

func TestCreateCategory_Validation(t *testing.T) {
	// 1. Initialize Mock and Service
	mockRepo := &MockCategoryRepository{}
	svc := NewProductCategoryService(mockRepo)
	ctx := context.Background()

	// 2. Define Table-Driven Scenarios [cite: 242]
	tests := []struct {
		name     string
		category model.ProductCategory
		wantErr  bool
	}{
		{
			name:     "Success - Valid Category",
			category: model.ProductCategory{Title: "Electronics", Description: "Gadgets"},
			wantErr:  false,
		},
		{
			name:     "Failure - Missing Title",
			category: model.ProductCategory{Title: "", Description: "Invalid data"},
			wantErr:  true,
		},
	}

	// 3. Execution Loop [cite: 242]
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.CreateCategory(ctx, tt.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s: CreateCategory() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}
func TestCategoryService_AdditionalMethods(t *testing.T) {
	mockRepo := &MockCategoryRepository{}
	svc := NewProductCategoryService(mockRepo)
	ctx := context.Background()

	// 1. Test ListCategories (Fixes the first red block)
	t.Run("ListCategories", func(t *testing.T) {
		_, err := svc.ListCategories(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	// 2. Test UpdateCategory (Table-Driven to cover the second red block)
	t.Run("UpdateCategory_Validation", func(t *testing.T) {
		tests := []struct {
			name     string
			id       string
			category model.ProductCategory
			wantErr  bool
		}{
			{
				name:     "Valid Update",
				id:       "cat-1",
				category: model.ProductCategory{Title: "Food"},
				wantErr:  false,
			},
			{
				name:     "Invalid Update - Missing Title",
				id:       "cat-1",
				category: model.ProductCategory{Title: ""},
				wantErr:  true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := svc.UpdateCategory(ctx, tt.id, tt.category)
				if (err != nil) != tt.wantErr {
					t.Errorf("%s: UpdateCategory() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				}
			})
		}
	})
}
func TestCategoryService_RepositoryErrors(t *testing.T) {
	mockRepo := &MockCategoryRepository{ShouldFail: true}
	svc := NewProductCategoryService(mockRepo)
	ctx := context.Background()

	t.Run("CreateCategory_DBError", func(t *testing.T) {
		err := svc.CreateCategory(ctx, model.ProductCategory{Title: "Valid"})
		if err == nil {
			t.Error("Expected error for CreateCategory, got nil")
		}
	})

	t.Run("ListCategories_DBError", func(t *testing.T) {
		_, err := svc.ListCategories(ctx)
		if err == nil {
			t.Error("Expected error for ListCategories, got nil")
		}
	})

	t.Run("UpdateCategory_DBError", func(t *testing.T) {
		err := svc.UpdateCategory(ctx, "id", model.ProductCategory{Title: "Valid"})
		if err == nil {
			t.Error("Expected error for UpdateCategory, got nil")
		}
	})

	t.Run("DeleteCategory_DBError", func(t *testing.T) {
		err := svc.DeleteCategory(ctx, "test-id")
		if err == nil {
			t.Error("Expected error for DeleteCategory, got nil")
		}
	})
	t.Run("GetByID_DBError", func(t *testing.T) {
		// 1. Setup a fresh mock forced to fail
		mockRepoFail := &MockCategoryRepository{ShouldFail: true}
		svcFail := NewProductCategoryService(mockRepoFail)

		// 2. Call the method
		_, err := svcFail.GetByID(ctx, "any-id")

		// 3. Assert the error exists
		if err == nil {
			t.Error("Expected error for GetByID, but got nil")
		}
	})
}
func TestCategoryService_FinalSuccessPaths(t *testing.T) {
	mockRepo := &MockCategoryRepository{ShouldFail: false}
	svc := NewProductCategoryService(mockRepo)
	ctx := context.Background()

	t.Run("GetByID_Success", func(t *testing.T) {
		_, err := svc.GetByID(ctx, "cat-123")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("DeleteCategory_Success", func(t *testing.T) {
		// Ensure you have a DeleteCategory method in category_service.go
		err := svc.DeleteCategory(ctx, "cat-123")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}
