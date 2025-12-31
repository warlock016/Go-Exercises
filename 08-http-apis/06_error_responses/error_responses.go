package error_responses

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// APIError represents a structured error response
type APIError struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

// WriteError writes a JSON error response
func WriteError(w http.ResponseWriter, status int, code, message string) {
	// TODO(human): Implement
	newErr := APIError{
		Code:    code,
		Message: message,
		Details: make(map[string]string),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(newErr)
}

// WriteErrorWithDetails writes a JSON error with additional details
func WriteErrorWithDetails(w http.ResponseWriter, status int, code, message string, details map[string]string) {
	// TODO(human): Implement
	newErr := APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(newErr)
}

// ValidationHandler demonstrates various error responses
func ValidationHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	w.Header().Set("Content-Type", "application/json")
	newErr := APIError{
		Details: make(map[string]string),
	}

	params := r.URL.Query()

	email := params.Get("email")
	if email == "" {
		newErr.Code = "MISSING_PARAM"
		newErr.Message = "email is required"
		newErr.Details["field"] = "email"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(newErr)
		return
	}

	age := params.Get("age")
	ageInt, err := strconv.ParseInt(age, 10, 64)

	if err != nil {
		newErr.Code = "MISSING_PARAM"
		newErr.Message = "age is required"
		newErr.Details["field"] = "age"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(newErr)
		return
	}

	if ageInt < 0 || ageInt > 150 {
		newErr.Code = "VALIDATION_ERROR"
		newErr.Message = "invalid age"
		newErr.Details["field"] = "age"
		newErr.Details["value"] = age
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(newErr)
		return
	}
}
