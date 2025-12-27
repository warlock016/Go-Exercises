# Exercise 03: JSON Response

**Learning Goal:** Master JSON encoding, struct tags, and Content-Type headers

**Difficulty:** Tier 1 - Introduction
**Estimated Time:** 20-25 minutes

## Problem Description

Build handlers that return structured JSON responses with proper field naming and content types. You'll work with structs, JSON tags, and learn how Go's encoding/json package works with http.ResponseWriter.

JSON is the standard format for HTTP APIs. Go makes it easy to convert structs to JSON, and you control field names using struct tags.

## Type Definitions

```go
type User struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    IsActive  bool   `json:"is_active"`
}

type Stats struct {
    TotalUsers   int `json:"total_users"`
    ActiveUsers  int `json:"active_users"`
    InactiveUsers int `json:"inactive_users"`
}
```

## Handler Signatures

```go
// GetUserHandler returns a single hardcoded user
// User ID=1, Name="Alice", Email="alice@example.com", IsActive=true
func GetUserHandler(w http.ResponseWriter, r *http.Request)

// ListUsersHandler returns array of 3 hardcoded users
// Users: Alice (1, active), Bob (2, inactive), Charlie (3, active)
func ListUsersHandler(w http.ResponseWriter, r *http.Request)

// StatsHandler returns statistics about users
// TotalUsers=3, ActiveUsers=2, InactiveUsers=1
func StatsHandler(w http.ResponseWriter, r *http.Request)
```

## Examples

### GetUserHandler
```
Request:  GET /user
Response: 200 OK
Content-Type: application/json
Body:     {
            "id": 1,
            "name": "Alice",
            "email": "alice@example.com",
            "is_active": true
          }
```

### ListUsersHandler
```
Request:  GET /users
Response: 200 OK
Content-Type: application/json
Body:     [
            {
              "id": 1,
              "name": "Alice",
              "email": "alice@example.com",
              "is_active": true
            },
            {
              "id": 2,
              "name": "Bob",
              "email": "bob@example.com",
              "is_active": false
            },
            {
              "id": 3,
              "name": "Charlie",
              "email": "charlie@example.com",
              "is_active": true
            }
          ]
```

### StatsHandler
```
Request:  GET /stats
Response: 200 OK
Content-Type: application/json
Body:     {
            "total_users": 3,
            "active_users": 2,
            "inactive_users": 1
          }
```

## Instructions

1. Implement `GetUserHandler`:
   - Create a User with ID=1, Name="Alice", Email="alice@example.com", IsActive=true
   - Set Content-Type header to "application/json"
   - Encode user as JSON to the response

2. Implement `ListUsersHandler`:
   - Create a slice of 3 Users (Alice, Bob, Charlie)
   - Return as JSON array

3. Implement `StatsHandler`:
   - Create Stats struct with correct counts
   - Return as JSON

4. Run tests with `go test -v`

## Hints

**Basic (start here):**
- Create struct: `user := User{ID: 1, Name: "Alice", ...}`
- Set header: `w.Header().Set("Content-Type", "application/json")`
- Encode JSON: `json.NewEncoder(w).Encode(user)`
- For array: `users := []User{user1, user2, user3}`
- JSON tag format: `` `json:"field_name"` ``

**Intermediate:**
- `json.NewEncoder(w)` writes directly to ResponseWriter
- Alternative: `json.Marshal()` returns bytes, then use `w.Write()`
- Struct tags control JSON field names (use snake_case for consistency)
- Omit fields with `,omitempty`: `` `json:"name,omitempty"` ``
- Always set Content-Type BEFORE encoding

**Complete solution pattern:**
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // 1. Create data structure
    data := MyStruct{
        Field1: "value",
        Field2: 42,
    }

    // 2. Set Content-Type header
    w.Header().Set("Content-Type", "application/json")

    // 3. Encode to JSON
    if err := json.NewEncoder(w).Encode(data); err != nil {
        // Encoding failed (rare, usually means struct isn't serializable)
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }
}
```

## Think About

1. What happens if you forget to set the Content-Type header? Will clients parse it correctly?
2. What's the difference between `json.NewEncoder(w).Encode()` and `json.Marshal()`?
3. Why do we use snake_case in JSON tags (is_active) vs Go's PascalCase (IsActive)?
4. What happens if you try to encode a struct with unexported fields?
5. How would you omit a field from JSON if it's empty/zero?

## What This Teaches

- **JSON encoding**: Converting Go structs to JSON
- **Struct tags**: Controlling JSON field names
- **Content-Type**: Why it matters for APIs
- **Arrays in JSON**: Encoding slices as JSON arrays
- **json.NewEncoder**: Streaming encoder that writes directly to http.ResponseWriter
- **Naming conventions**: Go (PascalCase) vs JSON (snake_case) field names
- **Hardcoded data**: Starting simple before adding databases
