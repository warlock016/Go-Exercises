# Exercise 10: HTTP Client

**Learning Goal:** Make HTTP requests using http.Client and test with httptest.Server

**Difficulty:** Tier 3 - Integration
**Estimated Time:** 35-40 minutes

## Problem Description

Build an HTTP client that makes GET and POST requests with proper error handling, timeouts, and JSON encoding/decoding. You'll test it using httptest.NewServer for offline testing.

## Type Definition

```go
type APIClient struct {
    BaseURL    string
    HTTPClient *http.Client
}
```

## Function Signatures

```go
// NewAPIClient creates a new API client with default timeout
func NewAPIClient(baseURL string) *APIClient

// Get performs a GET request and returns response body
func (c *APIClient) Get(path string) ([]byte, error)

// GetJSON performs a GET request and decodes JSON into v
func (c *APIClient) GetJSON(path string, v interface{}) error

// Post performs a POST request with JSON body
func (c *APIClient) Post(path string, body interface{}) error
```

## What This Teaches

- **HTTP client**: Making requests with http.Client
- **Timeout handling**: Setting reasonable timeouts
- **JSON marshaling**: Encoding request bodies
- **Error handling**: Network errors, status codes, decode errors
- **httptest.NewServer**: Testing clients without real servers
