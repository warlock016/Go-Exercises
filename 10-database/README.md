# Module 10: Database

**Focus Area:** SQL Database Operations with `database/sql`
**Prerequisites:** Module 05 (Structs), Module 06 (Error Handling), Module 07 (Interfaces)
**Estimated Time:** 10-12 hours
**Exercises:** 12

---

## Overview

Database operations are fundamental to most backend applications. This module teaches you how to work with SQL databases using Go's standard `database/sql` package. You'll learn connection management, query execution, transactions, and testing patterns.

**What You'll Learn:**
- Opening and managing database connections
- Executing queries and handling results
- Prepared statements and SQL injection prevention
- Transaction management with commit/rollback
- Handling NULL values with sql.Null types
- Connection pooling and performance tuning
- Repository pattern for clean architecture
- Testing database code effectively

---

## Module Structure

### Tier 1: Introduction (Exercises 01-03)
Build foundational database skills.

| Exercise | Concept | Time |
|----------|---------|------|
| 01_connection | Open, Ping, Close patterns | 30-35 min |
| 02_simple_queries | QueryRow, Query, Scan | 35-40 min |
| 03_exec_statements | INSERT, UPDATE, DELETE with Exec | 30-35 min |

### Tier 2: Application (Exercises 04-07)
Apply database patterns to practical problems.

| Exercise | Concept | Time |
|----------|---------|------|
| 04_prepared_statements | Prepare, reuse, prevent SQL injection | 35-40 min |
| 05_transactions | Begin, Commit, Rollback patterns | 40-45 min |
| 06_null_handling | sql.NullString, sql.NullInt64, etc. | 35-40 min |
| 07_scanning_structs | Scan into structs, custom Scanner | 45-50 min |

### Tier 3: Integration (Exercises 08-10)
Combine database techniques for production use.

| Exercise | Concept | Time |
|----------|---------|------|
| 08_migrations | Schema versioning patterns | 50-60 min |
| 09_connection_pool | SetMaxOpenConns, SetMaxIdleConns | 45-50 min |
| 10_context_timeouts | QueryContext, ExecContext with deadlines | 45-50 min |

### Tier 4: Mastery (Exercises 11-12)
Advanced patterns and testing.

| Exercise | Concept | Time |
|----------|---------|------|
| 11_repository_pattern | Clean architecture with abstraction | 55-65 min |
| 12_testing_databases | Test fixtures, cleanup, in-memory DBs | 60-75 min |

---

## Key Concepts

### Opening a Database Connection
```go
db, err := sql.Open("sqlite3", ":memory:")
if err != nil {
    return err
}
defer db.Close()

// Verify connection
if err := db.Ping(); err != nil {
    return err
}
```

### Query Patterns
```go
// Single row
var name string
err := db.QueryRow("SELECT name FROM users WHERE id = ?", id).Scan(&name)

// Multiple rows
rows, err := db.Query("SELECT id, name FROM users")
defer rows.Close()
for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
}
```

### Transactions
```go
tx, err := db.Begin()
if err != nil {
    return err
}
defer tx.Rollback() // No-op if committed

// Execute operations
tx.Exec("INSERT INTO users (name) VALUES (?)", name)

return tx.Commit()
```

---

## Setup

This module uses SQLite for simplicity (no external database required).

```bash
# The go.mod already includes the SQLite driver
go get github.com/mattn/go-sqlite3
```

---

## Testing Commands

### Run All Tests
```bash
cd 10-database
go test ./...
```

### Run Single Exercise
```bash
cd 10-database/01_connection
go test -v
```

### Run with Race Detection
```bash
go test -race ./...
```

---

## Success Criteria

By completing this module, you should be able to:
- Open and manage database connections properly
- Write queries that are safe from SQL injection
- Handle NULL values correctly
- Use transactions for atomic operations
- Configure connection pools for production use
- Write testable database code with proper abstractions
