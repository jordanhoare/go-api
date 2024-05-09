package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func RequestID(next http.Handler) http.Handler {
	// Placeholder for custom request id handling
	// RequestID wraps Chi's RequestID middleware
	// For now, just pass the request through
	return middleware.RequestID(next)
}
