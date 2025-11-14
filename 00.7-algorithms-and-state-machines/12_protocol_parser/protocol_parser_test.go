package protocolparser

import (
	"strings"
	"testing"
)

func TestSimpleGET(t *testing.T) {
	raw := "GET /index.html HTTP/1.1\nHost: example.com\n\n"

	req, err := ParseHTTPRequest(raw)
	if err != nil {
		t.Fatalf("ParseHTTPRequest() error = %v, want nil", err)
	}

	if req.Method != "GET" {
		t.Errorf("Method = %q, want %q", req.Method, "GET")
	}
	if req.Path != "/index.html" {
		t.Errorf("Path = %q, want %q", req.Path, "/index.html")
	}
	if req.Version != "HTTP/1.1" {
		t.Errorf("Version = %q, want %q", req.Version, "HTTP/1.1")
	}
	if req.Headers["Host"] != "example.com" {
		t.Errorf("Headers[Host] = %q, want %q", req.Headers["Host"], "example.com")
	}
	if req.Body != "" {
		t.Errorf("Body = %q, want empty", req.Body)
	}
}

func TestPOSTWithBody(t *testing.T) {
	raw := "POST /api/users HTTP/1.1\nHost: api.example.com\nContent-Type: application/json\n\n{\"name\": \"Alice\"}"

	req, err := ParseHTTPRequest(raw)
	if err != nil {
		t.Fatalf("ParseHTTPRequest() error = %v, want nil", err)
	}

	if req.Method != "POST" {
		t.Errorf("Method = %q, want %q", req.Method, "POST")
	}
	if req.Path != "/api/users" {
		t.Errorf("Path = %q, want %q", req.Path, "/api/users")
	}
	if req.Headers["Host"] != "api.example.com" {
		t.Errorf("Headers[Host] = %q, want %q", req.Headers["Host"], "api.example.com")
	}
	if req.Headers["Content-Type"] != "application/json" {
		t.Errorf("Headers[Content-Type] = %q, want %q", req.Headers["Content-Type"], "application/json")
	}
	wantBody := `{"name": "Alice"}`
	if req.Body != wantBody {
		t.Errorf("Body = %q, want %q", req.Body, wantBody)
	}
}

func TestMultipleHeaders(t *testing.T) {
	raw := "GET /data HTTP/1.1\nHost: example.com\nUser-Agent: Go Client\nAccept: application/json\n\n"

	req, err := ParseHTTPRequest(raw)
	if err != nil {
		t.Fatalf("ParseHTTPRequest() error = %v, want nil", err)
	}

	expectedHeaders := map[string]string{
		"Host":       "example.com",
		"User-Agent": "Go Client",
		"Accept":     "application/json",
	}

	if len(req.Headers) != len(expectedHeaders) {
		t.Errorf("Headers count = %d, want %d", len(req.Headers), len(expectedHeaders))
	}

	for key, want := range expectedHeaders {
		if got := req.Headers[key]; got != want {
			t.Errorf("Headers[%q] = %q, want %q", key, got, want)
		}
	}
}

func TestMultiLineBody(t *testing.T) {
	raw := "POST /submit HTTP/1.1\nHost: example.com\n\nLine 1\nLine 2\nLine 3"

	req, err := ParseHTTPRequest(raw)
	if err != nil {
		t.Fatalf("ParseHTTPRequest() error = %v, want nil", err)
	}

	wantBody := "Line 1\nLine 2\nLine 3"
	if req.Body != wantBody {
		t.Errorf("Body = %q, want %q", req.Body, wantBody)
	}
}

func TestWindowsLineEndings(t *testing.T) {
	raw := "GET /test HTTP/1.1\r\nHost: example.com\r\n\r\n"

	req, err := ParseHTTPRequest(raw)
	if err != nil {
		t.Fatalf("ParseHTTPRequest() error = %v, want nil", err)
	}

	if req.Method != "GET" {
		t.Errorf("Method = %q, want %q", req.Method, "GET")
	}
	if req.Headers["Host"] != "example.com" {
		t.Errorf("Headers[Host] = %q, want %q", req.Headers["Host"], "example.com")
	}
}

func TestHeaderWithSpaces(t *testing.T) {
	raw := "GET / HTTP/1.1\nUser-Agent:   Mozilla/5.0 (Windows NT)  \n\n"

	req, err := ParseHTTPRequest(raw)
	if err != nil {
		t.Fatalf("ParseHTTPRequest() error = %v, want nil", err)
	}

	// Should trim leading/trailing spaces from value
	want := "Mozilla/5.0 (Windows NT)"
	if req.Headers["User-Agent"] != want {
		t.Errorf("Headers[User-Agent] = %q, want %q", req.Headers["User-Agent"], want)
	}
}

func TestHeaderWithColonInValue(t *testing.T) {
	raw := "GET / HTTP/1.1\nTime: 12:30:45\n\n"

	req, err := ParseHTTPRequest(raw)
	if err != nil {
		t.Fatalf("ParseHTTPRequest() error = %v, want nil", err)
	}

	want := "12:30:45"
	if req.Headers["Time"] != want {
		t.Errorf("Headers[Time] = %q, want %q", req.Headers["Time"], want)
	}
}

func TestNoHeaders(t *testing.T) {
	raw := "GET / HTTP/1.1\n\n"

	req, err := ParseHTTPRequest(raw)
	if err != nil {
		t.Fatalf("ParseHTTPRequest() error = %v, want nil", err)
	}

	if req.Method != "GET" {
		t.Errorf("Method = %q, want %q", req.Method, "GET")
	}
	if len(req.Headers) != 0 {
		t.Errorf("Headers count = %d, want 0", len(req.Headers))
	}
	if req.Body != "" {
		t.Errorf("Body = %q, want empty", req.Body)
	}
}

func TestEmptyBodyAfterBlankLine(t *testing.T) {
	raw := "GET / HTTP/1.1\nHost: example.com\n\n"

	req, err := ParseHTTPRequest(raw)
	if err != nil {
		t.Fatalf("ParseHTTPRequest() error = %v, want nil", err)
	}

	if req.Body != "" {
		t.Errorf("Body = %q, want empty", req.Body)
	}
}

func TestDifferentMethods(t *testing.T) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}

	for _, method := range methods {
		raw := method + " / HTTP/1.1\n\n"
		req, err := ParseHTTPRequest(raw)
		if err != nil {
			t.Errorf("ParseHTTPRequest(%q) error = %v, want nil", method, err)
			continue
		}
		if req.Method != method {
			t.Errorf("Method = %q, want %q", req.Method, method)
		}
	}
}

func TestErrorEmptyRequest(t *testing.T) {
	raw := ""

	req, err := ParseHTTPRequest(raw)
	if err == nil {
		t.Errorf("ParseHTTPRequest(\"\") error = nil, want error")
	}
	if req != nil {
		t.Errorf("ParseHTTPRequest(\"\") req = %v, want nil", req)
	}
}

func TestErrorInvalidRequestLine(t *testing.T) {
	tests := []struct {
		name string
		raw  string
	}{
		{"missing path", "GET HTTP/1.1\n\n"},
		{"missing version", "GET /index.html\n\n"},
		{"only method", "GET\n\n"},
		{"too many parts", "GET / HTTP/1.1 Extra\n\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := ParseHTTPRequest(tt.raw)
			if err == nil {
				t.Errorf("ParseHTTPRequest(%q) error = nil, want error", tt.raw)
			}
			if req != nil {
				t.Errorf("ParseHTTPRequest(%q) req = %v, want nil", tt.raw, req)
			}
		})
	}
}

func TestErrorInvalidHeaderFormat(t *testing.T) {
	tests := []struct {
		name string
		raw  string
	}{
		{"no colon", "GET / HTTP/1.1\nInvalidHeader\n\n"},
		{"empty header line with text", "GET / HTTP/1.1\nJustText\n\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := ParseHTTPRequest(tt.raw)
			if err == nil {
				t.Errorf("ParseHTTPRequest(%q) error = nil, want error", tt.raw)
			}
			if req != nil {
				t.Errorf("ParseHTTPRequest(%q) req = %v, want nil", tt.raw, req)
			}
		})
	}
}

func TestComplexRequest(t *testing.T) {
	raw := `POST /api/v1/resource HTTP/1.1
Host: api.example.com
User-Agent: Go-http-client/1.1
Content-Type: application/json
Authorization: Bearer token123
Accept: application/json
Content-Length: 45

{"id": 1, "name": "Test", "active": true}`

	req, err := ParseHTTPRequest(raw)
	if err != nil {
		t.Fatalf("ParseHTTPRequest() error = %v, want nil", err)
	}

	// Check method, path, version
	if req.Method != "POST" {
		t.Errorf("Method = %q, want POST", req.Method)
	}
	if req.Path != "/api/v1/resource" {
		t.Errorf("Path = %q, want /api/v1/resource", req.Path)
	}
	if req.Version != "HTTP/1.1" {
		t.Errorf("Version = %q, want HTTP/1.1", req.Version)
	}

	// Check headers
	expectedHeaders := map[string]string{
		"Host":           "api.example.com",
		"User-Agent":     "Go-http-client/1.1",
		"Content-Type":   "application/json",
		"Authorization":  "Bearer token123",
		"Accept":         "application/json",
		"Content-Length": "45",
	}

	for key, want := range expectedHeaders {
		if got := req.Headers[key]; got != want {
			t.Errorf("Headers[%q] = %q, want %q", key, got, want)
		}
	}

	// Check body
	wantBody := `{"id": 1, "name": "Test", "active": true}`
	if req.Body != wantBody {
		t.Errorf("Body = %q, want %q", req.Body, wantBody)
	}
}

func TestBodyWithBlankLines(t *testing.T) {
	raw := "POST / HTTP/1.1\n\nLine 1\n\nLine 3\n"

	req, err := ParseHTTPRequest(raw)
	if err != nil {
		t.Fatalf("ParseHTTPRequest() error = %v, want nil", err)
	}

	// Body should include blank lines
	wantBody := "Line 1\n\nLine 3\n"
	if req.Body != wantBody {
		t.Errorf("Body = %q, want %q", req.Body, wantBody)
	}
}

func TestHTTPVersions(t *testing.T) {
	versions := []string{"HTTP/1.0", "HTTP/1.1", "HTTP/2.0"}

	for _, version := range versions {
		raw := "GET / " + version + "\n\n"
		req, err := ParseHTTPRequest(raw)
		if err != nil {
			t.Errorf("ParseHTTPRequest(%q) error = %v, want nil", version, err)
			continue
		}
		if req.Version != version {
			t.Errorf("Version = %q, want %q", req.Version, version)
		}
	}
}

func TestStringMethod(t *testing.T) {
	req := &HTTPRequest{
		Method:  "GET",
		Path:    "/test",
		Version: "HTTP/1.1",
		Headers: map[string]string{
			"Host": "example.com",
		},
		Body: "Test body",
	}

	str := req.String()

	// Should contain method, path, version
	if !strings.Contains(str, "GET") {
		t.Errorf("String() missing method, got: %s", str)
	}
	if !strings.Contains(str, "/test") {
		t.Errorf("String() missing path, got: %s", str)
	}
	if !strings.Contains(str, "HTTP/1.1") {
		t.Errorf("String() missing version, got: %s", str)
	}

	// Should contain headers
	if !strings.Contains(str, "Host") {
		t.Errorf("String() missing header key, got: %s", str)
	}
	if !strings.Contains(str, "example.com") {
		t.Errorf("String() missing header value, got: %s", str)
	}

	// Should indicate body exists
	if !strings.Contains(str, "Body") {
		t.Errorf("String() missing body indicator, got: %s", str)
	}
}

func TestStringMethodEmptyBody(t *testing.T) {
	req := &HTTPRequest{
		Method:  "GET",
		Path:    "/",
		Version: "HTTP/1.1",
		Headers: map[string]string{},
		Body:    "",
	}

	str := req.String()

	// Should indicate empty body
	if !strings.Contains(str, "empty") && !strings.Contains(str, "0 bytes") {
		t.Errorf("String() should indicate empty body, got: %s", str)
	}
}
