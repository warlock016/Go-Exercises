package custom_error_types

import (
	"strings"
	"testing"
	"time"
)

func TestValidationError(t *testing.T) {
	err := &ValidationError{
		Field:   "username",
		Value:   "ab",
		Message: "must be at least 3 characters",
	}

	errStr := err.Error()
	if !strings.Contains(errStr, "username") || !strings.Contains(errStr, "ab") {
		t.Errorf("ValidationError.Error() = %q, want to contain field and value", errStr)
	}
}

func TestHTTPError(t *testing.T) {
	err := &HTTPError{
		StatusCode: 404,
		Message:    "Not Found",
		URL:        "https://example.com/api",
	}

	errStr := err.Error()
	if !strings.Contains(errStr, "404") || !strings.Contains(errStr, "example.com") {
		t.Errorf("HTTPError.Error() = %q, want to contain status code and URL", errStr)
	}
}

func TestTimeoutError(t *testing.T) {
	err := &TimeoutError{
		Operation: "database-query",
		Duration:  30 * time.Second,
	}

	errStr := err.Error()
	if !strings.Contains(errStr, "database-query") || !strings.Contains(errStr, "30") {
		t.Errorf("TimeoutError.Error() = %q, want to contain operation and duration", errStr)
	}
}

func TestValidateUserInput(t *testing.T) {
	tests := []struct {
		name      string
		userName  string
		email     string
		age       int
		wantErr   bool
		wantField string
	}{
		{"valid input", "John", "john@example.com", 30, false, ""},
		{"empty name", "", "test@example.com", 25, true, "name"},
		{"invalid email", "John", "notanemail", 25, true, "email"},
		{"negative age", "John", "john@example.com", -1, true, "age"},
		{"age too high", "John", "john@example.com", 200, true, "age"},
		{"email without @", "John", "john.example.com", 25, true, "email"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUserInput(tt.userName, tt.email, tt.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUserInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				verr, ok := err.(*ValidationError)
				if !ok {
					t.Errorf("ValidateUserInput() error is not *ValidationError: %v", err)
					return
				}
				if verr.Field != tt.wantField {
					t.Errorf("ValidationError.Field = %q, want %q", verr.Field, tt.wantField)
				}
			}
		})
	}
}

func TestFetchResource(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		statusCode int
		wantData   string
		wantErr    bool
	}{
		{"success 200", "https://api.example.com", 200, "success", false},
		{"success 201", "https://api.example.com", 201, "success", false},
		{"client error 400", "https://api.example.com", 400, "", true},
		{"not found 404", "https://api.example.com/users", 404, "", true},
		{"server error 500", "https://api.example.com", 500, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchResource(tt.url, tt.statusCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.wantData {
				t.Errorf("FetchResource() = %q, want %q", got, tt.wantData)
			}
			if tt.wantErr {
				herr, ok := err.(*HTTPError)
				if !ok {
					t.Errorf("FetchResource() error is not *HTTPError: %v", err)
					return
				}
				if herr.StatusCode != tt.statusCode {
					t.Errorf("HTTPError.StatusCode = %d, want %d", herr.StatusCode, tt.statusCode)
				}
				if herr.URL != tt.url {
					t.Errorf("HTTPError.URL = %q, want %q", herr.URL, tt.url)
				}
			}
		})
	}
}

func TestPerformOperation(t *testing.T) {
	tests := []struct {
		name           string
		operation      string
		timeoutSeconds int
		wantErr        bool
	}{
		{"quick operation", "cache-lookup", 5, false},
		{"medium operation", "api-call", 30, false},
		{"timeout operation", "slow-query", 31, true},
		{"long timeout", "batch-process", 120, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PerformOperation(tt.operation, tt.timeoutSeconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("PerformOperation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				terr, ok := err.(*TimeoutError)
				if !ok {
					t.Errorf("PerformOperation() error is not *TimeoutError: %v", err)
					return
				}
				if terr.Operation != tt.operation {
					t.Errorf("TimeoutError.Operation = %q, want %q", terr.Operation, tt.operation)
				}
				expectedDuration := time.Duration(tt.timeoutSeconds) * time.Second
				if terr.Duration != expectedDuration {
					t.Errorf("TimeoutError.Duration = %v, want %v", terr.Duration, expectedDuration)
				}
			}
		})
	}
}

func TestExtractValidationError(t *testing.T) {
	valErr := &ValidationError{Field: "test", Value: 123, Message: "invalid"}
	httpErr := &HTTPError{StatusCode: 404, Message: "Not Found", URL: "test"}

	tests := []struct {
		name  string
		err   error
		want  *ValidationError
		wantOk bool
	}{
		{"validation error", valErr, valErr, true},
		{"http error", httpErr, nil, false},
		{"nil error", nil, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ExtractValidationError(tt.err)
			if ok != tt.wantOk {
				t.Errorf("ExtractValidationError() ok = %v, want %v", ok, tt.wantOk)
			}
			if tt.wantOk && got != tt.want {
				t.Errorf("ExtractValidationError() = %v, want %v", got, tt.want)
			}
		})
	}
}
