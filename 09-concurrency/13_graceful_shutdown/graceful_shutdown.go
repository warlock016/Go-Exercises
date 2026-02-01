package graceful_shutdown

import (
	"context"
	"os"
	"os/signal"
	"sync"
)

// Shutdownable is implemented by components that can shutdown
type Shutdownable interface {
	Shutdown(ctx context.Context) error
}

// Server represents a long-running service
type Server struct {
	// TODO(human): Define fields
	wg      sync.WaitGroup
	handler func(int) error
	start   chan struct{}
	stop    chan struct{}
}

// NewServer creates a new server
func NewServer(handler func(request int) error) *Server {
	// TODO(human): Implement
	return &Server{
		handler: handler,
		start:   make(chan struct{}),
		stop:    make(chan struct{}),
	}
}

// Start begins accepting requests on the channel
func (s *Server) Start(requests <-chan int) {
	// TODO(human): Implement
	go func() {
		defer close(s.start)
		for {
			select {
			case <-s.stop:
				return
			case v, ok := <-requests:
				if !ok {
					return
				}
				s.wg.Go(func() {
					err := s.handler(v)
					if err != nil {
						return
					}
				})
			}
		}
	}()
}

// Shutdown gracefully stops the server with timeout
func (s *Server) Shutdown(ctx context.Context) error {
	// TODO(human): Implement

	close(s.stop)
	done := make(chan struct{})

	go func() {
		<-s.start
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done: // this case would always trigger unless the ctx.Done signal is received before
		return nil
	}
}

// SignalHandler sets up OS signal handling and returns shutdown channel
func SignalHandler(signals ...os.Signal) <-chan struct{} {
	// TODO(human): Implement
	sigs := make(chan os.Signal, len(signals))
	result := make(chan struct{})

	for _, s := range signals {
		go func() {
			signal.Notify(sigs, s)
		}()
	}

	go func() {
		// sig := <-sigs
		<-sigs
		close(result)
	}()
	return result
}

// GracefulShutdown coordinates shutdown of multiple components
func GracefulShutdown(ctx context.Context, components ...Shutdownable) error {
	// TODO(human): Implement
	for _, c := range components {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := c.Shutdown(ctx); err != nil {
				return err // this would cause an early return, missing all subsequent shutdown functions
			} // return err only if !nil, else keep looping
		}
	}
	return nil
}

// Drainer drains a channel with timeout
func Drainer[T any](ctx context.Context, ch <-chan T, process func(T) error) error {
	// TODO(human): Implement
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case v, ok := <-ch:
			if !ok {
				return nil
			}
			if err := process(v); err != nil {
				return err
			}
		}
	}
	// return nil
}

// Coordinator manages graceful shutdown of multiple workers
type Coordinator struct {
	// TODO(human): Define fields
	worker map[string]func(context.Context) error
	cancel context.CancelFunc
	set    chan struct{}
}

// NewCoordinator creates a new Coordinator
func NewCoordinator() *Coordinator {
	// TODO(human): Implement
	return &Coordinator{
		worker: make(map[string]func(context.Context) error),
		set:    make(chan struct{}),
	}
}

// Add registers a worker with the coordinator
func (c *Coordinator) Add(name string, worker func(ctx context.Context) error) {
	// TODO(human): Implement
	if _, ok := c.worker[name]; !ok {
		c.worker[name] = worker
	}
}

// Run starts all workers and waits for completion or cancellation
func (c *Coordinator) Run(ctx context.Context) error {
	// TODO(human): Implement

	errCh := make(chan error, len(c.worker))

	wrkCtx, cancel := context.WithCancel(ctx)
	c.cancel = cancel
	close(c.set)
	defer cancel()

	for _, fn := range c.worker {
		go func(f func(context.Context) error) {
			errCh <- f(wrkCtx)
		}(fn)
	}

	var firstErr error
	for range c.worker {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case e := <-errCh:
			if e != nil && firstErr == nil {
				firstErr = e
			}
		}
	}

	return firstErr
}

// Shutdown gracefully stops all workers
func (c *Coordinator) Shutdown(ctx context.Context) error {
	// TODO(human): Implement

	// this forces Shitdown to wait for c.cancel to be assigned by Run.
	// Else we get a race condition, where we try calling cancel before it has been assigned in the struct
	<-c.set
	c.cancel()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
	// return nil
}
