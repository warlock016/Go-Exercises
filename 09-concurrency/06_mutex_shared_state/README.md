# Exercise 06: Mutex and Shared State

**Learning Goal:** Protect shared state from race conditions using sync.Mutex and sync.RWMutex

**Difficulty:** Tier 2 - Application
**Estimated Time:** 35-40 minutes

---

## Problem Description

When multiple goroutines access shared data, and at least one modifies it, you have a **data race**. Data races cause undefined behavior—crashes, corrupted data, or bugs that appear randomly.

Go provides `sync.Mutex` (mutual exclusion) to protect shared state. Only one goroutine can hold the lock at a time. `sync.RWMutex` allows multiple readers OR one writer, optimizing read-heavy workloads.

**Golden Rule:** If you share memory between goroutines, either use channels to coordinate access, or protect it with a mutex.

---

## Function Signatures

```go
// SafeCounter is a concurrent-safe counter
type SafeCounter struct {
    // Define fields
}

func NewSafeCounter() *SafeCounter
func (c *SafeCounter) Inc()
func (c *SafeCounter) Dec()
func (c *SafeCounter) Value() int

// SafeMap is a concurrent-safe string->int map
type SafeMap struct {
    // Define fields
}

func NewSafeMap() *SafeMap
func (m *SafeMap) Set(key string, value int)
func (m *SafeMap) Get(key string) (int, bool)
func (m *SafeMap) Delete(key string)
func (m *SafeMap) Len() int

// SafeCache uses RWMutex for read-heavy workloads
type SafeCache struct {
    // Define fields
}

func NewSafeCache() *SafeCache
func (c *SafeCache) Get(key string) (string, bool)    // Read lock
func (c *SafeCache) Set(key string, value string)     // Write lock
func (c *SafeCache) GetOrSet(key string, value string) string  // Careful with lock upgrade!
```

---

## Examples

### SafeCounter
```go
counter := NewSafeCounter()

var wg sync.WaitGroup
for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        counter.Inc()
    }()
}
wg.Wait()

fmt.Println(counter.Value()) // Always 1000, never less
```

### SafeMap
```go
m := NewSafeMap()

go m.Set("a", 1)
go m.Set("b", 2)
go m.Set("c", 3)

// After synchronization:
v, ok := m.Get("a") // v: 1, ok: true
m.Delete("a")
v, ok = m.Get("a")  // v: 0, ok: false
```

### SafeCache with RWMutex
```go
cache := NewSafeCache()
cache.Set("config", "production")

// Multiple readers can access simultaneously
for i := 0; i < 10; i++ {
    go func() {
        v, _ := cache.Get("config") // Read lock - concurrent OK
        fmt.Println(v)
    }()
}
```

---

## Instructions

1. Implement `SafeCounter` using sync.Mutex to protect an int
2. Implement all SafeCounter methods (Inc, Dec, Value)
3. Implement `SafeMap` using sync.Mutex to protect a map
4. Implement all SafeMap methods (Set, Get, Delete, Len)
5. Implement `SafeCache` using sync.RWMutex for better read performance
6. Implement GetOrSet carefully—this is tricky with RWMutex!
7. Run tests with `go test -v -race` (race detector!)

---

## Hints

### Basic
- `sync.Mutex`: `Lock()` and `Unlock()`
- `sync.RWMutex`: `RLock()`/`RUnlock()` for reads, `Lock()`/`Unlock()` for writes
- Always use `defer mu.Unlock()` right after `Lock()` to prevent deadlocks
- The zero value of Mutex is ready to use (no initialization needed)

### Intermediate
- For SafeMap, remember maps are reference types—initialize with `make()`
- For SafeCache.GetOrSet, you cannot "upgrade" RLock to Lock—you must release and reacquire
- Pattern: check-with-read-lock, release, lock-for-write, check-again, write
- Never copy a Mutex after first use (use pointers)

### Solution Pattern
```go
type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}
```

---

## Think About

1. Why use `defer mu.Unlock()` instead of calling Unlock() at the end?
2. When would RWMutex perform worse than Mutex?
3. What happens if you forget to unlock? What if you unlock twice?
4. Why can't you upgrade a read lock to a write lock?

---

## What This Teaches

- **sync.Mutex** - Basic mutual exclusion for any shared state
- **sync.RWMutex** - Optimized locking for read-heavy workloads
- **defer unlock pattern** - Guarantees unlock even on panic
- **Race detector** - Using `go test -race` to find data races
- **Lock granularity** - Holding locks only as long as necessary
- **Deadlock prevention** - Consistent lock ordering and avoiding nested locks
