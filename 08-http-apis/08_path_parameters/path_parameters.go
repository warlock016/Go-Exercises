package path_parameters

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

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
	elements := strings.Split(path, "/")
	for i, v := range elements {
		elements[i] = strings.TrimSpace(v)
	}

	template := strings.Split(pattern, "/")
	for i, v := range template {
		template[i] = strings.TrimSpace(v)
	}

	if len(elements) != len(template) {
		return ""
	}

	for i := range elements {
		if template[i] == elements[i] {
			continue
		}
		if strings.ContainsFunc(template[i], func(r rune) bool {
			switch r {
			case '{', '}':
				return true
			default:
				return false
			}
		}) {
			template[i] = strings.Trim(template[i], "{")
			template[i] = strings.Trim(template[i], "}")
			if paramName == template[i] {
				return elements[i]
			}
		}
	}
	return ""
}

// UserAPIHandler handles /users and /users/{id}
func UserAPIHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	var result string
	w.Header().Set("Content-Type", "application/json")

	switch r.URL.Path {
	case "/users":
		resp := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return

	default:
		resp := User{}
		result = ParsePathParam("/users/{id}", r.URL.Path, "id")
		val, err := strconv.ParseInt(result, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if val != 1 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// fmt.Printf("%s: %s: %d\n", r.URL.Path, result, val)

		resp.ID = int(val)
		resp.Name = "Alice"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}

}

// NestedPathHandler handles /orgs/{org}/projects/{project}
func NestedPathHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement

	params := []string{"org", "project"}
	resp := PathParams{}

	for _, param := range params {
		res := ParsePathParam("/orgs/{org}/projects/{project}", r.URL.Path, param)

		switch param {
		case "org":
			resp.Org = res
		case "project":
			resp.Project = res
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
