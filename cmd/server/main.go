package main

import (
	"context"
	"net/http"
	"time"

	"github.com/anvi23mth/inventory-system/internal/handler"
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

	db := client.Database("inventory_db")

	// 2. Initialize Layers (The Week 3 Way)
	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productSvc)

	// 3. Define Routes
	http.HandleFunc("/hello-world", handler.HelloWorld)

	// These will now point to the methods inside your productHandler instance
	http.HandleFunc("/products", productHandler.CreateProduct)
	http.HandleFunc("/products/", productHandler.HandleProductRequest)

	logger.Log.Info().Msg("Server started at :8080 with MongoDB")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Log.Fatal().Err(err).Msg("Server failed")
	}
}
