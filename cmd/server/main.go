package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/anvi23mth/inventory-system/internal/handler"
	"github.com/anvi23mth/inventory-system/internal/middleware"
	"github.com/anvi23mth/inventory-system/internal/repository"
	"github.com/anvi23mth/inventory-system/internal/service"
	"github.com/anvi23mth/inventory-system/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	logger.Init()

	// 1. Setup MongoDB Connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to connect to MongoDB")
	}

	// ... after database connection ...
	db := client.Database("inventory_db")

	// 2. Initialize Layers (The Week 3 Way)
	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(productRepo)
	// Use productSvc here (match the variable name above)
	productHandler := handler.NewProductHandler(productSvc)

	// 2. Create a New Router (Mux)
	mux := http.NewServeMux()

	// 3. Register your routes to the Mux
	mux.HandleFunc("/hello", handler.HelloWorld)
	mux.HandleFunc("/products/", productHandler.HandleProductRequest)
	mux.HandleFunc("/products", productHandler.CreateProduct) // For POST

	// 4. Wrap the Mux with your Middleware
	// Make sure you have imported your middleware package!
	wrappedMux := middleware.LoggingMiddleware(mux)

	// 5. Start the server
	log.Println("Server started at :8080 with Middleware")
	http.ListenAndServe(":8080", wrappedMux)
}
