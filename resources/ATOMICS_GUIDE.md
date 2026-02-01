# Atomics in Go: A Comprehensive Guide

## What Are Atomics?

Atomics are **low-level synchronization primitives** that provide thread-safe operations on single variables without using locks. They guarantee that read-modify-write operations happen as a single, indivisible unit.

```go
import "sync/atomic"

var counter atomic.Int64

// Safe for concurrent access - no locks needed
counter.Add(1)
val := counter.Load()
```

## Why Do Atomics Exist?

### The Problem: Non-Atomic Operations

What looks like one operation in Go is often multiple CPU instructions:

```go
var x int64 = 0

// This is NOT atomic - it's actually 3 operations:
x++
// 1. Read x from memory into CPU register
// 2. Increment the register
// 3. Write the register back to memory
```

With concurrent goroutines, this causes **race conditions**:

```
Goroutine A                 Goroutine B
─────────────────────────────────────────
Read x (0)
                            Read x (0)
Increment (1)
                            Increment (1)
Write x (1)
                            Write x (1)  ← Lost update!

Expected: x = 2
Actual:   x = 1
```

### The Solution: Atomic Instructions

CPUs provide special **atomic instructions** that complete as a single, uninterruptible unit:

```go
var counter atomic.Int64

// This is truly atomic - one CPU instruction
counter.Add(1)
// Uses CPU instruction like LOCK XADD on x86
// No other CPU can interfere mid-operation
```

---

## What Happens Under the Hood

### CPU-Level Mechanics

Modern CPUs provide atomic instructions through hardware support:

**1. Memory Barriers (Fences)**
```
┌─────────────────────────────────────────────────────────┐
│  CPU Core 1          Memory           CPU Core 2        │
│  ┌─────────┐        ┌──────┐         ┌─────────┐        │
│  │ Cache   │◄──────►│ RAM  │◄───────►│ Cache   │        │
│  └─────────┘        └──────┘         └─────────┘        │
│       │                                    │            │
│       └── Memory barrier forces ───────────┘            │
│           cache synchronization                         │
└─────────────────────────────────────────────────────────┘
```

Atomic operations include **memory barriers** that:
- Flush CPU caches to main memory
- Ensure all CPUs see the same value
- Prevent instruction reordering around the atomic

**2. Lock Prefix (x86)**

On x86 processors, atomic operations use the `LOCK` prefix:

```asm
; Regular increment (NOT atomic)
MOV EAX, [counter]    ; Load
INC EAX               ; Increment
MOV [counter], EAX    ; Store

; Atomic increment
LOCK INC [counter]    ; Single atomic instruction
```

The `LOCK` prefix:
- Locks the memory bus during the operation
- Prevents other CPUs from accessing that memory location
- Ensures the operation completes atomically

**3. Compare-And-Swap (CAS)**

The fundamental building block of many atomic operations:

```
CAS(address, expected, new):
    ATOMICALLY:
        if *address == expected:
            *address = new
            return true
        else:
            return false
```

Used to implement complex atomic operations:

```go
// Atomic Add using CAS (conceptually)
func atomicAdd(addr *int64, delta int64) int64 {
    for {
        old := atomic.LoadInt64(addr)
        new := old + delta
        if atomic.CompareAndSwapInt64(addr, old, new) {
            return new
        }
        // CAS failed, another goroutine modified it - retry
    }
}
```

---

## Go's Atomic Types (Go 1.19+)

### Type-Safe Atomic Types

```go
import "sync/atomic"

// Integer types
var i32 atomic.Int32
var i64 atomic.Int64
var u32 atomic.Uint32
var u64 atomic.Uint64
var uptr atomic.Uintptr

// Boolean
var flag atomic.Bool

// Pointer (generic)
var ptr atomic.Pointer[MyStruct]

// Any value (uses interface{})
var val atomic.Value
```

### Operations Available

| Type | Load | Store | Add | Swap | CompareAndSwap |
|------|------|-------|-----|------|----------------|
| `atomic.Int64` | ✅ | ✅ | ✅ | ✅ | ✅ |
| `atomic.Bool` | ✅ | ✅ | ❌ | ✅ | ✅ |
| `atomic.Pointer[T]` | ✅ | ✅ | ❌ | ✅ | ✅ |
| `atomic.Value` | ✅ | ✅ | ❌ | ✅ | ✅ |

---

## Operations Explained

### Load and Store

```go
var counter atomic.Int64

// Store: atomically write a value
counter.Store(42)

// Load: atomically read the value
val := counter.Load()  // val = 42
```

**When to use:**
- Reading/writing shared configuration
- Publishing data to other goroutines
- Simple flags

### Add

```go
var counter atomic.Int64

// Add: atomically add and return NEW value
newVal := counter.Add(5)   // Returns 5
newVal = counter.Add(-2)   // Returns 3 (can subtract too)
```

**When to use:**
- Counters (requests, errors, metrics)
- Sequence numbers
- Reference counting

### Swap

```go
var state atomic.Int64

// Swap: atomically set new value, return OLD value
oldVal := state.Swap(100)  // Returns previous value
```

**When to use:**
- Replacing values and needing to process the old one
- Taking ownership of a resource

### CompareAndSwap (CAS)

```go
var state atomic.Int64
state.Store(1)

// Only updates if current value matches expected
swapped := state.CompareAndSwap(1, 2)  // true, state is now 2
swapped = state.CompareAndSwap(1, 3)   // false, state stays 2
```

**When to use:**
- Implementing lock-free data structures
- Optimistic concurrency (try, retry if changed)
- State machines with valid transitions

---

## Practical Examples

### Example 1: Request Counter

```go
type Server struct {
    requestCount atomic.Int64
    errorCount   atomic.Int64
}

func (s *Server) HandleRequest() {
    s.requestCount.Add(1)

    if err := processRequest(); err != nil {
        s.errorCount.Add(1)
    }
}

func (s *Server) Stats() (requests, errors int64) {
    return s.requestCount.Load(), s.errorCount.Load()
}
```

### Example 2: Shutdown Flag

```go
type Worker struct {
    shutdown atomic.Bool
}

func (w *Worker) Run() {
    for !w.shutdown.Load() {
        // Do work...
    }
}

func (w *Worker) Stop() {
    w.shutdown.Store(true)
}
```

### Example 3: Lazy Initialization (Single Value)

```go
type Config struct {
    data atomic.Pointer[ConfigData]
}

func (c *Config) Get() *ConfigData {
    if ptr := c.data.Load(); ptr != nil {
        return ptr
    }

    // First access - load config
    newData := loadConfig()

    // Try to set it - might race with another goroutine
    if c.data.CompareAndSwap(nil, newData) {
        return newData
    }

    // Another goroutine won the race - use their value
    return c.data.Load()
}
```

### Example 4: State Machine with Valid Transitions

```go
const (
    StateIdle    int32 = 0
    StateRunning int32 = 1
    StateStopped int32 = 2
)

type Service struct {
    state atomic.Int32
}

func (s *Service) Start() error {
    // Only start if currently idle
    if !s.state.CompareAndSwap(StateIdle, StateRunning) {
        return errors.New("service not idle")
    }
    go s.run()
    return nil
}

func (s *Service) Stop() error {
    // Only stop if currently running
    if !s.state.CompareAndSwap(StateRunning, StateStopped) {
        return errors.New("service not running")
    }
    return nil
}
```

---

## Atomics vs Mutex: When to Use Which

### Use Atomics When:

| Scenario | Example |
|----------|---------|
| Single variable | Counter, flag, pointer |
| Simple operations | Load, store, add |
| High contention expected | Millions of increments/sec |
| Lock-free required | Real-time systems |

### Use Mutex When:

| Scenario | Example |
|----------|---------|
| Multiple variables together | Transfer between accounts |
| Complex operations | Read-modify-write with logic |
| Protecting data structures | Map, slice modifications |
| Conditional waits needed | sync.Cond scenarios |

### Performance Comparison

```go
// Atomic: ~10-20 nanoseconds per operation
counter.Add(1)

// Mutex: ~20-50 nanoseconds per operation (uncontended)
mu.Lock()
counter++
mu.Unlock()

// Under high contention:
// Atomic: Scales better (no waiting)
// Mutex: Can cause goroutine pile-up
```

### Decision Flowchart

```
Is it a single variable?
├── NO → Use Mutex
└── YES → Is the operation simple (load/store/add)?
          ├── NO → Use Mutex
          └── YES → Is performance critical?
                    ├── NO → Either works (Mutex is clearer)
                    └── YES → Use Atomic
```

---

## Memory Ordering and Happens-Before

### Go's Memory Model Guarantees

Atomic operations establish **happens-before** relationships:

```go
var data string
var ready atomic.Bool

// Goroutine 1
data = "hello"        // Write data
ready.Store(true)     // Atomic store with memory barrier

// Goroutine 2
if ready.Load() {     // Atomic load with memory barrier
    fmt.Println(data) // Guaranteed to see "hello"
}
```

The atomic store **synchronizes-with** the atomic load, creating a happens-before edge:

```
Goroutine 1                    Goroutine 2
─────────────────────────────────────────────
data = "hello"
    │
ready.Store(true) ────────────► ready.Load() == true
                                    │
                                fmt.Println(data)  // sees "hello"
```

### Common Mistake: Non-Atomic Flags

```go
var data string
var ready bool  // NOT atomic!

// Goroutine 1
data = "hello"
ready = true    // No memory barrier!

// Goroutine 2
if ready {      // Might see true...
    fmt.Println(data)  // ...but might NOT see "hello"!
}
```

Without atomics, the CPU might reorder operations or caches might be stale.

---

## Legacy Function-Based API

Before Go 1.19, atomics used functions instead of methods:

```go
// Old style (still works, but less type-safe)
var counter int64

atomic.AddInt64(&counter, 1)
val := atomic.LoadInt64(&counter)
atomic.StoreInt64(&counter, 0)

// New style (Go 1.19+, preferred)
var counter atomic.Int64

counter.Add(1)
val := counter.Load()
counter.Store(0)
```

**Prefer the new type-safe API** - it prevents accidentally using non-atomic operations on the variable.

---

## Common Pitfalls

### Pitfall 1: Mixing Atomic and Non-Atomic Access

```go
var counter atomic.Int64

counter.Add(1)      // Atomic
x := counter.Load() // Atomic
// WRONG: Can't do this
// y := int64(counter)  // Won't compile - correct!
```

The type-safe API prevents this mistake at compile time.

### Pitfall 2: Assuming Atomicity of Multiple Operations

```go
var a, b atomic.Int64

// NOT atomic as a unit!
a.Add(1)  // Atomic
b.Add(1)  // Atomic
// Another goroutine could see a incremented but not b

// If you need both together, use a mutex
```

### Pitfall 3: Over-Using CAS Loops

```go
// Potential live-lock under extreme contention
for {
    old := val.Load()
    if val.CompareAndSwap(old, old+1) {
        break
    }
    // If MANY goroutines compete, this can spin forever
}

// Solution: Add backoff or use mutex for high contention
```

---

## Summary

| Concept | Key Point |
|---------|-----------|
| **What** | Thread-safe operations on single variables |
| **Why** | Avoid race conditions without locks |
| **How** | CPU provides atomic instructions (LOCK, CAS) |
| **When** | Counters, flags, simple state, high performance |
| **Memory** | Includes barriers for visibility guarantees |
| **Types** | `atomic.Int64`, `atomic.Bool`, `atomic.Pointer[T]` |
| **Operations** | Load, Store, Add, Swap, CompareAndSwap |

### Quick Reference

```go
import "sync/atomic"

var counter atomic.Int64  // Zero value is 0, ready to use
var flag atomic.Bool      // Zero value is false
var ptr atomic.Pointer[T] // Zero value is nil

counter.Store(10)         // Write
val := counter.Load()     // Read
counter.Add(5)            // Add (returns new value)
old := counter.Swap(0)    // Replace (returns old value)
ok := counter.CompareAndSwap(expected, new)  // Conditional replace
```
