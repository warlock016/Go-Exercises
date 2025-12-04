# Exercise 15: API Client Builder

**Learning Goal:** Build a fluent HTTP client using method chaining, options, and functions

---

## 📝 Problem Description

This capstone exercise combines all patterns learned:
- Method chaining (builder pattern)
- Functional options
- Higher-order functions
- Error handling

You'll build a flexible HTTP client with a fluent API.

---

## 🎯 Type & Method Signatures

```go
type Client struct {
    baseURL string
    headers map[string]string
    timeout time.Duration
}

type Request struct {
    client  *Client
    method  string
    path    string
    headers map[string]string
    body    io.Reader
}

// Client builder
func NewClient(baseURL string, opts ...ClientOption) *Client

type ClientOption func(*Client)
func WithTimeout(timeout time.Duration) ClientOption
func WithHeader(key, value string) ClientOption

// Request builder
func (c *Client) NewRequest(method, path string) *Request
func (r *Request) WithHeader(key, value string) *Request
func (r *Request) WithBody(body io.Reader) *Request
func (r *Request) Do() (*http.Response, error)

// Convenience methods
func (c *Client) Get(path string) (*http.Response, error)
func (c *Client) Post(path string, body io.Reader) (*http.Response, error)
```

---

## 📖 Examples

```go
// Create client with options
client := NewClient("https://api.example.com",
    WithTimeout(10*time.Second),
    WithHeader("User-Agent", "MyApp/1.0"),
)

// Fluent request building
resp, err := client.NewRequest("GET", "/users").
    WithHeader("Accept", "application/json").
    Do()

// Convenience methods
resp, err := client.Get("/users")
resp, err := client.Post("/users", body)

// Complex request
resp, err := client.NewRequest("POST", "/data").
    WithHeader("Content-Type", "application/json").
    WithBody(strings.NewReader(`{"key":"value"}`)).
    Do()
```

---

## 📋 Instructions

1. Implement `Client` with options pattern
2. Implement `Request` with method chaining
3. Implement `Do()` to execute requests
4. Add convenience methods `Get()` and `Post()`
5. Combine base URL and path correctly
6. Merge client and request headers
7. Apply timeout
8. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Complete Solution</summary>

```go
package api_client_builder

import (
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	baseURL string
	headers map[string]string
	timeout time.Duration
	client  *http.Client
}

type Request struct {
	client  *Client
	method  string
	path    string
	headers map[string]string
	body    io.Reader
}

type ClientOption func(*Client)

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.timeout = timeout
	}
}

func WithHeader(key, value string) ClientOption {
	return func(c *Client) {
		c.headers[key] = value
	}
}

func NewClient(baseURL string, opts ...ClientOption) *Client {
	c := &Client{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		headers: make(map[string]string),
		timeout: 30 * time.Second,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.client = &http.Client{
		Timeout: c.timeout,
	}

	return c
}

func (c *Client) NewRequest(method, path string) *Request {
	return &Request{
		client:  c,
		method:  method,
		path:    path,
		headers: make(map[string]string),
	}
}

func (r *Request) WithHeader(key, value string) *Request {
	r.headers[key] = value
	return r
}

func (r *Request) WithBody(body io.Reader) *Request {
	r.body = body
	return r
}

func (r *Request) Do() (*http.Response, error) {
	url := r.client.baseURL + r.path
	req, err := http.NewRequest(r.method, url, r.body)
	if err != nil {
		return nil, err
	}

	// Apply client headers
	for k, v := range r.client.headers {
		req.Header.Set(k, v)
	}

	// Apply request headers (override client headers)
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	return r.client.client.Do(req)
}

func (c *Client) Get(path string) (*http.Response, error) {
	return c.NewRequest("GET", path).Do()
}

func (c *Client) Post(path string, body io.Reader) (*http.Response, error) {
	return c.NewRequest("POST", path).WithBody(body).Do()
}
```

</details>

---

## 🤔 Think About

1. How does this combine all patterns from the module?
2. Why use both client-level and request-level headers?
3. How would you add retry logic?
4. What about request/response middleware?

---

## 🎓 What This Teaches

- **Pattern combination**: Using multiple patterns together
- **Fluent APIs**: Creating expressive, chainable interfaces
- **HTTP clients**: Building custom HTTP client wrappers
- **Options pattern**: Flexible client configuration
- **Method chaining**: Builder pattern for requests
- **API design**: Creating intuitive, maintainable APIs

---

**Tier:** 4 - Mastery
**Estimated Time:** 60-75 minutes
