package api_client_builder

import (
	"io"
	"net/http"
	"strings"
	"time"
)

// Client represents an HTTP client with configuration
type Client struct {
	// TODO(human): Define fields
	baseURL string
	headers map[string]string
	timeout time.Duration
	client  *http.Client
}

// Request represents a single HTTP request
type Request struct {
	// TODO(human): Define fields
	client  *Client
	method  string
	path    string
	headers map[string]string
	body    io.Reader
}

// ClientOption configures a Client
type ClientOption func(*Client)

// WithTimeout sets the client timeout
func WithTimeout(timeout time.Duration) ClientOption {
	// TODO(human): Implement
	return func(c *Client) {
		c.timeout = timeout
	}
}

// WithHeader adds a default header to all requests
func WithHeader(key, value string) ClientOption {
	// TODO(human): Implement
	return func(c *Client) {
		c.headers[key] = value
	}
}

// NewClient creates a new HTTP client
func NewClient(baseURL string, opts ...ClientOption) *Client {
	// TODO(human): Implement

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

// NewRequest creates a new request builder
func (c *Client) NewRequest(method, path string) *Request {
	// TODO(human): Implement
	return &Request{
		client:  c,
		method:  method,
		path:    path,
		headers: make(map[string]string),
	}
}

// WithHeader adds a header to this request
func (r *Request) WithHeader(key, value string) *Request {
	// TODO(human): Implement
	new := r
	new.headers[key] = value
	return new
}

// WithBody sets the request body
func (r *Request) WithBody(body io.Reader) *Request {
	// TODO(human): Implement
	new := r
	new.body = body
	return new
}

// Do executes the request
func (r *Request) Do() (*http.Response, error) {
	// TODO(human): Implement
	url := r.client.baseURL
	if !strings.Contains(r.path, "/") {
		url += "/"
	}
	url += r.path

	req, err := http.NewRequest(r.method, url, r.body)
	if err != nil {
		return nil, err
	}

	for k, v := range r.client.headers {
		req.Header.Set(k, v)
	}

	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	return r.client.client.Do(req)
}

// Get performs a GET request
func (c *Client) Get(path string) (*http.Response, error) {
	// TODO(human): Implement
	return c.NewRequest("GET", path).Do()
}

// Post performs a POST request
func (c *Client) Post(path string, body io.Reader) (*http.Response, error) {
	// TODO(human): Implement
	return c.NewRequest("POST", path).WithBody(body).Do()
}
