package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/kythonlk/go-basic-backend/router/server"
)

func main() {
	server := server.NewServer()

	db, err := connectPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	// Create a simple table
	err = createTable(db)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

func connectPostgres() (*sql.DB, error) {
	connStr := "postgres://youruser:yourpassword@localhost:5432/yourdb?sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL using pgx successfully!")
	return db, nil
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	fmt.Println("Table 'users' created successfully!")
	return nil
}
