package path_parameters

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParsePathParam(t *testing.T) {
	tests := []struct {
		pattern   string
		path      string
		paramName string
		want      string
	}{
		{"/users/{id}", "/users/123", "id", "123"},
		{"/users/{id}", "/users/abc", "id", "abc"},
		{"/orgs/{org}/projects/{project}", "/orgs/golang/projects/http", "org", "golang"},
		{"/orgs/{org}/projects/{project}", "/orgs/golang/projects/http", "project", "http"},
		{"/items/{id}", "/items/", "id", ""},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := ParsePathParam(tt.pattern, tt.path, tt.paramName)
			if got != tt.want {
				t.Errorf("ParsePathParam(%q, %q, %q) = %q, want %q",
					tt.pattern, tt.path, tt.paramName, got, tt.want)
			}
		})
	}
}

func TestUserAPIHandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		wantStatus int
		wantCount  int
		wantUser   *User
	}{
		{
			name:       "list all users",
			url:        "/users",
			wantStatus: http.StatusOK,
			wantCount:  2,
		},
		{
			name:       "get user by id",
			url:        "/users/1",
			wantStatus: http.StatusOK,
			wantUser:   &User{ID: 1, Name: "Alice"},
		},
		{
			name:       "user not found",
			url:        "/users/999",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()

			UserAPIHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("UserAPIHandler() status = %d, want %d", w.Code, tt.wantStatus)
			}

			if tt.wantCount > 0 {
				var users []User
				if err := json.NewDecoder(w.Body).Decode(&users); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if len(users) != tt.wantCount {
					t.Errorf("UserAPIHandler() returned %d users, want %d", len(users), tt.wantCount)
				}
			}

			if tt.wantUser != nil {
				var got User
				if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if got != *tt.wantUser {
					t.Errorf("UserAPIHandler() = %+v, want %+v", got, *tt.wantUser)
				}
			}
		})
	}
}

func TestNestedPathHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/orgs/golang/projects/http", nil)
	w := httptest.NewRecorder()

	NestedPathHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("NestedPathHandler() status = %d, want %d", w.Code, http.StatusOK)
	}

	var got PathParams
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	want := PathParams{Org: "golang", Project: "http"}
	if got != want {
		t.Errorf("NestedPathHandler() = %+v, want %+v", got, want)
	}
}
