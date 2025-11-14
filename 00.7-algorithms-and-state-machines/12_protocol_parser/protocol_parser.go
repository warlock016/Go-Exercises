package protocolparser

import (
	"fmt"
	"strings"
)

// HTTPRequest represents a parsed HTTP request
type HTTPRequest struct {
	Method  string            // GET, POST, PUT, DELETE, etc.
	Path    string            // /index.html, /api/users, etc.
	Version string            // HTTP/1.1, HTTP/2.0, etc.
	Headers map[string]string // Header key-value pairs
	Body    string            // Request body content
}

// ParseHTTPRequest parses a raw HTTP request string into an HTTPRequest struct
//
// Expected format:
//   METHOD PATH VERSION
//   Header1: Value1
//   Header2: Value2
//
//   Body content here
//
// Example:
//   GET /index.html HTTP/1.1
//   Host: example.com
//   User-Agent: Go Client
//
//   Optional body content
func ParseHTTPRequest(raw string) (*HTTPRequest, error) {
	// TODO(human): Implement HTTP request parser using state machine
	//
	// Algorithm overview:
	// 1. Normalize line endings: replace \r\n with \n
	// 2. Split into lines by \n
	// 3. Parse request line (first line): METHOD PATH VERSION
	//    - Use strings.Fields() to split by whitespace
	//    - Verify exactly 3 parts, else return error
	// 4. Parse headers (subsequent lines until blank line):
	//    - Split each line by first ':' using strings.SplitN(line, ":", 2)
	//    - Trim spaces from key and value
	//    - Store in Headers map
	//    - If line is empty, transition to body parsing
	// 5. Parse body (everything after blank line):
	//    - Join remaining lines with \n
	// 6. Return populated HTTPRequest struct
	//
	// State machine:
	// STATE 1: READING_REQUEST_LINE (first line only)
	// STATE 2: READING_HEADERS (until blank line)
	// STATE 3: READING_BODY (all remaining content)
	//
	// Error cases to handle:
	// - Empty input → error
	// - Request line doesn't have 3 parts → error
	// - Header line missing ':' → error
	//
	// Hints:
	// - Initialize Headers map: make(map[string]string)
	// - Use strings.TrimSpace() on header values
	// - Use strings.Join() for multi-line body
	// - Track line index with variable i

	return nil, fmt.Errorf("not implemented")
}

// String returns a human-readable representation of the HTTP request
// Useful for debugging and testing
//
// Example output:
//   GET /index.html HTTP/1.1
//   Headers:
//     Host: example.com
//     User-Agent: Go Client
//   Body: (13 bytes)
func (r *HTTPRequest) String() string {
	// TODO(human): Implement string representation for debugging
	//
	// Format:
	// METHOD PATH VERSION
	// Headers:
	//   Key1: Value1
	//   Key2: Value2
	// Body: (N bytes) or empty
	//
	// Hints:
	// - Use strings.Builder for efficient string construction
	// - Use fmt.Sprintf for formatted output
	// - Iterate over Headers map (note: map iteration order is random)
	// - If body is empty, show "Body: empty"
	// - Otherwise show "Body: (N bytes)" where N = len(r.Body)

	var b strings.Builder

	// YOUR CODE HERE

	return b.String()
}
