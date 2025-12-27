package http_client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIClient_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test" {
			t.Errorf("Get() path = %q, want /test", r.URL.Path)
		}
		w.Write([]byte("test response"))
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	body, err := client.Get("/test")

	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if string(body) != "test response" {
		t.Errorf("Get() body = %q, want %q", string(body), "test response")
	}
}

func TestAPIClient_GetJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "hello"})
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)

	var result map[string]string
	err := client.GetJSON("/test", &result)

	if err != nil {
		t.Fatalf("GetJSON() error = %v", err)
	}

	if result["message"] != "hello" {
		t.Errorf("GetJSON() message = %q, want %q", result["message"], "hello")
	}
}

func TestAPIClient_Post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Post() method = %q, want POST", r.Method)
		}

		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)

		if body["name"] != "test" {
			t.Errorf("Post() body[name] = %q, want %q", body["name"], "test")
		}

		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	err := client.Post("/test", map[string]string{"name": "test"})

	if err != nil {
		t.Fatalf("Post() error = %v", err)
	}
}

func TestAPIClient_ErrorHandling(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	_, err := client.Get("/not-found")

	if err == nil {
		t.Errorf("Get() with 404 should return error")
	}
}
