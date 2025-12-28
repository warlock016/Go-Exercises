package io_copy_bridge

import (
	"io"
	"strings"
)

// ProcessingStats tracks bytes processed during copy operations
type ProcessingStats struct {
	BytesRead    int64
	BytesWritten int64
}

// ProcessOptions configures the processing pipeline
type ProcessOptions struct {
	Uppercase bool
	Limit     int
}

// ProgressWriter wraps a Writer and reports progress via callback
type ProgressWriter struct {
	Writer     io.Writer
	OnProgress func(bytesWritten int64)
	written    int64
}

// CountingReader wraps a Reader and counts bytes read
type CountingReader struct {
	Reader    io.Reader
	BytesRead int
}

// CountingWriter wraps a Writer and counts bytes written
type CountingWriter struct {
	Writer       io.Writer
	BytesWritten int
}

// UppercaseReader wraps a Reader and transforms bytes to uppercase
type UppercaseReader struct {
	Reader io.Reader
}

// LimitReader wraps a Reader and limits total bytes read
type LimitReader struct {
	Reader io.Reader
	Limit  int
	read   int
}

// Write implements io.Writer for ProgressWriter
func (w *ProgressWriter) Write(p []byte) (n int, err error) {
	// TODO(human): Implement
	n, err = w.Writer.Write(p)
	w.written += int64(n)
	w.OnProgress(w.written)
	return n, err
}

// Read implements io.Reader for CountingReader
func (r *CountingReader) Read(p []byte) (n int, err error) {
	// TODO(human): Implement
	n, err = r.Reader.Read(p)
	r.BytesRead += n
	return n, err
}

// Write implements io.Writer for CountingWriter
func (w *CountingWriter) Write(p []byte) (n int, err error) {
	// TODO(human): Implement
	n, err = w.Writer.Write(p)
	w.BytesWritten += n
	return n, err
}

// Read implements io.Reader for UppercaseReader
func (r *UppercaseReader) Read(p []byte) (n int, err error) {
	// TODO(human): Implement
	n, err = r.Reader.Read(p)
	upper := strings.ToUpper(string(p[:n]))
	copy(p, []byte(upper))
	return len([]byte(upper)), err
}

// Read implements io.Reader for LimitReader
func (r *LimitReader) Read(p []byte) (n int, err error) {
	// TODO(human): Implement
	remaining := r.Limit - r.read
	if remaining <= 0 {
		return 0, io.EOF
	}
	limit := min(remaining, len(p))
	n, err = r.Reader.Read(p[:limit])
	r.read += n
	return n, err
}

// CopyWithStats copies from src to dst and returns statistics
func CopyWithStats(dst io.Writer, src io.Reader) (*ProcessingStats, error) {
	// TODO(human): Implement

	wr, err := io.Copy(dst, src)
	return &ProcessingStats{
		BytesRead:    wr,
		BytesWritten: wr,
	}, err
}

// ProcessData copies with optional transformations (uppercase, limit)
func ProcessData(dst io.Writer, src io.Reader, opts ProcessOptions) (*ProcessingStats, error) {
	// TODO(human): Implement

	var pipeline io.Reader = src

	if opts.Uppercase {
		pipeline = &UppercaseReader{Reader: pipeline}
	}

	countingReader := &CountingReader{Reader: pipeline}
	pipeline = countingReader
	// pipeline = &CountingReader{Reader: pipeline}

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

// CopyN copies exactly n bytes from src to dst
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
	// TODO(human): Implement
	return io.CopyN(dst, src, n)
}
