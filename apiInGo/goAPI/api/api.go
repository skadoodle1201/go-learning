package api

import (
	"encoding/json"
	"net/http"
)

type GetCoinRequest struct {
	Username string
}

type GetCoinResponse struct {
	Username string
	Coins    int64
	Message  string
	Code     int
}

type ErrorResponse struct {
	Code    int
	Message string
}

func WriteError(w http.ResponseWriter, message string, code int) {
	resp := ErrorResponse{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		WriteError(w, "Internal Server Error", http.StatusInternalServerError)
	}
)
