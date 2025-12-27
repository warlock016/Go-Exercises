# Exercise 06: Interface Composition

**Concept:** Embedding interfaces to create larger interfaces
**Difficulty:** Medium
**Estimated Time:** 50 minutes

## Learning Goal

Learn how Go composes small interfaces into larger ones through embedding. This is a key pattern for building flexible abstractions.

## The Problem

Instead of creating large interfaces, Go encourages composing small ones:

```go
// Instead of this:
type LogCloser interface {
    Log(message string)
    Close() error
}

// Do this (compose smaller interfaces):
type Logger interface {
    Log(message string)
}

type Closer interface {
    Close() error
}

type LogCloser interface {
    Logger
    Closer
}
```

## Your Task

Create interfaces for logging and resource cleanup, then compose them.

### Interfaces

1. **Logger** - writes log messages
   - `Log(message string)`

2. **Closer** - cleans up resources
   - `Close() error`

3. **LogCloser** - combines both (via embedding)

### Type

**FileLogger** - implements LogCloser

**Fields:**
- `Filename string`
- `logs []string` - stored messages
- `closed bool` - whether Close() was called

## Function Signatures

```go
type Logger interface {
    Log(message string)
}

type Closer interface {
    Close() error
}

type LogCloser interface {
    Logger
    Closer
}

type FileLogger struct {
    Filename string
    logs     []string
    closed   bool
}

func (f *FileLogger) Log(message string)
func (f *FileLogger) Close() error
func (f *FileLogger) GetLogs() []string  // Helper for testing
```

## Examples

```go
logger := &FileLogger{Filename: "app.log"}
logger.Log("Starting")
logger.Log("Processing")
logger.Close()

// logger implements Logger, Closer, AND LogCloser
var l Logger = logger      // OK
var c Closer = logger      // OK
var lc LogCloser = logger  // OK
```

## What This Teaches

- Interface embedding
- Composing small interfaces
- Implementing multiple interfaces with one type
- Interface satisfaction through composition

---

**Next up:** Exercise 07 - Type Assertions
