package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func Logger(next http.Handler) http.Handler {
	// Placeholder for custom logging
	// For now, just pass the request through
	return middleware.Logger(next)
}
