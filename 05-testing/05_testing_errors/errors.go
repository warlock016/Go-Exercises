package errors

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ParseInt parses a string to an integer
func ParseInt(s string) (int, error) {
	if s == "" {
		return 0, errors.New("empty string")
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid integer: %w", err)
	}
	return n, nil
}

// Divide divides two numbers
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// ReadFile reads a file and returns its contents
func ReadFile(path string) (string, error) {
	if path == "" {
		return "", errors.New("empty path")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(data), nil
}
