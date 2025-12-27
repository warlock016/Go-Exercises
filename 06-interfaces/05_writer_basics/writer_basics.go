package writer_basics

import (
	"io"
	"strings"
)

// TODO(human): Define UpperWriter struct
type UpperWriter struct {
	Writer io.Writer
}

// TODO(human): Define LimitWriter struct
type LimitWriter struct {
	Writer  io.Writer
	Limit   int
	written int
}

// TODO(human): Implement Write() method for UpperWriter
func (w *UpperWriter) Write(p []byte) (n int, err error) {

	upper := strings.ToUpper(string(p))
	return w.Writer.Write([]byte(upper))
}

// TODO(human): Implement Write() method for LimitWriter
func (w *LimitWriter) Write(p []byte) (n int, err error) {

	remaining := w.Limit - w.written
	if remaining <= 0 {
		return 0, io.ErrShortWrite
	}

	toWrite := min(len(p), remaining)
	n, err = w.Writer.Write(p[:toWrite])
	w.written += n
	return n, err
}
