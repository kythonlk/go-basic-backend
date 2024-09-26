package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/kythonlk/go-basic-backend/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			render.JSON(w, r, map[string]string{"error": "missing auth token"})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			render.JSON(w, r, map[string]string{"error": "invalid auth token format"})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := tokenParts[1]
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			render.JSON(w, r, map[string]string{"error": "invalid or expired token"})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
