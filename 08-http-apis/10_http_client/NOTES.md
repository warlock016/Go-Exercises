# HTTP Client Implementation Notes

Notes from code review session - to be addressed in next coding session.

---

## Current Test Status

```
TestAPIClient_Get        FAIL  (error = EOF)
TestAPIClient_GetJSON    FAIL  (error = invalid character '\x00')
TestAPIClient_Post       PASS  (but not actually implemented - false positive)
TestAPIClient_ErrorHandling  PASS  (but needs proper status code checking)
```

---

## Issue 1: Using `resp.Body.Read()` Incorrectly

**Current code:**
```go
output := make([]byte, 100)
n, err := resp.Body.Read(output)
return output[:n], err  // Returns io.EOF as error!
```

**Problems:**
1. `Read()` returns `io.EOF` when finished — that's success, not an error
2. Fixed 100-byte buffer may be too small or too large
3. `Read()` may not read all data in one call (reads *up to* len bytes)

**Fix - use `io.ReadAll`:**
```go
import "io"

body, err := io.ReadAll(resp.Body)  // Reads everything into []byte
```

---

## Issue 2: Missing `resp.Body.Close()`

`resp.Body` is an `io.ReadCloser` — must be closed or connections leak.

**Pattern:**
```go
resp, err := c.HTTPClient.Get(url)
if err != nil {
    return nil, err
}
defer resp.Body.Close()  // Always close!

// Now read body...
```

---

## Issue 3: GetJSON Buffer Contains Null Bytes

**Current code:**
```go
output := make([]byte, 100)           // 100 bytes of zeros
resp.Body.Read(output)                 // Writes JSON into first ~20 bytes
json.Unmarshal(output, v)              // Tries to parse: {"message":"hello"}\x00\x00...
```

Null bytes after JSON cause `invalid character '\x00'`.

**Fix - use `json.NewDecoder`:**
```go
defer resp.Body.Close()
err = json.NewDecoder(resp.Body).Decode(v)  // Decodes directly from stream
```

---

## Issue 4: Missing Status Code Check

`TestAPIClient_ErrorHandling` expects error for 404, but code doesn't check status.

**Add status checking:**
```go
if resp.StatusCode >= 400 {
    return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
}
```

---

## Issue 5: Post Not Implemented

Need to:
1. Marshal `body` to JSON
2. Send POST request
3. Close response body
4. Check status code

**Pattern for POST with JSON body:**
```go
func (c *APIClient) Post(path string, body any) error {
    url, err := url.JoinPath(c.BaseURL, path)
    if err != nil {
        return err
    }

    // Marshal body to JSON
    jsonData, err := json.Marshal(body)
    if err != nil {
        return err
    }

    // Create request with JSON body
    resp, err := c.HTTPClient.Post(url, "application/json", bytes.NewReader(jsonData))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Check status
    if resp.StatusCode >= 400 {
        return fmt.Errorf("request failed with status %d", resp.StatusCode)
    }

    return nil
}
```

**Required import for bytes.NewReader:**
```go
import "bytes"
```

---

## Key Insight: resp.Body is a Stream

`resp.Body` is an `io.ReadCloser` — a one-way stream, not a buffer.

**Idiomatic patterns:**

```go
// For raw bytes:
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)

// For JSON:
defer resp.Body.Close()
err := json.NewDecoder(resp.Body).Decode(&v)
```

**Never use `Read()` directly** unless doing streaming or chunked processing.

---

## Complete Get() Example

```go
func (c *APIClient) Get(path string) ([]byte, error) {
    url, err := url.JoinPath(c.BaseURL, path)
    if err != nil {
        return nil, err
    }

    resp, err := c.HTTPClient.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
    }

    return io.ReadAll(resp.Body)
}
```

---

## Complete GetJSON() Example

```go
func (c *APIClient) GetJSON(path string, v any) error {
    url, err := url.JoinPath(c.BaseURL, path)
    if err != nil {
        return err
    }

    resp, err := c.HTTPClient.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return fmt.Errorf("request failed with status %d", resp.StatusCode)
    }

    return json.NewDecoder(resp.Body).Decode(v)
}
```

---

## Imports Needed

```go
import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "time"
)
```

---

## Summary Checklist

- [ ] Fix `Get()`: use `io.ReadAll`, add `defer resp.Body.Close()`, check status
- [ ] Fix `GetJSON()`: use `json.NewDecoder`, add `defer resp.Body.Close()`, check status
- [ ] Implement `Post()`: marshal JSON, use `http.Post`, close body, check status
- [ ] Ensure all methods return proper errors for 4xx/5xx status codes
