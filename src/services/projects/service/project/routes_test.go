package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/projects/types"
	"github.com/gorilla/mux"
	"github.com/golang-jwt/jwt/v4"
)

// Test responsible for verifying whether the creation of a project is successful
func TestHandleCreateProjectSuccess(test *testing.T) {
	mockStore := &mockProjectStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	project := types.RegisterProjectPayload{
		Name:        "Projeto 1",
		Description: "Descrição do projeto 1",
		MacroSetor:  "Macro setor 1",
		MicroSetor:  "Micro setor 1",
		ImageLink:   "https://example.com/image.jpg",
		UserId:      1,
	}

	projectData, _ := json.Marshal(project)

	req, _ := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(projectData))

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusCreated {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	expected := "Project created successfully."
	if responseRecorder.Body.String() != expected {
		test.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), expected)
	}
}

// Test responsible for checking whether the creation of a project returns an error when receiving invalid data
func TestHandlerCreateProjectBadRequest(test *testing.T) {
	mockStore := &mockProjectStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	badData := []byte(`{"name": "Incomplete Data"`)

	req, _ := http.NewRequest("POST", "/projects", bytes.NewBuffer(badData))
	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Test responsible for checking whether the creation of a project returns an error when receiving missing data
func TestHandlerCreateProjectMissingData(test *testing.T) {
	mockStore := &mockProjectStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	project := types.RegisterProjectPayload{
		Name:        "Projeto 1",
		Description: "Descrição do projeto 1",
		MacroSetor:  "Macro setor 1",
		MicroSetor:  "Micro setor 1",
		ImageLink:   "https://example.com/image.jpg",
	}

	projectData, _ := json.Marshal(project)

	req, _ := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(projectData))

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Test responsible for verifying whether the project listing is successful
func TestHandlerGetProjects(test *testing.T) {
	mockStore := &mockProjectStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	req, _ := http.NewRequest(http.MethodGet, "/projects", nil)

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// Test responsible for verifying that the project listing returns an error when there are no projects
func TestHandlerGetProjectsByUserID(test *testing.T) {
	mockStore := &mockProjectStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	// Cria um token JWT de exemplo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        "1",
		"expiresAt": 1717006225,
	})
	tokenString, _ := token.SignedString([]byte("secret"))

	// Cria a requisição com o header de autorização
	req, _ := http.NewRequest(http.MethodGet, "/projects/user", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// Test responsible for checking whether the listing of projects by user returns an error when receiving an invalid ID
func TestHandlerGetProjectsByUserIDBadRequest(test *testing.T) {
	mockStore := &mockProjectStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	req, _ := http.NewRequest(http.MethodGet, "/projects/user/100", nil)

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusNotFound {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Test responsible for listing projects by ID
func TestHandlerGetProjectByID(test *testing.T) {
	mockStore := &mockProjectStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	req, _ := http.NewRequest(http.MethodGet, "/projects/1", nil)

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// Test responsible for verifying the update of a project
func TestHandleUpdateProject(test *testing.T) {
	mockStore := &mockProjectStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	project := types.UpdateProjectPayload{
		ID:          1,
		Name:        "Projeto 1 Atualizado",
		Description: "Descrição do projeto 1 atualizada",
		MacroSetor:  "Macro setor 1 atualizado",
		MicroSetor:  "Micro setor 1 atualizado",
		ImageLink:   "https://example.com/image_updated.jpg",
		UserId:      1,
	}

	projectData, _ := json.Marshal(project)

	req, _ := http.NewRequest(http.MethodPut, "/projects/1", bytes.NewBuffer(projectData))

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// Test responsible for checking the update of a project returns an error when receiving invalid data
func TestHandleUpdateProjectBadRequest(test *testing.T) {
	mockStore := &mockProjectStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	badData := []byte(`{"name": "Incomplete Data"`)
	req, _ := http.NewRequest("PUT", "/projects/1", bytes.NewBuffer(badData))
	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Test responsible for checking the update of a project returns an error when receiving missing data
func TestHandleUpdateProjectMissingData(test *testing.T) {
	mockStore := &mockProjectStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	project := types.UpdateProjectPayload{
		ID:          1,
		Name:        "Projeto 1 Atualizado",
		Description: "Descrição do projeto 1 atualizada",
		MacroSetor:  "Macro setor 1 atualizado",
		MicroSetor:  "Micro setor 1 atualizado",
		ImageLink:   "https://example.com/image_updated.jpg",
	}

	projectData, _ := json.Marshal(project)

	req, _ := http.NewRequest(http.MethodPut, "/projects/1", bytes.NewBuffer(projectData))

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Test responsible for verifying the deletion of a project
func TestHandleDeleteProject(test *testing.T) {
	mockStore := &mockProjectStore{}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	req, _ := http.NewRequest(http.MethodDelete, "/projects/1", nil)

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusNoContent {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// ProjectStore interface mockup
type mockProjectStore struct{}

// ProjectStore interface methods
func (m *mockProjectStore) CreateProject(project types.Project) error {
	return nil
}

func (m *mockProjectStore) GetProjects() ([]types.Project, error) {
	return nil, nil
}

func (m *mockProjectStore) GetProjectsByUserID(id int) ([]types.Project, error) {
	if id == 1 {
		return []types.Project{
			{
				Id:          1,
				Name:        "Projeto 1",
				Description: "Descrição inicial do projeto 1",
				MacroSetor:  "Macro setor inicial",
				MicroSetor:  "Micro setor inicial",
				ImageLink:   "https://example.com/image_initial.jpg",
				UserId:      1,
			},
		}, nil
	}
	return nil, nil
}

func (m *mockProjectStore) GetProjectByID(id int) (*types.Project, error) {
	if id == 1 {
		return &types.Project{
			Id:          1,
			Name:        "Projeto 1",
			Description: "Descrição inicial do projeto 1",
			MacroSetor:  "Macro setor inicial",
			MicroSetor:  "Micro setor inicial",
			ImageLink:   "https://example.com/image_initial.jpg",
			UserId:      1,
		}, nil
	}
	return nil, nil
}

func (m *mockProjectStore) UpdateProject(project types.Project) error {
	if project.Id == 1 {
		return nil
	}
	return fmt.Errorf("project not found")
}

func (m *mockProjectStore) DeleteProject(id int) error {
	return nil
}
