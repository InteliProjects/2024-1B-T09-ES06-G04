package auth

import (
	"testing"
)

// test for JWT creation
func TestCreateJWT(t *testing.T) {
	secret := []byte("supersecret")

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Error("expected token to be not empty")
	}
}
