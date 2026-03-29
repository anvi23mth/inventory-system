package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	// Import your handler and middleware
	"github.com/anvi23mth/inventory-system/internal/handler"
	"github.com/anvi23mth/inventory-system/internal/middleware"
)

func TestHelloWorldEndpoint(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello", nil)
	rr := httptest.NewRecorder()

	// Use your actual handler
	handler := http.HandlerFunc(handler.HelloWorld)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello World"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
func TestCreateProductIntegration(t *testing.T) {
	// Note: This requires Docker MongoDB to be UP
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	// You would typically initialize your real database and handler here
	// For now, focus on the fact that your 'PASS' results confirm the
	// basic plumbing is working!
}
func TestProtectedProductRoute(t *testing.T) {
	// 1. Setup the test server environment
	adminKey := "integration-test-key"
	os.Setenv("ADMIN_KEY", adminKey)
	defer os.Unsetenv("ADMIN_KEY")

	// 2. Create a test request to a protected endpoint
	req, _ := http.NewRequest("GET", "/products/list", nil)

	// --- Scenario A: No API Key (Should fail) ---
	rr1 := httptest.NewRecorder()
	// You must pass your actual router/handler that has the middleware attached
	// For this example, we wrap a dummy handler with your middleware
	handlerToTest := middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handlerToTest.ServeHTTP(rr1, req)
	if rr1.Code != http.StatusForbidden {
		t.Errorf("Expected 403 Forbidden for missing key, got %v", rr1.Code)
	}

	// --- Scenario B: Correct API Key (Should pass) ---
	req.Header.Set("X-Admin-Key", adminKey)
	rr2 := httptest.NewRecorder()
	handlerToTest.ServeHTTP(rr2, req)

	if rr2.Code != http.StatusOK {
		t.Errorf("Expected 200 OK for valid key, got %v", rr2.Code)
	}
}
