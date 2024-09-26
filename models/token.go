package models

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kythonlk/go-basic-backend/cmd"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT with.
func GenerateToken(userID uuid.UUID, role string, expiresIn time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiresIn)

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(cmd.JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken checks JWT is valid.
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return cmd.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token or signature")
	}

	return claims, nil
}
