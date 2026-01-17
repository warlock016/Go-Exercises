# Exercise 01: Database Connection

## Learning Goal
Learn to open, verify, and close database connections using Go's `database/sql` package.

## Problem Description

Database connections are the foundation of all database operations. In Go, the `database/sql` package provides a generic interface that works with any SQL database through drivers.

Key concepts:
- `sql.Open()` returns a connection pool, not a single connection
- `db.Ping()` verifies the connection actually works
- `db.Close()` should be called when done (usually via defer)
- Connection strings vary by database driver

## Function Signatures

```go
// OpenDatabase opens a SQLite database at the given path
// Returns the database connection or an error
func OpenDatabase(path string) (*sql.DB, error)

// VerifyConnection checks if the database is reachable
// Returns nil if connection is healthy, error otherwise
func VerifyConnection(db *sql.DB) error

// CloseDatabase closes the database connection
func CloseDatabase(db *sql.DB) error

// WithDatabase opens a database, runs a function, then closes it
// This pattern ensures proper cleanup
func WithDatabase(path string, fn func(*sql.DB) error) error
```

## Examples

```go
// Basic usage
db, err := OpenDatabase(":memory:")  // SQLite in-memory database
if err != nil {
    log.Fatal(err)
}
defer CloseDatabase(db)

// Verify connection works
if err := VerifyConnection(db); err != nil {
    log.Fatal("database unreachable:", err)
}

// Using WithDatabase for automatic cleanup
err = WithDatabase("test.db", func(db *sql.DB) error {
    // Use db here
    return nil
})
```

## Instructions

1. Import the required packages (`database/sql` and SQLite driver)
2. Implement `OpenDatabase` using `sql.Open`
3. Implement `VerifyConnection` using `db.Ping()`
4. Implement `CloseDatabase` using `db.Close()`
5. Implement `WithDatabase` combining open, function call, and close

## Hints

### Basic
- SQLite driver name is "sqlite3"
- For in-memory SQLite, use path ":memory:"
- `sql.Open` doesn't actually connect - it prepares the pool

### Intermediate
- Always check errors from `sql.Open`
- `db.Ping()` forces an actual connection attempt
- In `WithDatabase`, use defer to ensure cleanup

### Solution Pattern
```go
func WithDatabase(path string, fn func(*sql.DB) error) error {
    db, err := OpenDatabase(path)
    if err != nil {
        return err
    }
    defer CloseDatabase(db)
    return fn(db)
}
```

## Think About

1. Why does `sql.Open` not return an error for an invalid path?
2. When would you use `:memory:` vs a file path?
3. What happens if you forget to close a database connection?
4. Why is `WithDatabase` pattern useful?

## What This Teaches

- Database connection lifecycle management
- The importance of verifying connections
- Resource cleanup patterns with defer
- Higher-order functions for resource management
