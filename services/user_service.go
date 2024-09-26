package services

import (
	"errors"
	"github.com/kythonlk/go-basic-backend/database"
	"github.com/kythonlk/go-basic-backend/utils"
	"net/http"
)

// RegisterUser handles new user registration.
func RegisterUser(r *http.Request) error {
	var newUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := utils.ParseJSON(r, &newUser); err != nil {
		return err
	}

	_, err := database.GetUserByEmail(newUser.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	_, err = database.CreateUser(newUser.Email, newUser.Password)
	if err != nil {
		return errors.New("could not create user")
	}

	go SendWelcomeEmail(newUser.Email)

	return nil
}
