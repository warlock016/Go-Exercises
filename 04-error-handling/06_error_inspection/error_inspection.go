package error_inspection

import (
	"errors"
	"fmt"
)

// CheckErrorType inspects an error and returns what type it is
func CheckErrorType(err error) string {
	// TODO(human): Implement
	if errors.Is(err, nil) {
		return "no_error"
	} else if errors.Is(err, ErrNotFound) {
		return "not_found"
	} else if errors.Is(err, ErrTimeout) {
		return "timeout"
	} else {
		return "unknown"
	}
}

// ExtractHTTPStatus extracts the HTTP status code from a wrapped HTTPError
func ExtractHTTPStatus(err error) (int, bool) {
	// TODO(human): Implement
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		return httpErr.StatusCode, true
	}
	return 0, false
}

// IsTemporaryError checks if error is a temporary/transient error
func IsTemporaryError(err error) bool {
	// TODO(human): Implement
	if errors.Is(err, ErrTimeout) {
		return true
	}

	var tempErr *TemporaryError
	if errors.As(err, &tempErr) {
		return true
	}

	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		return httpErr.StatusCode == 503
	}

	return false
}

// HandleBasedOnType performs different actions based on error type
func HandleBasedOnType(err error) string {
	// TODO(human): Implement
	if err == nil {
		return "no_error: success"
	}
	if errors.Is(err, ErrNotFound) {
		return "not_found: return 404"
	}
	if errors.Is(err, ErrTimeout) {
		return "timeout: retry operation"
	}
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		return fmt.Sprintf("http_error: status %d", httpErr.StatusCode)
	}
	var tempErr *TemporaryError
	if errors.As(err, &tempErr) {
		return "temporary: will retry"
	}
	return "unknown: log and alert"
}
