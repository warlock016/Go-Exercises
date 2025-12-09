package sentinel_errors

import (
	"errors"
	"strings"
)

// TODO(human): Define sentinel errors here using var block
var (
	ErrNotFound      = errors.New("resource not found")
	ErrAlreadyExists = errors.New("resource already exists")
	ErrUnauthorized  = errors.New("unauthorized access")
	ErrInvalidInput  = errors.New("invalid input")
	ErrTimeout       = errors.New("operation timed out")
)

// FindUser returns ErrNotFound if user doesn't exist
func FindUser(userID int, users map[int]string) (string, error) {
	// TODO(human): Implement
	if _, exists := users[userID]; !exists {
		return "", ErrNotFound
	}
	return users[userID], nil
}

// CreateUser returns ErrAlreadyExists if user exists, ErrInvalidInput if name is empty
func CreateUser(userID int, name string, users map[int]string) error {
	// TODO(human): Implement
	if len(name) == 0 {
		return ErrInvalidInput
	}
	if _, exists := users[userID]; exists {
		return ErrAlreadyExists
	} else {
		users[userID] = name
	}
	return nil
}

// DeleteUser returns ErrNotFound if user doesn't exist, ErrUnauthorized if userID is 0
func DeleteUser(userID int, users map[int]string) error {
	// TODO(human): Implement
	if userID == 0 {
		return ErrUnauthorized
	}

	if _, exists := users[userID]; !exists {
		return ErrNotFound
	} else {
		delete(users, userID)
	}
	return nil
}

// IsNotFoundError checks if error is ErrNotFound using errors.Is
func IsNotFoundError(err error) bool {
	// TODO(human): Implement
	return err == ErrNotFound
}

// IsTimeoutError checks if error is ErrTimeout
func IsTimeoutError(err error) bool {
	// TODO(human): Implement
	return err == ErrTimeout
}

// HandleError returns a user-friendly message based on the error type
func HandleError(err error) string {
	// TODO(human): Implement
	var result strings.Builder
	switch err {
	case ErrNotFound:
		result.WriteString("resource not found")
	case ErrInvalidInput:
		result.WriteString("invalid input")
	case ErrAlreadyExists:
		result.WriteString("resource already exists")
	case ErrUnauthorized:
		result.WriteString("unauthorized access")
	case ErrTimeout:
		result.WriteString("operation timed out")
	default:
		result.WriteString("unknown error")
	}
	return result.String()
}
