# Exercise 05: Status Codes

**Learning Goal:** Use semantic HTTP status codes for different response types

**Difficulty:** Tier 2 - Application
**Estimated Time:** 25-30 minutes

## Problem Description

Build handlers that return different HTTP status codes based on the operation: 200 OK for successful retrieval, 201 Created for new resources, 204 No Content for successful deletion, 404 Not Found for missing resources, 301 Moved Permanently for redirects, and 503 Service Unavailable for health checks.

Status codes communicate what happened to the client without them needing to parse the response body.

## Handler Signatures

```go
// ResourceHandler simulates CRUD operations on a resource with ID
// GET: returns 200 with JSON {"id": X, "data": "resource X"} if id > 0, else 404
// POST: returns 201 with JSON {"id": 42, "data": "created"}
// DELETE: returns 204 with no body
// PUT: returns 200 with JSON {"id": X, "data": "updated X"} if id > 0, else 404
func ResourceHandler(w http.ResponseWriter, r *http.Request)

// HealthHandler returns 200 if healthy=true query param, else 503
func HealthHandler(w http.ResponseWriter, r *http.Request)

// RedirectHandler returns 301 redirect to "/new-location"
func RedirectHandler(w http.ResponseWriter, r *http.Request)
```

## Examples

### ResourceHandler - GET
```
Request:  GET /resource?id=5
Response: 200 OK
Body:     {"id":5,"data":"resource 5"}

Request:  GET /resource?id=0
Response: 404 Not Found
Body:     resource not found
```

### ResourceHandler - POST
```
Request:  POST /resource
Response: 201 Created
Body:     {"id":42,"data":"created"}
```

### ResourceHandler - DELETE
```
Request:  DELETE /resource?id=5
Response: 204 No Content
Body:     (empty)
```

### ResourceHandler - PUT
```
Request:  PUT /resource?id=5
Response: 200 OK
Body:     {"id":5,"data":"updated 5"}

Request:  PUT /resource?id=0
Response: 404 Not Found
Body:     resource not found
```

### HealthHandler
```
Request:  GET /health?healthy=true
Response: 200 OK
Body:     {"status":"healthy"}

Request:  GET /health?healthy=false
Response: 503 Service Unavailable
Body:     {"status":"unhealthy"}

Request:  GET /health
Response: 503 Service Unavailable
Body:     {"status":"unhealthy"}
```

### RedirectHandler
```
Request:  GET /old-location
Response: 301 Moved Permanently
Header:   Location: /new-location
Body:     (empty or browser follows redirect)
```

## Instructions

1. Implement `ResourceHandler`:
   - Check `r.Method` to determine operation
   - Parse "id" query parameter as int
   - GET: return 200 with resource data if id > 0, else 404
   - POST: return 201 with created resource (id=42)
   - DELETE: return 204 with no body
   - PUT: return 200 with updated data if id > 0, else 404
   - For unsupported methods, return 405

2. Implement `HealthHandler`:
   - Check "healthy" query parameter
   - If "true", return 200 with {"status":"healthy"}
   - Otherwise (false or missing), return 503 with {"status":"unhealthy"}

3. Implement `RedirectHandler`:
   - Use `http.Redirect()` with 301 status to "/new-location"

4. Run tests with `go test -v`

## Hints

**Basic (start here):**
- Check method: `if r.Method == http.MethodGet { ... }`
- Get ID: `idStr := r.URL.Query().Get("id"); id, _ := strconv.Atoi(idStr)`
- Return 404: `http.Error(w, "resource not found", http.StatusNotFound)`
- Return 204: `w.WriteHeader(http.StatusNoContent)` (no body)
- Return 201: `w.WriteHeader(http.StatusCreated)` then encode JSON
- Redirect: `http.Redirect(w, r, "/new-location", http.StatusMovedPermanently)`

**Intermediate:**
- Use a switch statement for method routing
- 204 No Content should have truly empty body (no w.Write)
- For JSON responses, set Content-Type header
- Check query param: `healthy := r.URL.Query().Get("healthy") == "true"`
- Return 405 for unsupported methods

**Complete solution pattern:**
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // Return 200 or 404
        exists := checkIfExists()
        if !exists {
            http.Error(w, "not found", http.StatusNotFound)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)

    case http.MethodPost:
        // Return 201
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(created)

    case http.MethodDelete:
        // Return 204
        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}
```

## Think About

1. Why use 201 Created instead of 200 OK for POST requests?
2. Why does 204 No Content have no body? What if you want to return data?
3. What's the difference between 301 and 302 redirects?
4. Should health check failures be 503 or 500?
5. What happens if you write a body with status 204?

## What This Teaches

- **Status code semantics**: Each code has specific meaning
- **Method-based routing**: Different behavior per HTTP method
- **2xx Success codes**: 200, 201, 204 for different success types
- **4xx Client errors**: 404 for missing resources, 405 for wrong method
- **5xx Server errors**: 503 for service unavailable
- **3xx Redirects**: 301 for permanent redirects
- **No content responses**: 204 has truly empty body
- **RESTful patterns**: CRUD operations mapped to HTTP methods
