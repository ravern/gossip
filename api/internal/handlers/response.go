package handlers

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error errorResponseError
}

type errorResponseError struct {
	Message string
}

func jsonError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(&errorResponse{
		Error: errorResponseError{
			Message: err.Error(),
		},
	}); err != nil {
		internalServerError(w)
	}
}

func internalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
