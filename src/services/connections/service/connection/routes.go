package connection

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/connections/types"
	"github.com/Inteli-College/2024-1B-T09-ES06-G04/connections/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ConnectionStore
}

// NewHandler creates a new Handler instance with the provided connection store.
func NewHandler(store types.ConnectionStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registers the routes for handling different HTTP methods and paths related to connections.
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/connections", h.handleCreateConnection).Methods("POST")
	router.HandleFunc("/connections", h.handleGetConnectionByUserID).Methods("GET")
	router.HandleFunc("/connections/{ID}", h.handleUpdateConnection).Methods("PUT")
	router.HandleFunc("/ratings", h.handlerCreateRating).Methods("POST")
	router.HandleFunc("/ratings", h.handleGetRatingByID).Methods("GET")
	router.HandleFunc("/connections/true", h.handleGetAcceptedConnectionByUserID).Methods("GET")
}

// handleCreateConnection handles the creation of a new connection.
func (h *Handler) handleCreateConnection(w http.ResponseWriter, r *http.Request) {
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

	var payload types.CreateConnectionPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if payload.Feedback == "" || payload.ProjectID == 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing required fields"))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	connection := types.Connection{
		UserID:    userIDInt,
		Feedback:  payload.Feedback,
		Status:    payload.Status,
		ProjectID: payload.ProjectID,
	}

	if err := h.store.CreateConnection(connection); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "Connection created successfully.")
}

// handleGetConnectionByID retrieves a connection by its ID.
func (handler *Handler) handleGetConnectionByUserID(w http.ResponseWriter, r *http.Request) {

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

	connection, err := handler.store.GetConnectionByUserID(userIDInt)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if connection == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no connection found with ID %d", userIDInt))
		return
	}

	utils.WriteJSON(w, http.StatusOK, connection)
}

func (handler *Handler) handleUpdateConnection(w http.ResponseWriter, r *http.Request) {
	// Get the connection ID from the request URL
	vars := mux.Vars(r)
	connectionIDStr := vars["ID"]

	// Convert connection ID to integer
	connectionID, err := strconv.Atoi(connectionIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid connection ID: %w", err))
		return
	}

	// Get the user ID from the request headers
	userID := utils.GetIDFromHeaderRequest(r)
	if userID == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user ID"))
		return
	}

	// Convert user ID to integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %w", err))
		return
	}

	// Parse the request body into UpdateConnectionPayload struct
	var payload types.UpdateConnectionPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON payload: %w", err))
		return
	}

	// Validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Get the connections by ID
	connections, err := handler.store.GetConnectionsByID(connectionID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("connection with id %d not found", connectionID))
		return
	}

	// Check if the user is authorized to update the connection
	authorized := false
	for _, conn := range connections {
		if conn.UserID == userIDInt {
			authorized = true
			break
		}
	}
	if !authorized {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("unauthorized access to update connection"))
		return
	}

	// Update the connection
	updatedConn := types.Connection{
		ID:        connectionID,
		Feedback:  payload.Feedback,
		Status:    payload.Status,
		ProjectID: payload.ProjectID,
		UserID:    userIDInt,
	}

	err = handler.store.UpdateConnection(updatedConn)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("could not update connection: %w", err))
		return
	}

	// If the status is updated to true, create a record in user_connections
	if updatedConn.Status {
			// Get the project owner's user ID using the project_id from the updated connection
		projectOwnerID, err := handler.store.GetProjectOwnerID(updatedConn.ProjectID)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("could not get project owner ID: %w", err))
			return
		}

		// Create a record in user_connections
		err = handler.store.CreateUserConnection(projectOwnerID, updatedConn.ID)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("could not create user connection: %w", err))
			return
		}
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

// create rating
func (handler *Handler) handlerCreateRating(w http.ResponseWriter, r *http.Request) {

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

	var payload types.CreateRatingPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err = handler.store.CreateRating(types.Ratings{
		Rating:    payload.Rating,
		UserID:    userIDInt,
		ProjectID: payload.ProjectID,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Create a new connection if the rating is 1
	if payload.Rating == 1 {
		connection := types.Connection{
			UserID:    userIDInt,
			Feedback:  "",
			Status:    false,
			ProjectID: payload.ProjectID,
		}
		if err := handler.store.CreateConnection(connection); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	}

	utils.WriteJSON(w, http.StatusCreated, "Rating created successfully.")
}


// handleGetRatingByID retrieves ratings by its ID.
func (handler *Handler) handleGetRatingByID(w http.ResponseWriter, r *http.Request) {

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


	connection, err := handler.store.GetRatingByUserID(userIDInt)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if connection == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no connection found with ID %d", userIDInt))
		return
	}

	utils.WriteJSON(w, http.StatusOK, connection)
}

func (handler *Handler) handleGetAcceptedConnectionByUserID(w http.ResponseWriter, r *http.Request) {

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

	connection, err := handler.store.GetAcceptedConnectionByUserID(userIDInt)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if connection == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no connection found with ID %d", userIDInt))
		return
	}

	utils.WriteJSON(w, http.StatusOK, connection)
}
