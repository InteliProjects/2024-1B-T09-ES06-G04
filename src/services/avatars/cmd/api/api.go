package api 

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Inteli-College/2024-1B-T09-ES06-G04/avatars/service/avatar"
)

// API struct holds the router and the service
type ApiServer struct {
	addr string 
}

// NewApiServer creates a new ApiServer instance.
func NewApiServer(addr string) *ApiServer {
  return &ApiServer{
    addr: addr,
  }
}

// Run initializes and runs the server.
func (s *ApiServer) Run() error {
  router := mux.NewRouter()
  avatarHandler := avatar.NewHandler() 
  avatarHandler.RegisterRoutes(router) 
  log.Println("Servidor de Avatares ouvindo em", s.addr)
  return http.ListenAndServe(s.addr, router)
}
