package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/Inteli-College/2024-1B-T09-ES06-G04/avatars/config"
		"github.com/Inteli-College/2024-1B-T09-ES06-G04/avatars/service/avatar"
	)

// Main function to start the web server
func main() {
    cfg, err := config.LoadConfig(".")
    if err != nil {
        log.Fatal("Error loading config: ", err)
    }
    port := fmt.Sprintf(":%s", cfg.WebServerPort)
    router := mux.NewRouter()

    avatarHandler := avatar.NewHandler()
    avatarHandler.RegisterRoutes(router)

    log.Println("Server listening on", port)
    log.Fatal(http.ListenAndServe(port, router))
}
