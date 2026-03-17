package service

import (
	"errors"

	"github.com/anvi23mth/inventory-system/internal/model"
)

var products = make(map[string]model.Product)

// CREATE
func CreateProduct(p model.Product) (model.Product, error) {
	if p.ID == "" {
		return model.Product{}, errors.New("product ID is required")
	}
	products[p.ID] = p
	return p, nil
}

// READ (List)
func GetAllProducts() ([]model.Product, error) {
	var list []model.Product
	for _, p := range products {
		list = append(list, p)
	}
	return list, nil
}

// READ (Single)
func GetProductByID(id string) (model.Product, error) {
	p, exists := products[id]
	if !exists {
		return model.Product{}, errors.New("product not found")
	}
	return p, nil
}

// UPDATE (Modify)
func UpdateProduct(id string, p model.Product) (model.Product, error) {
	if _, exists := products[id]; !exists {
		return model.Product{}, errors.New("cannot update: product not found")
	}
	p.ID = id // Ensure ID remains consistent
	products[id] = p
	return p, nil
}

// DELETE
func DeleteProduct(id string) error {
	if _, exists := products[id]; !exists {
		return errors.New("cannot delete: product not found")
	}
	delete(products, id)
	return nil
}

func SeedData() {
	p := model.Product{ID: "seed_01", Name: "Sample Item", Price: 10.00, Quantity: 5}
	products[p.ID] = p
}
