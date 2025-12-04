package http_middleware_chain

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRecoveryMiddleware(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(log.Writer())

	panicHandler := func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	}

	wrapped := RecoveryMiddleware(panicHandler)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Should not panic
	wrapped(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status = %d, want 500", resp.StatusCode)
	}

	logOutput := buf.String()
	if !strings.Contains(logOutput, "Panic") || !strings.Contains(logOutput, "test panic") {
		t.Errorf("Expected panic to be logged, got: %s", logOutput)
	}
}

func TestRecoveryMiddlewareNoPanic(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}

	wrapped := RecoveryMiddleware(handler)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	wrapped(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status = %d, want 200", resp.StatusCode)
	}

	if string(body) != "success" {
		t.Errorf("Body = %q, want %q", string(body), "success")
	}
}

func TestCORSMiddleware(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}

	tests := []struct {
		name           string
		allowedOrigins []string
		requestOrigin  string
		wantHeader     string
	}{
		{"wildcard", []string{"*"}, "https://example.com", "*"},
		{"specific origin", []string{"https://example.com"}, "https://example.com", "https://example.com"},
		{"multiple origins match", []string{"https://a.com", "https://b.com"}, "https://b.com", "https://b.com"},
		{"origin not allowed", []string{"https://a.com"}, "https://b.com", ""},
		{"no origin header", []string{"*"}, "", "*"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			corsMiddleware := CORSMiddleware(tt.allowedOrigins)
			wrapped := corsMiddleware(handler)

			req := httptest.NewRequest("GET", "/", nil)
			if tt.requestOrigin != "" {
				req.Header.Set("Origin", tt.requestOrigin)
			}
			w := httptest.NewRecorder()

			wrapped(w, req)

			resp := w.Result()
			corsHeader := resp.Header.Get("Access-Control-Allow-Origin")

			if corsHeader != tt.wantHeader {
				t.Errorf("CORS header = %q, want %q", corsHeader, tt.wantHeader)
			}
		})
	}
}

func TestRequestIDMiddleware(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Handler can access request ID from context if needed
		w.Write([]byte("OK"))
	}

	wrapped := RequestIDMiddleware(handler)

	// Make multiple requests
	for i := 0; i < 3; i++ {
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		wrapped(w, req)

		resp := w.Result()
		requestID := resp.Header.Get("X-Request-ID")

		if requestID == "" {
			t.Error("Request ID header not set")
		}

		if !strings.HasPrefix(requestID, "req-") {
			t.Errorf("Request ID format incorrect: %q", requestID)
		}
	}
}

func TestBuildMiddlewareStack(t *testing.T) {
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

	stack := BuildMiddlewareStack(handler, middleware1, middleware2)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	stack(w, req)

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

func TestCompleteStack(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(log.Writer())

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	}

	stack := BuildMiddlewareStack(
		handler,
		RecoveryMiddleware,
		LoggingMiddleware,
		TimingMiddleware,
		RequestIDMiddleware,
	)

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	stack(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status = %d, want 200", resp.StatusCode)
	}

	if string(body) != "success" {
		t.Errorf("Body = %q, want %q", string(body), "success")
	}

	// Check request ID header
	if resp.Header.Get("X-Request-ID") == "" {
		t.Error("Request ID not set")
	}

	// Check logging
	logOutput := buf.String()
	if !strings.Contains(logOutput, "GET") || !strings.Contains(logOutput, "/test") {
		t.Error("Request not logged correctly")
	}
}

func TestStackWithRecovery(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(log.Writer())

	panicHandler := func(w http.ResponseWriter, r *http.Request) {
		panic("intentional panic")
	}

	stack := BuildMiddlewareStack(
		panicHandler,
		RecoveryMiddleware,
		LoggingMiddleware,
		RequestIDMiddleware,
	)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Should not panic
	stack(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status = %d, want 500", resp.StatusCode)
	}

	// Request ID should still be set (middleware before panic)
	if resp.Header.Get("X-Request-ID") == "" {
		t.Error("Request ID not set despite panic")
	}

	logOutput := buf.String()
	if !strings.Contains(logOutput, "Panic") {
		t.Error("Panic not logged")
	}
}

func TestStackWithCORS(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}

	cors := CORSMiddleware([]string{"https://example.com"})
	stack := BuildMiddlewareStack(
		handler,
		RecoveryMiddleware,
		cors,
	)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Origin", "https://example.com")
	w := httptest.NewRecorder()

	stack(w, req)

	resp := w.Result()

	if resp.Header.Get("Access-Control-Allow-Origin") != "https://example.com" {
		t.Error("CORS header not set")
	}
}
