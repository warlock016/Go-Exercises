package hello_handler

import (
	"fmt"
	"net/http"
)

// HelloHandler writes "Hello, World!" to the response
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Hello, World!"))
	fmt.Fprint(w, "Hello, World!")
}

// GreetHandler reads "name" from query params and writes "Hello, {name}!"
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", name)
}
