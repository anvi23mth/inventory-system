package middleware

import (
	"net/http"
	"net/http/httptest"
	"os" // Required to mock the .env key
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	// 1. Setup: Set the environment variable that your middleware expects
	adminKey := "test-secret-key"
	os.Setenv("ADMIN_KEY", adminKey)
	defer os.Unsetenv("ADMIN_KEY") // Clean up after the test finishes

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// 2. Initialize: Pass only the handler
	handlerToTest := AuthMiddleware(nextHandler)

	tests := []struct {
		name       string
		headerKey  string
		wantStatus int
	}{
		{
			name:       "Missing Key",
			headerKey:  "",
			wantStatus: http.StatusForbidden,
		},
		{
			name:       "Wrong Key",
			headerKey:  "wrong-password",
			wantStatus: http.StatusForbidden,
		},
		{
			name:       "Correct Key",
			headerKey:  "test-secret-key",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/products", nil)
			if tt.headerKey != "" {
				req.Header.Set("X-Admin-Key", tt.headerKey)
			}

			rr := httptest.NewRecorder()
			handlerToTest.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("%s: got status %v, want %v", tt.name, rr.Code, tt.wantStatus)
			}
		})
	}
}
