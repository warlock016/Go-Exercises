package io_copy_bridge

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestCountingReader(t *testing.T) {
	source := strings.NewReader("hello world")
	counter := &CountingReader{Reader: source}

	buf := make([]byte, 5)
	counter.Read(buf)
	counter.Read(buf)
	counter.Read(buf)

	if counter.BytesRead != 11 {
		t.Errorf("BytesRead = %d, want 11", counter.BytesRead)
	}
}

func TestCountingWriter(t *testing.T) {
	var buf bytes.Buffer
	counter := &CountingWriter{Writer: &buf}

	counter.Write([]byte("hello"))
	counter.Write([]byte(" world"))

	if counter.BytesWritten != 11 {
		t.Errorf("BytesWritten = %d, want 11", counter.BytesWritten)
	}

	if buf.String() != "hello world" {
		t.Errorf("Buffer = %q, want %q", buf.String(), "hello world")
	}
}

func TestUppercaseReader(t *testing.T) {
	source := strings.NewReader("hello world")
	upper := &UppercaseReader{Reader: source}

	result, _ := io.ReadAll(upper)

	if string(result) != "HELLO WORLD" {
		t.Errorf("UppercaseReader = %q, want %q", string(result), "HELLO WORLD")
	}
}

func TestLimitReader(t *testing.T) {
	source := strings.NewReader("hello world")
	limited := &LimitReader{Reader: source, Limit: 5}

	result, _ := io.ReadAll(limited)

	if string(result) != "hello" {
		t.Errorf("LimitReader = %q, want %q", string(result), "hello")
	}
}

func TestProgressWriter(t *testing.T) {
	var buf bytes.Buffer
	var progressCalls []int64

	progress := &ProgressWriter{
		Writer: &buf,
		OnProgress: func(n int64) {
			progressCalls = append(progressCalls, n)
		},
	}

	progress.Write([]byte("hello"))
	progress.Write([]byte(" world"))

	if buf.String() != "hello world" {
		t.Errorf("Buffer = %q, want %q", buf.String(), "hello world")
	}

	if len(progressCalls) != 2 {
		t.Errorf("Progress callbacks = %d, want 2", len(progressCalls))
	}

	// Check cumulative progress
	if progressCalls[0] != 5 {
		t.Errorf("First progress = %d, want 5", progressCalls[0])
	}
	if progressCalls[1] != 11 {
		t.Errorf("Second progress = %d, want 11", progressCalls[1])
	}
}

func TestCopyWithStats(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantRead     int64
		wantWritten  int64
	}{
		{"Basic", "hello world", 11, 11},
		{"Empty", "", 0, 0},
		{"Long", strings.Repeat("x", 1000), 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := strings.NewReader(tt.input)
			var dst bytes.Buffer

			stats, err := CopyWithStats(&dst, src)

			if err != nil {
				t.Errorf("CopyWithStats error: %v", err)
			}

			if stats.BytesRead != tt.wantRead {
				t.Errorf("BytesRead = %d, want %d", stats.BytesRead, tt.wantRead)
			}

			if stats.BytesWritten != tt.wantWritten {
				t.Errorf("BytesWritten = %d, want %d", stats.BytesWritten, tt.wantWritten)
			}

			if dst.String() != tt.input {
				t.Errorf("Output = %q, want %q", dst.String(), tt.input)
			}
		})
	}
}

func TestProcessDataNoOptions(t *testing.T) {
	src := strings.NewReader("hello world")
	var dst bytes.Buffer

	stats, err := ProcessData(&dst, src, ProcessOptions{})

	if err != nil {
		t.Errorf("ProcessData error: %v", err)
	}

	if dst.String() != "hello world" {
		t.Errorf("Output = %q, want %q", dst.String(), "hello world")
	}

	if stats.BytesWritten != 11 {
		t.Errorf("BytesWritten = %d, want 11", stats.BytesWritten)
	}
}

func TestProcessDataUppercase(t *testing.T) {
	src := strings.NewReader("hello world")
	var dst bytes.Buffer

	stats, err := ProcessData(&dst, src, ProcessOptions{Uppercase: true})

	if err != nil {
		t.Errorf("ProcessData error: %v", err)
	}

	if dst.String() != "HELLO WORLD" {
		t.Errorf("Output = %q, want %q", dst.String(), "HELLO WORLD")
	}

	if stats.BytesWritten != 11 {
		t.Errorf("BytesWritten = %d, want 11", stats.BytesWritten)
	}
}

func TestProcessDataLimit(t *testing.T) {
	src := strings.NewReader("hello world")
	var dst bytes.Buffer

	stats, err := ProcessData(&dst, src, ProcessOptions{Limit: 5})

	if err != nil {
		t.Errorf("ProcessData error: %v", err)
	}

	if dst.String() != "hello" {
		t.Errorf("Output = %q, want %q", dst.String(), "hello")
	}

	if stats.BytesWritten != 5 {
		t.Errorf("BytesWritten = %d, want 5", stats.BytesWritten)
	}
}

func TestProcessDataUppercaseAndLimit(t *testing.T) {
	src := strings.NewReader("hello world")
	var dst bytes.Buffer

	stats, err := ProcessData(&dst, src, ProcessOptions{
		Uppercase: true,
		Limit:     5,
	})

	if err != nil {
		t.Errorf("ProcessData error: %v", err)
	}

	if dst.String() != "HELLO" {
		t.Errorf("Output = %q, want %q", dst.String(), "HELLO")
	}

	if stats.BytesWritten != 5 {
		t.Errorf("BytesWritten = %d, want 5", stats.BytesWritten)
	}
}

func TestCopyN(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		n       int64
		want    string
		wantN   int64
	}{
		{"Partial", "hello world", 5, "hello", 5},
		{"Full", "hello", 10, "hello", 5},
		{"Zero", "hello", 0, "", 0},
		{"Exact", "hello", 5, "hello", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := strings.NewReader(tt.input)
			var dst bytes.Buffer

			n, _ := CopyN(&dst, src, tt.n)

			if dst.String() != tt.want {
				t.Errorf("CopyN output = %q, want %q", dst.String(), tt.want)
			}

			if n != tt.wantN {
				t.Errorf("CopyN n = %d, want %d", n, tt.wantN)
			}
		})
	}
}

func TestCopyNExact(t *testing.T) {
	// Test that CopyN stops exactly at n bytes
	src := strings.NewReader("hello world")
	var dst bytes.Buffer

	n, _ := CopyN(&dst, src, 5)

	if n != 5 {
		t.Errorf("CopyN returned %d, want 5", n)
	}

	if dst.String() != "hello" {
		t.Errorf("CopyN output = %q, want %q", dst.String(), "hello")
	}

	// Source should still have remaining data
	remaining, _ := io.ReadAll(src)
	if string(remaining) != " world" {
		t.Errorf("Remaining = %q, want %q", string(remaining), " world")
	}
}

func TestImplementsInterfaces(t *testing.T) {
	var _ io.Reader = &CountingReader{}
	var _ io.Reader = &UppercaseReader{}
	var _ io.Reader = &LimitReader{}
	var _ io.Writer = &CountingWriter{}
	var _ io.Writer = &ProgressWriter{}

	t.Log("All types implement their respective interfaces")
}

func TestIntegrationPipeline(t *testing.T) {
	// Full integration test: multiple transformations
	src := strings.NewReader("Go is an Amazing language for Systems Programming!")
	var dst bytes.Buffer

	stats, err := ProcessData(&dst, src, ProcessOptions{
		Uppercase: true,
		Limit:     20,
	})

	if err != nil {
		t.Errorf("Integration test error: %v", err)
	}

	expected := "GO IS AN AMAZING LAN"
	if dst.String() != expected {
		t.Errorf("Integration output = %q, want %q", dst.String(), expected)
	}

	if stats.BytesWritten != 20 {
		t.Errorf("Integration BytesWritten = %d, want 20", stats.BytesWritten)
	}
}
