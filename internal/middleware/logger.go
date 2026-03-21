package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Pass control to the next handler
		next.ServeHTTP(w, r)

		// Log the details after the handler finishes
		log.Printf(
			"METHOD: %s | PATH: %s | DURATION: %v",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
