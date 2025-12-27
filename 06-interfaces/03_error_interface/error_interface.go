package error_interface

import (
	"fmt"
)

// TODO(human): Define ValidationError struct
type ValidationError struct {
	Field   string
	Problem string
	MinLen  int
}

// TODO(human): Define NetworkError struct
type NetworkError struct {
	URL        string
	StatusCode int
	Message    string
}

// TODO(human): Implement Error() method for ValidationError
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s is %s (minimum %d characters)", e.Field, e.Problem, e.MinLen)
}

// TODO(human): Implement Error() method for NetworkError
func (e *NetworkError) Error() string {
	return fmt.Sprintf("network error: %s (URL: %s, status: %d)", e.Message, e.URL, e.StatusCode)
}

// Validate checks if a username meets the minimum length requirement
func Validate(username string) error {
	// TODO(human): Implement
	if len([]rune(username)) < 3 {
		return &ValidationError{
			Field:   "username",
			Problem: "too short",
			MinLen:  3,
		}
	}
	return nil
}

// FetchData simulates fetching data from a URL
func FetchData(url string, simulateFailure bool) (string, error) {
	// TODO(human): Implement

	// var result strings.Builder

	if simulateFailure {
		return "", &NetworkError{
			URL:        url,
			StatusCode: 404,
			Message:    "request failed",
		}
	}

	return fmt.Sprintf("data from %s", url), nil
}
