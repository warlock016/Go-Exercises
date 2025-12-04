package validation_errors

import (
	"testing"
)

func TestValidationErrors_Add(t *testing.T) {
	ve := &ValidationErrors{}
	ve.Add("name", "cannot be empty")
	ve.Add("email", "invalid format")

	if len(ve.Errors) != 2 {
		t.Errorf("Add() didn't add errors, got %d errors", len(ve.Errors))
	}
}

func TestValidationErrors_HasErrors(t *testing.T) {
	ve := &ValidationErrors{}
	if ve.HasErrors() {
		t.Error("HasErrors() = true for empty ValidationErrors")
	}

	ve.Add("field", "message")
	if !ve.HasErrors() {
		t.Error("HasErrors() = false after adding error")
	}
}

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name        string
		userName    string
		email       string
		age         int
		wantErr     bool
		wantErrCount int
	}{
		{"valid", "John", "john@example.com", 30, false, 0},
		{"all invalid", "", "bademail", 200, true, 3},
		{"invalid name", "", "john@example.com", 30, true, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUser(tt.userName, tt.email, tt.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				ve, ok := err.(*ValidationErrors)
				if !ok {
					t.Errorf("ValidateUser() error is not *ValidationErrors")
					return
				}
				if len(ve.Errors) != tt.wantErrCount {
					t.Errorf("ValidateUser() error count = %d, want %d", len(ve.Errors), tt.wantErrCount)
				}
			}
		})
	}
}
