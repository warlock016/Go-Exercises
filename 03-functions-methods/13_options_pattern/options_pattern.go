package options_pattern

import "time"

// Server represents a configurable server
type Server struct {
	// TODO(human): Define private fields
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

// Option is a function that configures a Server
type Option func(*Server)

// NewServer creates a new Server with the given options
func NewServer(opts ...Option) *Server {
	// TODO(human): Implement with defaults
	result := &Server{
		host:    "localhost",
		port:    8080,
		timeout: 30 * time.Second,
		maxConn: 100,
	}

	for _, opt := range opts {
		opt(result)
	}
	return result
}

// WithHost sets the server host
func WithHost(host string) Option {
	// TODO(human): Implement
	return func(s *Server) {
		s.host = host
	}
}

// WithPort sets the server port
func WithPort(port int) Option {
	// TODO(human): Implement
	return func(s *Server) {
		s.port = port
	}
}

// WithTimeout sets the server timeout
func WithTimeout(timeout time.Duration) Option {
	// TODO(human): Implement
	return func(s *Server) {
		s.timeout = timeout
	}
}

// WithMaxConnections sets the maximum number of connections
func WithMaxConnections(maxConn int) Option {
	// TODO(human): Implement
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

// Getter methods
func (s *Server) Host() string {
	// TODO(human): Implement
	return s.host
}

func (s *Server) Port() int {
	// TODO(human): Implement
	return s.port
}

func (s *Server) Timeout() time.Duration {
	// TODO(human): Implement
	return s.timeout
}

func (s *Server) MaxConnections() int {
	// TODO(human): Implement
	return s.maxConn
}
