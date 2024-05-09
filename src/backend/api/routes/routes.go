// src/backend/api/routes/routes.go

package routes

import (
	"net/http"
	"os"
	"path/filepath"

	"backend/api/handler"
	"backend/api/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.RateLimit)
	r.Use(middleware.RequestID)

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", handler.HealthHandler)
		r.Get("/hello", handler.HelloHandler)
	})

	// Static files
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	fs := http.FileServer(http.Dir(filesDir))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(filesDir, r.RequestURI)
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(filesDir, "index.html"))
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	return r
}
