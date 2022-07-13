package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, statusCode int, err error) {
	RespondWithJSON(w, statusCode, map[string]string{"error": err.Error()})
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
