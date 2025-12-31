package json_response

import (
	"encoding/json"
	"net/http"
)

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

// Stats represents user statistics
type Stats struct {
	TotalUsers    int `json:"total_users"`
	ActiveUsers   int `json:"active_users"`
	InactiveUsers int `json:"inactive_users"`
}

// GetUserHandler returns a single hardcoded user
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	res := User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		IsActive: true,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// ListUsersHandler returns array of 3 hardcoded users
func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	res := []User{{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		IsActive: true,
	}, {
		ID:       2,
		Name:     "Bob",
		Email:    "bob@example.com",
		IsActive: false,
	}, {
		ID:       3,
		Name:     "Charlie",
		Email:    "charlie@example.com",
		IsActive: true,
	}}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// StatsHandler returns statistics about users
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	res := Stats{
		TotalUsers:    3,
		ActiveUsers:   2,
		InactiveUsers: 1,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
