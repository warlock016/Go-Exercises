package api_error_format

import (
	"errors"
	"testing"
)

func TestNewProblemDetail(t *testing.T) {
	err := errors.New("test error")
	pd := NewProblemDetail(err, "/api/users/123")

	if pd.Instance == "" {
		t.Error("ProblemDetail.Instance is empty")
	}
	if pd.Status == 0 {
		t.Error("ProblemDetail.Status is 0")
	}
}

func TestValidationProblemDetail(t *testing.T) {
	errs := []string{"name is required", "email is invalid"}
	pd := ValidationProblemDetail(errs, "/api/users")

	if pd.Status != 400 {
		t.Errorf("ValidationProblemDetail.Status = %d, want 400", pd.Status)
	}
}
