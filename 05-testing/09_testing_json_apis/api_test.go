package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TODO(human): Test GetUserHandler
// 1. Create request: httptest.NewRequest("GET", "/user/123", nil)
// 2. Call handler
// 3. Check status is 200
// 4. Parse JSON response into User struct
// 5. Verify ID matches request
// 6. Check Content-Type header

// TODO(human): Test GetUserHandler with invalid ID
// Use path like "/user/abc"
// Should return 400 Bad Request

// TODO(human): Test CreateUserHandler with valid JSON
// 1. Create JSON body: strings.NewReader(`{"name":"Alice","email":"alice@example.com"}`)
// 2. Create POST request with body
// 3. Call handler
// 4. Check status is 201 Created
// 5. Parse response JSON
// 6. Verify user has ID assigned and name/email match

// TODO(human): Test CreateUserHandler with empty name
// Should return 400 Bad Request with error message

// TODO(human): Test CreateUserHandler with invalid JSON
// Body: strings.NewReader(`{invalid json}`)
// Should return 400 Bad Request
