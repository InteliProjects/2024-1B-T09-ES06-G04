package avatar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TestServeAvatarHandlerSuccess tests if the ServeAvatarHandler function returns the avatars correctly.
func TestServeAvatarHandlerSuccess(t *testing.T) {
	handler := NewHandler()

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	req, _ := http.NewRequest(http.MethodGet, "/avatars", nil)
	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var avatars map[string]string
	if err := json.NewDecoder(responseRecorder.Body).Decode(&avatars); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if len(avatars) != 20 {
		t.Errorf("expected 20 avatars, got %d", len(avatars))
	}

	for i := 1; i <= 20; i++ {
		if _, exists := avatars[fmt.Sprintf("icon %d", i)]; !exists {
			t.Errorf("missing avatar icon %d", i)
		}
	}
}

// TestGenerateRandomSeed tests if the generateRandomSeed function generates random seeds.
func TestGenerateRandomSeed(t *testing.T) {
	seed1 := generateRandomSeed()
	seed2 := generateRandomSeed()

	if seed1 == seed2 {
		t.Errorf("expected different seeds, got the same: %s", seed1)
	}
}

// TestValidateAvatarURL tests if the validateAvatarURL function correctly validates avatar URLs.
func TestValidateAvatarURL(t *testing.T) {
	validURL := "https://avatars.dicebear.com/api/personas/example.svg"
	if !validateAvatarURL(validURL) {
		t.Errorf("expected URL to be valid: %s", validURL)
	}

	// Using an obviously invalid URL that doesn't exist
	invalidURL := "https://teste"
	if validateAvatarURL(invalidURL) {
		t.Errorf("expected URL to be invalid: %s", invalidURL)
	}
}