package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/kythonlk/go-basic-backend/services"
	"net/http"
)

// AuthRoutes registers authentication routes.
func AuthRoutes(r chi.Router) {
	r.Post("/register", RegisterHandler)
	r.Post("/login", LoginHandler)
}

// RegisterHandler handles new user registration.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := services.RegisterUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("User registered successfully"))
}

// LoginHandler handles user login and token generation.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	token, err := services.LoginUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Write([]byte(token))
}
