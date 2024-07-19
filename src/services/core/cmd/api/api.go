package api

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/Inteli-College/2024-1B-T09-ES06-G04/core/service/user"
	"github.com/gorilla/mux"
)

// ApiServer struct defines the structure of the API server
type ApiServer struct {
	addr string
	db   *sql.DB
}

// NewApiServer creates a new instance of ApiServer.
func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

// newReverseProxy creates a new reverse proxy.
func newReverseProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ModifyResponse = func(resp *http.Response) error {
		return nil
	}

	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = url.Path + req.URL.Path
		req.Host = url.Host

		// Log headers
		log.Println("Request Headers:", req.Header)

		// Extract and verify the JWT
		authHeader := req.Header.Get("Authorization")
		if authHeader != "" {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Decode and verify the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Make sure that the token method conforms to "alg" value.
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrInvalidKey
				}
				return []byte("secret"), nil
			})

			if err != nil || !token.Valid {
				log.Println("Invalid JWT Token")
				return
			}

			// Extract claims
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				log.Println("JWT Claims:", claims)
			} else {
				log.Println("Invalid JWT Claims")
			}
		}
	}

	return proxy
}

// Run starts the API server.
func (s *ApiServer) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// redirect to projects service 
	target := "http://projects:8080"
	proxy := newReverseProxy(target)

	subrouter.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}).Methods(http.MethodGet, http.MethodPost)

	subrouter.HandleFunc("/projects/user", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)

	subrouter.HandleFunc("/projects/user/{userID}", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}).Methods(http.MethodGet)

	subrouter.HandleFunc("/projects/{id}", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	// redirect to the Python service
	targetPythonService := "http://model:5000"
	proxyPythonService := newReverseProxy(targetPythonService)

	subrouter.HandleFunc("/model-ratings", func(w http.ResponseWriter, r *http.Request) {
		proxyPythonService.ServeHTTP(w, r)
	}).Methods(http.MethodGet)

	// redirect to the Connections service 
	targetConnectionService := "http://connections:8081"
	proxyConnectionService := newReverseProxy(targetConnectionService)

	subrouter.HandleFunc("/connections", func(w http.ResponseWriter, r *http.Request) {
		proxyConnectionService.ServeHTTP(w, r)
	}).Methods(http.MethodGet, http.MethodPost, http.MethodPut)

	subrouter.HandleFunc("/connections/{ID}", func(w http.ResponseWriter, r *http.Request) {
		proxyConnectionService.ServeHTTP(w, r)
	}).Methods(http.MethodPut)

	subrouter.HandleFunc("/ratings", func(w http.ResponseWriter, r *http.Request) {
		proxyConnectionService.ServeHTTP(w, r)
	}).Methods(http.MethodGet, http.MethodPost)

	subrouter.HandleFunc("/connections/true", func(w http.ResponseWriter, r *http.Request) {
		proxyConnectionService.ServeHTTP(w, r)
	}).Methods(http.MethodGet, http.MethodPost)

	// redirect to the Avatars service 
	targetAvatarService := "http://:8083"
	proxyAvatarService := newReverseProxy(targetAvatarService)

	subrouter.HandleFunc("/avatars", func(w http.ResponseWriter, r *http.Request) {
		proxyAvatarService.ServeHTTP(w, r)
	}).Methods(http.MethodGet)


	log.Println("listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

