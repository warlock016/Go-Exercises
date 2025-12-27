package json_response

import "net/http"

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
}

// ListUsersHandler returns array of 3 hardcoded users
func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
}

// StatsHandler returns statistics about users
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
}
