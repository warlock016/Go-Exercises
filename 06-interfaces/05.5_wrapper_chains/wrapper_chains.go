package wrapper_chains

import (
	"io"
	"strings"
)

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

// TeeReader wraps a Reader and copies all data to a Writer
type TeeReader struct {
	Reader io.Reader
	Writer io.Writer
}

// CountingReader wraps a Reader and counts bytes read
type CountingReader struct {
	Reader    io.Reader
	BytesRead int
}

// Stats provides visibility into pipeline internals
type Stats struct {
	Counted *int
}

// Read implements io.Reader for UppercaseReader
func (r *UppercaseReader) Read(p []byte) (n int, err error) {
	// TODO(human): Implement
	upper := strings.ToUpper(string(p))
	return r.Reader.Read([]byte(upper))
}

// Read implements io.Reader for LimitReader
func (r *LimitReader) Read(p []byte) (n int, err error) {
	// TODO(human): Implement
	remaining := r.Limit - r.read
	if remaining <= 0 {
		return 0, io.ErrShortBuffer
	}

	return r.Reader.Read(p[:remaining])
}

// Read implements io.Reader for TeeReader
func (r *TeeReader) Read(p []byte) (n int, err error) {
	// TODO(human): Implement
	return r.Reader.Read(p)
}

// Read implements io.Reader for CountingReader
func (r *CountingReader) Read(p []byte) (n int, err error) {
	// TODO(human): Implement
	return 0, nil
}

// BuildPipeline creates a chained Reader: source → uppercase → count → limit
func BuildPipeline(source io.Reader, limit int) (io.Reader, *Stats) {
	// TODO(human): Implement
	return nil, nil
}
