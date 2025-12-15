package client

import (
	"fmt"
	"net/http"
	"net/url"
)

// admin/import/export purposes: "baseURL/api/v1" ["/api/v1"]
// query/discovery purposes: "baseURL/prometheus/api/v1" ["/prometheus/api/v1"]
func (c *HAClient) RequestStatus() (*http.Request, error) {
	// baseURL/prometheus/api/v1/status/tsdb
	url, err := url.JoinPath(c.BaseURL, "prometheus", "api", "v1", "status", "tsdb")
	if err != nil {
		return nil, fmt.Errorf("Join URL failed: %v", err)
	}

	result, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	return result, nil
}
func (c *HAClient) RequestLabels() (*http.Request, error) {
	result, err := url.JoinPath(c.BaseURL, "prometheus", "api", "v1", "labels")
	if err != nil {
		return nil, fmt.Errorf("unexpected URL join: %v", err)
	}

	request, err := http.NewRequest("GET", result, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}
	return request, nil
}
func (c *HAClient) RequestLabelValue(label string) (*http.Request, error) {
	if label == "" {
		return nil, fmt.Errorf("invalid input: empty label.")
	}
	result, err := url.JoinPath(c.BaseURL, "prometheus", "api", "v1", "label", label, "values")
	if err != nil {
		return nil, fmt.Errorf("unexpected URL join: %v", err)
	}
	request, err := http.NewRequest("GET", result, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}
	return request, nil
}
func (c *HAClient) RequestSeriesCount() (*http.Request, error) {
	result, err := url.JoinPath(c.BaseURL, "prometheus", "api", "v1", "series", "count")
	if err != nil {
		return nil, fmt.Errorf("Join URL failed: %v", err)
	}

	request, err := http.NewRequest("GET", result, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %s %w", result, err)
	}

	return request, nil
}
