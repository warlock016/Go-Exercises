package status_codes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestResourceHandler(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		url          string
		wantStatus   int
		wantResource *Resource
		wantError    string
	}{
		{
			name:         "GET existing resource",
			method:       http.MethodGet,
			url:          "/resource?id=5",
			wantStatus:   http.StatusOK,
			wantResource: &Resource{ID: 5, Data: "resource 5"},
		},
		{
			name:       "GET non-existing resource",
			method:     http.MethodGet,
			url:        "/resource?id=0",
			wantStatus: http.StatusNotFound,
			wantError:  "resource not found",
		},
		{
			name:       "GET missing id",
			method:     http.MethodGet,
			url:        "/resource",
			wantStatus: http.StatusNotFound,
			wantError:  "resource not found",
		},
		{
			name:         "POST creates resource",
			method:       http.MethodPost,
			url:          "/resource",
			wantStatus:   http.StatusCreated,
			wantResource: &Resource{ID: 42, Data: "created"},
		},
		{
			name:       "DELETE resource",
			method:     http.MethodDelete,
			url:        "/resource?id=5",
			wantStatus: http.StatusNoContent,
		},
		{
			name:         "PUT existing resource",
			method:       http.MethodPut,
			url:          "/resource?id=5",
			wantStatus:   http.StatusOK,
			wantResource: &Resource{ID: 5, Data: "updated 5"},
		},
		{
			name:       "PUT non-existing resource",
			method:     http.MethodPut,
			url:        "/resource?id=0",
			wantStatus: http.StatusNotFound,
			wantError:  "resource not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()

			ResourceHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("ResourceHandler() status = %d, want %d", w.Code, tt.wantStatus)
			}

			if tt.wantResource != nil {
				var got Resource
				if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				if got != *tt.wantResource {
					t.Errorf("ResourceHandler() = %+v, want %+v", got, *tt.wantResource)
				}
			}

			if tt.wantError != "" {
				gotError := strings.TrimSpace(w.Body.String())
				if !strings.Contains(gotError, tt.wantError) {
					t.Errorf("ResourceHandler() error = %q, want to contain %q", gotError, tt.wantError)
				}
			}

			// 204 should have empty body
			if tt.wantStatus == http.StatusNoContent {
				if w.Body.Len() > 0 {
					t.Errorf("ResourceHandler() DELETE should have empty body, got %q", w.Body.String())
				}
			}
		})
	}
}

func TestHealthHandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		wantStatus int
		wantHealth string
	}{
		{
			name:       "healthy service",
			url:        "/health?healthy=true",
			wantStatus: http.StatusOK,
			wantHealth: "healthy",
		},
		{
			name:       "unhealthy service",
			url:        "/health?healthy=false",
			wantStatus: http.StatusServiceUnavailable,
			wantHealth: "unhealthy",
		},
		{
			name:       "missing healthy param defaults to unhealthy",
			url:        "/health",
			wantStatus: http.StatusServiceUnavailable,
			wantHealth: "unhealthy",
		},
		{
			name:       "invalid healthy value defaults to unhealthy",
			url:        "/health?healthy=maybe",
			wantStatus: http.StatusServiceUnavailable,
			wantHealth: "unhealthy",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()

			HealthHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("HealthHandler() status = %d, want %d", w.Code, tt.wantStatus)
			}

			var got HealthStatus
			if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if got.Status != tt.wantHealth {
				t.Errorf("HealthHandler() status = %q, want %q", got.Status, tt.wantHealth)
			}
		})
	}
}

func TestRedirectHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/old-location", nil)
	w := httptest.NewRecorder()

	RedirectHandler(w, req)

	if w.Code != http.StatusMovedPermanently {
		t.Errorf("RedirectHandler() status = %d, want %d", w.Code, http.StatusMovedPermanently)
	}

	location := w.Header().Get("Location")
	if location != "/new-location" {
		t.Errorf("RedirectHandler() Location = %q, want %q", location, "/new-location")
	}
}
