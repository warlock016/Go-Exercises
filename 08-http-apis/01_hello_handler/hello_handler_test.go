package hello_handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name       string
		wantBody   string
		wantStatus int
	}{
		{
			name:       "returns hello world",
			wantBody:   "Hello, World!",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/hello", nil)
			w := httptest.NewRecorder()

			HelloHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("HelloHandler() status = %d, want %d", w.Code, tt.wantStatus)
			}

			gotBody := w.Body.String()
			if gotBody != tt.wantBody {
				t.Errorf("HelloHandler() body = %q, want %q", gotBody, tt.wantBody)
			}
		})
	}
}

func TestGreetHandler(t *testing.T) {
	tests := []struct {
		name       string
		queryParam string
		wantBody   string
		wantStatus int
	}{
		{
			name:       "greets by name",
			queryParam: "?name=Alice",
			wantBody:   "Hello, Alice!",
			wantStatus: http.StatusOK,
		},
		{
			name:       "greets stranger when no name",
			queryParam: "",
			wantBody:   "Hello, stranger!",
			wantStatus: http.StatusOK,
		},
		{
			name:       "greets Bob",
			queryParam: "?name=Bob",
			wantBody:   "Hello, Bob!",
			wantStatus: http.StatusOK,
		},
		{
			name:       "greets with unicode name",
			queryParam: "?name=世界",
			wantBody:   "Hello, 世界!",
			wantStatus: http.StatusOK,
		},
		{
			name:       "greets empty string as stranger",
			queryParam: "?name=",
			wantBody:   "Hello, stranger!",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/greet"+tt.queryParam, nil)
			w := httptest.NewRecorder()

			GreetHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("GreetHandler() status = %d, want %d", w.Code, tt.wantStatus)
			}

			gotBody := w.Body.String()
			if gotBody != tt.wantBody {
				t.Errorf("GreetHandler() body = %q, want %q", gotBody, tt.wantBody)
			}
		})
	}
}
