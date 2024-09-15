package types

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
)

var JwtKey = []byte(os.Getenv("JWT_TOKEN"))

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
