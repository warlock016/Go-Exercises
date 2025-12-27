package http_client

import "net/http"

// APIClient represents an HTTP API client
type APIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewAPIClient creates a new API client with 10 second timeout
func NewAPIClient(baseURL string) *APIClient {
	// TODO(human): Implement
	return nil
}

// Get performs a GET request and returns response body
func (c *APIClient) Get(path string) ([]byte, error) {
	// TODO(human): Implement
	return nil, nil
}

// GetJSON performs a GET request and decodes JSON into v
func (c *APIClient) GetJSON(path string, v interface{}) error {
	// TODO(human): Implement
	return nil
}

// Post performs a POST request with JSON body
func (c *APIClient) Post(path string, body interface{}) error {
	// TODO(human): Implement
	return nil
}
