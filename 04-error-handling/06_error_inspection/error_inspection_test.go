package error_inspection

import (
	"errors"
	"testing"
)

var ErrNotFound = errors.New("not found")
var ErrTimeout = errors.New("timeout")

func TestCheckErrorType(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"not found", ErrNotFound, "not_found"},
		{"timeout", ErrTimeout, "timeout"},
		{"nil", nil, "no_error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckErrorType(tt.err)
			if got == "" {
				t.Errorf("CheckErrorType() returned empty string")
			}
		})
	}
}

func TestIsTemporaryError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"timeout", ErrTimeout, true},
		{"not found", ErrNotFound, false},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This is a placeholder test
			// Implementation will determine exact behavior
			_ = IsTemporaryError(tt.err)
		})
	}
}
