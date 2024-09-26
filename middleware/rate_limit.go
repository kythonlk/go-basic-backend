package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/time/rate"
)

// RateLimiter limits the number of requests a user can make within a specific time frame.
func RateLimiter(rps int, burst int) func(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(rate.Limit(rps), burst)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func BasicRateLimiter() func(next http.Handler) http.Handler {
	return middleware.Throttle(100)
}
