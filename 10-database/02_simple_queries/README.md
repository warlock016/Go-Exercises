# Exercise 02: Simple Queries

## Learning Goal
Learn to execute SELECT queries and scan results using `QueryRow` and `Query`.

## Problem Description

Reading data from a database is the most common operation. Go provides two main methods:
- `QueryRow` for single-row results
- `Query` for multiple rows

Both require scanning results into Go variables.

## Function Signatures

```go
// GetUserByID retrieves a single user by ID
func GetUserByID(db *sql.DB, id int) (*User, error)

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *sql.DB) ([]User, error)

// GetUsersByAge retrieves users with age greater than minAge
func GetUsersByAge(db *sql.DB, minAge int) ([]User, error)

// CountUsers returns the total number of users
func CountUsers(db *sql.DB) (int, error)
```

## Hints

### Basic
- `QueryRow` returns a `*sql.Row` that you call `Scan` on
- `Query` returns `*sql.Rows` - iterate with `for rows.Next()`
- Always `defer rows.Close()` after Query

### Intermediate
- `QueryRow` returns `sql.ErrNoRows` if no match
- Check `rows.Err()` after the loop completes
- Use `?` placeholders for parameters (SQLite)

## What This Teaches
- Single vs multi-row queries
- Scanning into variables
- Error handling for empty results
