package http_handler_functions

import (
	"encoding/json"
	"net/http"
	"time"
)

// HelloHandler responds with "Hello, World!"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	// r.Response.Body
	w.Write([]byte("Hello, World!"))
}

// EchoHandler echoes the "message" query parameter
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	message := r.URL.Query().Get("message")
	if message == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing message parameter"))
		return
	}

	w.Write([]byte(message))
}

// StatusHandler returns a handler that responds with the given status code
func StatusHandler(code int) http.HandlerFunc {
	// TODO(human): Implement
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
	}
}

// JSONHandler responds with a JSON object
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	response := map[string]interface{}{
		"message":   "Hello, JSON!",
		"timestamp": time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// MethodFilter wraps a handler to only accept a specific HTTP method
func MethodFilter(method string, handler http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}
