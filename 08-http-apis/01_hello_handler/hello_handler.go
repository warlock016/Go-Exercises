package hello_handler

import "net/http"

// HelloHandler writes "Hello, World!" to the response
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
}

// GreetHandler reads "name" from query params and writes "Hello, {name}!"
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
}
