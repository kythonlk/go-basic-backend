package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DBPool is a globally accessible pool of database connections.
var DBPool *pgxpool.Pool

func ConnectDB(connectionString string) error {
	var err error
	DBPool, err = pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	err = DBPool.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	log.Println("Connected to the database successfully.")
	return nil
}

// CloseDB closes the database connection when the application shuts down.
func CloseDB() {
	DBPool.Close()
	log.Println("Database connection closed.")
}
