package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

type JSONResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp string      `json:"timestamp"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, status, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(statusCode)

	response := JSONResponse{
		Status:    status,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	json.NewEncoder(w).Encode(response)
}
