package sentinel_errors

// TODO(human): Define sentinel errors here using var block

// FindUser returns ErrNotFound if user doesn't exist
func FindUser(userID int, users map[int]string) (string, error) {
	// TODO(human): Implement
	return "", nil
}

// CreateUser returns ErrAlreadyExists if user exists, ErrInvalidInput if name is empty
func CreateUser(userID int, name string, users map[int]string) error {
	// TODO(human): Implement
	return nil
}

// DeleteUser returns ErrNotFound if user doesn't exist, ErrUnauthorized if userID is 0
func DeleteUser(userID int, users map[int]string) error {
	// TODO(human): Implement
	return nil
}

// IsNotFoundError checks if error is ErrNotFound using errors.Is
func IsNotFoundError(err error) bool {
	// TODO(human): Implement
	return false
}

// IsTimeoutError checks if error is ErrTimeout
func IsTimeoutError(err error) bool {
	// TODO(human): Implement
	return false
}

// HandleError returns a user-friendly message based on the error type
func HandleError(err error) string {
	// TODO(human): Implement
	return ""
}
