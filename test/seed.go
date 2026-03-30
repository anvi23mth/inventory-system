package test

import (
	"context"

	"github.com/anvi23mth/inventory-system/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// SeedDatabase populates the test DB with initial data for Integration Tests [cite: 255]
func SeedDatabase(db *mongo.Database) error {
	ctx := context.Background()
	categoriesCol := db.Collection("categories")

	// Clear existing data
	categoriesCol.Drop(ctx)

	seedData := []interface{}{
		model.ProductCategory{ID: "cat-1", Title: "Food", Description: "Edible items"},
		model.ProductCategory{ID: "cat-2", Title: "Kitchen", Description: "Appliances"},
	}

	_, err := categoriesCol.InsertMany(ctx, seedData)
	return err
}
