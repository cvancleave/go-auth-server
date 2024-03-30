package utils

import (
	"encoding/json"
	"net/http"
)

func SetCorsHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

func RespondJson(w http.ResponseWriter, code int, data any) {
	body, _ := json.Marshal(data)
	w.WriteHeader(code)
	w.Write(body)
}

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func RespondError(w http.ResponseWriter, code int, message string) {
	response := ErrorResponse{
		ErrorCode:    code,
		ErrorMessage: message,
	}
	body, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Write(body)
}
