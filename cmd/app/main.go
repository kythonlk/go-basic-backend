package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kythonlk/go-basic-backend/cmd"
	"github.com/kythonlk/go-basic-backend/database"
	// mi "github.com/kythonlk/go-basic-backend/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kythonlk/go-basic-backend/router"
)

func main() {
	// Initialize database connection
	if err := database.ConnectDB(cmd.ConnectionString); err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer database.CloseDB()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(mi.RateLimiter)

	router.PublicRoutes(r)
	router.AuthRoutes(r)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
