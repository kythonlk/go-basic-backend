package services

import (
	"errors"
	"github.com/kythonlk/go-basic-backend/database"
	"github.com/kythonlk/go-basic-backend/utils"
	"net/http"
)

// LoginUser handles user login, verifies credentials, and returns a JWT token.
func LoginUser(r *http.Request) (string, error) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := utils.ParseJSON(r, &credentials); err != nil {
		return "", err
	}

	user, err := database.GetUserByEmail(credentials.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPassword(credentials.Password, user.PasswordHash) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
