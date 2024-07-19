package avatar

import (
	"io"
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/gorilla/mux"
)

// This test is an integration test because it makes a real HTTP request to the server
func TestAvatarIntegration(t *testing.T) {
	resp, err := http.Get("http://localhost:8083/avatars")
	if err != nil {
		t.Fatalf("Failed to send request to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	t.Logf("Received response: %s", body)
}

// This test is an integration test because it makes a real HTTP request to the server
func TestServeAvatarHandlerNegative(t *testing.T) {
	router := mux.NewRouter()
	handler := NewHandler()
	handler.RegisterRoutes(router)

	req, err := http.NewRequest("GET", "/avatars/invalidparam=value", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code 400 for invalid parameters, got %d", rr.Code)
	}

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal("Failed to read response body", err)
	}
	t.Logf("Received error response: %s", body)
}
