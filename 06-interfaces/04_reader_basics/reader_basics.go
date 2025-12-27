package reader_basics

import (
	"io"
)

// TODO(human): Define RepeatReader struct
type RepeatReader struct {
	Byte  byte
	Count int
	read  int
}

// TODO(human): Define CountingReader struct
type CountingReader struct {
	Reader    io.Reader
	BytesRead int
}

// TODO(human): Implement Read() method for RepeatReader
func (r *RepeatReader) Read(p []byte) (n int, err error) {

	if r.read >= r.Count {
		return 0, io.EOF
	}

	count := 0

	for i := range r.Count - r.read {
		if i < len(p) {
			p[i] = r.Byte
			r.read++
			count++
		} else {
			break
		}
	}

	if r.read >= r.Count {
		return count, io.EOF
	}

	// return count, io.EOF
	return count, nil
}

// TODO(human): Implement Read() method for CountingReader
func (r *CountingReader) Read(p []byte) (n int, err error) {

	n, err = r.Reader.Read(p)
	r.BytesRead += n

	if n < len(p) {
		return n, io.EOF
	}
	return n, err
}
