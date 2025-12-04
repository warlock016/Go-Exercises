package http_error_responses

// ErrorToStatusCode maps errors to HTTP status codes
func ErrorToStatusCode(err error) int {
	// TODO(human): Implement
	return 0
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
	return ErrorResponse{}
}

// WriteErrorResponse simulates writing JSON error response
func WriteErrorResponse(err error) string {
	// TODO(human): Implement
	return ""
}
