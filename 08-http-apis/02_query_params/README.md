# Exercise 02: Query Parameters

**Learning Goal:** Parse and validate URL query parameters, return JSON responses

**Difficulty:** Tier 1 - Introduction
**Estimated Time:** 20-25 minutes

## Problem Description

Build handlers that parse multiple query parameters from URLs and return JSON responses. You'll learn to handle required vs optional parameters, provide defaults, and parse multiple values for the same parameter.

Query parameters appear after the `?` in a URL:
```
/search?q=golang&page=2&limit=10
         └─────────────────────┘
            query parameters
```

## Handler Signatures

```go
// SearchHandler expects:
//   - q (required): search query string
//   - page (optional, default 1): page number
//   - limit (optional, default 10): results per page
// Returns JSON: {"query": "...", "page": 1, "limit": 10}
// Returns 400 if "q" is missing
func SearchHandler(w http.ResponseWriter, r *http.Request)

// FilterHandler expects:
//   - tag (can be multiple): filter tags
// Returns JSON: {"tags": ["tag1", "tag2"]}
// If no tags, returns empty array: {"tags": []}
func FilterHandler(w http.ResponseWriter, r *http.Request)
```

## Examples

### SearchHandler
```
Request:  GET /search?q=golang
Response: 200 OK
Content-Type: application/json
Body:     {"query":"golang","page":1,"limit":10}

Request:  GET /search?q=golang&page=2
Response: 200 OK
Body:     {"query":"golang","page":2,"limit":10}

Request:  GET /search?q=golang&page=3&limit=25
Response: 200 OK
Body:     {"query":"golang","page":3,"limit":25}

Request:  GET /search
Response: 400 Bad Request
Body:     query parameter 'q' is required
```

### FilterHandler
```
Request:  GET /filter?tag=golang&tag=http&tag=json
Response: 200 OK
Content-Type: application/json
Body:     {"tags":["golang","http","json"]}

Request:  GET /filter?tag=backend
Response: 200 OK
Body:     {"tags":["backend"]}

Request:  GET /filter
Response: 200 OK
Body:     {"tags":[]}
```

## Instructions

1. Implement `SearchHandler`:
   - Check if "q" parameter exists, return 400 error if missing
   - Get "page" parameter, default to 1
   - Get "limit" parameter, default to 10
   - Create SearchResult struct with these values
   - Return as JSON with proper Content-Type header

2. Implement `FilterHandler`:
   - Get all "tag" parameter values (there can be multiple)
   - Create FilterResult struct with tags array
   - Return as JSON

3. Run tests with `go test -v`

## Hints

**Basic (start here):**
- Use `r.URL.Query().Get("param")` for single values
- Use `r.URL.Query()["param"]` for multiple values (returns []string)
- Set header: `w.Header().Set("Content-Type", "application/json")`
- Return error: `http.Error(w, "message", http.StatusBadRequest)`
- Encode JSON: `json.NewEncoder(w).Encode(data)`
- Query params are always strings, convert with `strconv.Atoi()`

**Intermediate:**
- Always set headers BEFORE writing the body
- `http.Error()` automatically sets Content-Type to text/plain
- For JSON errors, manually set status and encode JSON
- If parameter is missing, `Get()` returns empty string `""`
- Multiple values: `/path?tag=a&tag=b` gives `[]string{"a", "b"}`

**Complete solution pattern:**
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // 1. Get and validate required params
    required := r.URL.Query().Get("required")
    if required == "" {
        http.Error(w, "required param missing", http.StatusBadRequest)
        return
    }

    // 2. Get optional params with defaults
    optional := r.URL.Query().Get("optional")
    if optional == "" {
        optional = "default"
    }

    // 3. Convert string to int if needed
    num := 1
    if numStr := r.URL.Query().Get("num"); numStr != "" {
        num, _ = strconv.Atoi(numStr)  // Handle error in real code
    }

    // 4. Get multiple values
    tags := r.URL.Query()["tags"]

    // 5. Build response struct
    response := struct {
        Field string   `json:"field"`
        Tags  []string `json:"tags"`
    }{
        Field: required,
        Tags:  tags,
    }

    // 6. Write JSON response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
```

## Think About

1. What happens if you convert "abc" to int with `strconv.Atoi()`? Should you validate?
2. Why must you set the Content-Type header before writing the body?
3. What's the difference between `Query().Get("tag")` and `Query()["tag"]`?
4. How would you handle a page number of 0 or negative?
5. Should you validate that limit isn't too large (like 10000)?

## What This Teaches

- **Query parameter extraction**: Single vs multiple values
- **Default values**: Handling optional parameters
- **Validation**: Required parameters and error responses
- **JSON encoding**: Using json.NewEncoder with ResponseWriter
- **Content-Type header**: Proper JSON content type
- **Error responses**: Using http.Error for simple text errors
- **String conversion**: strconv.Atoi for numeric parameters
