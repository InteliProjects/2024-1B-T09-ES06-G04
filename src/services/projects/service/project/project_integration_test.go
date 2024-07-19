package project

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/projects/types"
	
)

// This test is an integration test get all projects
func TestProjectIntegration(t *testing.T) {
	resp, err := http.Get("http://localhost:8082/api/v1/projects")
	if err != nil {
		t.Fatalf("Failed to send request to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	t.Logf("Received response: %s", resp.Body)
}

// This test is an integration test get project by id
func TestProjectIntegrationById(t *testing.T) {
	resp, err := http.Get("http://localhost:8082/api/v1/projects/1")
	if err != nil {
		t.Fatalf("Failed to send request to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	t.Logf("Received response: %s", resp.Body)
}

// This test in an integration test get project by user id invalid
func TestProjectIntegrationByUserIdInvalid(t *testing.T) {
	resp, err := http.Get("http://localhost:8082/api/v1/projects/user/0")
	if err != nil {
		t.Fatalf("Failed to send request to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		t.Logf("User with ID %d not found", 0)
	} else {
		t.Fatalf("Expected status code 404, got %v", resp.StatusCode)
	}
}

// This test is an integration test create project
func TestCreateProject(t *testing.T) {
	projectPayload := types.RegisterProjectPayload{
		Name:        "Test Project",
		Description: "This is a test project",
		MacroSetor:  "Education",
		MicroSetor:  "E-learning",
		ImageLink:   "http://example.com/image.jpg",
		UserId:      1,
	}

	payloadBytes, err := json.Marshal(projectPayload)
	if err != nil {
		t.Fatalf("Failed to marshal project payload: %v", err)
	}

	resp, err := http.Post("http://localhost:8082/api/v1/projects", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatalf("Failed to send POST request to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code 201, got %v", resp.StatusCode)
	}

 	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	t.Logf("Received response: %s", body)
}

// This test is an integration test create project invalid user id
func TestCreateProjectInvalidUserId(t *testing.T) {
	projectPayload := types.RegisterProjectPayload{
		Name:        "Test Project",
		Description: "This is a test project",
		MacroSetor:  "Education",
		MicroSetor:  "E-learning",
		ImageLink:   "http://example.com/image.jpg",
		UserId:      0,
	}

	payloadBytes, err := json.Marshal(projectPayload)
	if err != nil {
		t.Fatalf("Failed to marshal project payload: %v", err)
	}

	resp, err := http.Post("http://localhost:8082/api/v1/projects", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatalf("Failed to send POST request to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		t.Logf("User with ID %d not found", 0)
	} else {
		t.Fatalf("Expected status code 404, got %v", resp.StatusCode)
	}
}

// This test is an integration test update project
func TestUpdateProject(t *testing.T) {
	updatePayload := types.UpdateProjectPayload{
		ID:           10,
		Name:        "Updated Project Name",
		Description: "Updated description",
		MacroSetor:  "Updated Macro Sector",
		MicroSetor:  "Updated Micro Sector",
		ImageLink:   "http://example.com/updated-image.jpg",
		UserId:      1, 
	}

	payloadBytes, err := json.Marshal(updatePayload)
	if err != nil {
		t.Fatalf("Failed to marshal update payload: %v", err)
	}

	req, err := http.NewRequest("PUT", "http://localhost:8082/api/v1/projects/8", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatalf("Failed to create PUT request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send PUT request to server: %v", err)
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

// This test is an integration test delete project
func TestDeleteProject(t *testing.T) {

	req, err := http.NewRequest("DELETE", "http://localhost:8082/api/v1/projects/9", nil)
	if err != nil {
		t.Fatalf("Failed to create DELETE request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send DELETE request to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		t.Fatalf("Cannot delete project: Project with ID %d not found. Make sure the project ID exists before running this test.", 10)
	} else if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("Expected status code 204 (No Content), got %v. Make sure the project with ID %d exists and can be deleted.", resp.StatusCode, 10)
	} else {
		t.Logf("Project with ID %d deleted successfully.", 10)
	}
}