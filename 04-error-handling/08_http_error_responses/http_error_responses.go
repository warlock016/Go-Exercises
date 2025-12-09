package http_error_responses

import (
	"encoding/json"
	"errors"
	"strconv"
)

// ErrorToStatusCode maps errors to HTTP status codes
func ErrorToStatusCode(err error) int {
	// TODO(human): Implement
	if err == nil {
		return 200
	}
	if errors.Is(err, ErrNotFound) {
		return 404
	}
	if errors.Is(err, ErrUnauthorized) {
		return 401
	}
	if errors.Is(err, ErrInvalidInput) {
		return 400
	}

	return 500 // internal server error
}

// ErrorResponse creates a JSON-serializable error response
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// NewErrorResponse creates an ErrorResponse from an error
func NewErrorResponse(err error) ErrorResponse {
	// TODO(human): Implement
	resp := ErrorResponse{
		Status: ErrorToStatusCode(err),
	}
	if err == nil {
		return ErrorResponse{}
	}
	if errors.Is(err, ErrNotFound) {
		resp.Code = "NOT_FOUND"
		resp.Message = "not found"
	}
	if errors.Is(err, ErrUnauthorized) {
		resp.Code = "UNAUTHORIZED"
		resp.Message = "unauthorized"
	}
	if errors.Is(err, ErrInvalidInput) {
		resp.Code = "INVALID_INPUT"
		resp.Message = "invalid input"
	}
	return resp
}

// WriteErrorResponse simulates writing JSON error response
func WriteErrorResponse(err error) string {
	// TODO(human): Implement
	result := ErrorResponse{
		Status:  ErrorToStatusCode(err),
		Message: "Status Code",
		Code:    strconv.FormatInt(int64(ErrorToStatusCode(err)), 32),
	}

	output, err := json.Marshal(result)
	if err != nil {
		return ""
	}
	// fmt.Println(output)
	return string(output)
}
