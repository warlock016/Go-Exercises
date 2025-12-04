package api_client_builder

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient("https://api.example.com")

	if client == nil {
		t.Fatal("NewClient returned nil")
	}
}

func TestClientWithOptions(t *testing.T) {
	client := NewClient("https://api.example.com",
		WithTimeout(5*time.Second),
		WithHeader("User-Agent", "TestClient/1.0"),
	)

	if client == nil {
		t.Fatal("NewClient returned nil")
	}
}

func TestGetRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/users" {
			t.Errorf("Path = %q, want /users", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	resp, err := client.Get("/users")

	if err != nil {
		t.Fatalf("Get returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status = %d, want 200", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "success" {
		t.Errorf("Body = %q, want %q", string(body), "success")
	}
}

func TestPostRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Method = %q, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) != "test data" {
			t.Errorf("Body = %q, want %q", string(body), "test data")
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	resp, err := client.Post("/data", strings.NewReader("test data"))

	if err != nil {
		t.Fatalf("Post returned error: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Status = %d, want 201", resp.StatusCode)
	}
}

func TestRequestBuilder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", r.Header.Get("Content-Type"))
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Accept = %q, want application/json", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	resp, err := client.NewRequest("GET", "/api").
		WithHeader("Content-Type", "application/json").
		WithHeader("Accept", "application/json").
		Do()

	if err != nil {
		t.Fatalf("Request returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status = %d, want 200", resp.StatusCode)
	}
}

func TestClientHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") != "TestClient/1.0" {
			t.Errorf("User-Agent = %q, want TestClient/1.0", r.Header.Get("User-Agent"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL,
		WithHeader("User-Agent", "TestClient/1.0"),
	)

	resp, err := client.Get("/test")

	if err != nil {
		t.Fatalf("Get returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status = %d, want 200", resp.StatusCode)
	}
}

func TestRequestHeadersOverrideClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		if userAgent != "RequestClient/2.0" {
			t.Errorf("User-Agent = %q, want RequestClient/2.0", userAgent)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL,
		WithHeader("User-Agent", "TestClient/1.0"),
	)

	resp, err := client.NewRequest("GET", "/test").
		WithHeader("User-Agent", "RequestClient/2.0").
		Do()

	if err != nil {
		t.Fatalf("Request returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status = %d, want 200", resp.StatusCode)
	}
}

func TestBaseURLHandling(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
		path    string
	}{
		{"no trailing slash", "https://api.example.com", "/users"},
		{"with trailing slash", "https://api.example.com/", "/users"},
		{"path without leading slash", "https://api.example.com", "users"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			// This test verifies URL construction doesn't fail
			client := NewClient(server.URL)
			_, err := client.Get(tt.path)

			if err != nil {
				t.Errorf("Request failed: %v", err)
			}
		})
	}
}

func TestRequestWithBody(t *testing.T) {
	expectedBody := `{"key":"value"}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		if string(body) != expectedBody {
			t.Errorf("Body = %q, want %q", string(body), expectedBody)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	resp, err := client.NewRequest("POST", "/data").
		WithBody(strings.NewReader(expectedBody)).
		Do()

	if err != nil {
		t.Fatalf("Request returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status = %d, want 200", resp.StatusCode)
	}
}

func TestComplexRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check method
		if r.Method != "PUT" {
			t.Errorf("Method = %q, want PUT", r.Method)
		}

		// Check headers
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type header missing")
		}
		if r.Header.Get("Authorization") != "Bearer token123" {
			t.Errorf("Authorization header missing")
		}

		// Check body
		body, _ := io.ReadAll(r.Body)
		if string(body) != `{"update":"data"}` {
			t.Errorf("Body incorrect")
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("updated"))
	}))
	defer server.Close()

	client := NewClient(server.URL,
		WithHeader("User-Agent", "TestClient/1.0"),
		WithTimeout(10*time.Second),
	)

	resp, err := client.NewRequest("PUT", "/resource/123").
		WithHeader("Content-Type", "application/json").
		WithHeader("Authorization", "Bearer token123").
		WithBody(strings.NewReader(`{"update":"data"}`)).
		Do()

	if err != nil {
		t.Fatalf("Request returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status = %d, want 200", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "updated" {
		t.Errorf("Body = %q, want %q", string(body), "updated")
	}
}

func TestMultipleRequests(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL)

	// Make multiple requests
	for i := 0; i < 3; i++ {
		resp, err := client.Get("/test")
		if err != nil {
			t.Fatalf("Request %d failed: %v", i, err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Request %d status = %d, want 200", i, resp.StatusCode)
		}
	}

	if requestCount != 3 {
		t.Errorf("Expected 3 requests, got %d", requestCount)
	}
}
