package error_wrapping

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// OpenAndReadFile opens a file and reads it, wrapping errors at each step
func OpenAndReadFile(filename string) (string, error) {
	// TODO(human): Implement
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open %s: %w", filename, err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", filename, err)
	}
	return string(bytes), nil
}

// ProcessUser validates and saves user, wrapping errors with context
func ProcessUser(name, email string) error {
	// TODO(human): Implement
	if name == "" {
		return fmt.Errorf("failed to process: %w", errors.New("empty name"))
	}
	if !strings.Contains(email, "@") {
		return fmt.Errorf("failed to process: %w", errors.New("invalid email"))
	}
	return nil
}

// FetchAndParse fetches URL and parses JSON, wrapping errors
func FetchAndParse(url string) (map[string]any, error) {
	// TODO(human): Implement
	if url == "" {
		return nil, fmt.Errorf("failed to fetch: %w", errors.New("empty URL"))
	}
	if !strings.Contains(url, "http") {
		return nil, fmt.Errorf("failed to fetch %w", errors.New("invalid URL scheme"))
	}
	return map[string]any{"url": url, "status": "ok"}, nil
}

// UnwrapOnce unwraps one level of error wrapping
func UnwrapOnce(err error) error {
	// TODO(human): Implement
	return errors.Unwrap(err)
}

// UnwrapAll repeatedly unwraps until reaching the root error
func UnwrapAll(err error) error {
	// TODO(human): Implement
	for {
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			return err
		}
		err = unwrapped
	}
}
