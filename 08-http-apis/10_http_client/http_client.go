package http_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// APIClient represents an HTTP API client
type APIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewAPIClient creates a new API client with 10 second timeout
func NewAPIClient(baseURL string) *APIClient {
	// TODO(human): Implement
	return &APIClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// Get performs a GET request and returns response body
func (c *APIClient) Get(path string) ([]byte, error) {
	// TODO(human): Implement

	url, err := url.JoinPath(c.BaseURL, path)
	if err != nil {
		return []byte{}, errors.New("failed to compose request url")
	}

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return []byte{}, fmt.Errorf("request failed with status %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != io.EOF && err != nil {
		return []byte{}, err
	}

	return body, nil
}

// GetJSON performs a GET request and decodes JSON into v
func (c *APIClient) GetJSON(path string, v any) error {
	// TODO(human): Implement

	url, err := url.JoinPath(c.BaseURL, path)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("request failed with status %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if len(body) == 0 {
		return err
	}
	err = json.Unmarshal(body, v)
	return err
}

// Post performs a POST request with JSON body
func (c *APIClient) Post(path string, body any) error {
	// TODO(human): Implement
	url, err := url.JoinPath(c.BaseURL, path)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	return nil
}
