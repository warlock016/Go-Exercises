package custom_error_types

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// TODO(human): Define ValidationError struct with Field, Value, Message
type ValidationError struct {
	Field   string
	Value   any
	Message string
}

// TODO(human): Implement Error() method on *ValidationError
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s (got %v)", e.Field, e.Message, e.Value)
}

// TODO(human): Define HTTPError struct with StatusCode, Message, URL
type HTTPError struct {
	StatusCode int
	Message    string
	URL        string
}

// TODO(human): Implement Error() method on *HTTPError
func (e *HTTPError) Error() string {
	return fmt.Sprintf("%s: %s (got %v)", e.URL, e.Message, e.StatusCode)
}

// TODO(human): Define TimeoutError struct with Operation, Duration
type TimeoutError struct {
	Operation string
	Duration  time.Duration
}

// TODO(human): Implement Error() method on *TimeoutError
func (e *TimeoutError) Error() string {
	return fmt.Sprintf("%s: (got %v)", e.Operation, e.Duration)
}

// ValidateUserInput checks user struct fields and returns ValidationError if invalid
func ValidateUserInput(name, email string, age int) error {
	// TODO(human): Implement
	if name == "" {
		return &ValidationError{
			Field:   "name",
			Value:   name,
			Message: "cannot be empty",
		}
	} else if !strings.Contains(email, "@") {
		return &ValidationError{
			Field:   "email",
			Value:   email,
			Message: "invalid email",
		}
	} else if age < 0 || age > 150 {
		return &ValidationError{
			Field:   "age",
			Value:   age,
			Message: "invalid age",
		}
	}
	return nil
}

// FetchResource simulates an HTTP request, returns HTTPError for bad status codes
func FetchResource(url string, statusCode int) (string, error) {
	// TODO(human): Implement
	if statusCode >= 400 {
		return "", &HTTPError{
			StatusCode: statusCode,
			Message:    "Request failed",
			URL:        url,
		}
	}

	return "success", nil
}

// PerformOperation simulates an operation with timeout
func PerformOperation(name string, timeoutSeconds int) error {
	// TODO(human): Implement

	if timeoutSeconds > 30 {
		return &TimeoutError{
			Operation: name,
			Duration:  time.Duration(timeoutSeconds) * time.Second,
		}
	}
	return nil
}

// ExtractValidationError attempts to extract ValidationError from error
func ExtractValidationError(err error) (*ValidationError, bool) {
	// TODO(human): Implement
	var verr *ValidationError
	if errors.As(err, &verr) {
		return verr, true
	}
	return nil, false
}
