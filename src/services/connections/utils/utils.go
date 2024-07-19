package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"encoding/base64"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func GetIDFromHeaderRequest(r *http.Request) string {
	// Verifica o header "Authorization" para obter o token JWT
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth != "" {
		// Remove o prefixo "Bearer " do token JWT, se presente
		token := strings.TrimPrefix(tokenAuth, "Bearer ")
		// Extrai o ID do usuário do token JWT
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

	// Se não houver token JWT no header "Authorization",
	// tenta obter o token da consulta da URL
	tokenQuery := r.URL.Query().Get("token")
	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}
