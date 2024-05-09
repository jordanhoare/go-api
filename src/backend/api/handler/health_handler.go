// src/backend/api/handler/health_handler.go

package handler

import (
	"encoding/json"
	"net/http"

	"backend/internal/model"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(model.HealthResponse{Status: "OK"})
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
