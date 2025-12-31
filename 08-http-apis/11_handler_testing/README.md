# Exercise 11: Handler Testing

**Learning Goal:** Write comprehensive tests for HTTP handlers

**Difficulty:** Tier 3 - Integration
**Estimated Time:** 30-35 minutes

## Problem Description

You're provided a complete UserService implementation with handlers. Your task is to write comprehensive tests covering success cases, error cases, headers, and edge cases using httptest patterns.

This reverses the usual flow - you write tests for working code to practice test-driven thinking.

## Provided Implementation

```go
type UserService struct {
    users map[int]User
}

// GET /users - list all
// POST /users - create (requires JSON body)
// GET /users/:id - get by ID
```

## Your Task

Write tests for:
1. ListUsers - returns array, correct content-type
2. CreateUser - accepts JSON, returns 201, validates required fields
3. GetUser - returns user, handles 404, validates ID format
4. Error cases - malformed JSON, missing fields, invalid IDs
5. Headers - Content-Type, status codes

## What This Teaches

- **Comprehensive testing**: Covering happy path and edge cases
- **Table-driven tests**: Organizing test cases efficiently
- **httptest patterns**: NewRequest, NewRecorder
- **Test organization**: One table per handler
- **Error testing**: Ensuring errors are handled correctly



## User notes:

```go
	a := "/users/1234"
    // 1. Trim removes all leading and trailing "/"
	b := strings.Trim(a, "/")
    // 2. Split separates the string around inner "/"
	res := strings.Split(b, "/")

    // short form: res := strings.Split(strings.Trim(a, "/"), "/")

```