package db

import (
	"database/sql"
	"log"
)

func setupTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			role TEXT NOT NULL DEFAULT 'user'
		)
	`)
	if err != nil {
		return err
	}
	log.Println("Creating tokens table")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tokens (
			username TEXT PRIMARY KEY,
			access_token TEXT,
			refresh_token TEXT,
			expires_at DATETIME
		)
	`)
	return err
}
