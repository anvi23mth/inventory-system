package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/anvi23mth/inventory-system/internal/model"
	"github.com/anvi23mth/inventory-system/internal/service"
)

// ProductHandler holds the service as a dependency
type ProductHandler struct {
	Service *service.ProductService
}

// NewProductHandler initializes the handler with the service
func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{Service: s}
}

// CreateProduct handles POST /products
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p model.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Note the use of r.Context() - required for MongoDB
	created, err := h.Service.CreateProduct(r.Context(), p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

// HandleProductRequest handles GET, PUT, and DELETE
func (h *ProductHandler) HandleProductRequest(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/products/")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		if id != "" && id != "list" {
			p, err := h.Service.GetProductByID(r.Context(), id)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(p)
		} else {
			list, _ := h.Service.ListProducts(r.Context())
			json.NewEncoder(w).Encode(list)
		}

	case http.MethodPut:
		var p model.Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		updated, err := h.Service.UpdateProduct(r.Context(), id, p)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		err := h.Service.DeleteProduct(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
