package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	//"strings"
	"testing"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/connections/types"
	"github.com/gorilla/mux"
)

// Test responsible for verifying if a connection creation is successful
// func TestHandleCreateConnectionSuccess(test *testing.T) {
// 	connectionStore := &mockConnectionStore{}
// 	handler := NewHandler(connectionStore)
//
// 	router := mux.NewRouter()
// 	handler.RegisterRoutes(router)
//
// 	connection := types.CreateConnectionPayload{
// 		Feedback:  "Feedback do projeto 1",
// 		Status:    true,
// 		ProjectID: 1,
// 		UserID: 1,
// 	}
//
// 	connectionData, _ := json.Marshal(connection)
//
// 	req, _ := http.NewRequest(http.MethodPost, "/connections", bytes.NewBuffer(connectionData))
//
// 	responseRecorder := httptest.NewRecorder()
//
// 	router.ServeHTTP(responseRecorder, req)
//
// 	if status := responseRecorder.Code; status != http.StatusCreated {
// 		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
// 	}
//
// 	expected := "Connection created successfully."
// 	if !strings.Contains(responseRecorder.Body.String(), expected) {
// 		test.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), expected)
// 	}
// }

// Test responsible for verifying if a connection creation returns an error when receiving invalid data
func TestHandlerCreateConnectionBadRequest(test *testing.T) {
	mockStore := &mockConnectionStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	badData := []byte(`{"name": "Incomplete Data"`)

	req, _ := http.NewRequest("POST", "/connections", bytes.NewBuffer(badData))
	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Test responsible for verifying if a connection creation returns an error when receiving missing data
func TestHandlerCreateConnectionMissingData(test *testing.T) {
	mockStore := &mockConnectionStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	connection := types.CreateConnectionPayload{
		Feedback: "Feedback da conexão 1",
		Status:   true,
	}

	connectionData, _ := json.Marshal(connection)

	req, _ := http.NewRequest(http.MethodPost, "/connections", bytes.NewBuffer(connectionData))

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Test responsible for verifying if the listing of accepted connections is successful
// func TestHandlerGetAcceptedConnections(test *testing.T) {
// 	mockStore := &mockConnectionStore{}
// 	handler := NewHandler(mockStore)

// 	router := mux.NewRouter()
// 	router.HandleFunc("/connections/{ID}/{connectionID}", handler.handleGetAcceptedConnections).Methods(http.MethodGet)

// 	req, _ := http.NewRequest(http.MethodGet, "/connections/1/1", nil)

// 	responseRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(responseRecorder, req)

// 	if status := responseRecorder.Code; status != http.StatusNotFound {
// 		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// }

// Test responsible for verifying if the listing of projects by user returns an error when receiving an invalid ID
func TestHandlerGetConnectionByIDBadRequest(test *testing.T) {
	mockStore := &mockConnectionStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	req, _ := http.NewRequest(http.MethodGet, "/connections/abc", nil)

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Test responsible for listing connections by ID
func TestHandlerGetConnectionByID(test *testing.T) {
	mockStore := &mockConnectionStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	req, _ := http.NewRequest(http.MethodGet, "/connections/1", nil)

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// Test responsible for verifying the update of a connection
func TestHandleUpdateProject(test *testing.T) {
	mockStore := &mockConnectionStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	connection := types.UpdateConnectionPayload{
		ID:        1,
		Feedback:  "Feedback do Projeto 1 Atualizado",
		Status:    false,
		ProjectID: 1,
	}

	connectionData, _ := json.Marshal(connection)

	req, _ := http.NewRequest(http.MethodPut, "/connections/1", bytes.NewBuffer(connectionData))

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// Teste responsável por verificar se a atualização de uma conexão retorna erro ao receber dados inválidos
func TestHandleUpdateConnectionBadRequest(test *testing.T) {
	mockStore := &mockConnectionStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	badData := []byte(`{"name": "Incomplete Data"`)
	req, _ := http.NewRequest("PUT", "/connections/1", bytes.NewBuffer(badData))
	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Teste responsável por verificar se a atualização de uma conexão retorna erro ao receber dados faltantes
func TestHandleUpdateConnectionMissingData(test *testing.T) {
	mockStore := &mockConnectionStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	connection := types.UpdateConnectionPayload{
		ID:       1,
		Feedback: "Feedback do Projeto 1 Atualizado",
		Status:   false,
	}

	connectionData, _ := json.Marshal(connection)

	req, _ := http.NewRequest(http.MethodPut, "/connections/1", bytes.NewBuffer(connectionData))

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Mock da interface ConnectionStore
type mockConnectionStore struct{}

// CreateRating implements types.ConnectionStore.
func (m *mockConnectionStore) CreateRating(rating types.Ratings) error {
	panic("unimplemented")
}

// CreateUserConnection implements types.ConnectionStore.
func (m *mockConnectionStore) CreateUserConnection(userID int, connectionID int) error {
	panic("unimplemented")
}

// GetAcceptedConnectionByUserID implements types.ConnectionStore.
func (m *mockConnectionStore) GetAcceptedConnectionByUserID(id int) ([]types.Connection, error) {
	panic("unimplemented")
}

// GetConnectionByUserID implements types.ConnectionStore.
func (m *mockConnectionStore) GetConnectionByUserID(id int) ([]types.Connection, error) {
	panic("unimplemented")
}

// GetConnectionsByID implements types.ConnectionStore.
func (m *mockConnectionStore) GetConnectionsByID(id int) ([]types.Connection, error) {
	panic("unimplemented")
}

// GetProjectOwnerID implements types.ConnectionStore.
func (m *mockConnectionStore) GetProjectOwnerID(projectID int) (int, error) {
	panic("unimplemented")
}

// GetRatingByUserID implements types.ConnectionStore.
func (m *mockConnectionStore) GetRatingByUserID(id int) ([]types.Ratings, error) {
	panic("unimplemented")
}

// Métodos da interface  ConnectionStore
func (m *mockConnectionStore) CreateConnection(Connection types.Connection) error {
	return nil
}

func (m *mockConnectionStore) GetAcceptedConnections(id int, idConnection int) ([]types.Connection, error) {
	return nil, nil
}

func (m *mockConnectionStore) GetConnectionByID(id int) (*types.Connection, error) {
	if id == 1 {
		return &types.Connection{
			ID:        1,
			Feedback:  "Feedback inicial da Conexão 1",
			Status:    true,
			ProjectID: 1,
		}, nil
	}
	return nil, nil
}

func (m *mockConnectionStore) UpdateConnection(connection types.Connection) error {
	if connection.ID == 1 {
		return nil
	}
	return fmt.Errorf("Connection not found")
}
