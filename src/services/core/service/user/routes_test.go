package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/core/types"
	"github.com/gorilla/mux"
)

// Test responsible for verifying whether a user's registration is successful
func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Name:         "madonna",
			Email:        "blabla",
			Password:     "senha123",
			CompanyName:  "spotify",
			Office:       "rj",
			LinkedinLink: "linklegal",
			Interest:     "saude e bem estar",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should correctly register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Name:         "madonna",
			Email:        "maddona@gmail.com",
			Password:     "senha123",
			CompanyName:  "spotify",
			Office:       "rj",
			LinkedinLink: "linklegal",
			Interest:     "saude e bem estar",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})

	t.Run("should correctly update the user", func(t *testing.T) {

		payload := types.User{
			ID:           1,
			Name:         "Novo Nome",
			Email:        "novoemail@example.com",
			Password:     "novasenha123",
			CompanyName:  "Nova Empresa",
			Office:       "Novo Escritório",
			LinkedinLink: "Novo Link do LinkedIn",
			Interest:     "Novo insteresse",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/users/{id}", handler.handleUpdateUser)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should correctly delete the user", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/users/{id}", handler.handleDeleteUser)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
}

// Test for successful user creation
func TestHandleRegisterSuccess(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	payload := types.RegisterUserPayload{
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		Password:     "strongpassword123",
		CompanyName:  "Doe Inc",
		Office:       "NY",
		LinkedinLink: "https://linkedin.com/in/johndoe",
		Interest:     "saude e bem estar",
	}

	payloadData, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(payloadData))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

// Test for error when creating a user with invalid data
func TestHandleRegisterInvalidData(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	invalidData := []byte(`{"email": "notanemail", "password": "123"}`) // Dados inválidos
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(invalidData))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Test for successful deletion of a user
func TestHandleDeleteUserSuccess(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// Test for error when deleting a user with invalid ID
func TestHandleDeleteUserInvalidID(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	req, _ := http.NewRequest(http.MethodDelete, "/users/abc", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Test for success in updating a user
func TestHandleUpdateUserSuccess(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	payload := types.User{
		ID:           1,
		Name:         "Jane Doe Updated",
		Email:        "jane.doe.updated@example.com",
		Password:     "updatedpassword123",
		CompanyName:  "Doe Updated Inc",
		Office:       "CA",
		LinkedinLink: "https://linkedin.com/in/janedoeupdated",
		Interest:     "saude e bem estar",
	}

	payloadData, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(payloadData))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// Test for error when updating a user with invalid ID
func TestHandleUpdateUserInvalidID(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	payload := types.User{
		Name:         "Jane Doe",
		Email:        "jane.doe@example.com",
		Password:     "password123",
		CompanyName:  "Doe Inc",
		Office:       "CA",
		LinkedinLink: "https://linkedin.com/in/janedoe",
		Interest:     "saude e bem estar",
	}

	payloadData, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/users/invalidID", bytes.NewBuffer(payloadData))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) (int, error) {
	return 0, nil
}

func (m *mockUserStore) DeleteUserByID(id int) error {
	return nil
}

func (m *mockUserStore) UpdateUser(user types.User) error {
	return nil
}
