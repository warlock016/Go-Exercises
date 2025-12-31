package status_codes

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Resource represents a resource with an ID
type Resource struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}

// HealthStatus represents the health check status
type HealthStatus struct {
	Status string `json:"status"`
}

var resources = make(map[int]Resource)

// ResourceHandler handles CRUD operations on resources
func ResourceHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		resources[42] = Resource{ID: 42, Data: "created"}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resources[42])
		return
	}
	// TODO(human): Implement
	id := r.URL.Query().Get("id")
	intId, _ := strconv.Atoi(id)
	val, ok := resources[intId]

	if id == "" || !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("resource not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {

	case "GET":
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(val)
		return

	case "PUT":
		val.Data = "updated 5"
		resources[intId] = val
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(val)
		return

	case "DELETE":
		delete(resources, intId)
		w.WriteHeader(http.StatusNoContent)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

// HealthHandler returns health status based on query parameter
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	w.Header().Set("Content-Type", "application/json")
	health := r.URL.Query().Get("healthy")
	state, err := strconv.ParseBool(health)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(HealthStatus{Status: "unhealthy"})
		return
	}

	switch state {
	case true:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HealthStatus{Status: "healthy"})
		return
	case false:
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(HealthStatus{Status: "unhealthy"})
		return
	}
}

// RedirectHandler redirects to a new location
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	w.Header().Set("Location", "/new-location")
	w.WriteHeader(http.StatusMovedPermanently)
}
