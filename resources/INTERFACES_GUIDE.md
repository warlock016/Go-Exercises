# Go Interfaces: Comprehensive Guide

This guide covers Go interfaces from fundamentals to advanced patterns, including internal mechanisms, design principles, and practical applications.

---

## Table of Contents

1. [High-Level Overview](#high-level-overview)
2. [Interface Values Internals](#interface-values-internals)
3. [The io.Reader and io.Writer Contracts](#the-ioreader-and-iowriter-contracts)
4. [Wrapper/Decorator Pattern](#wrapperdecorator-pattern)
5. [Wrapper Chains and Pipelines](#wrapper-chains-and-pipelines)
6. [io.Copy: Bridging Readers and Writers](#iocopy-bridging-readers-and-writers)
7. [Pointer vs Value Receivers](#pointer-vs-value-receivers)
8. [Interface Segregation Principle](#interface-segregation-principle)
9. [Polymorphic Collections](#polymorphic-collections)
10. [Type Assertions and Type Switches](#type-assertions-and-type-switches)
11. [Rules of Thumb](#rules-of-thumb)
12. [Common Pitfalls](#common-pitfalls)

---

## High-Level Overview

### What Is an Interface?

An interface in Go is a **contract** that specifies behavior without implementation. Any type that implements the required methods automatically satisfies the interface—no explicit declaration needed.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

// strings.Reader satisfies Reader (has Read method)
// os.File satisfies Reader (has Read method)
// bytes.Buffer satisfies Reader (has Read method)
// Your custom type can satisfy Reader too
```

### Why Interfaces Matter

| Benefit | Description |
|---------|-------------|
| **Decoupling** | Code depends on behavior, not concrete types |
| **Testability** | Easy to mock dependencies |
| **Composability** | Small interfaces combine into powerful abstractions |
| **Polymorphism** | Different types, same behavior contract |

### The Implicit Satisfaction Model

Unlike Java/C#, Go interfaces are satisfied **implicitly**:

```go
// No "implements" keyword needed
type MyReader struct{}

func (m MyReader) Read(p []byte) (int, error) {
    // Implementation
    return 0, nil
}

// MyReader now satisfies io.Reader automatically
var r io.Reader = MyReader{}  // Works!
```

---

## Interface Values Internals

Understanding what an interface value actually IS helps explain many behaviors.

### The Two-Word Structure

An interface value is **not** a simple pointer. It's a 16-byte structure containing two words:

```
┌─────────────────────────────────────────────────────────────────┐
│                     Interface Value (16 bytes)                  │
├─────────────────────────────┬───────────────────────────────────┤
│      Word 1: Type Info      │       Word 2: Data Pointer        │
│   (pointer to itable)       │   (pointer to concrete value)     │
└─────────────────────────────┴───────────────────────────────────┘
```

### The itable (Interface Table)

The type info pointer points to an **itable** containing:

```
┌─────────────────────────────────────────────────────────────────┐
│  itable for (io.Reader, *os.File):                              │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ concrete type: *os.File                                     ││
│  │ interface type: io.Reader                                   ││
│  │ method[0]: Read → (*os.File).Read (address: 0x4500)         ││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
```

### Method Call Sequence

When you call a method on an interface value:

```go
var r io.Reader = file
n, err := r.Read(buf)
```

The runtime:
1. Reads the itable pointer from word 1
2. Looks up the `Read` method address in the itable
3. Reads the data pointer from word 2
4. Calls the method with the data pointer as receiver

### nil Interface vs nil Concrete Value

This is a common source of confusion:

```go
var r io.Reader           // r is nil (both words are nil)
r == nil                  // true

var f *os.File = nil
var r io.Reader = f       // r is NOT nil! (type word is set)
r == nil                  // false (type info exists)

// The interface "knows" it holds a *os.File, even if that pointer is nil
```

**Rule:** An interface is nil only when BOTH type and value are nil.

---

## The io.Reader and io.Writer Contracts

These are Go's most important interfaces, used throughout the standard library.

### io.Reader Contract

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

**The contract:**

| Aspect | Requirement |
|--------|-------------|
| **Buffer ownership** | Caller provides buffer `p`, reader fills it |
| **Return value `n`** | Number of bytes actually read (0 to len(p)) |
| **Return value `err`** | nil on success, io.EOF at end, or other error |
| **Partial reads** | n < len(p) is valid and common |
| **EOF behavior** | May return n > 0 with EOF, or 0 with EOF on next call |

**Mental Model - The Lunchbox Analogy:**

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  Caller: "Here's an empty lunchbox (buffer). Fill it please."   │
│          ┌─────────────────────┐                                │
│          │ Empty buffer (p)    │                                │
│          │ [   |   |   |   ]   │  capacity: 4 bytes             │
│          └─────────────────────┘                                │
│                    │                                            │
│                    ▼                                            │
│  Reader: "I'll put food (data) in your lunchbox."               │
│          ┌─────────────────────┐                                │
│          │ Filled buffer       │                                │
│          │ [ H | e | l | l ]   │  n=4 bytes filled              │
│          └─────────────────────┘                                │
│                                                                 │
│  Reader: "Here's your lunchbox back. I put 4 items in it."      │
│          return 4, nil                                          │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### io.Writer Contract

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

**The contract:**

| Aspect | Requirement |
|--------|-------------|
| **Buffer ownership** | Caller provides data in `p`, writer consumes it |
| **Return value `n`** | Number of bytes successfully written |
| **Return value `err`** | nil on success, or error explaining failure |
| **Short writes** | n < len(p) requires non-nil error |
| **Buffer safety** | Writer must not modify `p` or retain references |

**Data Flow Comparison:**

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  Reader: Data flows FROM reader INTO caller's buffer            │
│                                                                 │
│     Source ──────────────────────▶ Buffer                       │
│     (Reader)        Read()         (Caller provides)            │
│                                                                 │
│  Writer: Data flows FROM caller's buffer INTO writer            │
│                                                                 │
│     Buffer ──────────────────────▶ Destination                  │
│     (Caller provides)  Write()     (Writer)                     │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Why Return (n int, err error)?

**Memory Efficiency:**
```go
// Bad: Reader allocates and returns data (allocations every call)
func Read() ([]byte, error)

// Good: Caller reuses buffer (zero allocations in loop)
func Read(p []byte) (int, error)
```

**Streaming Capability:**
```go
buf := make([]byte, 4096)  // One allocation
for {
    n, err := reader.Read(buf)  // Reuse same buffer
    if n > 0 {
        process(buf[:n])
    }
    if err == io.EOF {
        break
    }
}
```

---

## Wrapper/Decorator Pattern

Wrappers add behavior to existing Readers/Writers without modifying them.

### Anatomy of a Wrapper

```go
// A wrapper has three components:
type CountingReader struct {
    Reader    io.Reader  // 1. Embedded interface (the "inner" reader)
    BytesRead int        // 2. Additional state
}

// 3. Method that delegates + adds behavior
func (r *CountingReader) Read(p []byte) (n int, err error) {
    n, err = r.Reader.Read(p)  // Delegate to inner reader
    r.BytesRead += n           // Add counting behavior
    return n, err              // Pass through results
}
```

### Types of Wrappers

**Source Readers** - Generate or hold actual data:
```go
strings.NewReader("hello")  // Data comes from string
os.Open("file.txt")         // Data comes from file
bytes.NewBuffer(data)       // Data comes from byte slice
```

**Wrapper Readers** - Transform or observe data from another reader:
```go
// Counting: tracks bytes read
type CountingReader struct {
    Reader    io.Reader
    BytesRead int
}

// Transforming: modifies data as it passes through
type UppercaseReader struct {
    Reader io.Reader
}

// Limiting: restricts total bytes read
type LimitReader struct {
    Reader io.Reader
    Limit  int
    read   int
}

// Teeing: copies data to a writer while reading
type TeeReader struct {
    Reader io.Reader
    Writer io.Writer
}
```

### Wrapper Implementation Pattern

```go
func (r *WrapperReader) Read(p []byte) (n int, err error) {
    // 1. (Optional) Pre-processing
    //    e.g., limit buffer size for LimitReader

    // 2. Delegate to inner reader
    n, err = r.Reader.Read(p)

    // 3. (Optional) Post-processing
    //    e.g., count bytes, transform data

    // 4. Return results (usually pass through n and err)
    return n, err
}
```

### Error Handling in Wrappers

**Critical Rule:** Wrappers should pass through errors from inner readers, never invent errors.

```go
func (r *CountingReader) Read(p []byte) (n int, err error) {
    n, err = r.Reader.Read(p)
    r.BytesRead += n
    return n, err  // Pass through whatever error came from inner reader
}
```

**Exception:** Wrappers that enforce limits may return io.EOF when their limit is reached:

```go
func (r *LimitReader) Read(p []byte) (n int, err error) {
    remaining := r.Limit - r.read
    if remaining <= 0 {
        return 0, io.EOF  // Limit-enforced EOF is appropriate
    }
    // ... continue
}
```

---

## Wrapper Chains and Pipelines

Wrappers can be chained to create data processing pipelines.

### Building a Pipeline

```go
source := strings.NewReader("hello world")

// Chain: source → uppercase → count → limit
upper := &UppercaseReader{Reader: source}
counter := &CountingReader{Reader: upper}
limited := &LimitReader{Reader: counter, Limit: 5}

// Read from the end of the chain
result, _ := io.ReadAll(limited)  // "HELLO"
```

### Data Flow Through Chains

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  When caller calls: limited.Read(buf)                           │
│                                                                 │
│  Call chain (top-down):          Data returns (bottom-up):      │
│                                                                 │
│  Caller                          Caller receives "HELLO"        │
│    │                                   ▲                        │
│    │ limited.Read(buf)                 │                        │
│    ▼                                   │                        │
│  LimitReader                     Returns first 5 bytes          │
│    │ "Only pass 5 bytes"               ▲                        │
│    │ counter.Read(buf)                 │                        │
│    ▼                                   │                        │
│  CountingReader                  Counts 5, passes through       │
│    │ "Count what passes"               ▲                        │
│    │ upper.Read(buf)                   │                        │
│    ▼                                   │                        │
│  UppercaseReader                 Returns "HELLO WORLD"          │
│    │ "Transform to uppercase"          ▲                        │
│    │ source.Read(buf)                  │                        │
│    ▼                                   │                        │
│  strings.Reader                  Returns "hello world"          │
│    └── Contains "hello world"                                   │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Pipeline Order Matters

```go
// Order 1: Uppercase THEN Limit
upper := &UppercaseReader{Reader: source}
limited := &LimitReader{Reader: upper, Limit: 5}
// Result: "HELLO" (uppercase applied, then limited)

// Order 2: Limit THEN Uppercase
limited := &LimitReader{Reader: source, Limit: 5}
upper := &UppercaseReader{Reader: limited}
// Result: "HELLO" (same result, but different processing)

// Order 3: Count at different positions gives different results
// Count before limit: counts all bytes read from source
// Count after limit: counts only bytes that pass the limit
```

---

## io.Copy: Bridging Readers and Writers

`io.Copy` is the standard way to connect a Reader to a Writer.

### The Copy Pattern

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

**What it does internally:**

```go
// Simplified io.Copy implementation
func Copy(dst Writer, src Reader) (int64, error) {
    buf := make([]byte, 32*1024)  // 32KB buffer
    var written int64
    for {
        nr, readErr := src.Read(buf)
        if nr > 0 {
            nw, writeErr := dst.Write(buf[:nr])
            written += int64(nw)
            if writeErr != nil {
                return written, writeErr
            }
        }
        if readErr == io.EOF {
            return written, nil
        }
        if readErr != nil {
            return written, readErr
        }
    }
}
```

### Pipeline with io.Copy

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  ProcessData with transformations:                              │
│                                                                 │
│  ┌─────────┐    ┌───────────┐    ┌─────────┐    ┌───────────┐   │
│  │ Source  │───▶│ Uppercase │───▶│ Counter │───▶│  Limit    │   │
│  │ Reader  │    │  Reader   │    │ Reader  │    │  Reader   │   │
│  └─────────┘    └───────────┘    └─────────┘    └───────────┘   │
│                                                      │          │
│                       io.Copy connects them          │          │
│                              │                       │          │
│                              ▼                       ▼          │
│                    ┌─────────────────────────────────────┐      │
│                    │         io.Copy(dst, src)           │      │
│                    │   Reads from pipeline, writes to    │      │
│                    │         destination writer          │      │
│                    └─────────────────┬───────────────────┘      │
│                                      │                          │
│                                      ▼                          │
│                            ┌─────────────────┐                  │
│                            │   Destination   │                  │
│                            │     Writer      │                  │
│                            └─────────────────┘                  │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### io.Copy Variants

```go
// Copy all data
io.Copy(dst, src)

// Copy exactly n bytes (error if less available)
io.CopyN(dst, src, n)

// Copy with your own buffer (avoids allocation)
io.CopyBuffer(dst, src, buf)
```

---

## Pointer vs Value Receivers

Choosing between pointer and value receivers affects interface satisfaction.

### The Decision Framework

```
Does method modify any field?
    │
    ├── YES → Use pointer receiver (*T)
    │
    └── NO → Is struct large (>64 bytes or many fields)?
              │
              ├── YES → Consider pointer (*T) for efficiency
              │
              └── NO → Value receiver (T) is fine
```

### Examples by Category

**Value receivers** - Methods that only read:
```go
func (f File) Size() int64 {
    return f.FileSize  // Just reading
}

func (f File) Name() string {
    return f.FileName  // Just reading
}
```

**Pointer receivers** - Methods that modify state:
```go
func (c *CountingReader) Read(p []byte) (int, error) {
    n, err := c.Reader.Read(p)
    c.BytesRead += n  // Modifying c.BytesRead
    return n, err
}

func (f *FileLogger) Close() error {
    f.closed = true  // Modifying f.closed
    return f.file.Close()
}
```

### Interface Satisfaction Rules

```go
type Sizer interface {
    Size() int64
}

// With value receiver:
func (f File) Size() int64 { return f.FileSize }

var s Sizer = File{}   // Works: File satisfies Sizer
var s Sizer = &File{}  // Works: *File also satisfies (auto-deref)

// With pointer receiver:
func (f *File) Size() int64 { return f.FileSize }

var s Sizer = &File{}  // Works: *File satisfies Sizer
var s Sizer = File{}   // ERROR: File does not satisfy Sizer
```

**Why the asymmetry?**
- Given `*File`, Go can always get `File` (dereference)
- Given `File`, Go cannot always get `*File` (value might be in read-only memory, temporary, etc.)

### Consistency Rule

If ANY method needs a pointer receiver, use pointer receivers for ALL methods:

```go
type FileLogger struct {
    file    *os.File
    written int
    closed  bool
}

// ALL methods use *FileLogger for consistency
func (f *FileLogger) Write(p []byte) (int, error) { ... }
func (f *FileLogger) Close() error { ... }
func (f *FileLogger) BytesWritten() int { return f.written }  // Even this one
```

---

## Interface Segregation Principle

"Clients should not be forced to depend on interfaces they don't use."

### Monolithic vs Segregated

```go
// BAD: Monolithic interface
type Store interface {
    Create(id, value string) error
    Read(id string) (string, error)
    Update(id, value string) error
    Delete(id string) error
    List() []string
}

// Functions require full Store even if they only need one method
func ProcessItem(s Store) { /* only calls Read */ }

// GOOD: Segregated interfaces
type Creator interface { Create(id, value string) error }
type Reader interface { Read(id string) (string, error) }
type Updater interface { Update(id, value string) error }
type Deleter interface { Delete(id string) error }
type Lister interface { List() []string }

// Functions declare minimal dependencies
func ProcessItem(r Reader) { /* only needs Read */ }
```

### Benefits of Segregation

```go
// MigrateAll only needs what it uses
func MigrateAll(from Reader, fromList Lister, to Creator) error {
    for _, id := range fromList.List() {
        value, _ := from.Read(id)
        to.Create(id, value)
    }
    return nil
}

// Now these minimal types work:
type OnlyReader struct { data map[string]string }
func (r *OnlyReader) Read(id string) (string, error) { ... }

type OnlyLister struct { keys []string }
func (l *OnlyLister) List() []string { ... }

type OnlyCreator struct { data map[string]string }
func (c *OnlyCreator) Create(id, value string) error { ... }

// They don't implement full CRUD, but they work with MigrateAll!
```

### Standard Library Examples

```go
// io package uses small interfaces
type Reader interface { Read(p []byte) (n int, err error) }
type Writer interface { Write(p []byte) (n int, err error) }
type Closer interface { Close() error }

// Composed when needed
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

---

## Polymorphic Collections

Interfaces enable storing different concrete types in the same collection.

### How It Works

```go
type Prioritizable interface {
    Priority() int
}

type Task struct { priority int; description string }
func (t Task) Priority() int { return t.priority }

type Message struct { priority int; text string }
func (m Message) Priority() int { return m.priority }

// Both can live in the same slice
items := []Prioritizable{
    Task{priority: 5, description: "urgent"},
    Message{priority: 3, text: "hello"},
}
```

### Memory Layout

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  []Prioritizable (slice of interface values)                    │
│                                                                 │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ Slice header: ptr=0x1000, len=2, cap=2                      ││
│  └─────────────────────────────────────────────────────────────┘│
│           │                                                     │
│           ▼                                                     │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ items[0]: Interface value (16 bytes)                        ││
│  │ ┌──────────────────┬──────────────────┐                     ││
│  │ │ type: *TaskType  │ data: 0x2000 ────┼──▶ Task{5,"urgent"} ││
│  │ └──────────────────┴──────────────────┘                     ││
│  │                                                             ││
│  │ items[1]: Interface value (16 bytes)                        ││
│  │ ┌──────────────────┬──────────────────┐                     ││
│  │ │ type: *MsgType   │ data: 0x3000 ────┼──▶ Message{3,"hi"}  ││
│  │ └──────────────────┴──────────────────┘                     ││
│  └─────────────────────────────────────────────────────────────┘│
│                                                                 │
│  Key insight: All interface values are 16 bytes, regardless     │
│  of the concrete type they hold. The actual data lives          │
│  elsewhere in memory.                                           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Why You Can't Assign []T to []Interface

```go
tasks := []Task{{priority: 1}, {priority: 2}}
var items []Prioritizable = tasks  // COMPILE ERROR!
```

**Reason:** Memory layouts are incompatible:
- `[]Task` is a contiguous array of Task structs
- `[]Prioritizable` is a contiguous array of 16-byte interface values

You must convert element by element:
```go
items := make([]Prioritizable, len(tasks))
for i, t := range tasks {
    items[i] = t  // Each assignment creates an interface value
}
```

---

## Type Assertions and Type Switches

Extract concrete types from interface values.

### Type Assertion

```go
var r io.Reader = os.Stdin

// Unsafe assertion (panics if wrong type)
file := r.(*os.File)

// Safe assertion (returns ok=false if wrong type)
file, ok := r.(*os.File)
if ok {
    // file is *os.File
}
```

### Type Switch

```go
func describe(i interface{}) string {
    switch v := i.(type) {
    case string:
        return "string: " + v
    case int:
        return fmt.Sprintf("int: %d", v)
    case *os.File:
        return "file: " + v.Name()
    default:
        return "unknown type"
    }
}
```

### Type Switch Gotcha: Inner Conditionals

```go
func Add(a, b any) (any, bool) {
    switch a := a.(type) {
    case int:
        if bi, ok := b.(int); ok {
            return a + bi, true
        }
        // If ok is false, no return here!
        // Execution falls to END OF SWITCH, not to default
    case string:
        if bs, ok := b.(string); ok {
            return a + bs, true
        }
    default:
        return nil, false
    }
    return nil, false  // REQUIRED: cases may not return
}
```

**Why?** When a case matches but an inner `if` is false, execution continues to the end of the switch statement, NOT to the default case.

---

## Rules of Thumb

### Interface Design

| Rule | Rationale |
|------|-----------|
| Keep interfaces small (1-3 methods) | Easier to implement, compose, and mock |
| Accept interfaces, return structs | Flexible inputs, concrete outputs |
| Define interfaces where they're used | Consumer knows what it needs |
| Name single-method interfaces with -er suffix | Reader, Writer, Closer, Stringer |

### Receivers

| Situation | Use |
|-----------|-----|
| Method modifies receiver | Pointer receiver |
| Struct is large | Pointer receiver |
| Consistency with other methods | Match existing style |
| Method only reads, struct is small | Value receiver OK |

### Wrappers

| Rule | Rationale |
|------|-----------|
| Pass through errors from inner reader/writer | Don't hide problems |
| Don't invent errors (except limit-based EOF) | Wrappers transform, not validate |
| Store reference to inner interface, not concrete type | Enable chaining with any implementation |
| Use pointer receiver if wrapper has state | State changes must persist |

### io.Reader/Writer

| Rule | Rationale |
|------|-----------|
| Check n > 0 before processing data | Data may exist even with error |
| Handle io.EOF specially (it's not a failure) | EOF is expected end condition |
| Don't assume full buffer fills | Partial reads are valid |
| For writers: n < len(p) requires error | Writer contract |

---

## Common Pitfalls

### 1. nil Interface vs nil Concrete Value

```go
var f *os.File = nil
var r io.Reader = f
r == nil  // FALSE! Type info is set

// Fix: Check before assignment or use type assertion
if f != nil {
    r = f
}
```

### 2. Modifying with Value Receiver

```go
func (c CountingReader) Read(p []byte) (int, error) {
    n, err := c.Reader.Read(p)
    c.BytesRead += n  // This modification is LOST!
    return n, err
}

// Fix: Use pointer receiver
func (c *CountingReader) Read(p []byte) (int, error) { ... }
```

### 3. Uppercasing Before Reading

```go
// WRONG: Buffer is empty before Read!
func (r *UppercaseReader) Read(p []byte) (int, error) {
    upper := strings.ToUpper(string(p))  // p is empty!
    return r.Reader.Read([]byte(upper))
}

// CORRECT: Read first, then transform
func (r *UppercaseReader) Read(p []byte) (int, error) {
    n, err := r.Reader.Read(p)
    upper := strings.ToUpper(string(p[:n]))
    copy(p, []byte(upper))
    return n, err
}
```

### 4. Using io.Copy Inside Read()

```go
// WRONG: Consumes entire stream in one Read call!
func (r *TeeReader) Read(p []byte) (int, error) {
    io.Copy(r.Writer, r.Reader)  // Reads EVERYTHING
    return 0, io.EOF
}

// CORRECT: Write only the chunk just read
func (r *TeeReader) Read(p []byte) (int, error) {
    n, err := r.Reader.Read(p)
    if n > 0 {
        r.Writer.Write(p[:n])  // Just this chunk
    }
    return n, err
}
```

### 5. Buffer Size vs Cumulative Limit Confusion

```go
// Buffer size: Per-call maximum
buf := make([]byte, 5)  // Each Read gets at most 5 bytes

// Limit: Cumulative maximum
type LimitReader struct {
    Reader io.Reader
    Limit  int   // Total bytes allowed across ALL reads
    read   int   // Running total
}
```

### 6. Expecting Type Switch Default After Failed Case

```go
switch v := x.(type) {
case int:
    if v > 0 {
        return "positive"
    }
    // v <= 0: falls to END of switch, NOT to default!
default:
    return "not an int"
}
return "non-positive int"  // This line is reachable!
```

---

## Summary

Go interfaces enable powerful abstractions through:

1. **Implicit satisfaction** - No explicit declaration needed
2. **Small interfaces** - Easy to implement and compose
3. **Two-word structure** - Type info + data pointer enables polymorphism
4. **Wrapper pattern** - Add behavior without modification
5. **io.Copy bridge** - Connect readers to writers seamlessly

The key mental model: An interface value is an **envelope** that carries both "what type is inside" and "where is the data." This uniform 16-byte size enables polymorphic collections while the itable enables runtime method dispatch.
