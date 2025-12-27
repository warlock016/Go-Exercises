package writer_basics

import (
	"bytes"
	"io"
	"testing"
)

func TestUpperWriter(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Lowercase", "hello", "HELLO"},
		{"Mixed case", "HeLLo WoRLd", "HELLO WORLD"},
		{"Already upper", "HELLO", "HELLO"},
		{"With numbers", "test123", "TEST123"},
		{"Empty", "", ""},
		{"Special chars", "hello, world!", "HELLO, WORLD!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			uw := &UpperWriter{Writer: &buf}

			n, err := uw.Write([]byte(tt.input))
			if err != nil {
				t.Errorf("Write() error = %v", err)
			}
			if n != len(tt.input) {
				t.Errorf("Write() n = %d, want %d", n, len(tt.input))
			}
			if buf.String() != tt.want {
				t.Errorf("Write() result = %q, want %q", buf.String(), tt.want)
			}
		})
	}
}

func TestLimitWriter(t *testing.T) {
	tests := []struct {
		name        string
		limit       int
		input       string
		want        string
		wantN       int
	}{
		{"Within limit", 10, "hello", "hello", 5},
		{"Exact limit", 5, "hello", "hello", 5},
		{"Exceeds limit", 5, "hello world", "hello", 5},
		{"Zero limit", 0, "test", "", 0},
		{"Empty input", 10, "", "", 0},
		{"Multiple writes under limit", 10, "hello", "hello", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			lw := &LimitWriter{Writer: &buf, Limit: tt.limit}

			n, _ := lw.Write([]byte(tt.input))
			if n != tt.wantN {
				t.Errorf("Write() n = %d, want %d", n, tt.wantN)
			}
			if buf.String() != tt.want {
				t.Errorf("Write() result = %q, want %q", buf.String(), tt.want)
			}
		})
	}
}

func TestLimitWriterMultipleWrites(t *testing.T) {
	var buf bytes.Buffer
	lw := &LimitWriter{Writer: &buf, Limit: 10}

	// Write "hello" (5 bytes)
	n, err := lw.Write([]byte("hello"))
	if n != 5 || err != nil {
		t.Errorf("First write: n=%d, err=%v", n, err)
	}

	// Write " world" (6 bytes, but only 5 remaining)
	n, _ = lw.Write([]byte(" world"))
	if n != 5 {
		t.Errorf("Second write: n=%d, want 5", n)
	}

	// Total should be "hello worl" (10 bytes)
	if buf.String() != "hello worl" {
		t.Errorf("Buffer = %q, want %q", buf.String(), "hello worl")
	}

	// Next write should write 0 bytes
	n, _ = lw.Write([]byte("!"))
	if n != 0 {
		t.Errorf("Third write: n=%d, want 0", n)
	}
}

func TestUpperWriterWithMultipleWrites(t *testing.T) {
	var buf bytes.Buffer
	uw := &UpperWriter{Writer: &buf}

	uw.Write([]byte("hello "))
	uw.Write([]byte("world"))

	want := "HELLO WORLD"
	if buf.String() != want {
		t.Errorf("Multiple writes: got %q, want %q", buf.String(), want)
	}
}

func TestImplementsIOWriter(t *testing.T) {
	var _ io.Writer = &UpperWriter{}
	var _ io.Writer = &LimitWriter{}

	t.Log("✓ Both types implement io.Writer interface")
}

func TestIOCopyWithUpperWriter(t *testing.T) {
	input := bytes.NewReader([]byte("hello world"))
	var buf bytes.Buffer
	uw := &UpperWriter{Writer: &buf}

	n, err := io.Copy(uw, input)
	if err != nil {
		t.Errorf("io.Copy error: %v", err)
	}
	if n != 11 {
		t.Errorf("io.Copy n = %d, want 11", n)
	}
	if buf.String() != "HELLO WORLD" {
		t.Errorf("Result = %q, want %q", buf.String(), "HELLO WORLD")
	}
}
