package error_interface

import (
	"errors"
	"testing"
)

func TestValidationError(t *testing.T) {
	err := &ValidationError{
		Field:   "username",
		Problem: "too short",
		MinLen:  3,
	}

	want := "validation error: username is too short (minimum 3 characters)"
	got := err.Error()

	if got != want {
		t.Errorf("ValidationError.Error() = %q, want %q", got, want)
	}
}

func TestNetworkError(t *testing.T) {
	err := &NetworkError{
		URL:        "https://api.example.com",
		StatusCode: 404,
		Message:    "request failed",
	}

	want := "network error: request failed (URL: https://api.example.com, status: 404)"
	got := err.Error()

	if got != want {
		t.Errorf("NetworkError.Error() = %q, want %q", got, want)
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  bool
		errMsg   string
	}{
		{"Valid username", "alice", false, ""},
		{"Valid longer username", "bobsmith", false, ""},
		{"Exactly 3 chars", "joe", false, ""},
		{"Too short - 2 chars", "ab", true, "validation error: username is too short (minimum 3 characters)"},
		{"Too short - 1 char", "x", true, "validation error: username is too short (minimum 3 characters)"},
		{"Empty string", "", true, "validation error: username is too short (minimum 3 characters)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.username)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Validate(%q) expected error, got nil", tt.username)
					return
				}
				if err.Error() != tt.errMsg {
					t.Errorf("Validate(%q) error = %q, want %q", tt.username, err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("Validate(%q) expected no error, got %v", tt.username, err)
				}
			}
		})
	}
}

func TestValidateReturnsValidationError(t *testing.T) {
	err := Validate("ab")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Type assertion to check it's the right error type
	verr, ok := err.(*ValidationError)
	if !ok {
		t.Fatalf("expected *ValidationError, got %T", err)
	}

	if verr.Field != "username" {
		t.Errorf("Field = %q, want %q", verr.Field, "username")
	}
	if verr.Problem != "too short" {
		t.Errorf("Problem = %q, want %q", verr.Problem, "too short")
	}
	if verr.MinLen != 3 {
		t.Errorf("MinLen = %d, want %d", verr.MinLen, 3)
	}
}

func TestFetchData(t *testing.T) {
	tests := []struct {
		name            string
		url             string
		simulateFailure bool
		wantData        string
		wantErr         bool
		errMsg          string
	}{
		{
			name:            "Success",
			url:             "https://api.example.com",
			simulateFailure: false,
			wantData:        "data from https://api.example.com",
			wantErr:         false,
		},
		{
			name:            "Failure",
			url:             "https://api.example.com",
			simulateFailure: true,
			wantData:        "",
			wantErr:         true,
			errMsg:          "network error: request failed (URL: https://api.example.com, status: 404)",
		},
		{
			name:            "Different URL success",
			url:             "https://other.com/data",
			simulateFailure: false,
			wantData:        "data from https://other.com/data",
			wantErr:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := FetchData(tt.url, tt.simulateFailure)

			if tt.wantErr {
				if err == nil {
					t.Errorf("FetchData() expected error, got nil")
					return
				}
				if err.Error() != tt.errMsg {
					t.Errorf("FetchData() error = %q, want %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("FetchData() expected no error, got %v", err)
				}
				if data != tt.wantData {
					t.Errorf("FetchData() data = %q, want %q", data, tt.wantData)
				}
			}
		})
	}
}

func TestFetchDataReturnsNetworkError(t *testing.T) {
	url := "https://api.example.com"
	_, err := FetchData(url, true)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Type assertion to check it's the right error type
	nerr, ok := err.(*NetworkError)
	if !ok {
		t.Fatalf("expected *NetworkError, got %T", err)
	}

	if nerr.URL != url {
		t.Errorf("URL = %q, want %q", nerr.URL, url)
	}
	if nerr.StatusCode != 404 {
		t.Errorf("StatusCode = %d, want %d", nerr.StatusCode, 404)
	}
	if nerr.Message != "request failed" {
		t.Errorf("Message = %q, want %q", nerr.Message, "request failed")
	}
}

func TestErrorsImplementErrorInterface(t *testing.T) {
	// Verify our types satisfy error interface
	var _ error = &ValidationError{}
	var _ error = &NetworkError{}

	t.Log("✓ Both types implement error interface")
}

func TestErrorsAreErrors(t *testing.T) {
	// Verify we can use our errors with errors package
	verr := &ValidationError{Field: "test", Problem: "test", MinLen: 3}
	nerr := &NetworkError{URL: "test", StatusCode: 404, Message: "test"}

	if !errors.As(verr, &verr) {
		t.Error("ValidationError should work with errors.As")
	}

	if !errors.As(nerr, &nerr) {
		t.Error("NetworkError should work with errors.As")
	}
}
