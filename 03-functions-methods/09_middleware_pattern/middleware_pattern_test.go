package middleware_pattern

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLoggingMiddleware(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(log.Writer())

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}

	wrapped := LoggingMiddleware(handler)
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	wrapped(w, req)

	// Check handler was called
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if string(body) != "OK" {
		t.Error("Handler was not called correctly")
	}

	// Check logging occurred
	logOutput := buf.String()
	if !strings.Contains(logOutput, "GET") || !strings.Contains(logOutput, "/test") {
		t.Errorf("Expected log to contain 'GET /test', got: %s", logOutput)
	}
}

func TestTimingMiddleware(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(log.Writer())

	handler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("OK"))
	}

	wrapped := TimingMiddleware(handler)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	wrapped(w, req)

	// Check handler was called
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if string(body) != "OK" {
		t.Error("Handler was not called correctly")
	}

	// Check timing was logged
	logOutput := buf.String()
	if !strings.Contains(logOutput, "took") {
		t.Errorf("Expected timing log, got: %s", logOutput)
	}
}

func TestAuthMiddleware(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("authenticated"))
	}

	authMiddleware := AuthMiddleware("secret-token")
	wrapped := authMiddleware(handler)

	tests := []struct {
		name           string
		authHeader     string
		wantStatusCode int
		wantBody       string
	}{
		{"valid token", "Bearer secret-token", 200, "authenticated"},
		{"invalid token", "Bearer wrong-token", 401, ""},
		{"no auth header", "", 401, ""},
		{"malformed header", "secret-token", 401, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}
			w := httptest.NewRecorder()

			wrapped(w, req)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			if resp.StatusCode != tt.wantStatusCode {
				t.Errorf("Status = %d, want %d", resp.StatusCode, tt.wantStatusCode)
			}

			if string(body) != tt.wantBody {
				t.Errorf("Body = %q, want %q", string(body), tt.wantBody)
			}
		})
	}
}

func TestChain(t *testing.T) {
	// Track execution order
	var order []string

	middleware1 := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			order = append(order, "m1-before")
			next(w, r)
			order = append(order, "m1-after")
		}
	}

	middleware2 := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			order = append(order, "m2-before")
			next(w, r)
			order = append(order, "m2-after")
		}
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		order = append(order, "handler")
		w.Write([]byte("OK"))
	}

	chained := Chain(handler, middleware1, middleware2)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	chained(w, req)

	// Verify execution order: m1 wraps m2 wraps handler
	expected := []string{"m1-before", "m2-before", "handler", "m2-after", "m1-after"}
	if len(order) != len(expected) {
		t.Fatalf("Expected %d events, got %d: %v", len(expected), len(order), order)
	}

	for i, want := range expected {
		if order[i] != want {
			t.Errorf("Event %d: got %q, want %q", i, order[i], want)
		}
	}
}

func TestChainWithAuth(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	}

	auth := AuthMiddleware("token123")
	chained := Chain(handler, LoggingMiddleware, auth)

	// Test unauthorized
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	chained(w, req)

	if w.Result().StatusCode != 401 {
		t.Errorf("Unauthorized request status = %d, want 401", w.Result().StatusCode)
	}

	// Test authorized
	req2 := httptest.NewRequest("GET", "/", nil)
	req2.Header.Set("Authorization", "Bearer token123")
	w2 := httptest.NewRecorder()
	chained(w2, req2)

	resp := w2.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("Authorized request status = %d, want 200", resp.StatusCode)
	}

	if string(body) != "success" {
		t.Errorf("Body = %q, want %q", string(body), "success")
	}
}

func TestChainEmpty(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}

	// Chain with no middleware should return handler unchanged
	chained := Chain(handler)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	chained(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if string(body) != "OK" {
		t.Error("Handler was not called correctly")
	}
}
