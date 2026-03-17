package main

import (
	"net/http"

	"github.com/anvi23mth/inventory-system/internal/handler"
	"github.com/anvi23mth/inventory-system/internal/service"
	"github.com/anvi23mth/inventory-system/pkg/logger"
)

func main() {
	// 1. Initialize dependencies [cite: 152, 164]
	logger.Init()
	service.SeedData()

	// 2. Define Routes [cite: 161, 162, 178]

	// Basic check [cite: 162]
	http.HandleFunc("/hello-world", handler.HelloWorld)

	// CREATE: Use the specific handler for POST [cite: 178]
	http.HandleFunc("/products", handler.ProductHandler)

	// READ ALL, READ ONE, UPDATE, DELETE
	// The trailing slash "/" allows this to catch /products/list and /products/{id}
	http.HandleFunc("/products/", handler.ProductHandler)

	// 3. Start Server [cite: 159, 182]
	logger.Log.Info().Msg("Server started at :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Log.Fatal().Err(err).Msg("Server failed to start")
	}
}
