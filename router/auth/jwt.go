package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kythonlk/go-basic-backend/types"
)

func generateAccessToken(username, role string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(30) * time.Minute)
	claims := &types.Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(types.JwtKey)
}

func generateRefreshToken(username, role string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(7*24*60) * time.Minute)
	claims := &types.Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(types.JwtKey)
}
