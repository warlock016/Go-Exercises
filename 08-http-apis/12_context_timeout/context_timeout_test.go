package context_timeout

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSlowOperation(t *testing.T) {
	tests := []struct {
		name        string
		duration    time.Duration
		timeout     time.Duration
		expectError bool
	}{
		{
			name:        "completes before timeout",
			duration:    10 * time.Millisecond,
			timeout:     100 * time.Millisecond,
			expectError: false,
		},
		{
			name:        "cancelled by timeout",
			duration:    100 * time.Millisecond,
			timeout:     10 * time.Millisecond,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), tt.timeout)
			defer cancel()

			err := SlowOperation(ctx, tt.duration)

			if tt.expectError && err == nil {
				t.Errorf("SlowOperation() expected error but got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("SlowOperation() unexpected error: %v", err)
			}
		})
	}
}

func TestSlowHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/slow?duration=10ms", nil)
	w := httptest.NewRecorder()

	SlowHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("SlowHandler() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestTimeoutMiddleware(t *testing.T) {
	slowHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.Write([]byte("OK"))
	})

	handler := TimeoutMiddleware(10 * time.Millisecond)(slowHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusGatewayTimeout {
		t.Errorf("TimeoutMiddleware() status = %d, want %d", w.Code, http.StatusGatewayTimeout)
	}
}

func TestRequestIDMiddleware(t *testing.T) {
	var capturedID string

	handler := RequestIDMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("requestID")
		if id != nil {
			capturedID = id.(string)
		}
		w.Write([]byte("OK"))
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if capturedID == "" {
		t.Errorf("RequestIDMiddleware() did not set requestID in context")
	}
}
