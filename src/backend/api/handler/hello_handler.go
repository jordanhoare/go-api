package handler

import (
	"backend/internal/model"
	"encoding/json"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(model.HelloResponse{Message: "Hello world!"})
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
