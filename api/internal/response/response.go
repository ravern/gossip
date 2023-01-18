package response

import (
	"encoding/json"
	"net/http"
)

type OkResponse struct {
	Data interface{} `json:"data"`
}

func WriteJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&OkResponse{
		Data: payload,
	}); err != nil {
		InternalServerError(w)
	}
}

type ErrorResponse struct {
	Error ErrorResponseError `json:"error"`
}

type ErrorResponseError struct {
	Message string `json:"message"`
}

func JSONError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(&ErrorResponse{
		Error: ErrorResponseError{
			Message: err.Error(),
		},
	}); err != nil {
		InternalServerError(w)
	}
}

func InternalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
