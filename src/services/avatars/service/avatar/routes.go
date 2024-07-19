package avatar

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler struct defines the structure with BaseURL and Style fields
type Handler struct {
	BaseURL string
	Style   string 
}

// NewHandler initializes a new Handler with default values
func NewHandler() *Handler {
	return &Handler{
		BaseURL: "https://api.dicebear.com/8.x",
		Style:   "personas",                     
	}
}

// RegisterRoutes sets up the route for serving avatars
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/avatars", h.ServeAvatarHandler).Methods("GET") // Route for getting avatars
}

// ServeAvatarHandler is the handler that returns up to 50 avatar URLs generated by the DiceBear API.
func (h *Handler) ServeAvatarHandler(w http.ResponseWriter, r *http.Request) {
	const numAvatars = 20
	avatars := make(map[string]string, numAvatars)

	for i := 0; i < numAvatars; i++ {
		seed := generateRandomSeed()
		avatarURL := fmt.Sprintf("%s/%s/svg?seed=%s", h.BaseURL, h.Style, seed)
		log.Printf("Generated avatar URL: %s", avatarURL)

		if validateAvatarURL(avatarURL) {
			avatars[fmt.Sprintf("icon %d", i+1)] = avatarURL
		} else {
			log.Printf("Failed to validate avatar URL: %s", avatarURL)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(avatars); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// GenerateRandomSeed creates a random seed for use in avatar creation.
func generateRandomSeed() string {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Error generating random seed: %v", err)
	}
	return fmt.Sprintf("%x", b)
}

// ValidateAvatarURL checks whether the URL returns a 200 OK status (optional).
func validateAvatarURL(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error validating avatar URL: %v", err)
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}