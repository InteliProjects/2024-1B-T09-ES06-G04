package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/projects/service/project"
	"github.com/gorilla/mux"
)

// ApiServer represents the API server with an address and a database connection
type ApiServer struct {
	addr string
	db   *sql.DB
}

// NewApiServer creates a new instance of ApiServer with the provided address and database connection
func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

// Run starts the API server, sets up the router, registers routes for handling project-related requests, and listens for incoming HTTP requests
func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	projectStore := project.NewStore(s.db)
	projectHandler := project.NewHandler(projectStore)
	projectHandler.RegisterRoutes(subrouter)

	log.Println("listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
