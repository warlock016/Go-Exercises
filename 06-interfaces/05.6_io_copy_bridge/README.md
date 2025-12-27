# Exercise 05.6: The io.Copy Bridge

**Concept:** Connecting Readers and Writers with io.Copy
**Difficulty:** Medium
**Estimated Time:** 50 minutes

## Learning Goal

Understand how `io.Copy` bridges Readers and Writers, enabling powerful data processing pipelines. This is one of Go's most important patterns for I/O.

## The Missing Link

In exercises 04 and 05, you learned:
- **Readers** pull data FROM a source
- **Writers** push data TO a destination

But how do you connect them? Enter `io.Copy`:

```go
// io.Copy signature
func Copy(dst Writer, src Reader) (written int64, err error)
```

It reads from `src` and writes to `dst` until EOF.

## The io.Copy Pattern

```
┌─────────────────────────────────────────────────────────────────────────┐
│                                                                          │
│  WITHOUT io.Copy (manual loop):                                         │
│  ───────────────────────────────                                        │
│                                                                          │
│  buf := make([]byte, 4096)                                              │
│  for {                                                                   │
│      n, readErr := src.Read(buf)                                        │
│      if n > 0 {                                                          │
│          _, writeErr := dst.Write(buf[:n])                              │
│          if writeErr != nil { return writeErr }                         │
│      }                                                                   │
│      if readErr == io.EOF { break }                                     │
│      if readErr != nil { return readErr }                               │
│  }                                                                       │
│                                                                          │
│  WITH io.Copy (one line!):                                              │
│  ─────────────────────────                                              │
│                                                                          │
│  io.Copy(dst, src)  // Does all the above automatically                 │
│                                                                          │
└─────────────────────────────────────────────────────────────────────────┘
```

## Real-World Power

```go
// Copy file to file
src, _ := os.Open("input.txt")
dst, _ := os.Create("output.txt")
io.Copy(dst, src)

// Download from network to file
resp, _ := http.Get("https://example.com/data")
file, _ := os.Create("download.dat")
io.Copy(file, resp.Body)

// With transformations (using wrappers!)
src, _ := os.Open("input.txt")
upper := &UppercaseReader{Reader: src}
dst, _ := os.Create("output.txt")
limited := &LimitWriter{Writer: dst, Limit: 1000}
io.Copy(limited, upper)  // Uppercase + limit in one pipeline!
```

## Your Task

Create a data processing function that connects Readers and Writers with transformations.

### Types

#### ProcessingStats
Tracks what happened during processing.

```go
type ProcessingStats struct {
    BytesRead    int64
    BytesWritten int64
    Uppercase    bool
    Limited      bool
}
```

#### ProgressWriter
Wraps a Writer and reports progress via a callback.

**Fields:**
- `Writer io.Writer` - destination
- `OnProgress func(bytesWritten int64)` - callback for each write

### Functions

#### 1. CopyWithStats
Copy from Reader to Writer, returning statistics.

```go
func CopyWithStats(dst io.Writer, src io.Reader) (*ProcessingStats, error)
```

#### 2. ProcessData
Full pipeline: read → uppercase → count → write with limit.

```go
func ProcessData(dst io.Writer, src io.Reader, opts ProcessOptions) (*ProcessingStats, error)

type ProcessOptions struct {
    Uppercase  bool  // Transform to uppercase
    Limit      int   // Limit output bytes (0 = no limit)
}
```

#### 3. CopyN
Copy exactly n bytes (recreate io.CopyN for practice).

```go
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error)
```

## Function Signatures

```go
type ProcessingStats struct {
    BytesRead    int64
    BytesWritten int64
}

type ProcessOptions struct {
    Uppercase bool
    Limit     int
}

type ProgressWriter struct {
    Writer     io.Writer
    OnProgress func(bytesWritten int64)
    written    int64
}

func (w *ProgressWriter) Write(p []byte) (n int, err error)

func CopyWithStats(dst io.Writer, src io.Reader) (*ProcessingStats, error)
func ProcessData(dst io.Writer, src io.Reader, opts ProcessOptions) (*ProcessingStats, error)
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error)
```

## Examples

```go
// Basic copy with stats
src := strings.NewReader("hello world")
var dst bytes.Buffer
stats, _ := CopyWithStats(&dst, src)
// dst.String() = "hello world"
// stats.BytesRead = 11, stats.BytesWritten = 11

// Process with options
src := strings.NewReader("hello world")
var dst bytes.Buffer
stats, _ := ProcessData(&dst, src, ProcessOptions{
    Uppercase: true,
    Limit:     5,
})
// dst.String() = "HELLO"
// stats.BytesRead = 5, stats.BytesWritten = 5

// Copy exactly N bytes
src := strings.NewReader("hello world")
var dst bytes.Buffer
n, _ := CopyN(&dst, src, 5)
// dst.String() = "hello", n = 5

// Progress reporting
src := strings.NewReader("hello world")
var dst bytes.Buffer
progress := &ProgressWriter{
    Writer: &dst,
    OnProgress: func(n int64) {
        fmt.Printf("Written %d bytes so far\n", n)
    },
}
io.Copy(progress, src)
```

## The Pipeline Visualization

```
┌─────────────────────────────────────────────────────────────────────────┐
│                                                                          │
│  ProcessData with Uppercase=true, Limit=5:                              │
│                                                                          │
│  ┌─────────┐      ┌───────────┐      ┌─────────┐      ┌─────────┐      │
│  │  Source │─────▶│ Uppercase │─────▶│ Counter │─────▶│  Limit  │      │
│  │ Reader  │      │  Reader   │      │ Reader  │      │ Reader  │      │
│  └─────────┘      └───────────┘      └─────────┘      └─────────┘      │
│       │                                                    │            │
│       │                    io.Copy                         │            │
│       │           ┌───────────────────┐                    │            │
│       └───────────│  Connects them!   │◀───────────────────┘            │
│                   └─────────┬─────────┘                                 │
│                             │                                            │
│                             ▼                                            │
│                   ┌─────────────────┐                                   │
│                   │   Destination   │                                   │
│                   │     Writer      │                                   │
│                   └─────────────────┘                                   │
│                                                                          │
└─────────────────────────────────────────────────────────────────────────┘
```

## Hints

<details>
<summary>Hint 1: CopyWithStats Pattern</summary>

Use CountingReader and CountingWriter to track bytes:

```go
func CopyWithStats(dst io.Writer, src io.Reader) (*ProcessingStats, error) {
    countingSrc := &CountingReader{Reader: src}
    countingDst := &CountingWriter{Writer: dst}

    _, err := io.Copy(countingDst, countingSrc)

    return &ProcessingStats{
        BytesRead:    int64(countingSrc.BytesRead),
        BytesWritten: int64(countingDst.BytesWritten),
    }, err
}
```
</details>

<details>
<summary>Hint 2: ProcessData Pattern</summary>

Build the pipeline conditionally:

```go
func ProcessData(dst io.Writer, src io.Reader, opts ProcessOptions) (*ProcessingStats, error) {
    var pipeline io.Reader = src

    if opts.Uppercase {
        pipeline = &UppercaseReader{Reader: pipeline}
    }

    countingReader := &CountingReader{Reader: pipeline}
    pipeline = countingReader

    if opts.Limit > 0 {
        pipeline = &LimitReader{Reader: pipeline, Limit: opts.Limit}
    }

    countingWriter := &CountingWriter{Writer: dst}
    _, err := io.Copy(countingWriter, pipeline)

    return &ProcessingStats{
        BytesRead:    int64(countingReader.BytesRead),
        BytesWritten: int64(countingWriter.BytesWritten),
    }, err
}
```
</details>

<details>
<summary>Hint 3: CopyN Pattern</summary>

Use LimitReader to limit the source:

```go
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
    limited := &LimitReader{Reader: src, Limit: int(n)}
    return io.Copy(dst, limited)
}
```
</details>

## Think About

1. **Why use io.Copy?** Why not just write the loop yourself every time?
   - io.Copy handles edge cases, buffer management, and is optimized

2. **Which side to limit?** You can limit on Reader side (LimitReader) or Writer side (LimitWriter). When would you choose each?

3. **Error handling:** What happens if the Writer returns an error mid-copy? What about the Reader?

4. **Buffer size:** io.Copy uses an internal buffer (32KB by default). How might this affect behavior?

## What This Teaches

- io.Copy as the bridge between Reader and Writer
- Building conditional pipelines
- Combining read-side and write-side wrappers
- Real-world I/O patterns
- The power of interface composition

## Real-World Applications

```go
// HTTP file upload
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    file, _ := os.Create("upload.dat")
    limited := &LimitWriter{Writer: file, Limit: 10*1024*1024} // 10MB limit
    io.Copy(limited, r.Body)  // Safe: won't exceed 10MB
}

// Log rotation
func rotateLog(src io.Reader, dst io.Writer) {
    compressed := gzip.NewWriter(dst)
    io.Copy(compressed, src)
    compressed.Close()
}

// Data migration
func migrate(from, to Database) {
    reader := from.Export()
    writer := to.Import()
    io.Copy(writer, reader)
}
```

---

**Next up:** Exercise 06 - Interface Composition
