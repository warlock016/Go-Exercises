package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO(human): Test AuthMiddleware with valid token
// Set header: req.Header.Set("Authorization", "Bearer valid-token")
// Handler should be called

// TODO(human): Test AuthMiddleware with invalid token
// Should return 401 Unauthorized
// Handler should NOT be called

// TODO(human): Test AuthMiddleware with no token
// Should return 401 Unauthorized

// TODO(human): Test AddHeaderMiddleware
// Verify that X-Custom-Header is set in response
// Verify handler is still called
