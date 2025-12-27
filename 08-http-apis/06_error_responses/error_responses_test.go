package error_responses

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()

	WriteError(w, http.StatusBadRequest, "TEST_ERROR", "test message")

	if w.Code != http.StatusBadRequest {
		t.Errorf("WriteError() status = %d, want %d", w.Code, http.StatusBadRequest)
	}

	var got APIError
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if got.Code != "TEST_ERROR" {
		t.Errorf("WriteError() code = %q, want %q", got.Code, "TEST_ERROR")
	}
	if got.Message != "test message" {
		t.Errorf("WriteError() message = %q, want %q", got.Message, "test message")
	}
	if got.Details != nil {
		t.Errorf("WriteError() details should be omitted, got %v", got.Details)
	}
}

func TestWriteErrorWithDetails(t *testing.T) {
	w := httptest.NewRecorder()
	details := map[string]string{
		"field": "age",
		"value": "200",
	}

	WriteErrorWithDetails(w, http.StatusUnprocessableEntity, "VALIDATION_ERROR", "invalid age", details)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("WriteErrorWithDetails() status = %d, want %d", w.Code, http.StatusUnprocessableEntity)
	}

	var got APIError
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if got.Code != "VALIDATION_ERROR" {
		t.Errorf("WriteErrorWithDetails() code = %q, want %q", got.Code, "VALIDATION_ERROR")
	}
	if got.Details["field"] != "age" {
		t.Errorf("WriteErrorWithDetails() details[field] = %q, want %q", got.Details["field"], "age")
	}
}

func TestValidationHandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		wantStatus int
		wantCode   string
	}{
		{
			name:       "valid input",
			url:        "/validate?email=alice@example.com&age=25",
			wantStatus: http.StatusOK,
		},
		{
			name:       "missing email",
			url:        "/validate?age=25",
			wantStatus: http.StatusBadRequest,
			wantCode:   "MISSING_PARAM",
		},
		{
			name:       "empty email",
			url:        "/validate?email=&age=25",
			wantStatus: http.StatusBadRequest,
			wantCode:   "MISSING_PARAM",
		},
		{
			name:       "age too high",
			url:        "/validate?email=alice@example.com&age=200",
			wantStatus: http.StatusUnprocessableEntity,
			wantCode:   "VALIDATION_ERROR",
		},
		{
			name:       "negative age",
			url:        "/validate?email=alice@example.com&age=-5",
			wantStatus: http.StatusUnprocessableEntity,
			wantCode:   "VALIDATION_ERROR",
		},
		{
			name:       "age exactly 0",
			url:        "/validate?email=alice@example.com&age=0",
			wantStatus: http.StatusOK,
		},
		{
			name:       "age exactly 150",
			url:        "/validate?email=alice@example.com&age=150",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()

			ValidationHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("ValidationHandler() status = %d, want %d", w.Code, tt.wantStatus)
			}

			if tt.wantCode != "" {
				var got APIError
				if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				if got.Code != tt.wantCode {
					t.Errorf("ValidationHandler() code = %q, want %q", got.Code, tt.wantCode)
				}
			}
		})
	}
}
