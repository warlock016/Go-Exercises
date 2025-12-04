package error_inspection

// CheckErrorType inspects an error and returns what type it is
func CheckErrorType(err error) string {
	// TODO(human): Implement
	return ""
}

// ExtractHTTPStatus extracts the HTTP status code from a wrapped HTTPError
func ExtractHTTPStatus(err error) (int, bool) {
	// TODO(human): Implement
	return 0, false
}

// IsTemporaryError checks if error is a temporary/transient error
func IsTemporaryError(err error) bool {
	// TODO(human): Implement
	return false
}

// HandleBasedOnType performs different actions based on error type
func HandleBasedOnType(err error) string {
	// TODO(human): Implement
	return ""
}
