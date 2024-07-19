package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validate is an instance of validator used for payload validation
var Validate = validator.New()

// ParseJSON decodes the JSON payload from the request body into the provided struct
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

// WriteJSON writes the provided value as a JSON response with the specified status code.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// WriteError writes an error message as a JSON response with the specified status code
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

// GetIDFromHeaderRequest extracts the user ID from the Authorization header or URL query parameter
func GetIDFromHeaderRequest(r *http.Request) string {
	// Check the "Authorization" header for the JWT token
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth != "" {
		// Remove the "Bearer " prefix from the JWT token, if present
		token := strings.TrimPrefix(tokenAuth, "Bearer ")
		// Extract the user ID from the JWT token
		parts := strings.Split(token, ".")
		if len(parts) == 3 {
			claimsData, err := base64.RawURLEncoding.DecodeString(parts[1])
			if err != nil {
				return ""
			}
			var claims map[string]interface{}
			if err := json.Unmarshal(claimsData, &claims); err != nil {
				return ""
			}
			if id, ok := claims["id"].(string); ok {
				return id
			}
		}
	}

	// If there is no JWT token in the "Authorization" header, try to get the token from the URL query
	tokenQuery := r.URL.Query().Get("token")
	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}
