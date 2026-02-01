package concurrent_processor

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Job represents a unit of work
type Job struct {
	ID       string
	Payload  any
	Priority int
	Retries  int
	MaxRetry int
}

// JobResult contains the result of processing
type JobResult struct {
	JobID    string
	Success  bool
	Error    error
	Duration time.Duration
}

// Metrics tracks processor statistics
type Metrics struct {
	Submitted  int64
	Completed  int64
	Failed     int64
	Retried    int64
	InFlight   int64
	AvgLatency time.Duration
}

// ProcessorConfig configures the processor
type ProcessorConfig struct {
	Workers       int
	QueueSize     int
	RateLimit     int
	RetryDelay    time.Duration
	MaxRetryDelay time.Duration
}

// Processor is the main concurrent job processor
type Processor struct {
	// TODO(human): Define fields
	config  ProcessorConfig
	metrics Metrics
	queue   chan Job
	result  chan JobResult
	handler func(context.Context, Job) error
	mu      sync.Mutex // could also be RWMutex to provide read-heavy performance gains
	wg      sync.WaitGroup
	start   chan struct{}
	stop    chan struct{}
}

// NewProcessor creates a new processor with the given config and handler
func NewProcessor(config ProcessorConfig, handler func(context.Context, Job) error) *Processor {
	// TODO(human): Implement
	return &Processor{
		config:  config,
		queue:   make(chan Job, config.QueueSize),
		result:  make(chan JobResult, config.QueueSize),
		handler: handler,
		start:   make(chan struct{}),
		stop:    make(chan struct{}),
	}
}

// Submit adds a job to the processor
func (p *Processor) Submit(job Job) error {
	// Check if shutdown has started (non-blocking)
	select {
	case <-p.stop:
		return errors.New("processor stopped")
	default:
	}

	// Try to submit (queue may still close between check and send, so we check stop again)
	select {
	case <-p.stop:
		return errors.New("processor stopped")
	case p.queue <- job:
		p.mu.Lock()
		p.metrics.Submitted++
		p.mu.Unlock()
	}
	return nil
}

// SubmitBatch adds multiple jobs
func (p *Processor) SubmitBatch(jobs []Job) error {
	// TODO(human): Implement

	for _, job := range jobs {
		select {
		case <-p.stop:
			return errors.New("job channel previously closed")
		default:
			p.queue <- job
			p.mu.Lock()
			p.metrics.Submitted++
			p.mu.Unlock()
		}
	}
	return nil
}

// Results returns channel of completed job results
func (p *Processor) Results() <-chan JobResult {
	// TODO(human): Implement
	return p.result
}

// Metrics returns current metrics
func (p *Processor) Metrics() Metrics {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.metrics
}

// executeWithRetry handles ONLY retry logic - completely independent of rate limiting
func (p *Processor) executeWithRetry(ctx context.Context, job Job) error {
	// No retries configured: single attempt
	if job.MaxRetry <= 0 {
		return p.handler(ctx, job)
	}

	var lastErr error
	for attempt := range job.MaxRetry {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		err := p.handler(ctx, job)
		if err == nil {
			return nil // Success!
		}
		lastErr = err

		// Don't retry after last attempt
		if attempt+1 >= job.MaxRetry {
			break
		}

		// Record retry metric
		p.mu.Lock()
		p.metrics.Retried++
		p.mu.Unlock()

		// Exponential backoff
		backoff := min(p.config.RetryDelay*time.Duration(1<<attempt), p.config.MaxRetryDelay)
		timer := time.NewTimer(backoff)
		select {
		case <-timer.C:
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		}
	}
	return lastErr
}

// Start begins processing
func (p *Processor) Start(ctx context.Context) error {
	// Create shared rate limiter (if configured)
	var rateLimitCh <-chan time.Time
	var rateLimitMu sync.Mutex
	if p.config.RateLimit > 0 {
		ticker := time.NewTicker(time.Second / time.Duration(p.config.RateLimit))
		rateLimitCh = ticker.C
		// Note: ticker is not explicitly stopped - it lives for processor lifetime
	}

	for range p.config.Workers {
		p.wg.Go(func() {
			for {
				job, ok := <-p.queue
				if !ok {
					return
				}

				// Rate limiting: wait for token BEFORE processing
				if rateLimitCh != nil {
					rateLimitMu.Lock()
					<-rateLimitCh
					rateLimitMu.Unlock()
				}

				res := JobResult{
					JobID: job.ID,
				}

				p.mu.Lock()
				p.metrics.InFlight++
				p.mu.Unlock()

				start := time.Now()
				err := p.executeWithRetry(ctx, job)
				dur := time.Since(start)
				res.Duration = dur

				p.mu.Lock()
				p.metrics.InFlight--
				p.mu.Unlock()

				if err != nil {
					p.mu.Lock()
					p.metrics.Failed++
					p.mu.Unlock()
					res.Error = err
				} else {
					p.mu.Lock()
					p.metrics.Completed++
					p.mu.Unlock()
					res.Success = true
				}

				select {
				case p.result <- res:
				default:
				}
			}
		})
	}

	close(p.start)
	return nil
}

// Shutdown gracefully stops the processor
func (p *Processor) Shutdown(ctx context.Context) error {
	<-p.start
	close(p.stop)  // Signal no more submissions FIRST
	close(p.queue) // Then signal workers to drain

	// Wait for workers with timeout
	done := make(chan struct{})
	go func() {
		p.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Wait blocks until all submitted jobs are processed
func (p *Processor) Wait() error {
	// TODO(human): Implement

	<-p.start
	close(p.queue)
	p.wg.Wait()
	return nil
}
