package json_response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	w := httptest.NewRecorder()

	GetUserHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("GetUserHandler() status = %d, want %d", w.Code, http.StatusOK)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Content-Type = %q, want %q", contentType, "application/json")
	}

	var got User
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	want := User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		IsActive: true,
	}

	if got != want {
		t.Errorf("GetUserHandler() = %+v, want %+v", got, want)
	}
}

func TestListUsersHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	ListUsersHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ListUsersHandler() status = %d, want %d", w.Code, http.StatusOK)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Content-Type = %q, want %q", contentType, "application/json")
	}

	var got []User
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	want := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", IsActive: true},
		{ID: 2, Name: "Bob", Email: "bob@example.com", IsActive: false},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com", IsActive: true},
	}

	if len(got) != len(want) {
		t.Fatalf("ListUsersHandler() returned %d users, want %d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("ListUsersHandler() user[%d] = %+v, want %+v", i, got[i], want[i])
		}
	}
}

func TestStatsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	w := httptest.NewRecorder()

	StatsHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("StatsHandler() status = %d, want %d", w.Code, http.StatusOK)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Content-Type = %q, want %q", contentType, "application/json")
	}

	var got Stats
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	want := Stats{
		TotalUsers:    3,
		ActiveUsers:   2,
		InactiveUsers: 1,
	}

	if got != want {
		t.Errorf("StatsHandler() = %+v, want %+v", got, want)
	}
}

// Test that JSON field names are correct (snake_case)
func TestJSONFieldNames(t *testing.T) {
	user := User{
		ID:       1,
		Name:     "Test",
		Email:    "test@example.com",
		IsActive: true,
	}

	data, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("failed to marshal user: %v", err)
	}

	// Parse back as map to check field names
	var fields map[string]interface{}
	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	expectedFields := []string{"id", "name", "email", "is_active"}
	for _, field := range expectedFields {
		if _, exists := fields[field]; !exists {
			t.Errorf("JSON missing field %q, got fields: %v", field, fields)
		}
	}
}
