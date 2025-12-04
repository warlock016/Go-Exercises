package http_handler_functions

import "net/http"

// HelloHandler responds with "Hello, World!"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
}

// EchoHandler echoes the "message" query parameter
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
}

// StatusHandler returns a handler that responds with the given status code
func StatusHandler(code int) http.HandlerFunc {
	// TODO(human): Implement
	return nil
}

// JSONHandler responds with a JSON object
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
}

// MethodFilter wraps a handler to only accept a specific HTTP method
func MethodFilter(method string, handler http.HandlerFunc) http.HandlerFunc {
	// TODO(human): Implement
	return nil
}
