package middleware

import (
	"net/http"
)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Placeholder for rate limiting logic
		// Logger wraps Chi's built-in logger middleware
		// For now, just pass the request through
		next.ServeHTTP(w, r)
	})
}
