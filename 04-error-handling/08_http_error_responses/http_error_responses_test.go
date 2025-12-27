package http_error_responses

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"
)

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

func TestErrorToStatusCode_WrappedErrors(t *testing.T) {
	// Test that errors.Is() works with wrapped errors
	tests := []struct {
		name string
		err  error
		want int
	}{
		{"wrapped not found", fmt.Errorf("user 123: %w", ErrNotFound), 404},
		{"wrapped unauthorized", fmt.Errorf("token expired: %w", ErrUnauthorized), 401},
		{"wrapped invalid input", fmt.Errorf("field email: %w", ErrInvalidInput), 400},
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

func TestErrorToStatusCode_UnknownError(t *testing.T) {
	unknownErr := errors.New("some random error")
	got := ErrorToStatusCode(unknownErr)
	// Unknown errors should map to 500 (Internal Server Error)
	// Currently returns 0 - this test documents expected behavior
	if got != 500 {
		t.Errorf("ErrorToStatusCode(unknown) = %v, want 500", got)
	}
}

func TestNewErrorResponse(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		wantStatus int
		wantCode   string
	}{
		{"not found", ErrNotFound, 404, "NOT_FOUND"},
		{"unauthorized", ErrUnauthorized, 401, "UNAUTHORIZED"},
		{"invalid input", ErrInvalidInput, 400, "INVALID_INPUT"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := NewErrorResponse(tt.err)
			if resp.Status != tt.wantStatus {
				t.Errorf("NewErrorResponse().Status = %v, want %v", resp.Status, tt.wantStatus)
			}
			if resp.Code != tt.wantCode {
				t.Errorf("NewErrorResponse().Code = %v, want %v", resp.Code, tt.wantCode)
			}
			if resp.Message == "" {
				t.Error("NewErrorResponse().Message is empty")
			}
		})
	}
}

func TestWriteErrorResponse(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		wantContains []string
	}{
		{
			"not found",
			ErrNotFound,
			[]string{`"status":404`, `"message":`, `"code":`},
		},
		{
			"unauthorized",
			ErrUnauthorized,
			[]string{`"status":401`, `"message":`, `"code":`},
		},
		{
			"invalid input",
			ErrInvalidInput,
			[]string{`"status":400`, `"message":`, `"code":`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WriteErrorResponse(tt.err)
			if got == "" {
				t.Error("WriteErrorResponse() returned empty string")
				return
			}
			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("WriteErrorResponse() = %v, want to contain %v", got, want)
				}
			}
		})
	}
}

func TestWriteErrorResponse_ValidJSON(t *testing.T) {
	jsonStr := WriteErrorResponse(ErrNotFound)
	var resp ErrorResponse
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		t.Errorf("WriteErrorResponse() returned invalid JSON: %v", err)
	}
	if resp.Status != 404 {
		t.Errorf("Unmarshaled Status = %v, want 404", resp.Status)
	}
}
