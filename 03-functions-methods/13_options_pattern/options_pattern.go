package options_pattern

import "time"

// Server represents a configurable server
type Server struct {
	// TODO(human): Define private fields
}

// Option is a function that configures a Server
type Option func(*Server)

// NewServer creates a new Server with the given options
func NewServer(opts ...Option) *Server {
	// TODO(human): Implement with defaults
	return nil
}

// WithHost sets the server host
func WithHost(host string) Option {
	// TODO(human): Implement
	return nil
}

// WithPort sets the server port
func WithPort(port int) Option {
	// TODO(human): Implement
	return nil
}

// WithTimeout sets the server timeout
func WithTimeout(timeout time.Duration) Option {
	// TODO(human): Implement
	return nil
}

// WithMaxConnections sets the maximum number of connections
func WithMaxConnections(maxConn int) Option {
	// TODO(human): Implement
	return nil
}

// Getter methods
func (s *Server) Host() string {
	// TODO(human): Implement
	return ""
}

func (s *Server) Port() int {
	// TODO(human): Implement
	return 0
}

func (s *Server) Timeout() time.Duration {
	// TODO(human): Implement
	return 0
}

func (s *Server) MaxConnections() int {
	// TODO(human): Implement
	return 0
}
