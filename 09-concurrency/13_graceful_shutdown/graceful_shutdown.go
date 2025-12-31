package graceful_shutdown

import (
	"context"
	"os"
)

// Shutdownable is implemented by components that can shutdown
type Shutdownable interface {
	Shutdown(ctx context.Context) error
}

// Server represents a long-running service
type Server struct {
	// TODO(human): Define fields
}

// NewServer creates a new server
func NewServer(handler func(request int) error) *Server {
	// TODO(human): Implement
	return nil
}

// Start begins accepting requests on the channel
func (s *Server) Start(requests <-chan int) {
	// TODO(human): Implement
}

// Shutdown gracefully stops the server with timeout
func (s *Server) Shutdown(ctx context.Context) error {
	// TODO(human): Implement
	return nil
}

// SignalHandler sets up OS signal handling and returns shutdown channel
func SignalHandler(signals ...os.Signal) <-chan struct{} {
	// TODO(human): Implement
	return nil
}

// GracefulShutdown coordinates shutdown of multiple components
func GracefulShutdown(ctx context.Context, components ...Shutdownable) error {
	// TODO(human): Implement
	return nil
}

// Drainer drains a channel with timeout
func Drainer[T any](ctx context.Context, ch <-chan T, process func(T) error) error {
	// TODO(human): Implement
	return nil
}

// Coordinator manages graceful shutdown of multiple workers
type Coordinator struct {
	// TODO(human): Define fields
}

// NewCoordinator creates a new Coordinator
func NewCoordinator() *Coordinator {
	// TODO(human): Implement
	return nil
}

// Add registers a worker with the coordinator
func (c *Coordinator) Add(name string, worker func(ctx context.Context) error) {
	// TODO(human): Implement
}

// Run starts all workers and waits for completion or cancellation
func (c *Coordinator) Run(ctx context.Context) error {
	// TODO(human): Implement
	return nil
}

// Shutdown gracefully stops all workers
func (c *Coordinator) Shutdown(ctx context.Context) error {
	// TODO(human): Implement
	return nil
}
