# Exercise 05.5: Wrapper Chains

**Concept:** Chaining multiple wrapper Readers to build data transformation pipelines
**Difficulty:** Medium
**Estimated Time:** 55 minutes

## Learning Goal

Understand how wrapper Readers compose together to form data processing pipelines. This is how real Go programs process data: chain Readers together, each adding a layer of behavior.

## The Problem

In exercises 04 and 05, you learned wrappers in isolation. But the real power comes from **chaining** them:

```go
// Single wrapper (what you learned)
counted := &CountingReader{Reader: source}

// Chained wrappers (what this exercise teaches)
source := strings.NewReader("hello world")
upper := &UppercaseReader{Reader: source}      // Layer 1: transform to uppercase
counted := &CountingReader{Reader: upper}       // Layer 2: count bytes
limited := &LimitReader{Reader: counted, Limit: 5}  // Layer 3: limit output

// Data flows: source → upper → counted → limited → caller
```

## How Data Flows Through Chains

```
┌─────────────────────────────────────────────────────────────────────────┐
│                                                                         │
│  When caller calls: limited.Read(buf)                                   │
│                                                                         │
│  The call chain:                                                        │
│  ┌─────────────────────────────────────────────────────────────────┐    │
│  │ Caller                                                          │    │
│  │   │                                                             │    │
│  │   │ limited.Read(buf)                                           │    │
│  │   ▼                                                             │    │
│  │ LimitReader                                                     │    │
│  │   │ "I'll only return up to Limit bytes"                        │    │
│  │   │ counted.Read(buf)                                           │    │
│  │   ▼                                                             │    │
│  │ CountingReader                                                  │    │
│  │   │ "I'll count what passes through"                            │    │
│  │   │ upper.Read(buf)                                             │    │
│  │   ▼                                                             │    │
│  │ UppercaseReader                                                 │    │
│  │   │ "I'll uppercase the data"                                   │    │
│  │   │ source.Read(buf)                                            │    │
│  │   ▼                                                             │    │
│  │ strings.Reader                                                  │    │
│  │   └── "hello world" (the actual data)                           │    │
│  └─────────────────────────────────────────────────────────────────┘    │
│                                                                         │
│  Data returns: "hello world" → "HELLO WORLD" → counted → "HELLO"        │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

## Your Task

Implement three wrapper Readers and a function that chains them.

### Types

#### 1. UppercaseReader
Wraps a Reader and transforms all bytes to uppercase.

**Fields:**
- `Reader io.Reader` - the source Reader

**Behavior:**
- Read from wrapped Reader
- Convert bytes to uppercase (use `bytes.ToUpper`)
- Return the uppercase data

#### 2. LimitReader
Wraps a Reader and stops reading after N bytes total.

**Fields:**
- `Reader io.Reader` - the source Reader
- `Limit int` - maximum bytes to read (total across all calls)
- `read int` - bytes read so far (internal counter)

**Behavior:**
- Track cumulative bytes read
- Stop reading when limit reached
- Return `io.EOF` when limit hit

#### 3. TeeReader (Bonus)
Wraps a Reader and writes everything it reads to a Writer.

**Fields:**
- `Reader io.Reader` - the source Reader
- `Writer io.Writer` - where to copy data

**Behavior:**
- Read from wrapped Reader
- Write the same data to Writer
- Return what was read

### Function

**BuildPipeline(source io.Reader, limit int) (io.Reader, *Stats)**

Creates a pipeline: source → UppercaseReader → CountingReader → LimitReader

Returns:
- The final Reader (LimitReader) for the caller to read from
- A Stats struct to observe what happened

```go
type Stats struct {
    Counted *int  // Pointer to CountingReader's count
}
```

## Function Signatures

```go
type UppercaseReader struct {
    Reader io.Reader
}

type LimitReader struct {
    Reader io.Reader
    Limit  int
    read   int
}

type TeeReader struct {
    Reader io.Reader
    Writer io.Writer
}

type CountingReader struct {
    Reader    io.Reader
    BytesRead int
}

type Stats struct {
    Counted *int
}

func (r *UppercaseReader) Read(p []byte) (n int, err error)
func (r *LimitReader) Read(p []byte) (n int, err error)
func (r *TeeReader) Read(p []byte) (n int, err error)
func (r *CountingReader) Read(p []byte) (n int, err error)

func BuildPipeline(source io.Reader, limit int) (io.Reader, *Stats)
```

## Examples

```go
// Basic uppercase reader
source := strings.NewReader("hello")
upper := &UppercaseReader{Reader: source}
buf := make([]byte, 10)
n, _ := upper.Read(buf)
// buf[:n] = "HELLO"

// Chained pipeline
source := strings.NewReader("hello world")
pipeline, stats := BuildPipeline(source, 5)
result, _ := io.ReadAll(pipeline)
// result = "HELLO" (limited to 5 bytes, uppercased)
// *stats.Counted = 5 (bytes that passed through counter)

// Manual chain
source := strings.NewReader("go is awesome")
upper := &UppercaseReader{Reader: source}
limited := &LimitReader{Reader: upper, Limit: 10}
result, _ := io.ReadAll(limited)
// result = "GO IS AWES" (10 bytes, uppercased)
```

## Real-World Analogy

Think of a water treatment plant:

```
┌─────────────────────────────────────────────────────────────────────────┐
│                                                                          │
│  River  →  Filter  →  Chlorinate  →  Meter  →  Limit Valve  →  Faucet  │
│  (source)  (clean)    (treat)       (count)    (limit flow)   (output)  │
│                                                                          │
│  In code:                                                                │
│  source  →  upper  →  transform  →  count  →  limit  →  io.ReadAll     │
│                                                                          │
└─────────────────────────────────────────────────────────────────────────┘
```

Each stage adds behavior without knowing about the others.

## Hints

<details>
<summary>Hint 1: UppercaseReader Pattern</summary>

```go
func (r *UppercaseReader) Read(p []byte) (n int, err error) {
    n, err = r.Reader.Read(p)
    if n > 0 {
        // Transform the bytes that were read
        copy(p[:n], bytes.ToUpper(p[:n]))
    }
    return n, err
}
```
</details>

<details>
<summary>Hint 2: LimitReader Pattern</summary>

```go
func (r *LimitReader) Read(p []byte) (n int, err error) {
    if r.read >= r.Limit {
        return 0, io.EOF
    }

    // Calculate how much we can still read
    remaining := r.Limit - r.read
    if len(p) > remaining {
        p = p[:remaining]  // Shrink the buffer
    }

    n, err = r.Reader.Read(p)
    r.read += n
    return n, err
}
```
</details>

<details>
<summary>Hint 3: BuildPipeline Pattern</summary>

```go
func BuildPipeline(source io.Reader, limit int) (io.Reader, *Stats) {
    upper := &UppercaseReader{Reader: source}
    counter := &CountingReader{Reader: upper}
    limited := &LimitReader{Reader: counter, Limit: limit}

    return limited, &Stats{Counted: &counter.BytesRead}
}
```
</details>

## Think About

1. **Order matters:** What's the difference between `upper → limit` vs `limit → upper`?
   - One uppercases then limits; the other limits then uppercases
   - The data transformation order is reversed!

2. **Why pointers in Stats?** Because the counter updates after you create Stats. A pointer lets you see the updated value.

3. **Can wrappers be reused?** What happens if you try to read from the same pipeline twice?

4. **Error propagation:** If the innermost reader returns an error, how does it bubble up?

## What This Teaches

- How wrapper Readers chain together
- Data transformation pipelines
- The decorator pattern in action
- Why io.Reader is so powerful (composability!)
- Real-world patterns used in compression, encryption, logging

---

**Next up:** Exercise 05.6 - The io.Copy Bridge
