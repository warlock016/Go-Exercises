package middleware_basics

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func simpleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}

func TestLoggingMiddleware(t *testing.T) {
	handler := LoggingMiddleware(simpleHandler())

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("LoggingMiddleware() status = %d, want %d", w.Code, http.StatusOK)
	}

	if w.Body.String() != "OK" {
		t.Errorf("LoggingMiddleware() body = %q, want %q", w.Body.String(), "OK")
	}
}

func TestTimingMiddleware(t *testing.T) {
	handler := TimingMiddleware(simpleHandler())

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	header := w.Header().Get("X-Response-Time")
	if header == "" {
		t.Errorf("TimingMiddleware() missing X-Response-Time header")
	}

	// Parse as float to verify it's a valid number
	if _, err := strconv.ParseFloat(header, 64); err != nil {
		t.Errorf("TimingMiddleware() X-Response-Time = %q is not a valid number", header)
	}
}

func TestHeaderMiddleware(t *testing.T) {
	handler := HeaderMiddleware(simpleHandler())

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	header := w.Header().Get("X-Custom-Header")
	if header != "test-value" {
		t.Errorf("HeaderMiddleware() X-Custom-Header = %q, want %q", header, "test-value")
	}
}

func TestChain(t *testing.T) {
	handler := Chain(
		simpleHandler(),
		LoggingMiddleware,
		TimingMiddleware,
		HeaderMiddleware,
	)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	// Check all middleware ran
	if w.Header().Get("X-Response-Time") == "" {
		t.Errorf("Chain() missing X-Response-Time header")
	}

	if w.Header().Get("X-Custom-Header") != "test-value" {
		t.Errorf("Chain() X-Custom-Header = %q, want %q",
			w.Header().Get("X-Custom-Header"), "test-value")
	}

	if w.Body.String() != "OK" {
		t.Errorf("Chain() body = %q, want %q", w.Body.String(), "OK")
	}
}
