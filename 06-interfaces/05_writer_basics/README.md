# Exercise 05: Writer Basics

**Concept:** Implementing io.Writer interface
**Difficulty:** Medium
**Estimated Time:** 45 minutes

## Learning Goal

Master the `io.Writer` interface, the complement to io.Reader. Writers are used for files, network connections, buffers, and data transformation.

## The io.Writer Interface

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

**Contract:**
- Write len(p) bytes from p
- Return number of bytes written (n)
- Return error if n < len(p)
- Must not modify p, even temporarily

## Your Task

Implement two types that satisfy io.Writer:

### 1. UpperWriter
Converts bytes to uppercase before writing to underlying writer

**Fields:**
- `Writer io.Writer` - destination writer

### 2. LimitWriter
Stops writing after n bytes

**Fields:**
- `Writer io.Writer` - destination writer
- `Limit int` - max bytes to write
- `written int` - bytes written so far

## Function Signatures

```go
type UpperWriter struct {
    Writer io.Writer
}

type LimitWriter struct {
    Writer  io.Writer
    Limit   int
    written int
}

func (w *UpperWriter) Write(p []byte) (n int, err error)
func (w *LimitWriter) Write(p []byte) (n int, err error)
```

## Examples

```go
// UpperWriter
var buf bytes.Buffer
uw := &UpperWriter{Writer: &buf}
uw.Write([]byte("hello"))
// buf.String() = "HELLO"

// LimitWriter
var buf bytes.Buffer
lw := &LimitWriter{Writer: &buf, Limit: 5}
lw.Write([]byte("hello world"))
// buf.String() = "hello" (stopped at 5 bytes)
```

## Instructions

1. Open `writer_basics.go`
2. Import "io" and "bytes" packages
3. Define both structs
4. Implement Write() methods
5. Run `go test -v`

## Hints

<details>
<summary>Hint 1: UpperWriter Logic</summary>

Convert to uppercase before writing:

```go
func (w *UpperWriter) Write(p []byte) (n int, err error) {
    upper := bytes.ToUpper(p)
    return w.Writer.Write(upper)
}
```
</details>

<details>
<summary>Hint 2: LimitWriter Logic</summary>

Only write up to the limit:

```go
func (w *LimitWriter) Write(p []byte) (n int, err error) {
    remaining := w.Limit - w.written
    if remaining <= 0 {
        return 0, io.ErrShortWrite
    }

    toWrite := len(p)
    if toWrite > remaining {
        toWrite = remaining
    }

    n, err = w.Writer.Write(p[:toWrite])
    w.written += n
    return n, err
}
```
</details>

## What This Teaches

- io.Writer contract
- Writer composition/wrapping
- Data transformation in writers
- Limiting/controlling output

---

**Next up:** Exercise 06 - Interface Composition
