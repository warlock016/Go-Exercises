# Exercise 14: Semaphore and Sync Primitives

**Learning Goal:** Build and use synchronization primitives beyond basic mutexes

**Difficulty:** Tier 4 - Mastery
**Estimated Time:** 50-55 minutes

---

## Problem Description

Beyond `sync.Mutex`, Go provides (and you can build) more sophisticated synchronization primitives:

- **Semaphore** - Limit concurrent access to N resources
- **sync.Once** - Execute initialization exactly once
- **sync.Cond** - Wait for/signal conditions
- **Barrier** - Wait for N goroutines to reach a point

These primitives are building blocks for complex concurrent systems.

---

## Function Signatures

```go
// Semaphore limits concurrent access
type Semaphore struct {
    // Define fields
}

func NewSemaphore(permits int) *Semaphore
func (s *Semaphore) Acquire()
func (s *Semaphore) TryAcquire() bool
func (s *Semaphore) Release()
func (s *Semaphore) Available() int

// WeightedSemaphore allows different weights per acquisition
type WeightedSemaphore struct {
    // Define fields
}

func NewWeightedSemaphore(maxWeight int64) *WeightedSemaphore
func (ws *WeightedSemaphore) Acquire(ctx context.Context, weight int64) error
func (ws *WeightedSemaphore) TryAcquire(weight int64) bool
func (ws *WeightedSemaphore) Release(weight int64)

// Barrier waits for N goroutines to arrive before proceeding
type Barrier struct {
    // Define fields
}

func NewBarrier(count int) *Barrier
func (b *Barrier) Wait() // Blocks until all participants arrive

// SingleFlight deduplicates concurrent calls for same key
type SingleFlight struct {
    // Define fields
}

func NewSingleFlight() *SingleFlight
func (sf *SingleFlight) Do(key string, fn func() (interface{}, error)) (interface{}, error)

// LazyInit provides lazy initialization with sync.Once semantics
type LazyInit[T any] struct {
    // Define fields
}

func (l *LazyInit[T]) Get(init func() T) T
func (l *LazyInit[T]) IsInitialized() bool
```

---

## Examples

### Semaphore
```go
sem := NewSemaphore(3) // Allow 3 concurrent

for i := 0; i < 10; i++ {
    go func(id int) {
        sem.Acquire()
        defer sem.Release()
        // Only 3 goroutines can be here at once
        doWork()
    }(i)
}
```

### Barrier
```go
barrier := NewBarrier(3)

for i := 0; i < 3; i++ {
    go func(id int) {
        fmt.Printf("Worker %d preparing...\n", id)
        time.Sleep(time.Duration(id) * 100 * time.Millisecond)

        barrier.Wait() // All three wait here

        fmt.Printf("Worker %d proceeding!\n", id)
    }(i)
}
// All workers print "preparing" before any prints "proceeding"
```

### SingleFlight
```go
sf := NewSingleFlight()

// Multiple concurrent calls with same key
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        result, _ := sf.Do("expensive-key", func() (interface{}, error) {
            time.Sleep(100 * time.Millisecond)
            return expensiveComputation(), nil
        })
        // All 10 get same result, but computation runs only once
        use(result)
    }()
}
```

### LazyInit
```go
var config LazyInit[*Config]

// Multiple goroutines can call Get - init runs once
cfg := config.Get(func() *Config {
    return loadConfig()
})
```

---

## Instructions

1. Implement `Semaphore` using buffered channel or mutex+condition
2. Implement `WeightedSemaphore` for variable-weight acquisitions
3. Implement `Barrier` using mutex and condition variable
4. Implement `SingleFlight` to deduplicate concurrent calls
5. Implement `LazyInit` using sync.Once pattern
6. Handle edge cases: panic in init, release without acquire
7. Run tests with `go test -v -race`

---

## Hints

### Basic
- Simple semaphore: buffered channel of size N, send to acquire, receive to release
- `sync.Once` ensures `Do()` is called exactly once
- `sync.Cond` needs a `Locker` (usually `*sync.Mutex`)
- `sync.Cond.Wait()` atomically unlocks and waits for signal

### Intermediate
- For WeightedSemaphore, track total weight with mutex
- Barrier: count arrivals, last one broadcasts to wake others
- SingleFlight: map of in-flight calls, mutex to protect map
- `sync.Cond.Broadcast()` wakes all waiters, `Signal()` wakes one

### Solution Pattern
```go
type Semaphore struct {
    ch chan struct{}
}

func NewSemaphore(permits int) *Semaphore {
    return &Semaphore{ch: make(chan struct{}, permits)}
}

func (s *Semaphore) Acquire() {
    s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
    <-s.ch
}
```

---

## Think About

1. Why is a channel-based semaphore simpler than mutex-based?
2. When would you need WeightedSemaphore vs regular Semaphore?
3. What happens if a barrier Wait() panics?
4. How does SingleFlight differ from caching?

---

## What This Teaches

- **Semaphore pattern** - Limiting concurrent access to resources
- **sync.Cond** - Condition variables for complex synchronization
- **sync.Once** - One-time initialization pattern
- **Barrier synchronization** - Coordinating multiple goroutines
- **Request deduplication** - Reducing redundant work
- **Building primitives** - Understanding how sync tools work internally
