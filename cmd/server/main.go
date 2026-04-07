// package main
package main

import (
	"context"
	"log"
	"net/http"
	"os" // Added for environment variables
	"time"

	"github.com/anvi23mth/inventory-system/internal/handler"
	"github.com/anvi23mth/inventory-system/internal/middleware"
	"github.com/anvi23mth/inventory-system/internal/repository"
	"github.com/anvi23mth/inventory-system/internal/service"
	"github.com/anvi23mth/inventory-system/pkg/logger"
	"github.com/joho/godotenv" // Added for .env support
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	logger.Init()
	godotenv.Load()
	// 1. Load Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system defaults")
	}

	// Pull settings from .env
	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	// 2. Setup MongoDB Connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to connect to MongoDB")
	}

	db := client.Database(dbName)

	// 3. Initialize Layers
	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productSvc)

	// 4. Create a New Router (Mux)
	mux := http.NewServeMux()

	// --- PUBLIC ROUTES ---
	mux.HandleFunc("/hello", handler.HelloWorld)
	mux.HandleFunc("/products/", productHandler.HandleProductRequest)

	// --- SECURE ROUTES (Week 4 Addition) ---
	// Wrap CreateProduct with AuthMiddleware so only admins can add items
	mux.Handle("/products", middleware.AuthMiddleware(http.HandlerFunc(productHandler.CreateProduct)))

	// 5. Wrap everything with Logging Middleware
	wrappedMux := middleware.LoggingMiddleware(mux)

	// 6. Start the server
	log.Printf("Server started at :%s with Logging and Auth Middleware", port)
	if err := http.ListenAndServe(":"+port, wrappedMux); err != nil {
		log.Fatal(err)
	}
}

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/anvi23mth/inventory-system/internal/handler"
// 	"github.com/anvi23mth/inventory-system/internal/middleware"
// 	"github.com/anvi23mth/inventory-system/internal/repository"
// 	"github.com/anvi23mth/inventory-system/internal/service"
// 	"github.com/anvi23mth/inventory-system/pkg/logger"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {
// 	logger.Init()

// 	// 1. Setup MongoDB Connection
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		logger.Log.Fatal().Err(err).Msg("Failed to connect to MongoDB")
// 	}

// 	// ... after database connection ...
// 	db := client.Database("inventory_db")

// 	// 2. Initialize Layers (The Week 3 Way)
// 	productRepo := repository.NewProductRepository(db)
// 	productSvc := service.NewProductService(productRepo)
// 	// Use productSvc here (match the variable name above)
// 	productHandler := handler.NewProductHandler(productSvc)

// 	// 2. Create a New Router (Mux)
// 	mux := http.NewServeMux()

// 	// 3. Register your routes to the Mux
// 	mux.HandleFunc("/hello", handler.HelloWorld)
// 	mux.HandleFunc("/products/", productHandler.HandleProductRequest)
// 	mux.HandleFunc("/products", productHandler.CreateProduct) // For POST

// 	// 4. Wrap the Mux with your Middleware
// 	// Make sure you have imported your middleware package!
// 	wrappedMux := middleware.LoggingMiddleware(mux)

// 	// 5. Start the server
// 	log.Println("Server started at :8080 with Middleware")
// 	http.ListenAndServe(":8080", wrappedMux)
// }
