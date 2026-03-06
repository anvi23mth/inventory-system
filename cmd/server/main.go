package main

import (
	"net/http"

	"github.com/yourusername/inventory-system/internal/handler"
	"github.com/yourusername/inventory-system/pkg/logger"
)

func main() {

	logger.Init()

	http.HandleFunc("/hello-world", handler.HelloWorld)

	logger.Log.Info().Msg("Server started at :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Server failed")
	}
}