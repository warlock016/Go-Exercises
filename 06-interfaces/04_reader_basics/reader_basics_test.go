package reader_basics

import (
	"io"
	"strings"
	"testing"
)

func TestRepeatReader(t *testing.T) {
	tests := []struct {
		name      string
		byte      byte
		count     int
		bufSize   int
		wantData  string
		wantCount int
		wantErr   error
	}{
		{"Repeat A 5 times", 'A', 5, 10, "AAAAA", 5, io.EOF},
		{"Repeat X 3 times", 'X', 3, 10, "XXX", 3, io.EOF},
		{"Exact buffer size", 'B', 5, 5, "BBBBB", 5, io.EOF},
		{"Zero count", 'C', 0, 10, "", 0, io.EOF},
		{"Small buffer", 'D', 10, 3, "DDD", 3, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RepeatReader{Byte: tt.byte, Count: tt.count}
			buf := make([]byte, tt.bufSize)

			n, err := r.Read(buf)

			if n != tt.wantCount {
				t.Errorf("Read() n = %d, want %d", n, tt.wantCount)
			}

			if string(buf[:n]) != tt.wantData {
				t.Errorf("Read() data = %q, want %q", string(buf[:n]), tt.wantData)
			}

			if err != tt.wantErr {
				t.Errorf("Read() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepeatReaderMultipleCalls(t *testing.T) {
	r := &RepeatReader{Byte: 'A', Count: 10}
	buf := make([]byte, 3)

	// First read: AAA
	n, err := r.Read(buf)
	if n != 3 || string(buf[:n]) != "AAA" || err != nil {
		t.Errorf("First read: n=%d, data=%q, err=%v", n, string(buf[:n]), err)
	}

	// Second read: AAA
	n, err = r.Read(buf)
	if n != 3 || string(buf[:n]) != "AAA" || err != nil {
		t.Errorf("Second read: n=%d, data=%q, err=%v", n, string(buf[:n]), err)
	}

	// Third read: AAA
	n, err = r.Read(buf)
	if n != 3 || string(buf[:n]) != "AAA" || err != nil {
		t.Errorf("Third read: n=%d, data=%q, err=%v", n, string(buf[:n]), err)
	}

	// Fourth read: A (only 1 left)
	n, err = r.Read(buf)
	if n != 1 || string(buf[:n]) != "A" || err != io.EOF {
		t.Errorf("Fourth read: n=%d, data=%q, err=%v", n, string(buf[:n]), err)
	}

	// Fifth read: EOF
	n, err = r.Read(buf)
	if n != 0 || err != io.EOF {
		t.Errorf("Fifth read: n=%d, err=%v, want n=0, err=io.EOF", n, err)
	}
}

func TestCountingReader(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		bufSize      int
		wantData     string
		wantN        int
		wantTotal    int
	}{
		{"Read hello", "hello", 10, "hello", 5, 5},
		{"Read with exact buffer", "test", 4, "test", 4, 4},
		{"Read with large buffer", "hi", 10, "hi", 2, 2},
		{"Empty string", "", 10, "", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := strings.NewReader(tt.input)
			cr := &CountingReader{Reader: sr}
			buf := make([]byte, tt.bufSize)

			n, _ := cr.Read(buf)

			if n != tt.wantN {
				t.Errorf("Read() n = %d, want %d", n, tt.wantN)
			}

			if string(buf[:n]) != tt.wantData {
				t.Errorf("Read() data = %q, want %q", string(buf[:n]), tt.wantData)
			}

			if cr.BytesRead != tt.wantTotal {
				t.Errorf("BytesRead = %d, want %d", cr.BytesRead, tt.wantTotal)
			}
		})
	}
}

func TestCountingReaderMultipleReads(t *testing.T) {
	sr := strings.NewReader("hello world")
	cr := &CountingReader{Reader: sr}
	buf := make([]byte, 5)

	// First read: "hello"
	n, _ := cr.Read(buf)
	if n != 5 || string(buf[:n]) != "hello" || cr.BytesRead != 5 {
		t.Errorf("First read: n=%d, data=%q, total=%d", n, string(buf[:n]), cr.BytesRead)
	}

	// Second read: " worl"
	n, _ = cr.Read(buf)
	if n != 5 || string(buf[:n]) != " worl" || cr.BytesRead != 10 {
		t.Errorf("Second read: n=%d, data=%q, total=%d", n, string(buf[:n]), cr.BytesRead)
	}

	// Third read: "d"
	n, err := cr.Read(buf)
	if n != 1 || string(buf[:n]) != "d" || cr.BytesRead != 11 || err != io.EOF {
		t.Errorf("Third read: n=%d, data=%q, total=%d, err=%v", n, string(buf[:n]), cr.BytesRead, err)
	}
}

func TestCountingReaderWithRepeatReader(t *testing.T) {
	// CountingReader wrapping RepeatReader!
	rr := &RepeatReader{Byte: 'X', Count: 8}
	cr := &CountingReader{Reader: rr}
	buf := make([]byte, 3)

	// Read 3 times
	cr.Read(buf) // XXX
	cr.Read(buf) // XXX
	cr.Read(buf) // XX

	if cr.BytesRead != 8 {
		t.Errorf("BytesRead = %d, want 8", cr.BytesRead)
	}
}

func TestImplementsIOReader(t *testing.T) {
	var _ io.Reader = &RepeatReader{}
	var _ io.Reader = &CountingReader{}

	t.Log("✓ Both types implement io.Reader interface")
}

func TestIOCopyWithRepeatReader(t *testing.T) {
	// Test that our reader works with standard library functions
	r := &RepeatReader{Byte: 'Z', Count: 5}
	var buf strings.Builder

	n, err := io.Copy(&buf, r)
	if err != nil {
		t.Errorf("io.Copy error: %v", err)
	}
	if n != 5 {
		t.Errorf("io.Copy n = %d, want 5", n)
	}
	if buf.String() != "ZZZZZ" {
		t.Errorf("io.Copy result = %q, want %q", buf.String(), "ZZZZZ")
	}
}
