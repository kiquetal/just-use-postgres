package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Use DATABASE_URL environment variable or default to a local connection string
	// Format: postgres://username:password@localhost:5432/dbname
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// Default fallback for local development (matching docker-compose.yaml)
		connStr = "postgres://postgres:password@localhost:5432/exercises?sslmode=disable"
	}

	// Connect to the database
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	// Verify the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	// Example: Simple query
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Database version: %s\n", version)
}
