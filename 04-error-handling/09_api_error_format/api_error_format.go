package api_error_format

import (
	"errors"
	"strings"
)

// ProblemDetail represents RFC 7807 Problem Details
type ProblemDetail struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

var ErrNotFound = errors.New("not found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrInvalidInput = errors.New("invalid input")

// NewProblemDetail creates a problem detail from an error
func NewProblemDetail(err error, requestPath string) ProblemDetail {
	// TODO(human): Implement
	result := ProblemDetail{
		Instance: requestPath,
	}

	if err == nil {
		return ProblemDetail{
			Type:   "200",
			Title:  "Found",
			Status: 200,
			Detail: "",
		}
	} else if errors.Is(err, ErrNotFound) {
		result.Type = "about:blank"
		result.Title = "Not Found"
		result.Status = 404
	} else if errors.Is(err, ErrUnauthorized) {
		result.Type = "about:blank"
		result.Title = "Unauthorized"
		result.Status = 401
	} else if errors.Is(err, ErrInvalidInput) {
		result.Type = "about:blank"
		result.Title = "Invalid"
		result.Status = 400
	} else {
		result.Type = "about:blank"
		result.Title = "Internal"
		result.Status = 500
	}
	result.Detail = err.Error()
	return result
}

// ValidationProblemDetail creates a problem detail for validation errors
func ValidationProblemDetail(errors []string, requestPath string) ProblemDetail {
	// TODO(human): Build a ProblemDetail with:
	// - Status: always 400
	// - Title: should contain "Validation"
	// - Detail: combine all error strings (hint: strings.Join)
	// - Instance: the requestPath
	// - Type: "about:blank"
	return ProblemDetail{
		Status:   400,
		Title:    "Validation",
		Detail:   strings.Join(errors, ","),
		Instance: requestPath,
		Type:     "about:blank",
	}
}
