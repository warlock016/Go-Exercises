# io.Reader and io.Writer: A Practical Guide

Understanding Go's fundamental I/O interfaces — how they work under the hood and when to use each implementation.

> **Related:** [INTERFACES_GUIDE.md](INTERFACES_GUIDE.md) covers the wrapper/decorator patterns built on these interfaces.

---

## Table of Contents

1. [Why Interfaces Instead of Raw Bytes?](#why-interfaces-instead-of-raw-bytes)
2. [The io.Reader Interface](#the-ioreader-interface)
3. [The io.Writer Interface](#the-iowriter-interface)
4. [Common Implementations](#common-implementations)
5. [How Reading Actually Works](#how-reading-actually-works)
6. [Practical Patterns](#practical-patterns)
7. [HTTP Context: Request and Response Bodies](#http-context-request-and-response-bodies)
8. [Common Pitfalls](#common-pitfalls)
9. [Quick Reference](#quick-reference)

---

## Why Interfaces Instead of Raw Bytes?

### The Problem with `[]byte`

Imagine if all I/O functions took `[]byte` directly:

```go
// Hypothetical: if everything used []byte
func ProcessData(data []byte) error
func SendRequest(body []byte) (*Response, error)
```

**Problems:**
1. **Memory:** A 1GB file upload requires 1GB in memory all at once
2. **Inflexibility:** Data might come from files, networks, or generated on-the-fly
3. **No streaming:** Can't process data as it arrives
4. **Coupling:** Every function must handle the full data

### The Solution: Abstract the Source

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

`io.Reader` is a **contract**: "I can provide bytes when you ask."

The consumer doesn't care WHERE bytes come from — it just calls `Read()`.

```
┌─────────────────────────────────────────────────────────────────┐
│                     io.Reader Abstraction                       │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│   ┌──────────────┐                                              │
│   │ String       │──┐                                           │
│   └──────────────┘  │                                           │
│   ┌──────────────┐  │    ┌──────────────┐    ┌──────────────┐   │
│   │ []byte       │──┼───►│  io.Reader   │───►│  Consumer    │   │
│   └──────────────┘  │    │  interface   │    │  (doesn't    │   │
│   ┌──────────────┐  │    └──────────────┘    │  care about  │   │
│   │ File         │──┤                        │  source)     │   │
│   └──────────────┘  │                        └──────────────┘   │
│   ┌──────────────┐  │                                           │
│   │ Network      │──┤                                           │
│   └──────────────┘  │                                           │
│   ┌──────────────┐  │                                           │
│   │ Compressed   │──┘                                           │
│   └──────────────┘                                              │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## The io.Reader Interface

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### How `Read()` Works

1. **You provide a buffer** (`p []byte`) — space for data to be written into
2. **Reader fills the buffer** with up to `len(p)` bytes
3. **Returns `n`** — how many bytes were actually written
4. **Returns `err`** — `nil` for success, `io.EOF` when done, or an error

### Key Behaviors

```go
reader := strings.NewReader("Hello")
buf := make([]byte, 3)

// First read: gets up to 3 bytes
n, err := reader.Read(buf)
// n=3, buf=["H","e","l"], err=nil

// Second read: gets remaining bytes
n, err = reader.Read(buf)
// n=2, buf=["l","o","l"], err=nil (old "l" still there)

// Third read: no more data
n, err = reader.Read(buf)
// n=0, err=io.EOF
```

### Important Rules

| Rule | Explanation |
|------|-------------|
| `Read` may return less than `len(p)` bytes | Even if more data exists — this is normal |
| `n > 0` can occur with `err != nil` | Always process bytes before checking error |
| `io.EOF` means success | It signals "done reading," not failure |
| Buffer is reused | Previous contents may remain after short reads |

---

## The io.Writer Interface

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### How `Write()` Works

1. **You provide data** (`p []byte`) — bytes to write
2. **Writer consumes the data** (to file, network, buffer, etc.)
3. **Returns `n`** — how many bytes were written
4. **Returns `err`** — `nil` for success, or an error

### Key Behaviors

```go
var buf bytes.Buffer

n, err := buf.Write([]byte("Hello"))
// n=5, err=nil, buf now contains "Hello"

n, err = buf.Write([]byte(" World"))
// n=6, err=nil, buf now contains "Hello World"
```

---

## Common Implementations

### Readers: Where Data Comes FROM

| Type | Package | Creates From | Use Case |
|------|---------|--------------|----------|
| `*strings.Reader` | `strings` | `string` | Testing, fixed text |
| `*bytes.Reader` | `bytes` | `[]byte` | Testing, in-memory data |
| `*bytes.Buffer` | `bytes` | `[]byte` | Read/write buffer |
| `*os.File` | `os` | File path | File I/O |
| `resp.Body` | `net/http` | HTTP response | HTTP clients |
| `r.Body` | `net/http` | HTTP request | HTTP handlers |
| `*gzip.Reader` | `compress/gzip` | Compressed stream | Decompression |

### Writers: Where Data Goes TO

| Type | Package | Writes To | Use Case |
|------|---------|-----------|----------|
| `*bytes.Buffer` | `bytes` | In-memory buffer | Collecting output |
| `*os.File` | `os` | File on disk | File I/O |
| `http.ResponseWriter` | `net/http` | HTTP response | HTTP handlers |
| `*gzip.Writer` | `compress/gzip` | Compressed stream | Compression |
| `os.Stdout` | `os` | Standard output | Console output |

### Creating Readers

```go
// From string
r1 := strings.NewReader("hello world")

// From []byte
data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
r2 := bytes.NewReader(data)

// From file
r3, err := os.Open("file.txt")
if err != nil { ... }
defer r3.Close()

// From buffer (also implements Writer)
buf := bytes.NewBufferString("initial content")
```

### Creating Writers

```go
// To buffer
var buf bytes.Buffer
buf.Write([]byte("hello"))

// To file
f, err := os.Create("output.txt")
if err != nil { ... }
defer f.Close()
f.Write([]byte("content"))

// To HTTP response (in handler)
func handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("response"))
}
```

---

## How Reading Actually Works

### Lazy Evaluation

Data is NOT read when you create a reader — only when `Read()` is called:

```go
// Step 1: Create reader (NO DATA READ YET)
reader := strings.NewReader("Hello World")
// The string is stored, position is 0

// Step 2: Someone calls Read (NOW data is read)
buf := make([]byte, 5)
n, _ := reader.Read(buf)  // Copies "Hello" into buf
// Position advances to 5

// Step 3: Another Read
n, _ = reader.Read(buf)   // Copies " Worl" into buf
// Position advances to 10
```

### Visual: Internal State Changes

```
┌─────────────────────────────────────────────────────────────────┐
│  strings.Reader state during reads                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  After creation:                                                │
│  ┌─────────────────────────────────────────┐                    │
│  │  s: "Hello World"                       │                    │
│  │  i: 0  ▲                                │                    │
│  │        └── current position             │                    │
│  └─────────────────────────────────────────┘                    │
│                                                                 │
│  After Read(buf) where len(buf)=5:                              │
│  ┌─────────────────────────────────────────┐                    │
│  │  s: "Hello World"                       │                    │
│  │  i: 5      ▲                            │                    │
│  │            └── position advanced        │                    │
│  │  buf = "Hello"                          │                    │
│  └─────────────────────────────────────────┘                    │
│                                                                 │
│  After another Read(buf):                                       │
│  ┌─────────────────────────────────────────┐                    │
│  │  s: "Hello World"                       │                    │
│  │  i: 10          ▲                       │                    │
│  │                 └── near end            │                    │
│  │  buf = " Worl"                          │                    │
│  └─────────────────────────────────────────┘                    │
│                                                                 │
│  After final Read(buf):                                         │
│  ┌─────────────────────────────────────────┐                    │
│  │  s: "Hello World"                       │                    │
│  │  i: 11           ▲                      │                    │
│  │                  └── at end             │                    │
│  │  n = 1, buf[0] = "d", err = io.EOF      │                    │
│  └─────────────────────────────────────────┘                    │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Practical Patterns

### Pattern 1: Read All Content

```go
// When you need everything in memory
data, err := io.ReadAll(reader)
if err != nil {
    return err
}
// data is []byte containing entire content
```

### Pattern 2: Copy Between Reader and Writer

```go
// Copy from reader to writer
written, err := io.Copy(writer, reader)
// Efficiently streams data without loading all into memory
```

### Pattern 3: Decode JSON Directly

```go
// Instead of reading then unmarshaling:
data, _ := io.ReadAll(r.Body)
json.Unmarshal(data, &result)  // Two steps, loads all into memory

// Decode directly from reader:
json.NewDecoder(r.Body).Decode(&result)  // One step, streams
```

### Pattern 4: Encode JSON Directly

```go
// Instead of marshaling then writing:
data, _ := json.Marshal(result)
w.Write(data)  // Two steps

// Encode directly to writer:
json.NewEncoder(w).Encode(result)  // One step, streams
```

### Pattern 5: Limited Reading

```go
// Read at most N bytes (prevents memory attacks)
limited := io.LimitReader(reader, 1024*1024)  // 1MB max
data, err := io.ReadAll(limited)
```

---

## HTTP Context: Request and Response Bodies

### Request Body (Handler Perspective)

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // r.Body is io.ReadCloser — a Reader you must close

    // Option 1: Decode JSON directly
    var input MyStruct
    err := json.NewDecoder(r.Body).Decode(&input)

    // Option 2: Read all bytes
    data, err := io.ReadAll(r.Body)

    // Note: r.Body can only be read ONCE (it's a stream)
}
```

### Response Writer (Handler Perspective)

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // w is http.ResponseWriter — implements io.Writer

    // Option 1: Write raw bytes
    w.Write([]byte("Hello"))

    // Option 2: Encode JSON directly
    json.NewEncoder(w).Encode(result)

    // Option 3: Use fmt
    fmt.Fprintf(w, "User: %s", name)
}
```

### Testing: Creating Request Bodies

```go
// From string (e.g., JSON)
body := strings.NewReader(`{"name": "Alice"}`)
req := httptest.NewRequest("POST", "/users", body)

// From struct (marshal first)
data, _ := json.Marshal(myStruct)
body := bytes.NewReader(data)
req := httptest.NewRequest("POST", "/users", body)

// Empty body (for GET requests)
req := httptest.NewRequest("GET", "/users", nil)
```

### The Complete Request/Response Flow

```
┌──────────────────────────────────────────────────────────────────┐
│  HTTP Handler I/O Flow                                           │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  REQUEST (reading from client):                                  │
│                                                                  │
│  Client ──► Network ──► r.Body (io.ReadCloser) ──► Your Handler  │
│                              │                                   │
│                              ▼                                   │
│                    json.NewDecoder(r.Body).Decode(&input)        │
│                              │                                   │
│                              ▼                                   │
│                    Bytes flow from network into your struct      │
│                                                                  │
│  ──────────────────────────────────────────────────────────────  │
│                                                                  │
│  RESPONSE (writing to client):                                   │
│                                                                  │
│  Your Handler ──► w (http.ResponseWriter) ──► Network ──► Client │
│       │                                                          │
│       ▼                                                          │
│  json.NewEncoder(w).Encode(result)                               │
│       │                                                          │
│       ▼                                                          │
│  Bytes flow from your struct to network                          │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

## Common Pitfalls

### Pitfall 1: Reading Body Twice

```go
// WRONG: Body is a stream, can only read once
data1, _ := io.ReadAll(r.Body)
data2, _ := io.ReadAll(r.Body)  // data2 is empty!

// FIX: Read once, reuse the data
data, _ := io.ReadAll(r.Body)
// Use data multiple times
```

### Pitfall 2: Forgetting to Close

```go
// WRONG: Resource leak
resp, _ := http.Get(url)
data, _ := io.ReadAll(resp.Body)
// resp.Body never closed!

// FIX: Always close
resp, _ := http.Get(url)
defer resp.Body.Close()
data, _ := io.ReadAll(resp.Body)
```

### Pitfall 3: Using Read() Directly

```go
// WRONG: Read() may not read everything
buf := make([]byte, 100)
n, _ := reader.Read(buf)  // n might be < 100 even if more data exists

// FIX: Use io.ReadAll or io.ReadFull
data, _ := io.ReadAll(reader)  // Reads everything

// Or if you need exactly N bytes:
buf := make([]byte, 100)
_, err := io.ReadFull(reader, buf)  // Reads exactly 100 or errors
```

### Pitfall 4: Ignoring Bytes Before Error

```go
// WRONG: Ignoring n when err != nil
n, err := reader.Read(buf)
if err != nil {
    return err  // Might discard valid data!
}

// FIX: Process bytes first
n, err := reader.Read(buf)
if n > 0 {
    process(buf[:n])  // Handle the data
}
if err != nil {
    if err == io.EOF {
        return nil  // Normal completion
    }
    return err  // Actual error
}
```

### Pitfall 5: Buffer Pollution

```go
// CAREFUL: Buffer retains old data after short reads
buf := make([]byte, 10)
reader.Read(buf)  // Reads "Hello" (5 bytes)
// buf = ['H','e','l','l','o', 0, 0, 0, 0, 0]

reader.Read(buf)  // Reads "!" (1 byte)
// buf = ['!','e','l','l','o', 0, 0, 0, 0, 0]
//        ▲   └─────┬─────┘
//        │         └── OLD DATA still there!
//        └── new data

// FIX: Always use buf[:n]
n, _ := reader.Read(buf)
process(buf[:n])  // Only use the bytes that were actually read
```

---

## Quick Reference

### Creating Readers

| From | Code |
|------|------|
| String | `strings.NewReader("text")` |
| []byte | `bytes.NewReader(data)` |
| File | `os.Open("path")` |
| Buffer | `bytes.NewBuffer(data)` |

### Creating Writers

| To | Code |
|----|------|
| Buffer | `var buf bytes.Buffer` |
| File | `os.Create("path")` |
| Discard | `io.Discard` |

### Reading Patterns

| Need | Use |
|------|-----|
| All content | `io.ReadAll(r)` |
| Exactly N bytes | `io.ReadFull(r, buf)` |
| Limited amount | `io.LimitReader(r, max)` |
| JSON decode | `json.NewDecoder(r).Decode(&v)` |

### Writing Patterns

| Need | Use |
|------|-----|
| Raw bytes | `w.Write(data)` |
| String | `io.WriteString(w, s)` |
| Formatted | `fmt.Fprintf(w, format, args...)` |
| JSON encode | `json.NewEncoder(w).Encode(v)` |
| Copy stream | `io.Copy(w, r)` |

### HTTP Specifics

| Context | Reader | Writer |
|---------|--------|--------|
| Handler | `r.Body` | `w` (ResponseWriter) |
| Client | `resp.Body` | Request body via `bytes.Reader` |
| Testing | `strings.NewReader` | `httptest.NewRecorder().Body` |

---

## Summary

| Concept | Key Point |
|---------|-----------|
| `io.Reader` | "I can provide bytes on demand" |
| `io.Writer` | "I can consume bytes you give me" |
| Why interfaces? | Abstraction, streaming, memory efficiency |
| Lazy evaluation | Data read only when `Read()` is called |
| Streams are one-way | Can only read/write once (generally) |
| Always close | `defer body.Close()` for resources |
| Use helpers | `io.ReadAll`, `io.Copy`, `json.NewDecoder` |
