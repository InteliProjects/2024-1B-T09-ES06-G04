package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/core/service/auth"
	"github.com/Inteli-College/2024-1B-T09-ES06-G04/core/types"
	"github.com/Inteli-College/2024-1B-T09-ES06-G04/core/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Responsible for handling requests related to users
type Handler struct {
	store types.UserStore
}

// NewHandler creates a new instance of Handler with the provided UserStore
func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registers routes to handle different HTTP methods and paths related to users
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/users/{id}",(h.handleDeleteUser)).Methods("DELETE")
	router.HandleFunc("/users/{id}", h.handleUpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", h.handleGetUserById).Methods("GET")
}

// handleGetUserById handles getting a user by ID
func (h *Handler) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDStr := params["id"]

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.store.GetUserByID(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}


// handleLogin handles authentication of a user
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	secret := []byte("secret")
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})

}

// handleRegister handles registering a new user
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
	}

	if err := utils.Validate.Struct(payload); err != nil {
			errors := err.(validator.ValidationErrors)
			utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
			return
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
			return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
	}

	// Now the CreateUser function returns the ID of the created user and an error
	userID, err := h.store.CreateUser(types.User{
			Name:         payload.Name,
			Password:     hashedPassword,
			Email:        payload.Email,
			CompanyName:  payload.CompanyName,
			Office:       payload.Office,
			LinkedinLink: payload.LinkedinLink,
			Interest:     payload.Interest,
	})
	if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
	}

	// Send the ID of the created user as part of the JSON response
	utils.WriteJSON(w, http.StatusCreated, map[string]int{"userID": userID})
}

// handleDeleteUser handles deleting a user
func (h *Handler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDStr := params["id"]

	 userID, err := strconv.Atoi(userIDStr)
	 if err != nil {
	 	utils.WriteError(w, http.StatusBadRequest, err)
	 	return
	 }

	err = h.store.DeleteUserByID(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

// handleUpdateUser handles updating a user
func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	id, err := strconv.Atoi(userID)
	if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
	}

	var user types.User
	if err := utils.ParseJSON(r, &user); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
	}

	user.ID = id

	err = h.store.UpdateUser(user)
	if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
