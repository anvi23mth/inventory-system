package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/anvi23mth/inventory-system/internal/model"
	"github.com/anvi23mth/inventory-system/internal/service"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	// Extracts ID from /products/{id}
	id := strings.TrimPrefix(r.URL.Path, "/products/")

	switch r.Method {
	case http.MethodGet:
		if id != "" && id != "list" {
			getProduct(w, id)
		} else {
			listProducts(w)
		}
	case http.MethodPut: // UPDATE/MODIFY
		updateProduct(w, r, id)
	case http.MethodDelete: // DELETE
		deleteProduct(w, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func listProducts(w http.ResponseWriter) {
	pList, _ := service.GetAllProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pList)
}

func getProduct(w http.ResponseWriter, id string) {
	p, err := service.GetProductByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}
	json.NewEncoder(w).Encode(p)
}

func updateProduct(w http.ResponseWriter, r *http.Request, id string) {
	var p model.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	updated, err := service.UpdateProduct(id, p)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func deleteProduct(w http.ResponseWriter, id string) {
	err := service.DeleteProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 No Content success
}
