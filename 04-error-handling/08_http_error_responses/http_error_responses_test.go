package http_error_responses

import (
	"errors"
	"testing"
)

var ErrNotFound = errors.New("not found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrInvalidInput = errors.New("invalid input")

func TestErrorToStatusCode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{"not found", ErrNotFound, 404},
		{"unauthorized", ErrUnauthorized, 401},
		{"invalid input", ErrInvalidInput, 400},
		{"nil error", nil, 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrorToStatusCode(tt.err)
			if got != tt.want {
				t.Errorf("ErrorToStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewErrorResponse(t *testing.T) {
	resp := NewErrorResponse(ErrNotFound)
	if resp.Status == 0 {
		t.Error("NewErrorResponse() Status is 0")
	}
	if resp.Message == "" {
		t.Error("NewErrorResponse() Message is empty")
	}
}
