# Exercise 08: Path Parameters

**Learning Goal:** Extract parameters from URL paths without external routers

**Difficulty:** Tier 3 - Integration
**Estimated Time:** 35-40 minutes

## Problem Description

Build handlers that extract IDs and other values from URL paths like `/users/123` or `/orgs/golang/projects/http`. You'll implement path parsing logic to handle RESTful URL patterns.

## Function Signatures

```go
// ParsePathParam extracts a parameter from a URL path
// Example: ParsePathParam("/users/{id}", "/users/123", "id") returns "123"
func ParsePathParam(pattern, path, paramName string) string

// UserAPIHandler handles /users and /users/{id}
// GET /users - returns array of all users
// GET /users/123 - returns single user with ID 123 (404 if not found)
func UserAPIHandler(w http.ResponseWriter, r *http.Request)

// NestedPathHandler handles /orgs/{org}/projects/{project}
// Returns JSON: {"org": "...", "project": "..."}
func NestedPathHandler(w http.ResponseWriter, r *http.Request)
```

## Examples

### UserAPIHandler
```
Request:  GET /users
Response: 200 OK
Body:     [{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]

Request:  GET /users/1
Response: 200 OK
Body:     {"id":1,"name":"Alice"}

Request:  GET /users/999
Response: 404 Not Found
```

### NestedPathHandler
```
Request:  GET /orgs/golang/projects/http
Response: 200 OK
Body:     {"org":"golang","project":"http"}
```

## What This Teaches

- **Path parsing**: Manual URL path parameter extraction
- **RESTful patterns**: Collection vs single resource
- **String manipulation**: Splitting and matching paths
- **Pattern matching**: Simple routing without frameworks
