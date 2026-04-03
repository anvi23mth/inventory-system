package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/anvi23mth/inventory-system/internal/handler"
	"github.com/anvi23mth/inventory-system/internal/middleware"
	"github.com/anvi23mth/inventory-system/internal/repository"
	"github.com/anvi23mth/inventory-system/internal/service"
	"github.com/anvi23mth/inventory-system/pkg/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// enableCORS allows the browser frontend to talk to the Go server
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Admin-Key")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	logger.Init()

	// 1. Load Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system defaults")
	}

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

	// 4. Create Router
	mux := http.NewServeMux()

	// --- PUBLIC ROUTES ---
	// Serve frontend static files
	mux.Handle("/", http.FileServer(http.Dir("./frontend")))
	mux.HandleFunc("/hello", handler.HelloWorld)
	mux.HandleFunc("/products/", productHandler.HandleProductRequest)

	// --- SECURE ROUTES ---
	mux.Handle("/products", middleware.AuthMiddleware(http.HandlerFunc(productHandler.CreateProduct)))

	// 5. Wrap with CORS first, then Logging
	wrappedMux := middleware.LoggingMiddleware(enableCORS(mux))

	// 6. Start the server
	log.Printf("Server started at :%s", port)
	if err := http.ListenAndServe(":"+port, wrappedMux); err != nil {
		log.Fatal(err)
	}
}

// // package main
// package main

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"os" // Added for environment variables
// 	"time"

// 	"github.com/anvi23mth/inventory-system/internal/handler"
// 	"github.com/anvi23mth/inventory-system/internal/middleware"
// 	"github.com/anvi23mth/inventory-system/internal/repository"
// 	"github.com/anvi23mth/inventory-system/internal/service"
// 	"github.com/anvi23mth/inventory-system/pkg/logger"
// 	"github.com/joho/godotenv" // Added for .env support
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {
// 	logger.Init()
// 	godotenv.Load()
// 	// 1. Load Environment Variables
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No .env file found, using system defaults")
// 	}

// 	// Pull settings from .env
// 	port := os.Getenv("PORT")
// 	mongoURI := os.Getenv("MONGO_URI")
// 	dbName := os.Getenv("DB_NAME")

// 	// 2. Setup MongoDB Connection
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
// 	if err != nil {
// 		logger.Log.Fatal().Err(err).Msg("Failed to connect to MongoDB")
// 	}

// 	db := client.Database(dbName)

// 	// 3. Initialize Layers
// 	productRepo := repository.NewProductRepository(db)
// 	productSvc := service.NewProductService(productRepo)
// 	productHandler := handler.NewProductHandler(productSvc)

// 	// 4. Create a New Router (Mux)
// 	mux := http.NewServeMux()

// 	// --- PUBLIC ROUTES ---
// 	mux.HandleFunc("/hello", handler.HelloWorld)
// 	mux.HandleFunc("/products/", productHandler.HandleProductRequest)

// 	// --- SECURE ROUTES (Week 4 Addition) ---
// 	// Wrap CreateProduct with AuthMiddleware so only admins can add items
// 	mux.Handle("/products", middleware.AuthMiddleware(http.HandlerFunc(productHandler.CreateProduct)))

// 	// 5. Wrap everything with Logging Middleware
// 	wrappedMux := middleware.LoggingMiddleware(mux)

// 	// 6. Start the server
// 	log.Printf("Server started at :%s with Logging and Auth Middleware", port)
// 	if err := http.ListenAndServe(":"+port, wrappedMux); err != nil {
// 		log.Fatal(err)
// 	}
// }
