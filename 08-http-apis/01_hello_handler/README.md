# Exercise 01: Hello Handler

**Learning Goal:** Understand http.HandlerFunc signature and basic response writing

**Difficulty:** Tier 1 - Introduction
**Estimated Time:** 15-20 minutes

## Problem Description

Create your first HTTP handlers that respond to requests with simple text messages. You'll learn the fundamental pattern of reading from the request and writing to the response.

Every HTTP handler in Go has the same signature:
```go
func HandlerName(w http.ResponseWriter, r *http.Request)
```

The `w` (ResponseWriter) is where you write your response. The `r` (Request) contains information about the incoming request.

## Handler Signatures

```go
// HelloHandler writes "Hello, World!" to the response
func HelloHandler(w http.ResponseWriter, r *http.Request)

// GreetHandler reads "name" from query params and writes "Hello, {name}!"
// If name is missing, write "Hello, stranger!"
func GreetHandler(w http.ResponseWriter, r *http.Request)
```

## Examples

### HelloHandler
```
Request:  GET /hello
Response: 200 OK
Body:     Hello, World!
```

### GreetHandler
```
Request:  GET /greet?name=Alice
Response: 200 OK
Body:     Hello, Alice!

Request:  GET /greet
Response: 200 OK
Body:     Hello, stranger!

Request:  GET /greet?name=Bob
Response: 200 OK
Body:     Hello, Bob!
```

## Instructions

1. Implement `HelloHandler` to write "Hello, World!" using `w.Write()`
2. Implement `GreetHandler` to:
   - Get the "name" query parameter from `r.URL.Query().Get("name")`
   - If name is empty, use "stranger"
   - Write "Hello, {name}!" to the response
3. Run tests with `go test -v`

## Hints

**Basic (start here):**
- Use `w.Write([]byte("your text"))` to write response body
- Access query parameters with `r.URL.Query().Get("paramName")`
- String concatenation: `"Hello, " + name + "!"`
- Or use `fmt.Sprintf("Hello, %s!", name)`

**Intermediate:**
- `w.Write()` returns `(int, error)` but you can ignore return values for now
- Query params are always strings (even if they look like numbers)
- Empty string `""` is the zero value for missing parameters

**Complete solution pattern:**
```go
func GreetHandler(w http.ResponseWriter, r *http.Request) {
    // Get query parameter
    param := r.URL.Query().Get("paramName")

    // Handle empty case
    if param == "" {
        param = "default"
    }

    // Build response
    response := fmt.Sprintf("Template %s", param)

    // Write response
    w.Write([]byte(response))
}
```

## Think About

1. What happens if you don't call `w.Write()`? (Hint: Try it and check the response)
2. What HTTP status code is sent by default if you only call `w.Write()`?
3. Can you call `w.Write()` multiple times? What happens?
4. What's the difference between `Get("name")` returning `""` vs the parameter not being in the URL at all?

## What This Teaches

- **Handler function signature**: The fundamental pattern for all HTTP handlers
- **ResponseWriter**: How to write data back to the client
- **Request.URL.Query()**: How to extract query parameters from the URL
- **Default behavior**: Go automatically sets status 200 if you don't specify otherwise
- **Testing pattern**: Using httptest to test handlers without running a server
