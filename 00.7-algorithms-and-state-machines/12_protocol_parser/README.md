# Exercise 12: Protocol Parser (HTTP Request)

**Difficulty:** Hard
**Time:** 60-70 minutes
**Concepts:** Multi-state parsing, protocol design, line-based parsing, state transitions

## Learning Goal

Implement a parser for HTTP request messages using a multi-state state machine. You'll learn how network protocols encode information in structured text format and how to parse that structure reliably. This exercise simulates real-world parsing challenges like handling headers, bodies, and malformed data.

## The Problem

HTTP (Hypertext Transfer Protocol) is the foundation of web communication. Every time your browser loads a webpage, it sends an HTTP request like this:

```
GET /index.html HTTP/1.1
Host: example.com
User-Agent: Mozilla/5.0
Content-Length: 13

Hello, World!
```

Your task is to parse this text format into a structured Go data type. The parser must handle:
1. **Request line**: Method, path, and HTTP version
2. **Headers**: Key-value pairs (can be multiple)
3. **Body**: Optional content after headers

The parser is a state machine with these states:
```
START → READING_METHOD → READING_PATH → READING_VERSION →
        READING_HEADERS → READING_BODY → DONE
```

## Your Task

Implement two functions:

```go
type HTTPRequest struct {
    Method  string            // GET, POST, PUT, DELETE, etc.
    Path    string            // /index.html, /api/users, etc.
    Version string            // HTTP/1.1, HTTP/2.0, etc.
    Headers map[string]string // Key-value pairs
    Body    string            // Request body (may be empty)
}

func ParseHTTPRequest(raw string) (*HTTPRequest, error)

func (r *HTTPRequest) String() string  // For debugging/display
```

## HTTP Request Format

An HTTP request has three parts:

### 1. Request Line (first line)
```
<METHOD> <PATH> <VERSION>
```
Example: `GET /hello HTTP/1.1`

### 2. Headers (key: value pairs, one per line)
```
<Key>: <Value>
<Key>: <Value>
...
```
Example:
```
Host: example.com
Content-Type: application/json
```

### 3. Blank Line (separates headers from body)
```

```
A single empty line indicates end of headers.

### 4. Body (everything after blank line)
```
<body content>
```
Example: `{"name": "John", "age": 30}`

## Examples

### Example 1: Simple GET Request

```go
raw := "GET /index.html HTTP/1.1\nHost: example.com\n\n"

req, err := ParseHTTPRequest(raw)
// req.Method = "GET"
// req.Path = "/index.html"
// req.Version = "HTTP/1.1"
// req.Headers = map[string]string{"Host": "example.com"}
// req.Body = ""
// err = nil
```

### Example 2: POST with Body

```go
raw := "POST /api/users HTTP/1.1\nHost: api.example.com\nContent-Type: application/json\n\n{\"name\": \"Alice\"}"

req, err := ParseHTTPRequest(raw)
// req.Method = "POST"
// req.Path = "/api/users"
// req.Version = "HTTP/1.1"
// req.Headers = map[string]string{
//     "Host": "api.example.com",
//     "Content-Type": "application/json",
// }
// req.Body = "{\"name\": \"Alice\"}"
// err = nil
```

### Example 3: Multiple Headers

```go
raw := "GET /data HTTP/1.1\nHost: example.com\nUser-Agent: Go Client\nAccept: application/json\n\n"

req, err := ParseHTTPRequest(raw)
// req.Headers = map[string]string{
//     "Host": "example.com",
//     "User-Agent": "Go Client",
//     "Accept": "application/json",
// }
```

### Example 4: Malformed Request (Error Cases)

```go
// Missing path
raw := "GET HTTP/1.1\n\n"
req, err := ParseHTTPRequest(raw)
// req = nil
// err != nil (e.g., "invalid request line")

// Missing version
raw := "GET /index.html\n\n"
req, err := ParseHTTPRequest(raw)
// err != nil

// Invalid header format (no colon)
raw := "GET / HTTP/1.1\nInvalidHeader\n\n"
req, err := ParseHTTPRequest(raw)
// err != nil (e.g., "invalid header format")
```

## Edge Cases

```go
// Empty request
ParseHTTPRequest("")  // → error

// Only request line
ParseHTTPRequest("GET / HTTP/1.1\n\n")  // → valid (no headers, no body)

// Body with multiple lines
raw := "POST / HTTP/1.1\n\nLine 1\nLine 2\nLine 3"
// req.Body = "Line 1\nLine 2\nLine 3"

// Windows line endings (\r\n)
raw := "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"
// Should handle both \n and \r\n

// Header with spaces in value
raw := "GET / HTTP/1.1\nUser-Agent: Mozilla/5.0 (Windows NT)\n\n"
// req.Headers["User-Agent"] = "Mozilla/5.0 (Windows NT)"

// Empty body after blank line
ParseHTTPRequest("GET / HTTP/1.1\n\n")  // req.Body = ""
```

## Instructions

1. Open `protocol_parser.go`
2. Implement `ParseHTTPRequest` function using a state machine
3. Implement the `String()` method for HTTPRequest (for debugging)
4. Run `go test -v` to verify your solution
5. Handle all error cases and edge cases

## Hints

### Hint 1: Split into Lines First

**Basic Hint:** HTTP requests are line-based. Split by newlines first, then process line-by-line.

**Intermediate Hint:** Use `strings.Split(raw, "\n")` to get lines. Handle both `\n` and `\r\n` by using `strings.ReplaceAll(raw, "\r\n", "\n")` first.

```go
func ParseHTTPRequest(raw string) (*HTTPRequest, error) {
    // Normalize line endings
    raw = strings.ReplaceAll(raw, "\r\n", "\n")
    lines := strings.Split(raw, "\n")

    // Process lines with state machine...
}
```

### Hint 2: State Machine Design

**Basic Hint:** Use a state variable that tracks which part of the request you're parsing. Start with READING_REQUEST_LINE, then READING_HEADERS, then READING_BODY.

**Intermediate Hint:** Define state constants:
```go
const (
    STATE_REQUEST_LINE = iota
    STATE_HEADERS
    STATE_BODY
)

state := STATE_REQUEST_LINE
```

**Advanced Hint:** State transitions:
- **STATE_REQUEST_LINE**: Parse first line into method/path/version, transition to STATE_HEADERS
- **STATE_HEADERS**: Parse each line as "Key: Value"
  - If line is empty → transition to STATE_BODY
  - If line has no colon → return error
- **STATE_BODY**: Accumulate all remaining lines into body string

### Hint 3: Parsing the Request Line

**Basic Hint:** The first line has three parts separated by spaces: METHOD, PATH, VERSION.

**Intermediate Hint:** Use `strings.Fields(line)` to split by whitespace, then verify you have exactly 3 parts.

```go
parts := strings.Fields(lines[0])
if len(parts) != 3 {
    return nil, fmt.Errorf("invalid request line")
}
method, path, version := parts[0], parts[1], parts[2]
```

### Hint 4: Parsing Headers

**Basic Hint:** Each header line has format "Key: Value". Split by the first colon.

**Intermediate Hint:** Use `strings.SplitN(line, ":", 2)` to split into key and value, limiting to 2 parts (in case value contains colons).

```go
parts := strings.SplitN(line, ":", 2)
if len(parts) != 2 {
    return nil, fmt.Errorf("invalid header format: %s", line)
}
key := strings.TrimSpace(parts[0])
value := strings.TrimSpace(parts[1])
headers[key] = value
```

**Advanced Hint:** After splitting by `:`, trim spaces from both key and value using `strings.TrimSpace()`. This handles headers like `Host:  example.com  ` correctly.

### Hint 5: Parsing the Body

**Basic Hint:** The body is everything after the first blank line. It can span multiple lines.

**Intermediate Hint:** Once you encounter a blank line (empty string after splitting), switch to STATE_BODY. Join all remaining lines with `\n`:

```go
// In STATE_HEADERS:
if line == "" {
    state = STATE_BODY
    continue
}

// In STATE_BODY:
bodyLines := lines[i:]  // All remaining lines
body := strings.Join(bodyLines, "\n")
```

**Advanced Hint:** Use an index variable to track current line. When switching to BODY state, slice lines from current index to end, then join with newlines.

### Hint 6: Error Handling

**Basic Hint:** Return errors for these cases:
- Empty input
- Request line doesn't have 3 parts
- Header line doesn't contain ":"
- Any malformed structure

**Intermediate Hint:** Use descriptive error messages:
```go
if raw == "" {
    return nil, fmt.Errorf("empty request")
}
if len(parts) != 3 {
    return nil, fmt.Errorf("invalid request line: expected METHOD PATH VERSION")
}
if len(headerParts) != 2 {
    return nil, fmt.Errorf("invalid header format: %s", line)
}
```

### Hint 7: Complete Algorithm

```go
func ParseHTTPRequest(raw string) (*HTTPRequest, error) {
    // 1. Normalize and split
    raw = strings.ReplaceAll(raw, "\r\n", "\n")
    lines := strings.Split(raw, "\n")

    if len(lines) == 0 || lines[0] == "" {
        return nil, fmt.Errorf("empty request")
    }

    // 2. Parse request line
    parts := strings.Fields(lines[0])
    if len(parts) != 3 {
        return nil, fmt.Errorf("invalid request line")
    }

    req := &HTTPRequest{
        Method:  parts[0],
        Path:    parts[1],
        Version: parts[2],
        Headers: make(map[string]string),
    }

    // 3. Parse headers and body
    i := 1
    // Parse headers
    for i < len(lines) && lines[i] != "" {
        // Split by first colon
        headerParts := strings.SplitN(lines[i], ":", 2)
        if len(headerParts) != 2 {
            return nil, fmt.Errorf("invalid header format")
        }
        key := strings.TrimSpace(headerParts[0])
        value := strings.TrimSpace(headerParts[1])
        req.Headers[key] = value
        i++
    }

    // Skip blank line
    if i < len(lines) && lines[i] == "" {
        i++
    }

    // Parse body (everything remaining)
    if i < len(lines) {
        req.Body = strings.Join(lines[i:], "\n")
    }

    return req, nil
}
```

## Think About

1. **Why is HTTP text-based instead of binary?**
   - Human-readable for debugging
   - Easy to implement in any language
   - Firewall and proxy inspection
   - However, HTTP/2 is binary for efficiency!

2. **What happens if a header value contains a colon?**
   - Example: `Time: 12:30:45`
   - Solution: Split on first colon only (use `SplitN(line, ":", 2)`)

3. **How would you handle duplicate headers?**
   - HTTP allows it: `Set-Cookie: A\nSet-Cookie: B`
   - Options: Store as array, append values, or keep last

4. **What about case sensitivity?**
   - HTTP headers are case-insensitive ("Host" == "host")
   - Solution: Normalize keys to lowercase or use case-insensitive map

5. **How do real HTTP parsers handle this?**
   - More sophisticated: streaming parsers, chunk encoding, compression
   - Go's `net/http` package has a production-grade parser

6. **What security concerns exist?**
   - Header injection: Attacker adds `\n` in value to inject headers
   - Buffer overflow: Extremely long headers/body
   - Validation: Ensure method is valid, path doesn't contain dangerous chars

## What This Teaches

- **Protocol parsing** - Understanding structured text formats
- **Multi-state machines** - Coordinating multiple parsing phases
- **Line-based parsing** - Common pattern in network protocols
- **Error handling** - Validating input and returning descriptive errors
- **String manipulation** - Splitting, trimming, joining
- **Data modeling** - Designing structs to represent protocol messages
- **Real-world patterns** - How HTTP, SMTP, FTP, etc. work

## Common Mistakes to Avoid

1. **Not handling empty lines correctly**
   ```go
   // BUG - blank line ends headers, not just any empty string
   for _, line := range lines {
       if line == "" {
           break  // This breaks on first empty line!
       }
   }

   // FIX - use state tracking
   ```

2. **Forgetting to initialize Headers map**
   ```go
   // BUG - nil map panic
   req := &HTTPRequest{}
   req.Headers["Host"] = "example.com"  // Panic!

   // FIX
   req := &HTTPRequest{
       Headers: make(map[string]string),
   }
   ```

3. **Not trimming spaces in headers**
   ```go
   // BUG - keeps spaces in value
   parts := strings.Split("Host: example.com", ":")
   // parts[1] = " example.com" (leading space!)

   // FIX
   value := strings.TrimSpace(parts[1])
   ```

4. **Splitting by ":" instead of first ":"**
   ```go
   // BUG - breaks if value contains ":"
   parts := strings.Split("Time: 12:30:45", ":")
   // parts = ["Time", " 12", "30", "45"] (wrong!)

   // FIX
   parts := strings.SplitN(line, ":", 2)
   // parts = ["Time", " 12:30:45"] (correct!)
   ```

5. **Not handling Windows line endings**
   ```go
   // BUG - treats \r as part of value
   raw := "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"
   lines := strings.Split(raw, "\n")
   // lines[0] = "GET / HTTP/1.1\r" (has \r!)

   // FIX
   raw = strings.ReplaceAll(raw, "\r\n", "\n")
   ```

6. **Losing newlines in body**
   ```go
   // BUG - body becomes single line
   body := ""
   for _, line := range bodyLines {
       body += line  // No newlines between!
   }

   // FIX
   body := strings.Join(bodyLines, "\n")
   ```

## Challenge Extensions (Optional)

After completing the basic version, try these:

1. **Validate HTTP methods**
   ```go
   validMethods := []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH"}
   // Return error if method not in list
   ```

2. **Parse query parameters from path**
   ```go
   // "/search?q=golang&page=2" → Path="/search", Query={"q":"golang", "page":"2"}
   ```

3. **Handle chunked transfer encoding**
   ```go
   // If header "Transfer-Encoding: chunked", parse body in chunks
   ```

4. **Case-insensitive headers**
   ```go
   // Normalize all header keys to lowercase or Title-Case
   ```

5. **Support HTTP response parsing**
   ```go
   // "HTTP/1.1 200 OK\nContent-Type: text/html\n\n<html>..."
   type HTTPResponse struct {
       Version string
       StatusCode int
       StatusText string
       Headers map[string]string
       Body string
   }
   ```

6. **Streaming parser**
   ```go
   // Parse incrementally as data arrives (don't require full request upfront)
   ```

## Real-World Applications

- **Web servers**: Parse incoming HTTP requests (nginx, Apache, Go's `net/http`)
- **API clients**: Build HTTP requests (curl, Postman, SDKs)
- **Proxies**: Inspect and modify HTTP traffic
- **Testing tools**: Generate test requests, mock servers
- **Security tools**: Analyze HTTP traffic for vulnerabilities
- **Load testers**: Generate thousands of requests per second
- **Protocol debugging**: Wireshark, tcpdump

## After Completing

You now understand:
- How HTTP requests are structured
- How to implement multi-state parsing
- Line-based protocol parsing patterns
- Error handling in parsers
- The importance of validation and edge cases
- Why protocols use text vs binary formats

This parsing pattern applies to many protocols:
- **SMTP** (email): Similar structure with commands and headers
- **FTP** (file transfer): Command-response protocol
- **IRC** (chat): Message parsing
- **Custom protocols**: You can design your own!

---

**Ready to implement?** Open `protocol_parser.go` and look for the `TODO(human)` markers!
