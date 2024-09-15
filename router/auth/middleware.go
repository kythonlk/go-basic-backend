package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"

	"github.com/kythonlk/go-basic-backend/types"
)

type contextKey string

const (
	usernameKey contextKey = "username"
	roleKey     contextKey = "role"
)

func RoleBasedMiddleware(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value("role")
		if userRole == nil || userRole != role {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims := &types.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return types.JwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, usernameKey, claims.Username)
		ctx = context.WithValue(ctx, roleKey, claims.Role)
		next(w, r.WithContext(ctx))
	}
}
