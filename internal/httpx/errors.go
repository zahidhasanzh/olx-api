package httpx

import (
	"encoding/json"
	"net/http"
)

type Code string

const (
	CodeInvalidId     Code = "invalid_id"
	NotFound          Code = "not_found"
	CodeInternalError Code = "internal_error"
	MalFormedJson     Code = "malformed_json"
	CodeValidationError  Code = "validation_failed"
	Unauthenticated   Code = "unauthenticated"
	Forbidden         Code = "forbidden"
	Conflict          Code = "conflict"
	RateLimited       Code = "rate_limited"
)

type errorEnvelope struct {
	Error errorPyload `json:"error"`
}

type errorPyload struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field", omitempty`
}

func Error(w http.ResponseWriter, status int, message string, code Code) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(errorEnvelope{errorPyload{
		Code:    code,
		Message: message,
	}})

}

func ValidationError(w http.ResponseWriter, status int, message string, code Code, field string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(errorEnvelope{errorPyload{
		Code:    code,
		Message: message,
		Field:   field,
	}})

}
