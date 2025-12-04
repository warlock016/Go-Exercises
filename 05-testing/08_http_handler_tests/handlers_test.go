package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TODO(human): Test HelloHandler
// Use httptest.NewRequest and httptest.NewRecorder
// func TestHelloHandler(t *testing.T) {
//     req := httptest.NewRequest("GET", "/hello", nil)
//     rec := httptest.NewRecorder()
//
//     HelloHandler(rec, req)
//
//     // Check status code
//     if rec.Code != http.StatusOK {
//         t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
//     }
//
//     // Check body
//     if rec.Body.String() != "Hello, World!" {
//         t.Errorf("body = %q, want %q", rec.Body.String(), "Hello, World!")
//     }
//
//     // Check Content-Type header
//     contentType := rec.Header().Get("Content-Type")
//     if contentType != "text/plain" {
//         t.Errorf("Content-Type = %q, want %q", contentType, "text/plain")
//     }
// }

// TODO(human): Test HelloHandler with wrong method (POST)
// Should return 405 Method Not Allowed

// TODO(human): Test JSONHandler
// Check status code, body contains {"message": "Hello, JSON!"}
// Check Content-Type is "application/json"

// TODO(human): Test EchoHandler
// Create request with body: strings.NewReader("echo this")
// req := httptest.NewRequest("POST", "/echo", strings.NewReader("echo this"))
// Check that response body equals request body

// TODO(human): Test EchoHandler with GET method
// Should return 405 Method Not Allowed
