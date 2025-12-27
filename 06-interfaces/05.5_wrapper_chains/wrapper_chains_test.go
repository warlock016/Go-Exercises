package wrapper_chains

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestUppercaseReader(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		bufSize int
		want    string
	}{
		{"Lowercase", "hello", 10, "HELLO"},
		{"Mixed case", "HeLLo WoRLd", 20, "HELLO WORLD"},
		{"Already upper", "HELLO", 10, "HELLO"},
		{"With numbers", "test123", 10, "TEST123"},
		{"Empty", "", 10, ""},
		{"Special chars", "hello!", 10, "HELLO!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := strings.NewReader(tt.input)
			upper := &UppercaseReader{Reader: source}
			buf := make([]byte, tt.bufSize)

			n, _ := upper.Read(buf)
			got := string(buf[:n])

			if got != tt.want {
				t.Errorf("UppercaseReader.Read() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestUppercaseReaderMultipleReads(t *testing.T) {
	source := strings.NewReader("hello world")
	upper := &UppercaseReader{Reader: source}
	buf := make([]byte, 5)

	// First read
	n, err := upper.Read(buf)
	if n != 5 || string(buf[:n]) != "HELLO" || err != nil {
		t.Errorf("First read: n=%d, data=%q, err=%v", n, string(buf[:n]), err)
	}

	// Second read
	n, err = upper.Read(buf)
	if n != 5 || string(buf[:n]) != " WORL" {
		t.Errorf("Second read: n=%d, data=%q, err=%v", n, string(buf[:n]), err)
	}

	// Third read
	n, err = upper.Read(buf)
	if n != 1 || string(buf[:n]) != "D" || err != io.EOF {
		t.Errorf("Third read: n=%d, data=%q, err=%v", n, string(buf[:n]), err)
	}
}

func TestLimitReader(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		limit   int
		bufSize int
		want    string
	}{
		{"Within limit", "hello", 10, 20, "hello"},
		{"Exact limit", "hello", 5, 20, "hello"},
		{"Exceeds limit", "hello world", 5, 20, "hello"},
		{"Zero limit", "hello", 0, 20, ""},
		{"Large limit", "hi", 100, 20, "hi"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := strings.NewReader(tt.input)
			limited := &LimitReader{Reader: source, Limit: tt.limit}

			result, _ := io.ReadAll(limited)
			got := string(result)

			if got != tt.want {
				t.Errorf("LimitReader = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestLimitReaderMultipleReads(t *testing.T) {
	source := strings.NewReader("hello world")
	limited := &LimitReader{Reader: source, Limit: 7}
	buf := make([]byte, 3)

	// Read 1: "hel"
	n, err := limited.Read(buf)
	if n != 3 || string(buf[:n]) != "hel" || err != nil {
		t.Errorf("Read 1: n=%d, data=%q, err=%v", n, string(buf[:n]), err)
	}

	// Read 2: "lo "
	n, err = limited.Read(buf)
	if n != 3 || string(buf[:n]) != "lo " || err != nil {
		t.Errorf("Read 2: n=%d, data=%q, err=%v", n, string(buf[:n]), err)
	}

	// Read 3: "w" (only 1 byte left in limit)
	n, err = limited.Read(buf)
	if n != 1 || string(buf[:n]) != "w" {
		t.Errorf("Read 3: n=%d, data=%q, err=%v", n, string(buf[:n]), err)
	}

	// Read 4: should return 0, EOF (limit reached)
	n, err = limited.Read(buf)
	if n != 0 || err != io.EOF {
		t.Errorf("Read 4: n=%d, err=%v, want n=0, err=io.EOF", n, err)
	}
}

func TestCountingReader(t *testing.T) {
	source := strings.NewReader("hello world")
	counter := &CountingReader{Reader: source}

	buf := make([]byte, 5)
	counter.Read(buf) // Read 5
	counter.Read(buf) // Read 5
	counter.Read(buf) // Read 1

	if counter.BytesRead != 11 {
		t.Errorf("BytesRead = %d, want 11", counter.BytesRead)
	}
}

func TestTeeReader(t *testing.T) {
	source := strings.NewReader("hello world")
	var copied bytes.Buffer
	tee := &TeeReader{Reader: source, Writer: &copied}

	result, _ := io.ReadAll(tee)

	if string(result) != "hello world" {
		t.Errorf("TeeReader result = %q, want %q", string(result), "hello world")
	}

	if copied.String() != "hello world" {
		t.Errorf("TeeReader copied = %q, want %q", copied.String(), "hello world")
	}
}

func TestBuildPipeline(t *testing.T) {
	source := strings.NewReader("hello world")
	pipeline, stats := BuildPipeline(source, 5)

	result, _ := io.ReadAll(pipeline)

	// Should be uppercase and limited to 5 bytes
	if string(result) != "HELLO" {
		t.Errorf("Pipeline result = %q, want %q", string(result), "HELLO")
	}

	// Counter should have counted 5 bytes
	if stats.Counted == nil || *stats.Counted != 5 {
		t.Errorf("Pipeline stats.Counted = %v, want 5", stats.Counted)
	}
}

func TestBuildPipelineFullData(t *testing.T) {
	source := strings.NewReader("go")
	pipeline, stats := BuildPipeline(source, 100)

	result, _ := io.ReadAll(pipeline)

	if string(result) != "GO" {
		t.Errorf("Pipeline result = %q, want %q", string(result), "GO")
	}

	if stats.Counted == nil || *stats.Counted != 2 {
		t.Errorf("Pipeline stats.Counted = %v, want 2", stats.Counted)
	}
}

func TestChainOrder(t *testing.T) {
	// Test that order matters: uppercase THEN limit
	source := strings.NewReader("hello")

	// Chain: source → uppercase → limit
	upper := &UppercaseReader{Reader: source}
	limited := &LimitReader{Reader: upper, Limit: 3}

	result, _ := io.ReadAll(limited)

	// Should get first 3 uppercase letters
	if string(result) != "HEL" {
		t.Errorf("uppercase→limit = %q, want %q", string(result), "HEL")
	}
}

func TestImplementsIOReader(t *testing.T) {
	var _ io.Reader = &UppercaseReader{}
	var _ io.Reader = &LimitReader{}
	var _ io.Reader = &TeeReader{}
	var _ io.Reader = &CountingReader{}

	t.Log("All types implement io.Reader interface")
}

func TestDeepChain(t *testing.T) {
	// source → upper → count → limit → upper again
	source := strings.NewReader("test data here")

	layer1 := &UppercaseReader{Reader: source}
	layer2 := &CountingReader{Reader: layer1}
	layer3 := &LimitReader{Reader: layer2, Limit: 8}

	result, _ := io.ReadAll(layer3)

	if string(result) != "TEST DAT" {
		t.Errorf("Deep chain = %q, want %q", string(result), "TEST DAT")
	}

	if layer2.BytesRead != 8 {
		t.Errorf("Counter in chain = %d, want 8", layer2.BytesRead)
	}
}
