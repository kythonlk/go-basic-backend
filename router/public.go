
package router

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

// PublicRoutes registers public routes such as health check.
func PublicRoutes(r chi.Router) {
	r.Get("/health", HealthCheckHandler)
}

// HealthCheckHandler provides a simple status response.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
