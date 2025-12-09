package api_error_format

import (
	"errors"
	"strings"
	"testing"
)

func TestNewProblemDetail_AllFieldsPopulated(t *testing.T) {
	err := errors.New("user not found in database")
	pd := NewProblemDetail(err, "/api/users/123")

	// All RFC 7807 fields should be populated
	if pd.Type == "" {
		t.Error("ProblemDetail.Type is empty - should be a URI identifying the error type")
	}
	if pd.Title == "" {
		t.Error("ProblemDetail.Title is empty - should be a short summary")
	}
	if pd.Status == 0 {
		t.Error("ProblemDetail.Status is 0 - should be an HTTP status code")
	}
	if pd.Detail == "" {
		t.Error("ProblemDetail.Detail is empty - should contain the error message")
	}
	if pd.Instance == "" {
		t.Error("ProblemDetail.Instance is empty - should be the request path")
	}
}

func TestNewProblemDetail_InstanceMatchesRequestPath(t *testing.T) {
	tests := []struct {
		name        string
		requestPath string
	}{
		{"root path", "/"},
		{"api endpoint", "/api/users/123"},
		{"nested path", "/api/v1/orders/456/items"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := NewProblemDetail(errors.New("error"), tt.requestPath)
			if pd.Instance != tt.requestPath {
				t.Errorf("Instance = %q, want %q", pd.Instance, tt.requestPath)
			}
		})
	}
}

func TestNewProblemDetail_DetailContainsErrorMessage(t *testing.T) {
	errMsg := "user with ID 999 was not found"
	err := errors.New(errMsg)
	pd := NewProblemDetail(err, "/api/users/999")

	if !strings.Contains(pd.Detail, errMsg) {
		t.Errorf("Detail = %q, should contain error message %q", pd.Detail, errMsg)
	}
}

func TestNewProblemDetail_StatusCodesForKnownErrors(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		wantStatus int
	}{
		{"not found error", ErrNotFound, 404},
		{"unauthorized error", ErrUnauthorized, 401},
		{"invalid input error", ErrInvalidInput, 400},
		{"unknown error", errors.New("something broke"), 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := NewProblemDetail(tt.err, "/api/test")
			if pd.Status != tt.wantStatus {
				t.Errorf("Status = %d, want %d", pd.Status, tt.wantStatus)
			}
		})
	}
}

func TestNewProblemDetail_TypeIsValidURI(t *testing.T) {
	pd := NewProblemDetail(ErrNotFound, "/api/users/123")

	// Type should be a URI - either "about:blank" or a URL
	if pd.Type != "about:blank" && !strings.HasPrefix(pd.Type, "http") {
		t.Errorf("Type = %q, should be 'about:blank' or start with 'http'", pd.Type)
	}
}

func TestValidationProblemDetail_Status400(t *testing.T) {
	errs := []string{"name is required"}
	pd := ValidationProblemDetail(errs, "/api/users")

	if pd.Status != 400 {
		t.Errorf("Status = %d, want 400", pd.Status)
	}
}

func TestValidationProblemDetail_AllFieldsPopulated(t *testing.T) {
	errs := []string{"name is required", "email is invalid"}
	pd := ValidationProblemDetail(errs, "/api/users")

	if pd.Type == "" {
		t.Error("Type is empty")
	}
	if pd.Title == "" {
		t.Error("Title is empty")
	}
	if pd.Detail == "" {
		t.Error("Detail is empty")
	}
	if pd.Instance == "" {
		t.Error("Instance is empty")
	}
}

func TestValidationProblemDetail_DetailContainsAllErrors(t *testing.T) {
	errs := []string{"name is required", "email is invalid", "age must be positive"}
	pd := ValidationProblemDetail(errs, "/api/users")

	for _, errMsg := range errs {
		if !strings.Contains(pd.Detail, errMsg) {
			t.Errorf("Detail = %q, should contain %q", pd.Detail, errMsg)
		}
	}
}

func TestValidationProblemDetail_InstanceMatchesPath(t *testing.T) {
	pd := ValidationProblemDetail([]string{"error"}, "/api/products/new")

	if pd.Instance != "/api/products/new" {
		t.Errorf("Instance = %q, want %q", pd.Instance, "/api/products/new")
	}
}

func TestValidationProblemDetail_EmptyErrors(t *testing.T) {
	// Edge case: what happens with no errors?
	pd := ValidationProblemDetail([]string{}, "/api/users")

	// Should still return a valid structure
	if pd.Status != 400 {
		t.Errorf("Status = %d, want 400", pd.Status)
	}
}

func TestValidationProblemDetail_TitleIndicatesValidation(t *testing.T) {
	pd := ValidationProblemDetail([]string{"error"}, "/api/users")

	// Title should indicate this is a validation error
	titleLower := strings.ToLower(pd.Title)
	if !strings.Contains(titleLower, "validation") {
		t.Errorf("Title = %q, should contain 'validation'", pd.Title)
	}
}
