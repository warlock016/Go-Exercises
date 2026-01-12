# Mutex Guide

A comprehensive reference for Go's synchronization primitives: `sync.Mutex` and `sync.RWMutex`.

---

## Table of Contents

1. [Overview](#overview)
2. [sync.Mutex — Exclusive Lock](#syncmutex--exclusive-lock)
3. [sync.RWMutex — Reader/Writer Lock](#syncrwmutex--readerwriter-lock)
4. [When to Use Each](#when-to-use-each)
5. [Common Patterns](#common-patterns)
6. [Gotchas and Pitfalls](#gotchas-and-pitfalls)
7. [Performance Considerations](#performance-considerations)
8. [Quick Reference](#quick-reference)

---

## Overview

Go provides two mutex types for protecting shared state:

| Type | Use Case | Concurrent Access |
|------|----------|-------------------|
| `sync.Mutex` | General purpose | One goroutine at a time |
| `sync.RWMutex` | Read-heavy workloads | Multiple readers OR one writer |

**The Cardinal Rule:**
- `Lock()` / `Unlock()` = "I might write" (exclusive access)
- `RLock()` / `RUnlock()` = "I promise to only read" (shared access)

---

## sync.Mutex — Exclusive Lock

### Basic Usage

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

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}
```

### Behavior

```
Goroutine 1: [====LOCK====]
Goroutine 2:               [====LOCK====]
Goroutine 3:                              [====LOCK====]

All operations serialized, one at a time
```

- Only ONE goroutine can hold the lock
- All others block until it's released
- No distinction between reads and writes

### When to Use Mutex

- Counters with frequent updates
- Maps with mixed read/write patterns
- Any shared state where writes are common
- Simple cases where RWMutex overhead isn't worth it

---

## sync.RWMutex — Reader/Writer Lock

### Basic Usage

```go
type SafeCache struct {
    mu    sync.RWMutex
    cache map[string]string
}

// Read operation - allows concurrent readers
func (c *SafeCache) Get(key string) (string, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    v, ok := c.cache[key]
    return v, ok
}

// Write operation - exclusive access
func (c *SafeCache) Set(key string, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.cache[key] = value
}
```

### Behavior

```
Goroutine 1 (read):  [==RLOCK==]        [==RLOCK==]
Goroutine 2 (read):  [==RLOCK==]        [==RLOCK==]
Goroutine 3 (read):  [==RLOCK==]        [==RLOCK==]
Goroutine 4 (write):            [LOCK]             [LOCK]
                     ↑          ↑
                     Readers    Writer waits for readers,
                     concurrent then gets exclusive access
```

### Lock Compatibility Matrix

| Holding | Lock() | RLock() |
|---------|--------|---------|
| Nothing | ✅ Granted | ✅ Granted |
| RLock (1+) | ❌ Waits | ✅ Granted |
| Lock | ❌ Waits | ❌ Waits |

### When to Use RWMutex

- Caches (read 100x, write 1x)
- Configuration that rarely changes
- Read-heavy data stores
- When reads outnumber writes by 10:1 or more

---

## When to Use Each

### Decision Flowchart

```
Is shared state accessed concurrently?
├── No → No mutex needed
└── Yes → What's the read/write ratio?
    ├── Mostly writes (>30% writes) → sync.Mutex
    ├── Mixed → sync.Mutex (simpler)
    └── Mostly reads (<10% writes) → sync.RWMutex
```

### Comparison Table

| Scenario | Recommendation | Why |
|----------|----------------|-----|
| Counter with Inc/Dec | `Mutex` | All operations are writes |
| Cache with Get/Set | `RWMutex` | Gets are reads, Sets are rare |
| Map with frequent updates | `Mutex` | Writes too common for RWMutex benefit |
| Config loaded once | `RWMutex` | Write once, read forever |
| Session store | `RWMutex` | Mostly checking if session exists |

### Performance Rule of Thumb

```
RWMutex overhead ≈ 2x Mutex for Lock/Unlock
RWMutex benefit = concurrent reads

Break-even: ~5-10 concurrent readers per write
```

If you have fewer concurrent readers, Mutex might actually be faster due to lower overhead.

---

## Common Patterns

### Pattern 1: Protect-and-Return

```go
func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value  // Safe: defer runs after return value is computed
}
```

### Pattern 2: Check-Lock-Check (Double-Checked Locking)

Used when you might need to write but want to avoid write lock when possible:

```go
func (c *SafeCache) GetOrSet(key, value string) string {
    // Fast path: read lock to check
    c.mu.RLock()
    if v, ok := c.cache[key]; ok {
        c.mu.RUnlock()
        return v
    }
    c.mu.RUnlock()

    // Slow path: upgrade to write lock
    c.mu.Lock()
    defer c.mu.Unlock()

    // Double-check: another goroutine might have set it
    if v, ok := c.cache[key]; ok {
        return v
    }

    c.cache[key] = value
    return value
}
```

**Why double-check?** Between `RUnlock()` and `Lock()`, another goroutine could have set the value.

### Pattern 3: Lock Embedding

```go
type SafeMap struct {
    sync.Mutex  // Embedded, not named field
    m map[string]int
}

func (s *SafeMap) Set(key string, value int) {
    s.Lock()  // Direct method call
    defer s.Unlock()
    s.m[key] = value
}
```

### Pattern 4: Copy-on-Read for Complex Data

```go
func (c *SafeCache) GetAll() map[string]string {
    c.mu.RLock()
    defer c.mu.RUnlock()

    // Return a COPY, not the original
    copy := make(map[string]string, len(c.cache))
    for k, v := range c.cache {
        copy[k] = v
    }
    return copy
}
```

### Pattern 5: Atomic with sync/atomic

For simple counters, consider `sync/atomic` instead of mutex:

```go
import "sync/atomic"

type AtomicCounter struct {
    value atomic.Int64
}

func (c *AtomicCounter) Inc() {
    c.value.Add(1)
}

func (c *AtomicCounter) Value() int64 {
    return c.value.Load()
}
```

---

## Gotchas and Pitfalls

### Pitfall 1: Writing with RLock

```go
// ❌ WRONG - Data race!
func (c *SafeCache) Set(key, value string) {
    c.mu.RLock()         // RLock doesn't protect writes!
    defer c.mu.RUnlock()
    c.cache[key] = value // Other RLock holders reading concurrently!
}

// ✅ CORRECT
func (c *SafeCache) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.cache[key] = value
}
```

### Pitfall 2: Lock Not Unlocked

```go
// ❌ WRONG - Deadlock if early return
func (c *SafeCounter) Inc() error {
    c.mu.Lock()
    if c.value >= 100 {
        return errors.New("max reached")  // Lock never released!
    }
    c.value++
    c.mu.Unlock()
    return nil
}

// ✅ CORRECT - defer ensures unlock
func (c *SafeCounter) Inc() error {
    c.mu.Lock()
    defer c.mu.Unlock()
    if c.value >= 100 {
        return errors.New("max reached")
    }
    c.value++
    return nil
}
```

### Pitfall 3: Nested Locks (Deadlock)

```go
// ❌ WRONG - Deadlock!
func (c *SafeCache) Process() {
    c.mu.Lock()
    // ... some work ...
    c.mu.Lock()  // Already holding lock! Deadlock!
    c.mu.Unlock()
    c.mu.Unlock()
}
```

Go mutexes are NOT reentrant. Same goroutine cannot lock twice.

### Pitfall 4: Lock Upgrade Deadlock

```go
// ❌ WRONG - Deadlock!
func (c *SafeCache) GetOrSet(key, value string) string {
    c.mu.RLock()
    if _, ok := c.cache[key]; !ok {
        c.mu.Lock()  // Deadlock! Still holding RLock!
        c.cache[key] = value
        c.mu.Unlock()
    }
    v := c.cache[key]
    c.mu.RUnlock()
    return v
}

// ✅ CORRECT - Release RLock before acquiring Lock
func (c *SafeCache) GetOrSet(key, value string) string {
    c.mu.RLock()
    if v, ok := c.cache[key]; ok {
        c.mu.RUnlock()
        return v
    }
    c.mu.RUnlock()  // Release first!

    c.mu.Lock()
    defer c.mu.Unlock()
    // ... check and set ...
}
```

### Pitfall 5: Copying Mutex

```go
// ❌ WRONG - Copied mutex is broken!
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c Counter) Value() int {  // Value receiver copies the mutex!
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

// ✅ CORRECT - Use pointer receiver
func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}
```

---

## Performance Considerations

### Benchmark: Mutex vs RWMutex

| Scenario | Mutex | RWMutex | Winner |
|----------|-------|---------|--------|
| 100% writes | ~20ns | ~40ns | Mutex |
| 50% reads | ~20ns | ~30ns | Mutex |
| 90% reads | ~20ns | ~15ns | RWMutex |
| 99% reads (10 concurrent) | ~20ns | ~5ns | RWMutex |

### Lock Contention Tips

1. **Minimize critical section size**
   ```go
   // ❌ Lock held too long
   c.mu.Lock()
   result := expensiveComputation()
   c.cache[key] = result
   c.mu.Unlock()

   // ✅ Lock only for shared state access
   result := expensiveComputation()
   c.mu.Lock()
   c.cache[key] = result
   c.mu.Unlock()
   ```

2. **Consider sharding for high contention**
   ```go
   type ShardedMap struct {
       shards [16]struct {
           mu sync.Mutex
           m  map[string]int
       }
   }

   func (s *ShardedMap) getShard(key string) int {
       return int(key[0]) % 16
   }
   ```

3. **Use sync.Map for specific patterns**
   - Keys written once, read many times
   - Disjoint sets of keys per goroutine

---

## Quick Reference

### sync.Mutex

```go
var mu sync.Mutex

mu.Lock()    // Acquire exclusive lock (blocks if held)
mu.Unlock()  // Release lock

// Always use defer for safety
mu.Lock()
defer mu.Unlock()
```

### sync.RWMutex

```go
var mu sync.RWMutex

// Exclusive write lock
mu.Lock()
mu.Unlock()

// Shared read lock
mu.RLock()
mu.RUnlock()

// Common pattern
func Read() {
    mu.RLock()
    defer mu.RUnlock()
    // read operations
}

func Write() {
    mu.Lock()
    defer mu.Unlock()
    // write operations
}
```

### Decision Checklist

- [ ] All operations are writes → `Mutex`
- [ ] Reads outnumber writes 10:1+ → `RWMutex`
- [ ] Simple use case → `Mutex` (lower complexity)
- [ ] Never write while holding `RLock`
- [ ] Never upgrade `RLock` → `Lock` without releasing first
- [ ] Always use `defer` for Unlock
- [ ] Always use pointer receivers with mutexes

---

## Related Resources

- [CONCURRENCY_DEBUGGING_GUIDE.md](CONCURRENCY_DEBUGGING_GUIDE.md) - Debugging concurrent code
- [CHANNELS_GUIDE.md](CHANNELS_GUIDE.md) - Channel-based synchronization
- [sync package documentation](https://pkg.go.dev/sync)

---

**Last Updated:** January 2026
**Module:** 09 Concurrency - Mutex and Shared State
