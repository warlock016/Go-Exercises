package error_middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoveryMiddleware(t *testing.T) {
	panicHandler := func(w http.ResponseWriter, r *http.Request) {
		panic("something went wrong")
	}

	wrapped := RecoveryMiddleware(panicHandler)
	if wrapped == nil {
		t.Skip("RecoveryMiddleware not implemented")
	}

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Should not panic
	wrapped(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("RecoveryMiddleware status = %d, want 500", w.Code)
	}
}

func TestCombinedMiddleware(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	wrapped := CombinedMiddleware(handler)
	if wrapped == nil {
		t.Skip("CombinedMiddleware not implemented")
	}

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	wrapped(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("CombinedMiddleware status = %d, want 200", w.Code)
	}
}
