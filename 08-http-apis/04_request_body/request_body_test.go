package request_body

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	tests := []struct {
		name         string
		body         string
		wantStatus   int
		wantResponse *CreateUserResponse
		wantError    string
	}{
		{
			name:       "valid user creation",
			body:       `{"name":"Alice","email":"alice@example.com"}`,
			wantStatus: http.StatusCreated,
			wantResponse: &CreateUserResponse{
				ID:      42,
				Name:    "Alice",
				Email:   "alice@example.com",
				Message: "User created successfully",
			},
		},
		{
			name:       "valid user with unicode",
			body:       `{"name":"世界","email":"world@example.com"}`,
			wantStatus: http.StatusCreated,
			wantResponse: &CreateUserResponse{
				ID:      42,
				Name:    "世界",
				Email:   "world@example.com",
				Message: "User created successfully",
			},
		},
		{
			name:       "malformed JSON",
			body:       `{"name":"Alice","email":}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid JSON",
		},
		{
			name:       "empty JSON object",
			body:       `{}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "name and email are required",
		},
		{
			name:       "missing name",
			body:       `{"email":"alice@example.com"}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "name is required",
		},
		{
			name:       "missing email",
			body:       `{"name":"Alice"}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "email is required",
		},
		{
			name:       "empty name",
			body:       `{"name":"","email":"alice@example.com"}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "name is required",
		},
		{
			name:       "empty email",
			body:       `{"name":"Alice","email":""}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "email is required",
		},
		{
			name:       "both empty",
			body:       `{"name":"","email":""}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "name and email are required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			CreateUserHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("CreateUserHandler() status = %d, want %d", w.Code, tt.wantStatus)
			}

			if tt.wantResponse != nil {
				var got CreateUserResponse
				if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				if got != *tt.wantResponse {
					t.Errorf("CreateUserHandler() = %+v, want %+v", got, *tt.wantResponse)
				}

				contentType := w.Header().Get("Content-Type")
				if contentType != "application/json" {
					t.Errorf("Content-Type = %q, want %q", contentType, "application/json")
				}
			}

			if tt.wantError != "" {
				gotError := strings.TrimSpace(w.Body.String())
				if !strings.Contains(gotError, tt.wantError) {
					t.Errorf("CreateUserHandler() error = %q, want to contain %q", gotError, tt.wantError)
				}
			}
		})
	}
}
