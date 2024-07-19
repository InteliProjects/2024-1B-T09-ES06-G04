package project

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/projects/types"
	"github.com/Inteli-College/2024-1B-T09-ES06-G04/projects/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Responsible for handling requests related to users
type Handler struct {
	store types.ProjectStore
}

// NewHandler creates a new instance of Handler with the provided UserStore
func NewHandler(store types.ProjectStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registers routes to handle different HTTP methods and paths related to users
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/projects", h.handleGetProjects).Methods("GET")
	router.HandleFunc("/projects", h.handleCreateProject).Methods("POST")
	router.HandleFunc("/projects/user", h.handleGetProjectsByUserID).Methods("GET")
	router.HandleFunc("/projects/{id}", h.handleGetProjectByID).Methods("GET")
	router.HandleFunc("/projects/{id}", h.handleUpdateProject).Methods("PUT")
	router.HandleFunc("/projects/{id}", h.handleDeleteProject).Methods("DELETE")
}

// Responsible for creating a new project
func (handler *Handler) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterProjectPayload

	// Decodes the request payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validates the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Create the project in the database
	err = handler.store.CreateProject(types.Project{
		Name:        payload.Name,
		Description: payload.Description,
		MacroSetor:  payload.MacroSetor,
		MicroSetor:  payload.MicroSetor,
		ImageLink:   payload.ImageLink,
		UserId:      payload.UserId,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond successfully
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Project created successfully."))
}

// Responsible for getting all projects
func (handler *Handler) handleGetProjects(w http.ResponseWriter, r *http.Request) {
	// Search projects in the database
	projects, err := handler.store.GetProjects()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with the projects found
	utils.WriteJSON(w, http.StatusOK, projects)
}

// Responsible for getting all projects from a user
func (handler *Handler) handleGetProjectsByUserID(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from JWT header
	userID := utils.GetIDFromHeaderRequest(r)
	if userID == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user ID"))
		return
	}
 
	// Converts user ID to integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Searches for user projects in the database
	projects, err := handler.store.GetProjectsByUserID(userIDInt)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with the projects found
	if len(projects) == 0 {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no projects found for user with id %d", userIDInt))
		return
	}

	// Respond with the projects found
	utils.WriteJSON(w, http.StatusOK, projects)
}

// Responsible for getting a project by ID
func (handler *Handler) handleGetProjectByID(w http.ResponseWriter, r *http.Request) {
	// Extract project ID from URL
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	// Validates the project ID
	 if err != nil {
	 	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid project ID"))
	 	return
	 }

	// Search the project in the database
	project, err := handler.store.GetProjectByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Reply with the project found
	if project == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no project found with ID %d", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, project)
}

// Responsible for updating a project
func (handler *Handler) handleUpdateProject(w http.ResponseWriter, r *http.Request) {
	var payload types.UpdateProjectPayload

	// Decodes the request payload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validates the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Update the project in the database
	_, err := handler.store.GetProjectByID(payload.ID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("project with id %d not found", payload.ID))
		return
	}

	// Update the project in the database
	err = handler.store.UpdateProject(types.Project{
		Id:          payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
		MacroSetor:  payload.MacroSetor,
		MicroSetor:  payload.MicroSetor,
		ImageLink:   payload.ImageLink,
		UserId:      payload.UserId,
	})

	// Respond successfully
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, nil)
}

// Responsible for deleting a project
func (handler *Handler) handleDeleteProject(w http.ResponseWriter, r *http.Request) {
	// Extract project ID from URL
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	// Validates the project ID
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid project ID"))
		return
	}

	// Deletes the project from the database
	err = handler.store.DeleteProject(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond successfully
	w.WriteHeader(http.StatusNoContent)
}
