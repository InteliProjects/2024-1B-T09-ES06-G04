package api

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/Inteli-College/2024-1B-T09-ES06-G04/connections/service/connection"
	"github.com/gorilla/mux"
)

// ApiServer struct defines the server with an address and a database connection
type ApiServer struct {
	addr string
	db   *sql.DB
}

// NewApiServer initializes a new ApiServer instance
func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

// Run starts the HTTP server and sets up the routes
func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	connectionStore := connection.NewConnection(s.db)
	connectionHandler := connection.NewHandler(connectionStore)
	connectionHandler.RegisterRoutes(subrouter)

	log.Println("listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
