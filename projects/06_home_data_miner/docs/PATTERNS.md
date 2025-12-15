# Go Patterns Guide - HTTP API Clients

This guide covers Go patterns and idioms relevant to building HTTP API clients. Examples are generic and intentionally incomplete - use them as reference patterns for your implementation.

---

## Table of Contents

1. [URL Construction](#1-url-construction)
2. [HTTP Request Building](#2-http-request-building)
3. [Response Handling](#3-response-handling)
4. [JSON Parsing Patterns](#4-json-parsing-patterns)
5. [Error Handling](#5-error-handling)
6. [Client Struct Pattern](#6-client-struct-pattern)
7. [Configuration Patterns](#7-configuration-patterns)
8. [Testing HTTP Clients](#8-testing-http-clients)

---

## 1. URL Construction

### The `net/url` Package

Go's `net/url` package provides tools for URL manipulation.

#### Parsing URLs

```go
import "net/url"

// Parse a URL string into components
u, err := url.Parse("http://example.com:8080/api/v1/query")
if err != nil {
    // Handle malformed URL
}

// Access components
u.Scheme   // "http"
u.Host     // "example.com:8080"
u.Path     // "/api/v1/query"
```

#### Building Query Parameters with `url.Values`

`url.Values` is a `map[string][]string` with helper methods:

```go
params := url.Values{}

// Set a single value (replaces any existing)
params.Set("query", "some_metric{label=\"value\"}")

// Add a value (allows multiple values for same key)
params.Add("match[]", "metric_one")
params.Add("match[]", "metric_two")

// Get the encoded string
encoded := params.Encode()
// Result: "match%5B%5D=metric_one&match%5B%5D=metric_two&query=some_metric..."
```

**Key Insight:** `Encode()` handles all URL encoding automatically - special characters like `{`, `}`, `"`, `=` become `%7B`, `%7D`, `%22`, `%3D`.

#### Combining Base URL with Query Parameters

**Pattern 1: String concatenation (simple)**
```go
baseURL := "http://example.com:8080"
endpoint := "/api/v1/query"
params := url.Values{}
params.Set("query", "metric_name")

fullURL := baseURL + endpoint + "?" + params.Encode()
```

**Pattern 2: Using url.URL struct (safer)**
```go
u, _ := url.Parse("http://example.com:8080/api/v1/query")
params := url.Values{}
params.Set("query", "metric_name")
u.RawQuery = params.Encode()

fullURL := u.String()
```

**Pattern 3: Path joining with `url.JoinPath` (Go 1.19+)**
```go
base := "http://example.com:8080"
fullPath, err := url.JoinPath(base, "api", "v1", "query")
// Result: "http://example.com:8080/api/v1/query"
// Note: Does NOT include query parameters - add those separately
```

### Common Pitfall: Query String in JoinPath

```go
// WRONG - JoinPath will escape the ? and =
url.JoinPath(base, "api/v1/query?foo=bar")

// CORRECT - Build path, then add query params separately
path, _ := url.JoinPath(base, "api/v1/query")
fullURL := path + "?" + params.Encode()
```

---

## 2. HTTP Request Building

### Creating Requests with `http.NewRequest`

```go
import "net/http"

// GET request (no body)
req, err := http.NewRequest("GET", fullURL, nil)

// POST request with body
body := strings.NewReader(`{"key": "value"}`)
req, err := http.NewRequest("POST", fullURL, body)
```

### Setting Headers

```go
// Single header
req.Header.Set("Authorization", "Bearer token123")

// Multiple values for same header
req.Header.Add("Accept", "application/json")
req.Header.Add("Accept", "text/plain")

// Common headers for JSON APIs
req.Header.Set("Accept", "application/json")
req.Header.Set("Content-Type", "application/json")  // Only for POST/PUT with JSON body
```

**Note:** For GET requests with query parameters, `Content-Type` is typically not needed since there's no request body.

### Using Context for Timeouts/Cancellation

```go
import "context"

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
```

---

## 3. Response Handling

### Basic Response Flow

```go
client := &http.Client{Timeout: 10 * time.Second}

resp, err := client.Do(req)
if err != nil {
    // Network error, timeout, etc.
    return err
}
defer resp.Body.Close()  // CRITICAL: Always close the body

// Check status code BEFORE reading body
if resp.StatusCode != http.StatusOK {
    // Handle error response
}

// Read body
body, err := io.ReadAll(resp.Body)
if err != nil {
    return err
}
```

### Why `defer resp.Body.Close()` Matters

```
Without Close():              With Close():
┌────────────────┐           ┌────────────────┐
│ Request 1      │           │ Request 1      │
│ Connection open├──┐        │ Connection open│──┐
└────────────────┘  │        └───────┬────────┘  │
                    │                │ Close()   │
┌────────────────┐  │        ┌───────▼────────┐  │
│ Request 2      │  │        │ Request 2      │  │
│ NEW connection │──┤        │ REUSE connection│◀─┘
└────────────────┘  │        └───────┬────────┘
       ...          │                │ Close()
┌────────────────┐  │               ...
│ Request N      │  │
│ EXHAUSTED!     │◀─┘        Connection pool healthy
└────────────────┘
```

### Status Code Checking Pattern

```go
// Pattern: Check for success range
if resp.StatusCode < 200 || resp.StatusCode >= 300 {
    body, _ := io.ReadAll(resp.Body)
    return fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, body)
}

// Or check specific codes
switch resp.StatusCode {
case http.StatusOK:
    // Process success
case http.StatusBadRequest:
    // Handle 400
case http.StatusNotFound:
    // Handle 404
default:
    // Handle unexpected
}
```

---

## 4. JSON Parsing Patterns

### Basic Unmarshaling

```go
import "encoding/json"

type Response struct {
    Status string `json:"status"`
    Data   []Item `json:"data"`
}

var result Response
err := json.Unmarshal(body, &result)
if err != nil {
    return fmt.Errorf("parsing JSON: %w", err)
}
```

### Struct Tags Explained

```go
type Example struct {
    // Field name in Go    -> JSON key mapping
    UserID   int    `json:"user_id"`          // "user_id" in JSON
    Name     string `json:"name"`             // "name" in JSON
    Email    string `json:"email,omitempty"`  // Omit if empty when marshaling
    Internal string `json:"-"`                // Never marshal/unmarshal
}
```

### Handling Dynamic JSON with `json.RawMessage`

When part of the JSON structure varies:

```go
type APIResponse struct {
    Status string          `json:"status"`
    Data   json.RawMessage `json:"data"`  // Defer parsing
}

// First pass: get the wrapper
var resp APIResponse
json.Unmarshal(body, &resp)

// Second pass: parse Data based on what we expect
var items []Item
json.Unmarshal(resp.Data, &items)
```

### Handling Mixed Type Arrays

VictoriaMetrics returns values as `[timestamp, "string_value"]`:

```go
// Option 1: Use []interface{} (requires type assertions)
type Result struct {
    Value []interface{} `json:"value"`
}
// Access: timestamp := result.Value[0].(float64)
//         value := result.Value[1].(string)

// Option 2: Use [2]interface{} for fixed size
type Result struct {
    Value [2]interface{} `json:"value"`
}

// Option 3: Custom unmarshaler (advanced)
type TimeValue struct {
    Timestamp int64
    Value     string
}
func (tv *TimeValue) UnmarshalJSON(data []byte) error {
    // Custom parsing logic
}
```

### Map for Unknown Keys

When the JSON has arbitrary keys:

```go
// For: {"__name__": "metric", "label1": "value1", "label2": "value2"}
type Series map[string]string

var series Series
json.Unmarshal(data, &series)
// Access: series["__name__"], series["label1"], etc.
```

---

## 5. Error Handling

### Error Wrapping with Context

```go
// Add context to errors using %w verb
data, err := fetchData(url)
if err != nil {
    return fmt.Errorf("fetching from %s: %w", url, err)
}

// Caller can unwrap to check underlying error
if errors.Is(err, context.DeadlineExceeded) {
    // Handle timeout
}
```

### Custom Error Types

```go
type APIError struct {
    StatusCode int
    Message    string
    ErrorType  string  // e.g., "bad_data", "timeout"
}

func (e *APIError) Error() string {
    return fmt.Sprintf("API error %d (%s): %s", e.StatusCode, e.ErrorType, e.Message)
}

// Usage
return &APIError{
    StatusCode: resp.StatusCode,
    Message:    "invalid query syntax",
    ErrorType:  "bad_data",
}
```

### Error Handling Decision Tree

```
Error occurred
     │
     ├── Is it recoverable?
     │        │
     │        ├── Yes (timeout, rate limit) → Retry with backoff
     │        │
     │        └── No (bad query, not found) → Return error to caller
     │
     └── Should caller know specifics?
              │
              ├── Yes → Use custom error type or wrap with context
              │
              └── No → Wrap with fmt.Errorf
```

---

## 6. Client Struct Pattern

### Why Use a Client Struct?

```
Without struct:                   With struct:

DoQuery(url, timeout, ...)       client.Query(promQL)
DoSeries(url, timeout, ...)      client.GetSeries(match)
DoLabels(url, timeout, ...)      client.GetLabels()

Config passed everywhere ❌       Config stored once ✓
No connection reuse ❌            Connection pooling ✓
```

### Basic Client Pattern

```go
type Client struct {
    httpClient *http.Client
    baseURL    string
}

func NewClient(baseURL string, timeout time.Duration) *Client {
    return &Client{
        httpClient: &http.Client{Timeout: timeout},
        baseURL:    baseURL,
    }
}

// Methods use the stored configuration
func (c *Client) Query(promQL string) (*QueryResult, error) {
    // Build URL using c.baseURL
    // Execute using c.httpClient
}
```

### Private Helper Method Pattern

```go
// Lowercase = private to package
func (c *Client) doRequest(endpoint string, params url.Values) ([]byte, error) {
    // 1. Build full URL
    // 2. Create request
    // 3. Execute request
    // 4. Check status
    // 5. Read body
    // 6. Return bytes or error
}

// Public methods use the helper
func (c *Client) Query(promQL string) (*QueryResult, error) {
    params := url.Values{}
    params.Set("query", promQL)

    body, err := c.doRequest("/api/v1/query", params)
    if err != nil {
        return nil, err
    }

    // Parse JSON specific to this endpoint
    // ...
}
```

### Functional Options Pattern (Advanced)

For flexible configuration:

```go
type ClientOption func(*Client)

func WithTimeout(d time.Duration) ClientOption {
    return func(c *Client) {
        c.httpClient.Timeout = d
    }
}

func WithHeader(key, value string) ClientOption {
    return func(c *Client) {
        c.defaultHeaders[key] = value
    }
}

func NewClient(baseURL string, opts ...ClientOption) *Client {
    c := &Client{
        httpClient: &http.Client{Timeout: 30 * time.Second},
        baseURL:    baseURL,
    }
    for _, opt := range opts {
        opt(c)
    }
    return c
}

// Usage
client := NewClient("http://localhost:8428",
    WithTimeout(10*time.Second),
    WithHeader("X-Custom", "value"),
)
```

---

## 7. Configuration Patterns

### Environment-Based Configuration

```go
type Config struct {
    BaseURL string
    Timeout time.Duration
}

func LoadConfig(envPath string) (*Config, error) {
    // 1. Load .env file
    // 2. Read required values
    // 3. Validate
    // 4. Parse durations, etc.
    // 5. Return config or error
}
```

### Validation Pattern

```go
func (c *Config) Validate() error {
    if c.BaseURL == "" {
        return errors.New("BaseURL is required")
    }

    u, err := url.Parse(c.BaseURL)
    if err != nil {
        return fmt.Errorf("invalid BaseURL: %w", err)
    }

    if u.Scheme != "http" && u.Scheme != "https" {
        return fmt.Errorf("BaseURL must use http or https scheme")
    }

    if c.Timeout <= 0 {
        c.Timeout = 30 * time.Second  // Default
    }

    return nil
}
```

### Duration Parsing

```go
import "time"

// From string
timeout, err := time.ParseDuration("10s")  // 10 * time.Second
timeout, err := time.ParseDuration("5m")   // 5 * time.Minute

// Valid suffixes: ns, us, ms, s, m, h
```

---

## 8. Testing HTTP Clients

### Using `httptest` for Mock Servers

```go
import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestQuery(t *testing.T) {
    // Create a mock server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Verify request
        if r.URL.Path != "/api/v1/query" {
            t.Errorf("unexpected path: %s", r.URL.Path)
        }

        // Check query params
        query := r.URL.Query().Get("query")
        if query == "" {
            t.Error("missing query parameter")
        }

        // Return mock response
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"success","data":...}`))
    }))
    defer server.Close()

    // Use mock server URL
    client := NewClient(server.URL, 5*time.Second)
    result, err := client.Query("test_metric")

    // Assert results
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    // ... more assertions
}
```

### Table-Driven HTTP Tests

```go
func TestQueryErrors(t *testing.T) {
    tests := []struct {
        name           string
        serverStatus   int
        serverResponse string
        wantErr        bool
    }{
        {
            name:           "success",
            serverStatus:   200,
            serverResponse: `{"status":"success","data":[]}`,
            wantErr:        false,
        },
        {
            name:           "bad request",
            serverStatus:   400,
            serverResponse: `{"status":"error","error":"bad query"}`,
            wantErr:        true,
        },
        {
            name:           "server error",
            serverStatus:   500,
            serverResponse: ``,
            wantErr:        true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.WriteHeader(tt.serverStatus)
                w.Write([]byte(tt.serverResponse))
            }))
            defer server.Close()

            client := NewClient(server.URL, 5*time.Second)
            _, err := client.Query("test")

            if (err != nil) != tt.wantErr {
                t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### Testing URL Construction (Unit Test)

```go
func TestBuildURL(t *testing.T) {
    tests := []struct {
        name     string
        base     string
        endpoint string
        params   map[string]string
        want     string
    }{
        {
            name:     "simple query",
            base:     "http://localhost:8428",
            endpoint: "/api/v1/query",
            params:   map[string]string{"query": "up"},
            want:     "http://localhost:8428/api/v1/query?query=up",
        },
        {
            name:     "special characters",
            base:     "http://localhost:8428",
            endpoint: "/api/v1/query",
            params:   map[string]string{"query": `metric{label="value"}`},
            want:     "http://localhost:8428/api/v1/query?query=metric%7Blabel%3D%22value%22%7D",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := buildURL(tt.base, tt.endpoint, tt.params)
            if got != tt.want {
                t.Errorf("buildURL() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

---

## Quick Reference: Import Paths

```go
import (
    // URL handling
    "net/url"

    // HTTP client
    "net/http"

    // Reading response body
    "io"

    // JSON parsing
    "encoding/json"

    // Context for timeouts
    "context"

    // Time and durations
    "time"

    // Error handling
    "errors"
    "fmt"

    // Testing HTTP
    "net/http/httptest"
    "testing"
)
```

---

## Summary: Request Flow Checklist

When implementing an API method:

- [ ] Build query parameters with `url.Values`
- [ ] Construct full URL (base + endpoint + encoded params)
- [ ] Create request with `http.NewRequest`
- [ ] Set any required headers
- [ ] Execute with `client.Do(req)`
- [ ] Defer `resp.Body.Close()`
- [ ] Check status code
- [ ] Read body with `io.ReadAll`
- [ ] Unmarshal JSON into typed struct
- [ ] Return data or wrapped error
