package http_handler_functions

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	HelloHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if string(body) != "Hello, World!" {
		t.Errorf("HelloHandler() = %q, want %q", string(body), "Hello, World!")
	}
}

func TestEchoHandler(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		wantBody       string
		wantStatusCode int
	}{
		{"with message", "?message=test", "test", 200},
		{"with spaces", "?message=hello%20world", "hello world", 200},
		{"empty message", "?message=", "missing message parameter", 400},
		{"no message", "", "missing message parameter", 400},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/echo"+tt.query, nil)
			w := httptest.NewRecorder()

			EchoHandler(w, req)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			if resp.StatusCode != tt.wantStatusCode {
				t.Errorf("EchoHandler() status = %d, want %d", resp.StatusCode, tt.wantStatusCode)
			}

			if string(body) != tt.wantBody {
				t.Errorf("EchoHandler() = %q, want %q", string(body), tt.wantBody)
			}
		})
	}
}

func TestStatusHandler(t *testing.T) {
	tests := []struct {
		name string
		code int
	}{
		{"OK", 200},
		{"Not Found", 404},
		{"Internal Server Error", 500},
		{"Created", 201},
		{"No Content", 204},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := StatusHandler(tt.code)
			req := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()

			handler(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.code {
				t.Errorf("StatusHandler(%d) status = %d, want %d", tt.code, resp.StatusCode, tt.code)
			}
		})
	}
}

func TestJSONHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	JSONHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	// Check Content-Type
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Content-Type = %q, want %q", contentType, "application/json")
	}

	// Parse JSON
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Check fields
	if msg, ok := data["message"].(string); !ok || msg != "Hello, JSON!" {
		t.Errorf("message = %v, want %q", data["message"], "Hello, JSON!")
	}

	if _, ok := data["timestamp"]; !ok {
		t.Error("timestamp field missing")
	}
}

func TestMethodFilter(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	}

	tests := []struct {
		name           string
		allowedMethod  string
		requestMethod  string
		wantStatusCode int
		wantBody       string
	}{
		{"POST allowed", "POST", "POST", 200, "success"},
		{"GET denied", "POST", "GET", 405, ""},
		{"PUT denied", "POST", "PUT", 405, ""},
		{"GET allowed", "GET", "GET", 200, "success"},
		{"DELETE allowed", "DELETE", "DELETE", 200, "success"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered := MethodFilter(tt.allowedMethod, handler)
			req := httptest.NewRequest(tt.requestMethod, "/", nil)
			w := httptest.NewRecorder()

			filtered(w, req)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			if resp.StatusCode != tt.wantStatusCode {
				t.Errorf("MethodFilter status = %d, want %d", resp.StatusCode, tt.wantStatusCode)
			}

			if string(body) != tt.wantBody {
				t.Errorf("MethodFilter body = %q, want %q", string(body), tt.wantBody)
			}
		})
	}
}

func TestMethodFilterPreservesHandler(t *testing.T) {
	// Test that MethodFilter doesn't modify the original handler behavior
	called := false
	handler := func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("custom response"))
	}

	filtered := MethodFilter("POST", handler)
	req := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()

	filtered(w, req)

	if !called {
		t.Error("Handler was not called")
	}

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 201 {
		t.Errorf("Status = %d, want 201", resp.StatusCode)
	}

	if string(body) != "custom response" {
		t.Errorf("Body = %q, want %q", string(body), "custom response")
	}
}
