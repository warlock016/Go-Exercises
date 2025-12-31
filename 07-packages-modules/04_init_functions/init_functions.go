package initfuncs

import "strings"

// Handler is a function that processes a string
type Handler func(string) string

// TODO(human): Declare package-level handlers map and initialized flag
var handlers = make(map[string]Handler)
var initialized bool

// TODO(human): Implement init() function
func init() {
	Register("uppercase", func(s string) string {
		return strings.ToUpper(s)
	})

	Register("lowercase", func(s string) string {
		return strings.ToLower(s)
	})
	initialized = true
}

// Register adds a handler to the registry
func Register(name string, handler Handler) {
	// TODO(human): Implement
	handlers[name] = handler
}

// GetHandler retrieves a handler by name
func GetHandler(name string) (Handler, bool) {
	// TODO(human): Implement
	if handler, ok := handlers[name]; ok {
		return handler, true
	}
	return nil, false
}

// IsInitialized returns true if the package has been initialized
func IsInitialized() bool {
	// TODO(human): Implement
	return initialized
}
