# HTTP Protocol in Go: A Complete Guide

A comprehensive reference for understanding HTTP and building handlers in Go.

---

## What Is HTTP?

HTTP (HyperText Transfer Protocol) is a **text-based request/response protocol**. Every HTTP message (request or response) has the same fundamental structure:

```
┌─────────────────────────────────────────────────────────────────┐
│                        HTTP MESSAGE                             │
├─────────────────────────────────────────────────────────────────┤
│  START LINE     (Request: method + path, Response: status)      │
├─────────────────────────────────────────────────────────────────┤
│  HEADERS        (Key: Value pairs, metadata about the message)  │
├─────────────────────────────────────────────────────────────────┤
│  BLANK LINE     (CRLF - signals end of headers)                 │
├─────────────────────────────────────────────────────────────────┤
│  BODY           (Optional: the actual content/payload)          │
└─────────────────────────────────────────────────────────────────┘
```

**Key insight:** HTTP is just structured text. Understanding this structure makes Go's `http.ResponseWriter` operations intuitive.

---

## What HTTP Actually Looks Like on the Wire

### A Real HTTP Request (what the client sends):

```
GET /user?id=123 HTTP/1.1        ← START LINE (method, path, version)
Host: api.example.com            ← HEADER
Accept: application/json         ← HEADER
Authorization: Bearer xyz123     ← HEADER
User-Agent: Mozilla/5.0          ← HEADER
                                 ← BLANK LINE (end of headers)
                                 ← BODY (empty for GET)
```

### A Real HTTP Response (what your server sends):

```
HTTP/1.1 200 OK                  ← START LINE (version, status code, reason)
Content-Type: application/json   ← HEADER
Content-Length: 42               ← HEADER
Date: Thu, 05 Dec 2025 10:30:00  ← HEADER
                                 ← BLANK LINE (end of headers)
{"id": 123, "name": "Alice"}     ← BODY (the actual data)
```

---

## The Three Parts You Control in Go

When you write an HTTP handler, you're constructing a **response**. Go's `http.ResponseWriter` gives you control over all three parts:

```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    // 1. HEADERS - metadata about your response
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Custom-Header", "my-value")

    // 2. STATUS LINE - the response code
    w.WriteHeader(http.StatusOK)  // 200

    // 3. BODY - the actual content
    w.Write([]byte(`{"message": "hello"}`))
}
```

### Critical Order Rule

```
┌────────────────────────────────────────────────────────────────┐
│  You MUST set headers BEFORE WriteHeader()                     │
│  You MUST call WriteHeader() BEFORE Write()                    │
│  (Or let Write() implicitly call WriteHeader(200))             │
└────────────────────────────────────────────────────────────────┘

w.Header().Set(...)   ← Can modify headers
w.WriteHeader(201)    ← LOCKS headers, sets status
w.Header().Set(...)   ← TOO LATE! Headers already sent
w.Write(body)         ← Sends body
```

---

## Headers: The Metadata Layer

Headers are **key-value pairs** that describe the message. Think of them as an envelope's label - they tell you about the contents without opening it.

### Common Request Headers (what client sends)

| Header | Purpose | Example |
|--------|---------|---------|
| `Host` | Which server to connect to | `api.example.com` |
| `Accept` | What formats client understands | `application/json` |
| `Content-Type` | Format of request body | `application/json` |
| `Content-Length` | Size of request body | `1234` |
| `Authorization` | Authentication credentials | `Bearer token123` |
| `User-Agent` | Client software identifier | `Mozilla/5.0...` |
| `Cookie` | Session data | `session=abc123` |
| `Accept-Encoding` | Supported compression | `gzip, deflate` |
| `Accept-Language` | Preferred languages | `en-US,en;q=0.9` |
| `Referer` | Previous page URL | `https://example.com/page` |
| `Origin` | Request origin (CORS) | `https://mysite.com` |

### Common Response Headers (what server sends)

| Header | Purpose | Example |
|--------|---------|---------|
| `Content-Type` | Format of response body | `application/json` |
| `Content-Length` | Size of body in bytes | `1234` |
| `Set-Cookie` | Tell client to store cookie | `session=xyz; HttpOnly` |
| `Cache-Control` | Caching instructions | `max-age=3600` |
| `Location` | Redirect URL (with 3xx status) | `/new-location` |
| `Access-Control-Allow-Origin` | CORS permission | `*` or `https://allowed.com` |
| `Content-Encoding` | Compression used | `gzip` |
| `ETag` | Resource version identifier | `"abc123"` |
| `Last-Modified` | When resource changed | `Wed, 21 Oct 2025 07:28:00 GMT` |

### Reading Request Headers in Go

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Get a specific header (case-insensitive)
    auth := r.Header.Get("Authorization")
    contentType := r.Header.Get("Content-Type")

    // Get all values for a header (some can have multiple)
    accepts := r.Header.Values("Accept")

    // Check if header exists
    if r.Header.Get("X-Custom") == "" {
        // Header not present (or empty)
    }

    // Iterate all headers
    for name, values := range r.Header {
        for _, value := range values {
            fmt.Printf("%s: %s\n", name, value)
        }
    }
}
```

### Setting Response Headers in Go

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Set a header (overwrites if exists)
    w.Header().Set("Content-Type", "application/json")

    // Add a header (allows multiple values)
    w.Header().Add("Set-Cookie", "a=1; HttpOnly")
    w.Header().Add("Set-Cookie", "b=2; HttpOnly")

    // Delete a header
    w.Header().Del("X-Unwanted")

    // Get current value before sending
    current := w.Header().Get("Content-Type")
}
```

---

## Body: The Payload

The body is the **actual data** being transferred. It can be:

- **Empty** (common for GET requests, redirects, 204 responses)
- **Text** (HTML, JSON, XML, plain text)
- **Binary** (images, files, protobuf)

### Reading Request Body

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Method 1: Read all bytes
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()  // Always close!

    // Method 2: JSON decode directly
    var data MyStruct
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Method 3: Form data
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid form", http.StatusBadRequest)
        return
    }
    username := r.FormValue("username")

    // Method 4: Multipart form (file uploads)
    if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max
        http.Error(w, "Invalid multipart", http.StatusBadRequest)
        return
    }
    file, header, err := r.FormFile("upload")
}
```

### Writing Response Body

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Method 1: Write raw bytes
    w.Write([]byte("Hello, World!"))

    // Method 2: Use fmt (writes to w which implements io.Writer)
    fmt.Fprintf(w, "User: %s, Age: %d", username, age)

    // Method 3: JSON encoder (writes directly to w)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(myStruct)

    // Method 4: Serve a file
    http.ServeFile(w, r, "path/to/file.html")

    // Method 5: Stream data
    flusher, ok := w.(http.Flusher)
    if ok {
        for data := range dataChannel {
            fmt.Fprintf(w, "data: %s\n\n", data)
            flusher.Flush()  // Send immediately
        }
    }
}
```

---

## Status Codes: The Response Categories

```
┌─────────────────────────────────────────────────────────────────┐
│  1xx - Informational (rare, protocol-level)                     │
│  2xx - Success (request worked)                                 │
│  3xx - Redirect (go somewhere else)                             │
│  4xx - Client Error (your fault)                                │
│  5xx - Server Error (our fault)                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Complete Status Code Reference

#### 2xx Success

| Code | Constant | When to Use |
|------|----------|-------------|
| 200 | `http.StatusOK` | Success, returning data |
| 201 | `http.StatusCreated` | Created new resource (POST) |
| 202 | `http.StatusAccepted` | Request accepted, processing async |
| 204 | `http.StatusNoContent` | Success, no body to return (DELETE) |

#### 3xx Redirect

| Code | Constant | When to Use |
|------|----------|-------------|
| 301 | `http.StatusMovedPermanently` | URL changed forever (SEO) |
| 302 | `http.StatusFound` | Temporary redirect |
| 303 | `http.StatusSeeOther` | Redirect after POST (PRG pattern) |
| 304 | `http.StatusNotModified` | Cached version is still valid |
| 307 | `http.StatusTemporaryRedirect` | Temporary, keep method |
| 308 | `http.StatusPermanentRedirect` | Permanent, keep method |

#### 4xx Client Error

| Code | Constant | When to Use |
|------|----------|-------------|
| 400 | `http.StatusBadRequest` | Malformed request syntax |
| 401 | `http.StatusUnauthorized` | Need authentication |
| 403 | `http.StatusForbidden` | Authenticated but not allowed |
| 404 | `http.StatusNotFound` | Resource doesn't exist |
| 405 | `http.StatusMethodNotAllowed` | Wrong HTTP method |
| 409 | `http.StatusConflict` | Resource conflict (duplicate) |
| 415 | `http.StatusUnsupportedMediaType` | Wrong Content-Type |
| 422 | `http.StatusUnprocessableEntity` | Validation failed |
| 429 | `http.StatusTooManyRequests` | Rate limited |

#### 5xx Server Error

| Code | Constant | When to Use |
|------|----------|-------------|
| 500 | `http.StatusInternalServerError` | Unexpected server error |
| 501 | `http.StatusNotImplemented` | Feature not implemented |
| 502 | `http.StatusBadGateway` | Upstream server error |
| 503 | `http.StatusServiceUnavailable` | Server overloaded/maintenance |
| 504 | `http.StatusGatewayTimeout` | Upstream timeout |

---

## Complete Request/Response Flow

```
CLIENT                                             SERVER
  │                                                   │
  │  ┌─────────────────────────────────────────────┐  │
  │  │ GET /api/user?id=123 HTTP/1.1               │  │
  │  │ Host: example.com                           │  │
  │  │ Accept: application/json                    │──┼──►  r *http.Request
  │  │ Authorization: Bearer token123              │  │     - r.Method = "GET"
  │  │                                             │  │     - r.URL.Path = "/api/user"
  │  │ (no body for GET)                           │  │     - r.URL.Query().Get("id") = "123"
  │  └─────────────────────────────────────────────┘  │     - r.Header.Get("Authorization")
  │                                                   │
  │                                                   │     Your Handler Runs:
  │                                                   │     - Read request info
  │                                                   │     - Do business logic
  │                                                   │     - Write response
  │                                                   │
  │  ┌─────────────────────────────────────────────┐  │
  │  │ HTTP/1.1 200 OK                             │  │
  │  │ Content-Type: application/json              │◄─┼──  w http.ResponseWriter
  │  │ Content-Length: 42                          │  │     - w.Header().Set(...)
  │  │                                             │  │     - w.WriteHeader(200)
  │  │ {"id": 123, "name": "Alice"}                │  │     - w.Write(jsonBytes)
  │  └─────────────────────────────────────────────┘  │
  │                                                   │
```

---

## Go's http.ResponseWriter Interface

```go
type ResponseWriter interface {
    // Header returns the header map that will be sent
    // Modify this BEFORE calling WriteHeader or Write
    Header() http.Header

    // WriteHeader sends the status code
    // Can only be called once! Subsequent calls are ignored
    // If you don't call it, Write() calls WriteHeader(200) automatically
    WriteHeader(statusCode int)

    // Write writes data to the response body
    // First call implicitly calls WriteHeader(200) if not already called
    Write([]byte) (int, error)
}
```

---

## Go's http.Request Struct

```go
type Request struct {
    // Method: GET, POST, PUT, DELETE, PATCH, etc.
    Method string

    // URL contains path, query params, fragment
    URL *url.URL
    // r.URL.Path = "/api/users"
    // r.URL.Query().Get("id") = "123"
    // r.URL.RawQuery = "id=123&sort=name"

    // Proto: "HTTP/1.1" or "HTTP/2.0"
    Proto string

    // Headers from client
    Header http.Header

    // Body (for POST/PUT/PATCH)
    Body io.ReadCloser

    // ContentLength: -1 if unknown
    ContentLength int64

    // Host: from Host header or URL
    Host string

    // RemoteAddr: client IP:port
    RemoteAddr string  // "192.168.1.1:54321"

    // Context for cancellation/timeouts/values
    // ctx := r.Context()

    // Form data (after ParseForm)
    Form url.Values
    PostForm url.Values

    // Multipart form (after ParseMultipartForm)
    MultipartForm *multipart.Form
}
```

### Extracting Request Information

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Method
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Path
    path := r.URL.Path  // "/api/users/123"

    // Query parameters
    id := r.URL.Query().Get("id")           // Single value
    tags := r.URL.Query()["tags"]           // Multiple values

    // Path parameters (with router like chi or gorilla/mux)
    // id := chi.URLParam(r, "id")

    // Headers
    auth := r.Header.Get("Authorization")

    // Cookies
    cookie, err := r.Cookie("session")
    if err == nil {
        sessionID := cookie.Value
    }

    // Client IP (may be behind proxy)
    clientIP := r.RemoteAddr
    // Or check X-Forwarded-For header for proxied requests

    // Context (for timeouts, cancellation, values)
    ctx := r.Context()
    userID := ctx.Value("userID")
}
```

---

## Common Handler Patterns

### Pattern 1: Simple Text Response

```go
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
    // Implicit: Content-Type: text/plain, Status: 200
}
```

### Pattern 2: JSON Response

```go
func JSONHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{
        "message": "success",
        "count":   42,
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(data); err != nil {
        // Log error, but headers already sent
        log.Printf("JSON encode error: %v", err)
    }
}
```

### Pattern 3: Error Response

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "missing id parameter", http.StatusBadRequest)
        return  // IMPORTANT: return after error!
    }

    // Continue processing...
}
```

### Pattern 4: JSON Error Response

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    if err := validate(r); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "error": err.Error(),
        })
        return
    }
}
```

### Pattern 5: Redirect

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Method 1: Manual
    w.Header().Set("Location", "/new-path")
    w.WriteHeader(http.StatusFound)  // 302

    // Method 2: Helper function (preferred)
    http.Redirect(w, r, "/new-path", http.StatusFound)
}
```

### Pattern 6: Method Filtering

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handleGet(w, r)
    case http.MethodPost:
        handlePost(w, r)
    default:
        w.Header().Set("Allow", "GET, POST")
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
```

### Pattern 7: Reading JSON Body

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Validate
    if input.Name == "" {
        http.Error(w, "name is required", http.StatusUnprocessableEntity)
        return
    }

    // Process...
}
```

### Pattern 8: Setting Cookies

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    cookie := &http.Cookie{
        Name:     "session",
        Value:    "abc123",
        Path:     "/",
        MaxAge:   86400,           // 1 day in seconds
        HttpOnly: true,            // Not accessible via JavaScript
        Secure:   true,            // HTTPS only
        SameSite: http.SameSiteLaxMode,
    }
    http.SetCookie(w, cookie)

    w.Write([]byte("Cookie set!"))
}
```

---

## Handler Factories (Closures + HTTP)

Combine closures with HTTP handlers for dependency injection:

```go
// Factory injects dependencies
func MakeUserHandler(db *sql.DB, logger *log.Logger) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Query().Get("id")

        user, err := db.Query("SELECT * FROM users WHERE id = ?", id)
        if err != nil {
            logger.Printf("Database error: %v", err)
            http.Error(w, "Internal error", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(user)
    }
}

// Usage
func main() {
    db := connectDB()
    logger := log.New(os.Stdout, "[API] ", log.LstdFlags)

    http.HandleFunc("/user", MakeUserHandler(db, logger))
    http.ListenAndServe(":8080", nil)
}
```

---

## Middleware Pattern

Middleware wraps handlers to add cross-cutting functionality:

```go
// Middleware signature
type Middleware func(http.Handler) http.Handler

// Logging middleware
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        next.ServeHTTP(w, r)  // Call the next handler

        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

// Auth middleware
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return  // Don't call next
        }

        // Validate token...

        next.ServeHTTP(w, r)
    })
}

// Chain middleware
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        h = middlewares[i](h)
    }
    return h
}

// Usage
handler := Chain(
    myHandler,
    LoggingMiddleware,
    AuthMiddleware,
)
```

---

## Visual Summary: What Goes Where

```
                     HTTP RESPONSE
    ┌──────────────────────────────────────────────┐
    │           STATUS LINE                         │
    │  ┌────────────────────────────────────────┐  │
    │  │  HTTP/1.1 200 OK                       │  │ ← w.WriteHeader(200)
    │  └────────────────────────────────────────┘  │
    ├──────────────────────────────────────────────┤
    │           HEADERS                            │
    │  ┌────────────────────────────────────────┐  │
    │  │  Content-Type: application/json        │  │ ← w.Header().Set(...)
    │  │  Content-Length: 42                    │  │ ← Auto-calculated by Go
    │  │  Date: Thu, 05 Dec 2025 10:30:00 GMT   │  │ ← Auto-added by Go
    │  │  X-Custom: my-value                    │  │ ← w.Header().Set(...)
    │  └────────────────────────────────────────┘  │
    ├──────────────────────────────────────────────┤
    │           (blank line)                       │ ← Automatic
    ├──────────────────────────────────────────────┤
    │           BODY                               │
    │  ┌────────────────────────────────────────┐  │
    │  │  {"id": 123, "name": "Alice"}          │  │ ← w.Write(bytes) or
    │  │                                        │  │   json.NewEncoder(w).Encode(data)
    │  └────────────────────────────────────────┘  │
    └──────────────────────────────────────────────┘
```

---

## Content-Type Reference

| Content-Type | Use For | Go Example |
|--------------|---------|------------|
| `application/json` | JSON data | `json.NewEncoder(w).Encode(data)` |
| `text/plain` | Plain text | `w.Write([]byte("hello"))` |
| `text/html` | HTML pages | `w.Write([]byte("<h1>Hi</h1>"))` |
| `application/xml` | XML data | `xml.NewEncoder(w).Encode(data)` |
| `application/octet-stream` | Binary files | `io.Copy(w, file)` |
| `image/png` | PNG images | `io.Copy(w, imageFile)` |
| `application/pdf` | PDF files | `io.Copy(w, pdfFile)` |
| `multipart/form-data` | File uploads | (request only) |
| `application/x-www-form-urlencoded` | Form data | (request only) |

---

## HTTP Methods Semantics

| Method | Purpose | Body? | Idempotent? | Safe? |
|--------|---------|-------|-------------|-------|
| GET | Retrieve resource | No | Yes | Yes |
| HEAD | Get headers only | No | Yes | Yes |
| POST | Create resource | Yes | No | No |
| PUT | Replace resource | Yes | Yes | No |
| PATCH | Partial update | Yes | No | No |
| DELETE | Remove resource | Optional | Yes | No |
| OPTIONS | Get allowed methods | No | Yes | Yes |

**Idempotent:** Multiple identical requests have same effect as single request.
**Safe:** Request doesn't modify server state.

---

## Testing HTTP Handlers

```go
func TestHelloHandler(t *testing.T) {
    // Create a request
    req := httptest.NewRequest(http.MethodGet, "/hello", nil)

    // Create a response recorder
    w := httptest.NewRecorder()

    // Call the handler
    HelloHandler(w, req)

    // Check status code
    if w.Code != http.StatusOK {
        t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
    }

    // Check body
    if w.Body.String() != "Hello, World!" {
        t.Errorf("got body %q, want %q", w.Body.String(), "Hello, World!")
    }

    // Check headers
    contentType := w.Header().Get("Content-Type")
    if contentType != "text/plain; charset=utf-8" {
        t.Errorf("got Content-Type %q", contentType)
    }
}

func TestJSONHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodPost, "/api/user",
        strings.NewReader(`{"name": "Alice"}`))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    UserHandler(w, req)

    // Parse JSON response
    var response map[string]interface{}
    json.NewDecoder(w.Body).Decode(&response)

    if response["name"] != "Alice" {
        t.Errorf("unexpected response: %v", response)
    }
}
```

---

## Common Mistakes

### 1. Writing After WriteHeader
```go
// WRONG - headers already sent
w.WriteHeader(200)
w.Header().Set("Content-Type", "application/json")  // Too late!
```

### 2. Forgetting to Return After Error
```go
// WRONG - continues execution after error
if err != nil {
    http.Error(w, "error", 500)
    // Missing return!
}
// This code still runs!
```

### 3. Multiple WriteHeader Calls
```go
// WRONG - second WriteHeader is ignored
w.WriteHeader(200)
w.WriteHeader(201)  // Ignored! Still 200
```

### 4. Not Closing Request Body
```go
// WRONG - may leak resources
body, _ := io.ReadAll(r.Body)
// Missing: defer r.Body.Close()
```

### 5. Assuming Content-Type
```go
// WRONG - may be form data, not JSON
json.NewDecoder(r.Body).Decode(&data)
// Should check Content-Type header first
```

---

## Summary

| Concept | Key Point |
|---------|-----------|
| HTTP Message | Start line + Headers + Blank line + Body |
| ResponseWriter | Controls headers, status, and body |
| Order | Set headers → WriteHeader → Write |
| Status Codes | 2xx success, 3xx redirect, 4xx client error, 5xx server error |
| Headers | Metadata key-value pairs (before body) |
| Body | The actual payload (text, JSON, binary) |
| Handler | `func(w http.ResponseWriter, r *http.Request)` |
| Middleware | `func(http.Handler) http.Handler` |
| Testing | Use `httptest.NewRequest` and `httptest.NewRecorder` |

---

## Resources

- [Go net/http Package Documentation](https://pkg.go.dev/net/http)
- [HTTP/1.1 Specification (RFC 7230-7235)](https://tools.ietf.org/html/rfc7230)
- [MDN HTTP Reference](https://developer.mozilla.org/en-US/docs/Web/HTTP)
- [Go by Example - HTTP Servers](https://gobyexample.com/http-servers)
- [Go by Example - HTTP Clients](https://gobyexample.com/http-clients)
