package api_client_builder

import (
	"io"
	"net/http"
	"time"
)

// Client represents an HTTP client with configuration
type Client struct {
	// TODO(human): Define fields
}

// Request represents a single HTTP request
type Request struct {
	// TODO(human): Define fields
}

// ClientOption configures a Client
type ClientOption func(*Client)

// WithTimeout sets the client timeout
func WithTimeout(timeout time.Duration) ClientOption {
	// TODO(human): Implement
	return nil
}

// WithHeader adds a default header to all requests
func WithHeader(key, value string) ClientOption {
	// TODO(human): Implement
	return nil
}

// NewClient creates a new HTTP client
func NewClient(baseURL string, opts ...ClientOption) *Client {
	// TODO(human): Implement
	return nil
}

// NewRequest creates a new request builder
func (c *Client) NewRequest(method, path string) *Request {
	// TODO(human): Implement
	return nil
}

// WithHeader adds a header to this request
func (r *Request) WithHeader(key, value string) *Request {
	// TODO(human): Implement
	return nil
}

// WithBody sets the request body
func (r *Request) WithBody(body io.Reader) *Request {
	// TODO(human): Implement
	return nil
}

// Do executes the request
func (r *Request) Do() (*http.Response, error) {
	// TODO(human): Implement
	return nil, nil
}

// Get performs a GET request
func (c *Client) Get(path string) (*http.Response, error) {
	// TODO(human): Implement
	return nil, nil
}

// Post performs a POST request
func (c *Client) Post(path string, body io.Reader) (*http.Response, error) {
	// TODO(human): Implement
	return nil, nil
}
