# Module 08: HTTP APIs

**Focus:** Building HTTP handlers, REST API patterns, request/response handling, testing with httptest
**Prerequisites:** Basic Go syntax, structs, JSON encoding, error handling
**Estimated Time:** 8-12 hours

## Overview

This module teaches you to build HTTP APIs in Go using the standard library. You'll learn handlers, routing patterns, JSON responses, error handling, middleware, and HTTP client testing - all without external frameworks.

**Key Learning Goals:**
- Master http.ResponseWriter and http.Request
- Understand HTTP status codes and when to use them
- Parse query parameters, JSON bodies, and path variables
- Build structured error responses
- Implement method-based routing
- Write middleware for cross-cutting concerns
- Test handlers using httptest package

## Important: All Exercises Are Offline

All exercises use the `net/http/httptest` package. You'll never make real network calls - everything runs in-memory with mock requests and responses. This makes tests fast, reliable, and independent of external services.

## Module Structure

### Tier 1: Introduction (Exercises 1-3)
Basic handlers, query parameters, and JSON responses.

**01_hello_handler** - Basic http.HandlerFunc patterns
**02_query_params** - URL query string parsing
**03_json_response** - JSON encoding and Content-Type headers

### Tier 2: Application (Exercises 4-7)
Request body parsing, status codes, error handling, and method routing.

**04_request_body** - JSON request body parsing
**05_status_codes** - Semantic HTTP status codes
**06_error_responses** - Structured error responses
**07_method_routing** - HTTP method-based routing

### Tier 3: Integration (Exercises 8-12)
Advanced patterns including path parameters, middleware, clients, and context.

**08_path_parameters** - URL path parameter extraction
**09_middleware_basics** - Handler wrapping and chaining
**10_http_client** - Making HTTP requests with testing
**11_handler_testing** - Comprehensive handler testing patterns
**12_context_timeout** - Request context and timeouts

## Testing Commands

### Run all module tests
```bash
cd 08-http-apis
go test ./...
```

### Run single exercise
```bash
cd 08-http-apis/01_hello_handler
go test -v
```

### Run with coverage
```bash
go test ./... -cover
```

### Run specific test
```bash
go test -v -run TestHelloHandler
```

## Key Concepts

### HTTP Request/Response Cycle
Every HTTP handler receives a Request and writes to a ResponseWriter:

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // 1. Read request (method, path, query, headers, body)
    // 2. Process business logic
    // 3. Write response (headers, status, body)
}
```

### Status Codes Guide
- **2xx Success**: 200 OK, 201 Created, 204 No Content
- **3xx Redirect**: 301 Moved Permanently, 302 Found
- **4xx Client Error**: 400 Bad Request, 401 Unauthorized, 403 Forbidden, 404 Not Found, 422 Unprocessable Entity
- **5xx Server Error**: 500 Internal Server Error, 503 Service Unavailable

### httptest Package
Testing without real servers:

```go
// Create mock request
req := httptest.NewRequest(http.MethodGet, "/path?id=123", nil)

// Create mock response recorder
w := httptest.NewRecorder()

// Call handler
MyHandler(w, req)

// Assert response
if w.Code != http.StatusOK {
    t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
}
```

## Resources

Before starting, review these guides:

- **[HTTP_PROTOCOL_GUIDE.md](../resources/HTTP_PROTOCOL_GUIDE.md)** - Complete HTTP protocol reference
- **[HTTP_ERROR_ARCHITECTURE.md](../resources/HTTP_ERROR_ARCHITECTURE.md)** - Error handling patterns
- [Go net/http Package](https://pkg.go.dev/net/http)
- [Go net/http/httptest Package](https://pkg.go.dev/net/http/httptest)

## Progress Tracking

Mark exercises as complete in your PROGRESS.md file as you finish them.

- [ ] 01_hello_handler
- [ ] 02_query_params
- [ ] 03_json_response
- [ ] 04_request_body
- [ ] 05_status_codes
- [ ] 06_error_responses
- [ ] 07_method_routing
- [ ] 08_path_parameters
- [ ] 09_middleware_basics
- [ ] 10_http_client
- [ ] 11_handler_testing
- [ ] 12_context_timeout

## What You'll Build

By the end of this module, you'll have built handlers for:
- Simple text and JSON responses
- Query parameter parsing and validation
- JSON request body handling
- RESTful resource management with all HTTP methods
- Structured error responses with appropriate status codes
- Middleware for logging, timing, and request modification
- HTTP clients with proper error handling
- Context-aware handlers that respect timeouts

These are the foundational skills for building production-ready HTTP APIs in Go.
