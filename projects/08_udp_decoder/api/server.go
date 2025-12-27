package api

import (
	"context"
	"net/http"

	"github.com/warlock16/udp_decoder/device"
)

// Server handles REST API requests
type Server struct {
	addr     string
	registry *device.Registry
	mux      *http.ServeMux
	server   *http.Server
}

// NewServer creates a new API server
func NewServer(addr string, registry *device.Registry) *Server {
	// TODO(human): Create server with mux, set up routes
	return nil
}

// Routes configures all API endpoints
func (s *Server) Routes() {
	// TODO(human): Register all handlers with mux
}

// Start runs the HTTP server (blocks until context cancelled)
func (s *Server) Start(ctx context.Context) error {
	// TODO(human): Start server, handle graceful shutdown
	return nil
}

// Shutdown gracefully stops the server
func (s *Server) Shutdown(ctx context.Context) error {
	// TODO(human): Call server.Shutdown
	return nil
}
