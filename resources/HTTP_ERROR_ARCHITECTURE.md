# HTTP Error Architecture Blueprint

A unified mental model for understanding when and how to handle errors in HTTP applications.

> **Prerequisites:** Familiarity with [ERROR_HANDLING_PATTERNS.md](ERROR_HANDLING_PATTERNS.md) and [HTTP_PROTOCOL_GUIDE.md](HTTP_PROTOCOL_GUIDE.md)

---

## The HTTP Error Lifecycle

Every HTTP request passes through multiple layers. Errors can occur at any layer, and each layer has different error characteristics:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           HTTP REQUEST LIFECYCLE                            │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   CLIENT                                                                    │
│     │                                                                       │
│     ▼                                                                       │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │                    MIDDLEWARE CHAIN                                 │   │
│   │  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐             │   │
│   │  │ Recovery │→ │ Logging  │→ │   Auth   │→ │RateLimit │→ ...        │   │
│   │  │          │  │          │  │          │  │          │             │   │
│   │  │ panic→500│  │ log req  │  │ 401/403  │  │   429    │             │   │
│   │  └──────────┘  └──────────┘  └──────────┘  └──────────┘             │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                    │                                        │
│                                    ▼                                        │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │                         HANDLER LAYER                               │   │
│   │                                                                     │   │
│   │  • Parse request (400 - bad JSON, missing params)                   │   │
│   │  • Validate input (400/422 - validation errors)                     │   │
│   │  • Check method (405 - wrong HTTP method)                           │   │
│   │  • Check content-type (415 - wrong media type)                      │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                    │                                        │
│                                    ▼                                        │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │                        SERVICE LAYER                                │   │
│   │                                                                     │   │
│   │  • Business logic errors (400/409/422 - depends on case)            │   │
│   │  • Resource not found (404)                                         │   │
│   │  • Permission denied (403)                                          │   │
│   │  • Conflict/duplicate (409)                                         │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                    │                                        │
│                                    ▼                                        │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │                      REPOSITORY LAYER                               │   │
│   │                                                                     │   │
│   │  • Database errors (500 - connection, query failure)                │   │
│   │  • Network timeouts (504 - upstream timeout)                        │   │
│   │  • External service down (502/503 - dependency failure)             │   │
│   │  • Temporary failures (503 - retryable)                             │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Key Insight:** Errors bubble UP from lower layers to the handler, which transforms them into HTTP responses.

---

## Error Classification Matrix

| Error Type              | Status | Retryable? | Who's Fault? | Go Pattern       | Response Format    |
|-------------------------|--------|------------|--------------|------------------|--------------------|
| **Bad Request**         | 400    | No         | Client       | Wrapped error    | Simple JSON        |
| **Unauthorized**        | 401    | No*        | Client       | Sentinel         | Simple JSON        |
| **Forbidden**           | 403    | No         | Client       | Sentinel         | Simple JSON        |
| **Not Found**           | 404    | No         | Client       | Sentinel         | Simple JSON        |
| **Method Not Allowed**  | 405    | No         | Client       | N/A (handler)    | Text               |
| **Conflict**            | 409    | Maybe      | Client       | Custom type      | JSON with details  |
| **Validation Failed**   | 422    | No         | Client       | ValidationErrors | RFC 7807           |
| **Rate Limited**        | 429    | Yes        | Client       | Custom type      | JSON + Retry-After |
| **Internal Error**      | 500    | No         | Server       | Catch-all        | Simple JSON        |
| **Bad Gateway**         | 502    | Yes        | Server       | Wrapped error    | Simple JSON        |
| **Service Unavailable** | 503    | Yes        | Server       | Custom type      | JSON + Retry-After |
| **Gateway Timeout**     | 504    | Yes        | Server       | Wrapped error    | Simple JSON        |

*401 might be retryable after getting new credentials

### Retryable vs Non-Retryable

```
┌─────────────────────────────────────────────────────────────────┐
│                    IS THIS ERROR RETRYABLE?                     │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  RETRYABLE (transient - might succeed later)                    │
│  ├── 429 Too Many Requests (wait for rate limit reset)          │
│  ├── 502 Bad Gateway (upstream might recover)                   │
│  ├── 503 Service Unavailable (server overloaded)                │
│  ├── 504 Gateway Timeout (network/upstream slow)                │
│  └── Connection refused, timeouts, temporary network issues     │
│                                                                 │
│  NOT RETRYABLE (permanent - same request will always fail)      │
│  ├── 400 Bad Request (fix your request)                         │
│  ├── 401 Unauthorized (get valid credentials first)             │
│  ├── 403 Forbidden (you don't have permission)                  │
│  ├── 404 Not Found (resource doesn't exist)                     │
│  ├── 405 Method Not Allowed (use correct HTTP method)           │
│  ├── 409 Conflict (resolve the conflict first)                  │
│  ├── 422 Validation Error (fix the input data)                  │
│  └── 500 Internal Error (server bug - retrying won't help)      │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Decision Trees

### Which Status Code Should I Use?

```
START: An error occurred
         │
         ▼
    Is it the client's fault?
         │
    ┌────┴────┐
   YES        NO
    │          │
    ▼          ▼
┌───────┐  Is it a known
│Client │  server issue?
│ Error │      │
└───┬───┘  ┌───┴───┐
    │     YES      NO
    │      │       │
    │      ▼       ▼
    │  ┌───────┐ ┌───────┐
    │  │Server │ │  500  │
    │  │ Error │ │Catch- │
    │  └───┬───┘ │  all  │
    │      │     └───────┘
    ▼      ▼
```

**Client Errors (4xx):**
```
Is authentication missing/invalid?
├── YES → 401 Unauthorized
└── NO  → Is user authenticated but not allowed?
          ├── YES → 403 Forbidden
          └── NO  → Does the resource exist?
                    ├── NO  → 404 Not Found
                    └── YES → Is the request malformed?
                              ├── YES → 400 Bad Request
                              └── NO  → Is validation failing?
                                        ├── YES → 422 Unprocessable Entity
                                        └── NO  → Is there a conflict?
                                                  ├── YES → 409 Conflict
                                                  └── NO  → 400 Bad Request
```

**Server Errors (5xx):**
```
Is an upstream/external service failing?
├── YES → Did it timeout?
│         ├── YES → 504 Gateway Timeout
│         └── NO  → 502 Bad Gateway
└── NO  → Is the server temporarily overloaded?
          ├── YES → 503 Service Unavailable
          └── NO  → 500 Internal Server Error
```

### Which Go Error Pattern Should I Use?

```
START: I need to define an error
         │
         ▼
    Is it a well-known, expected condition?
    (e.g., "not found", "unauthorized")
         │
    ┌────┴────┐
   YES        NO
    │          │
    ▼          ▼
┌─────────┐  Do I need to attach
│Sentinel │  extra information?
│  Error  │      │
│         │  ┌───┴───┐
│var Err  │ YES      NO
│NotFound │  │       │
│= errors │  ▼       ▼
│.New()   │┌──────┐┌──────────┐
└─────────┘│Custom││ Wrapped  │
           │ Type ││  Error   │
           │      ││          │
           │struct││fmt.Errorf│
           │with  ││("ctx: %w"│
           │fields││, err)    │
           └──────┘└──────────┘
```

---

## Error Response Design Patterns

### Pattern 1: Simple Text Error (Quick & Dirty)

**When to use:** Internal APIs, simple endpoints, rapid prototyping

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "missing id parameter", http.StatusBadRequest)
        return
    }
}
```

**Pros:** Simple, built-in, no dependencies
**Cons:** Not machine-parseable, limited information

---

### Pattern 2: Simple JSON Error

**When to use:** REST APIs, when clients need to parse errors

```go
type ErrorResponse struct {
    Error   string `json:"error"`
    Code    string `json:"code,omitempty"`
    Message string `json:"message"`
}

func writeError(w http.ResponseWriter, status int, code, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(ErrorResponse{
        Error:   http.StatusText(status),
        Code:    code,
        Message: message,
    })
}

// Usage
func Handler(w http.ResponseWriter, r *http.Request) {
    user, err := findUser(id)
    if errors.Is(err, ErrNotFound) {
        writeError(w, http.StatusNotFound, "USER_NOT_FOUND",
            fmt.Sprintf("user %s not found", id))
        return
    }
}
```

**Response:**
```json
{
    "error": "Not Found",
    "code": "USER_NOT_FOUND",
    "message": "user 123 not found"
}
```

---

### Pattern 3: RFC 7807 Problem Details

**When to use:** Public APIs, when you want industry-standard error format

```go
type ProblemDetails struct {
    Type     string `json:"type"`
    Title    string `json:"title"`
    Status   int    `json:"status"`
    Detail   string `json:"detail,omitempty"`
    Instance string `json:"instance,omitempty"`
}

func writeProblem(w http.ResponseWriter, r *http.Request, status int, title, detail string) {
    w.Header().Set("Content-Type", "application/problem+json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(ProblemDetails{
        Type:     "about:blank",
        Title:    title,
        Status:   status,
        Detail:   detail,
        Instance: r.URL.Path,
    })
}
```

**Response:**
```json
{
    "type": "about:blank",
    "title": "Not Found",
    "status": 404,
    "detail": "User with ID 123 was not found",
    "instance": "/api/users/123"
}
```

---

### Pattern 4: Validation Errors (Multiple Fields)

**When to use:** Form validation, when multiple things can be wrong

```go
type ValidationError struct {
    Type    string            `json:"type"`
    Title   string            `json:"title"`
    Status  int               `json:"status"`
    Errors  map[string]string `json:"errors"`
}

func writeValidationError(w http.ResponseWriter, errors map[string]string) {
    w.Header().Set("Content-Type", "application/problem+json")
    w.WriteHeader(http.StatusUnprocessableEntity)
    json.NewEncoder(w).Encode(ValidationError{
        Type:   "https://example.com/errors/validation",
        Title:  "Validation Failed",
        Status: 422,
        Errors: errors,
    })
}

// Usage
func Handler(w http.ResponseWriter, r *http.Request) {
    errs := make(map[string]string)

    if input.Email == "" {
        errs["email"] = "email is required"
    }
    if input.Age < 0 {
        errs["age"] = "age must be positive"
    }

    if len(errs) > 0 {
        writeValidationError(w, errs)
        return
    }
}
```

**Response:**
```json
{
    "type": "https://example.com/errors/validation",
    "title": "Validation Failed",
    "status": 422,
    "errors": {
        "email": "email is required",
        "age": "age must be positive"
    }
}
```

---

### Pattern 5: Rate Limit Error (with Retry-After)

**When to use:** Rate limiting, temporary service unavailability

```go
func writeRateLimitError(w http.ResponseWriter, retryAfter int) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Retry-After", strconv.Itoa(retryAfter))
    w.WriteHeader(http.StatusTooManyRequests)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "error":       "rate limit exceeded",
        "retry_after": retryAfter,
    })
}
```

**Response Headers:**
```
HTTP/1.1 429 Too Many Requests
Retry-After: 60
Content-Type: application/json
```

---

## Request Lifecycle Integration

### Where Errors Commonly Occur

```go
func UserHandler(w http.ResponseWriter, r *http.Request) {
    // ─────────────────────────────────────────────────────────
    // HANDLER LAYER: Request parsing & validation
    // ─────────────────────────────────────────────────────────

    // Method check → 405
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Content-Type check → 415
    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
        return
    }

    // Body parsing → 400
    var input CreateUserInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        writeError(w, http.StatusBadRequest, "INVALID_JSON", err.Error())
        return
    }

    // Validation → 422
    if errs := validateInput(input); len(errs) > 0 {
        writeValidationError(w, errs)
        return
    }

    // ─────────────────────────────────────────────────────────
    // SERVICE LAYER: Business logic
    // ─────────────────────────────────────────────────────────

    user, err := userService.Create(r.Context(), input)
    if err != nil {
        // Map service errors to HTTP responses
        switch {
        case errors.Is(err, ErrDuplicateEmail):
            // Conflict → 409
            writeError(w, http.StatusConflict, "DUPLICATE_EMAIL",
                "A user with this email already exists")
        case errors.Is(err, ErrNotFound):
            // Not found → 404
            writeError(w, http.StatusNotFound, "NOT_FOUND", err.Error())
        default:
            // ─────────────────────────────────────────────────
            // REPOSITORY LAYER: Infrastructure errors bubble up
            // ─────────────────────────────────────────────────

            // Log the actual error (don't expose to client)
            log.Printf("unexpected error: %v", err)

            // Generic error → 500
            writeError(w, http.StatusInternalServerError,
                "INTERNAL_ERROR", "An unexpected error occurred")
        }
        return
    }

    // Success → 201
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
```

---

## Middleware Error Handling Order

**Critical:** Middleware order matters. Errors flow UP through the chain.

```
┌─────────────────────────────────────────────────────────────────────────┐
│                     RECOMMENDED MIDDLEWARE ORDER                         │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                          │
│  REQUEST FLOW (top to bottom):                                          │
│                                                                          │
│  1. Recovery      ← FIRST: Catches panics from ALL subsequent layers    │
│       │                                                                  │
│       ▼                                                                  │
│  2. Request ID    ← Adds ID for error correlation                       │
│       │                                                                  │
│       ▼                                                                  │
│  3. Logging       ← Logs request start, captures duration               │
│       │                                                                  │
│       ▼                                                                  │
│  4. Rate Limit    ← Rejects before expensive operations                 │
│       │                                                                  │
│       ▼                                                                  │
│  5. Auth          ← Validates credentials                               │
│       │                                                                  │
│       ▼                                                                  │
│  6. CORS          ← Handles preflight, sets headers                     │
│       │                                                                  │
│       ▼                                                                  │
│  7. Handler       ← Business logic                                      │
│                                                                          │
│  ERROR FLOW (bottom to top):                                            │
│  Handler panics → bubbles up → Recovery catches → 500 response          │
│                                                                          │
└─────────────────────────────────────────────────────────────────────────┘
```

### Complete Middleware Chain Example

```go
func main() {
    handler := http.HandlerFunc(myHandler)

    // Apply in REVERSE order (last applied runs first)
    wrapped := Chain(handler,
        RecoveryMiddleware,     // 1st to run
        RequestIDMiddleware,    // 2nd
        LoggingMiddleware,      // 3rd
        RateLimitMiddleware,    // 4th
        AuthMiddleware,         // 5th
    )

    http.ListenAndServe(":8080", wrapped)
}

func Chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        h = middlewares[i](h)
    }
    return h
}
```

### Recovery Middleware (Always First)

```go
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                // Log with request ID for correlation
                reqID := r.Header.Get("X-Request-ID")
                log.Printf("[%s] PANIC: %v\n%s", reqID, err, debug.Stack())

                // Return JSON error (not stack trace!)
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(map[string]string{
                    "error":      "internal server error",
                    "request_id": reqID,
                })
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

---

## Client vs Server Perspective

### Server: Creating Error Responses

```go
// Define your error hierarchy
var (
    ErrNotFound     = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrForbidden    = errors.New("forbidden")
    ErrConflict     = errors.New("conflict")
)

// Centralized error-to-status mapping
func errorToStatus(err error) int {
    switch {
    case errors.Is(err, ErrNotFound):
        return http.StatusNotFound
    case errors.Is(err, ErrUnauthorized):
        return http.StatusUnauthorized
    case errors.Is(err, ErrForbidden):
        return http.StatusForbidden
    case errors.Is(err, ErrConflict):
        return http.StatusConflict
    default:
        return http.StatusInternalServerError
    }
}

// Centralized error response writer
func handleError(w http.ResponseWriter, r *http.Request, err error) {
    status := errorToStatus(err)

    // Log internal errors, not client errors
    if status >= 500 {
        log.Printf("internal error: %v", err)
    }

    // Don't expose internal details for 5xx
    message := err.Error()
    if status >= 500 {
        message = "an internal error occurred"
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(map[string]string{
        "error": message,
    })
}
```

### Client: Parsing Error Responses

```go
type APIError struct {
    StatusCode int
    Code       string
    Message    string
    Retryable  bool
}

func (e *APIError) Error() string {
    return fmt.Sprintf("[%d] %s: %s", e.StatusCode, e.Code, e.Message)
}

func callAPI(url string) (*Response, error) {
    resp, err := http.Get(url)
    if err != nil {
        // Network error - usually retryable
        return nil, &APIError{
            StatusCode: 0,
            Code:       "NETWORK_ERROR",
            Message:    err.Error(),
            Retryable:  true,
        }
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        // Parse error response
        var errResp struct {
            Error   string `json:"error"`
            Message string `json:"message"`
        }
        json.NewDecoder(resp.Body).Decode(&errResp)

        return nil, &APIError{
            StatusCode: resp.StatusCode,
            Code:       errResp.Error,
            Message:    errResp.Message,
            Retryable:  isRetryableStatus(resp.StatusCode),
        }
    }

    // Parse success response...
}

func isRetryableStatus(status int) bool {
    switch status {
    case 429, 502, 503, 504:
        return true
    default:
        return false
    }
}
```

---

## Quick Reference Tables

### Status Code → Go Error Pattern

| Status | Recommended Pattern | Example |
|--------|--------------------| --------|
| 400 | Wrapped or custom | `fmt.Errorf("parsing: %w", err)` |
| 401 | Sentinel | `var ErrUnauthorized = errors.New(...)` |
| 403 | Sentinel | `var ErrForbidden = errors.New(...)` |
| 404 | Sentinel | `var ErrNotFound = errors.New(...)` |
| 409 | Custom type | `type ConflictError struct { Field string }` |
| 422 | ValidationErrors | `type ValidationErrors struct { Errors []FieldError }` |
| 429 | Custom type | `type RateLimitError struct { RetryAfter int }` |
| 500 | Catch-all | Log original, return generic |
| 503 | Custom type | `type ServiceUnavailableError struct { RetryAfter int }` |

### Common Scenario → Complete Solution

| Scenario | Status | Response Format | Headers |
|----------|--------|-----------------|---------|
| Invalid JSON body | 400 | Simple JSON | - |
| Missing required field | 422 | RFC 7807 with field errors | - |
| Invalid auth token | 401 | Simple JSON | `WWW-Authenticate` |
| No permission | 403 | Simple JSON | - |
| Resource doesn't exist | 404 | Simple JSON | - |
| Duplicate resource | 409 | JSON with conflict details | - |
| Too many requests | 429 | JSON | `Retry-After` |
| Database timeout | 504 | Simple JSON | - |
| Unexpected panic | 500 | Simple JSON (no details) | - |

### Response Format Decision

```
Do you need machine-parseable errors?
├── NO  → http.Error() (plain text)
└── YES → Is it a public API?
          ├── YES → RFC 7807 Problem Details
          └── NO  → Simple JSON { "error": "...", "message": "..." }

Are there multiple validation errors?
├── YES → RFC 7807 with "errors" extension field
└── NO  → Simple error response
```

---

## Best Practices

### Do

1. **Centralize error handling** - One function maps errors to responses
2. **Log 5xx errors, not 4xx** - Client errors are expected; server errors need investigation
3. **Include request IDs** - For error correlation in logs
4. **Use appropriate status codes** - Not everything is a 400 or 500
5. **Set Content-Type header** - Before writing error response
6. **Return after writing error** - Don't continue processing

### Don't

1. **Don't expose stack traces** - Log them, don't send to client
2. **Don't expose internal error messages** - For 5xx, use generic messages
3. **Don't forget Retry-After** - For 429 and 503 responses
4. **Don't use 200 for errors** - Some old APIs do this; don't
5. **Don't mix response formats** - Be consistent across your API

---

## Resources

- [ERROR_HANDLING_PATTERNS.md](ERROR_HANDLING_PATTERNS.md) - Go error fundamentals
- [HTTP_PROTOCOL_GUIDE.md](HTTP_PROTOCOL_GUIDE.md) - HTTP protocol details
- [RFC 7807 - Problem Details](https://tools.ietf.org/html/rfc7807)
- [HTTP Status Codes (MDN)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)
