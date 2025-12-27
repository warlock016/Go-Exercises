# Exercise 04: Request Body

**Learning Goal:** Parse JSON request bodies, validate input, and return appropriate status codes

**Difficulty:** Tier 2 - Application
**Estimated Time:** 25-30 minutes

## Problem Description

Build a handler that accepts JSON in the request body, parses it into a struct, validates the input, and returns a JSON response with status code 201 Created on success or 400 Bad Request on errors.

Reading request bodies is fundamental to REST APIs. POST, PUT, and PATCH requests typically send data in the body as JSON.

## Type Definitions

```go
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type CreateUserResponse struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Message string `json:"message"`
}
```

## Handler Signature

```go
// CreateUserHandler accepts JSON body with name and email
// Validates that both fields are non-empty
// Returns 201 Created with user data (hardcoded ID=42)
// Returns 400 Bad Request if JSON is malformed or validation fails
func CreateUserHandler(w http.ResponseWriter, r *http.Request)
```

## Examples

### Success Case
```
Request:  POST /users
Content-Type: application/json
Body:     {"name":"Alice","email":"alice@example.com"}

Response: 201 Created
Content-Type: application/json
Body:     {
            "id": 42,
            "name": "Alice",
            "email": "alice@example.com",
            "message": "User created successfully"
          }
```

### Malformed JSON
```
Request:  POST /users
Content-Type: application/json
Body:     {"name":"Alice","email":}

Response: 400 Bad Request
Body:     invalid JSON
```

### Missing Name
```
Request:  POST /users
Content-Type: application/json
Body:     {"email":"alice@example.com"}

Response: 400 Bad Request
Body:     name is required
```

### Missing Email
```
Request:  POST /users
Content-Type: application/json
Body:     {"name":"Alice"}

Response: 400 Bad Request
Body:     email is required
```

### Both Missing
```
Request:  POST /users
Content-Type: application/json
Body:     {}

Response: 400 Bad Request
Body:     name and email are required
```

## Instructions

1. Implement `CreateUserHandler`:
   - Decode JSON from `r.Body` into `CreateUserRequest`
   - If decoding fails, return 400 with message "invalid JSON"
   - Validate that name is not empty
   - Validate that email is not empty
   - If both missing, return 400 with "name and email are required"
   - If only name missing, return 400 with "name is required"
   - If only email missing, return 400 with "email is required"
   - On success, return 201 with CreateUserResponse (ID=42, include message)

2. Run tests with `go test -v`

## Hints

**Basic (start here):**
- Decode body: `json.NewDecoder(r.Body).Decode(&request)`
- Check decode error: `if err != nil { http.Error(...); return }`
- Validate: `if request.Name == "" { ... }`
- Return 201: `w.WriteHeader(http.StatusCreated)` before encoding JSON
- Return 400: `http.Error(w, "message", http.StatusBadRequest)`

**Intermediate:**
- `r.Body` is an `io.ReadCloser`, close it with `defer r.Body.Close()`
- Decode errors can be JSON syntax errors or EOF
- Order matters: set status code BEFORE writing body
- For JSON response, use `w.WriteHeader()` then `json.NewEncoder(w).Encode()`
- For text error, use `http.Error()` which sets status automatically

**Complete solution pattern:**
```go
func CreateHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Parse request body
    var input CreateRequest
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }

    // 2. Validate required fields
    if input.Field == "" {
        http.Error(w, "field is required", http.StatusBadRequest)
        return
    }

    // 3. Process (hardcode ID for now)
    response := CreateResponse{
        ID:    42,
        Field: input.Field,
        Message: "created successfully",
    }

    // 4. Return 201 with JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}
```

## Think About

1. Should you check the Content-Type header to ensure the client sent JSON?
2. What happens if the client sends extra JSON fields not in your struct?
3. How would you handle very large request bodies? (Hint: limit with `http.MaxBytesReader`)
4. Why return different error messages for different validation failures vs a generic message?
5. Should validation errors be 400 Bad Request or 422 Unprocessable Entity?

## What This Teaches

- **Request body parsing**: Using json.NewDecoder with r.Body
- **Error handling**: Different errors need different status codes
- **Validation**: Input validation before processing
- **201 Created**: Proper status for resource creation
- **Status before body**: Must call WriteHeader before Write
- **Resource IDs**: Placeholder IDs before adding database
- **Error messages**: Clear, actionable error messages for clients
