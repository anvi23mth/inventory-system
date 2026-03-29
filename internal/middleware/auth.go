package middleware

import (
	"net/http"
	"os"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Matches the key in your .env
		adminKey := os.Getenv("ADMIN_KEY")
		if adminKey == "" {
			// Optional: Log a warning that the server is insecure
		}
		userKey := r.Header.Get("X-Admin-Key")

		if userKey != adminKey {
			http.Error(w, "Forbidden: Invalid API Key", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
