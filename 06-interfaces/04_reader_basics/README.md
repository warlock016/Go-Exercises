# Exercise 04: Reader Basics

**Concept:** Implementing io.Reader interface
**Difficulty:** Medium
**Estimated Time:** 45 minutes

## Learning Goal

Master the `io.Reader` interface, one of Go's most important and widely-used interfaces. Understanding Reader unlocks file I/O, network programming, and data processing patterns.

## The io.Reader Interface

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

**Contract:**
- Fill the byte slice `p` with data
- Return the number of bytes read (`n`)
- Return `io.EOF` when no more data available
- Can return both data and error in same call

## Your Task

Implement two types that satisfy io.Reader:

### 1. RepeatReader
Yields the same byte `n` times

**Fields:**
- `Byte byte` - the byte to repeat
- `Count int` - how many times to repeat
- `read int` - internal counter (how many read so far)

**Example:** RepeatReader{Byte: 'A', Count: 5} yields "AAAAA"

### 2. CountingReader
Wraps another Reader and counts bytes read

**Fields:**
- `Reader io.Reader` - the wrapped reader
- `BytesRead int` - total bytes read

**Example:** Wraps strings.NewReader("hello"), tracks that 5 bytes were read

## Function Signatures

```go
type RepeatReader struct {
    Byte  byte
    Count int
    read  int
}

type CountingReader struct {
    Reader    io.Reader
    BytesRead int
}

func (r *RepeatReader) Read(p []byte) (n int, err error)
func (r *CountingReader) Read(p []byte) (n int, err error)
```

## Examples

```go
// RepeatReader
r := &RepeatReader{Byte: 'A', Count: 5}
buf := make([]byte, 10)
n, err := r.Read(buf)
// n = 5, buf[:n] = "AAAAA", err = io.EOF

// CountingReader
sr := strings.NewReader("hello")
cr := &CountingReader{Reader: sr}
buf := make([]byte, 10)
n, err := cr.Read(buf)
// n = 5, buf[:n] = "hello", cr.BytesRead = 5
```

## Instructions

1. Open `reader_basics.go`
2. Import "io" package
3. Define both structs
4. Implement Read() for RepeatReader
5. Implement Read() for CountingReader
6. Run `go test -v`

## Hints

<details>
<summary>Hint 1: RepeatReader.Read Logic</summary>

```go
func (r *RepeatReader) Read(p []byte) (n int, err error) {
    if r.read >= r.Count {
        return 0, io.EOF
    }

    for i := 0; i < len(p) && r.read < r.Count; i++ {
        p[i] = r.Byte
        r.read++
        n++
    }

    if r.read >= r.Count {
        err = io.EOF
    }
    return n, err
}
```
</details>

<details>
<summary>Hint 2: CountingReader.Read Logic</summary>

```go
func (r *CountingReader) Read(p []byte) (n int, err error) {
    n, err = r.Reader.Read(p)
    r.BytesRead += n
    return n, err
}
```
</details>

## Think About

1. **Why use pointer receiver?**
   - Read modifies internal state (read counter, BytesRead)
   - Multiple calls must see previous state changes

2. **Can you return data AND error?**
   - Yes! Read can return (5, io.EOF) - read 5 bytes and no more data
   - This is valid and common

3. **What if p is larger than remaining data?**
   - Return only what's available
   - Don't wait for buffer to fill

## What This Teaches

- io.Reader contract
- Pointer receivers for stateful types
- Wrapping readers (decorator pattern)
- Buffer management
- EOF handling

---

**Next up:** Exercise 05 - Writer Basics
