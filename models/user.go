package models

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User struct defines the user model.
type User struct {
	ID               uuid.UUID `json:"id"`
	Email            string    `json:"email"`
	PasswordHash     string    `json:"-"`
	IsVerified       bool      `json:"is_verified"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	LastLogin        time.Time `json:"last_login"`
	ResetToken       string    `json:"-"`
	ResetTokenExpiry time.Time `json:"-"`
	Role             string    `json:"role"`
}

// SetPassword encrypts the password before storing it.
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password: " + err.Error())
	}
	u.PasswordHash = string(hash)
	return nil
}

// Check Password compares the password with the stored password.
func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return errors.New("incorrect password")
	}
	return nil
}
