package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func JsonError(w http.ResponseWriter, message string, statusCode int) {

	response := Response{
		Message: message,
		Status:  statusCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(response)
}
