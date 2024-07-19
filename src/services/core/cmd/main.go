package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/core/cmd/api"
	"github.com/Inteli-College/2024-1B-T09-ES06-G04/core/config"
	"github.com/Inteli-College/2024-1B-T09-ES06-G04/core/db"
)

// main function to set up configuration, database connection, and start the server
func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := db.NewPostgresStorage(dsn)

	initStorage(db)

	port := fmt.Sprintf(":%s", cfg.WebServerPort)
	server := api.NewApiServer(port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// initStorage function to check the database connection
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("database connection pool established")
}
