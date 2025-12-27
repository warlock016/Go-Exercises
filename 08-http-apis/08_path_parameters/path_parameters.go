package path_parameters

import "net/http"

// User represents a user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PathParams represents extracted path parameters
type PathParams struct {
	Org     string `json:"org"`
	Project string `json:"project"`
}

// ParsePathParam extracts a parameter from a URL path
func ParsePathParam(pattern, path, paramName string) string {
	// TODO(human): Implement
	return ""
}

// UserAPIHandler handles /users and /users/{id}
func UserAPIHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
}

// NestedPathHandler handles /orgs/{org}/projects/{project}
func NestedPathHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
}
