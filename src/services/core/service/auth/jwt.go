package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/core/types"
	"github.com/Inteli-College/2024-1B-T09-ES06-G04/core/utils"
	"github.com/golang-jwt/jwt"
)

// contextKey is a custom type for the context key.
type contextKey string

// UserKey is the context key for storing user ID.
const UserKey contextKey = "id"

// WithJWTAuth is a middleware function that handles JWT authentication.
func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract token from the request
		tokenString := utils.GetTokenFromRequest(r)

		// Validate the token
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		// Extract user ID from the token claims
		claims := token.Claims.(jwt.MapClaims)
		str := claims["id"].(string)
		userID, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert userID to int: %v", err)
			permissionDenied(w)
			return
		}

		// Get user from the store using the extracted user ID
		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		// Add the user to the context
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		// Call the handler function if the token is valid
		handlerFunc(w, r)
	}
}

// CreateJWT creates a new JWT token with the provided secret and user ID.
func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(604800)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// validateJWT validates the provided JWT token.
func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})
}

// permissionDenied sends a permission denied response.
func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

// GetUserIDFromContext retrieves the user ID stored in the context.
func GetUserIDFromContext(ctx context.Context) int {
	id, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return id 
}

