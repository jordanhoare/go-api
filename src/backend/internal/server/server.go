// src/backend/internal/server/server.go

package server

import (
	"fmt"
	"net/http"
	"time"

	"backend/api/routes"
	"backend/internal/config"
	"backend/internal/database"
)

type Server struct {
	Port         int
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Db           database.Service
}

func NewServer() *http.Server {
	serverCfg := config.LoadServerConfig()
	dbCfg := config.LoadDatabaseConfig()
	dbService := database.New(dbCfg)

	server := &Server{
		Port:         serverCfg.Port,
		IdleTimeout:  serverCfg.IdleTimeout,
		ReadTimeout:  serverCfg.ReadTimeout,
		WriteTimeout: serverCfg.WriteTimeout,
		Db:           dbService,
	}

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", server.Port),
		Handler:      routes.NewRouter(),
		IdleTimeout:  server.IdleTimeout,
		ReadTimeout:  server.ReadTimeout,
		WriteTimeout: server.WriteTimeout,
	}

	return httpServer
}
