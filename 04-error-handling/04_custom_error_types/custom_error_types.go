package custom_error_types

// TODO(human): Define ValidationError struct with Field, Value, Message

// TODO(human): Implement Error() method on *ValidationError

// TODO(human): Define HTTPError struct with StatusCode, Message, URL

// TODO(human): Implement Error() method on *HTTPError

// TODO(human): Define TimeoutError struct with Operation, Duration

// TODO(human): Implement Error() method on *TimeoutError

// ValidateUserInput checks user struct fields and returns ValidationError if invalid
func ValidateUserInput(name, email string, age int) error {
	// TODO(human): Implement
	return nil
}

// FetchResource simulates an HTTP request, returns HTTPError for bad status codes
func FetchResource(url string, statusCode int) (string, error) {
	// TODO(human): Implement
	return "", nil
}

// PerformOperation simulates an operation with timeout
func PerformOperation(name string, timeoutSeconds int) error {
	// TODO(human): Implement
	return nil
}

// ExtractValidationError attempts to extract ValidationError from error
func ExtractValidationError(err error) (*ValidationError, bool) {
	// TODO(human): Implement
	return nil, false
}
