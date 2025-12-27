# Exercise 07: Method Routing

**Learning Goal:** Route requests based on HTTP method, return 405 for unsupported methods

**Difficulty:** Tier 2 - Application
**Estimated Time:** 30-35 minutes

## Problem Description

Build a handler that routes to different logic based on HTTP method (GET/POST/PUT/DELETE), maintains in-memory state, and returns 405 Method Not Allowed with proper Allow header for unsupported methods.

## Type Definitions

```go
type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type ItemStore struct {
    items map[int]Item
    nextID int
}
```

## Handler Signature

```go
// ItemHandler handles CRUD operations on items
// GET /items?id=X - returns item with ID X (404 if not found)
// POST /items with JSON body {"name":"..."} - creates item, returns 201
// PUT /items?id=X with JSON body {"name":"..."} - updates item, returns 200 (404 if not found)
// DELETE /items?id=X - deletes item, returns 204
// Other methods - returns 405 with Allow header
func ItemHandler(store *ItemStore) http.HandlerFunc
```

## Instructions

1. Implement ItemStore methods: NewItemStore(), Get(), Create(), Update(), Delete()
2. Implement ItemHandler that routes based on r.Method
3. Return 405 with "Allow" header listing supported methods
4. Run tests with `go test -v`

## What This Teaches

- **Method-based routing**: Switch on r.Method
- **Stateful handlers**: Using closures to inject dependencies
- **405 Method Not Allowed**: Proper error for wrong HTTP method
- **Allow header**: Communicating supported methods
- **In-memory storage**: Simple CRUD without database
