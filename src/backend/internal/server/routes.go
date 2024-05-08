// src/backend/internal/server/routes.go

package server

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/health", s.healthHandler)
	r.Get("/api/hello", s.helloHandler)
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	fs := http.FileServer(http.Dir(filesDir))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(filepath.Join(filesDir, r.RequestURI)); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(filesDir, "index.html"))
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) helloHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Hello world!",
	}
	jsonResp, _ := json.Marshal(data)
	_, _ = w.Write(jsonResp)
}
