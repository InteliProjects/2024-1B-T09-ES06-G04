package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// NewPostgresStorage initializes a new PostgreSQL storage with the given connection string
func NewPostgresStorage(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
