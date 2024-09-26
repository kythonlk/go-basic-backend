package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kythonlk/go-basic-backend/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser inserts a new user into the database.
func CreateUser(email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	userID := uuid.New()
	query := `INSERT INTO users (id, email, password_hash) VALUES ($1, $2, $3)`

	_, err = DBPool.Exec(context.Background(), query, userID, email, string(hashedPassword))
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return &models.User{
		ID:           userID,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}, nil
}

// GetUserByEmail fetches a user by their email address.
func GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, password_hash FROM users WHERE email = $1`
	row := DBPool.QueryRow(context.Background(), query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &user, nil
}

// GetUserByID fetches a user by their UUID.
func GetUserByID(userID uuid.UUID) (*models.User, error) {
	query := `SELECT id, email, password_hash FROM users WHERE id = $1`
	row := DBPool.QueryRow(context.Background(), query, userID)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &user, nil
}
