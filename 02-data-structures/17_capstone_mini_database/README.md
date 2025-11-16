# Exercise 17: Capstone - Mini Database

## 🎯 Learning Goal
Build a comprehensive in-memory key-value database that combines all data structure concepts from this module: maps, slices, structs, custom types, performance optimization, and complex state management. This capstone demonstrates real-world application of data structures in a production-like system.

## 📝 Problem Description

You'll implement a fully-featured in-memory database similar to Redis or memcached, but simplified. This database will support:
- **CRUD Operations** - Create, Read, Update, Delete records
- **Indexing** - Secondary indexes for fast lookups by different fields
- **Transactions** - Begin, commit, rollback with ACID-like properties
- **Range Queries** - Get all keys, search by prefix
- **Statistics** - Track operations, estimate memory usage

This exercise integrates everything you've learned:
- Maps for key-value storage
- Slices for maintaining ordered data and transaction logs
- Structs for complex data modeling
- Performance optimization techniques
- State management and error handling

## 🔧 Type Definitions and Function Signatures

Implement in `mini_database.go`:

```go
// Record represents a database record with metadata
type Record struct {
	Key       string
	Value     string
	CreatedAt int64  // Unix timestamp
	UpdatedAt int64  // Unix timestamp
}

// Operation represents a single database operation for transaction rollback
type Operation struct {
	OpType   string  // "set", "delete"
	Key      string
	OldValue *Record // Previous value (nil if key didn't exist)
}

// Database is the main in-memory database
type Database struct {
	// TODO(human): Add fields for:
	// - data map[string]Record (main storage)
	// - indexes map[string]map[string][]string (secondary indexes: indexName -> indexValue -> []keys)
	// - transactionLog []Operation (for rollback)
	// - inTransaction bool (transaction state flag)
	// - stats struct with Reads, Writes, Deletes counters
}

// NewDatabase creates a new empty database
func NewDatabase() *Database

// CRUD Operations

// Set stores a key-value pair (creates or updates)
func (db *Database) Set(key, value string) error

// Get retrieves a value by key (returns error if not found)
func (db *Database) Get(key string) (string, error)

// Delete removes a key-value pair (returns error if not found)
func (db *Database) Delete(key string) error

// Exists checks if a key exists in the database
func (db *Database) Exists(key string) bool

// Range Queries

// GetAll returns all keys in the database (order not guaranteed)
func (db *Database) GetAll() []string

// GetByPrefix returns all keys that start with the given prefix
func (db *Database) GetByPrefix(prefix string) []string

// Count returns the number of records in the database
func (db *Database) Count() int

// Indexing

// CreateIndex creates a secondary index on a field
// For simplicity, we'll index by extracting a value from the record value
// using a provided extractor function
func (db *Database) CreateIndex(indexName string) error

// AddToIndex adds a record to an index with the given index value
func (db *Database) AddToIndex(indexName, indexValue, key string) error

// QueryIndex returns all keys that have the given value in the index
func (db *Database) QueryIndex(indexName, indexValue string) []string

// Transactions

// Begin starts a new transaction
func (db *Database) Begin() error

// Commit commits the current transaction (clears transaction log)
func (db *Database) Commit() error

// Rollback rolls back all operations since Begin()
func (db *Database) Rollback() error

// Statistics

// GetStats returns database statistics
type Stats struct {
	RecordCount int
	Reads       int64
	Writes      int64
	Deletes     int64
	IndexCount  int
}

func (db *Database) GetStats() Stats

// MemoryUsage estimates the database memory usage in bytes (approximate)
func (db *Database) MemoryUsage() int64
```

## 💡 Examples

```go
// Create database
db := NewDatabase()

// Basic CRUD
db.Set("user:1", "Alice")
db.Set("user:2", "Bob")
value, _ := db.Get("user:1")  // "Alice"
db.Delete("user:1")
exists := db.Exists("user:1")  // false

// Range queries
db.Set("user:1", "Alice")
db.Set("user:2", "Bob")
db.Set("product:1", "Laptop")
users := db.GetByPrefix("user:")  // ["user:1", "user:2"]
all := db.GetAll()  // ["user:1", "user:2", "product:1"]

// Indexing
db.CreateIndex("usernames")
db.Set("user:1", "Alice")
db.AddToIndex("usernames", "Alice", "user:1")
keys := db.QueryIndex("usernames", "Alice")  // ["user:1"]

// Transactions
db.Begin()
db.Set("key1", "value1")
db.Set("key2", "value2")
db.Delete("key3")
db.Commit()  // Changes are permanent

// Rollback example
db.Set("important", "original")
db.Begin()
db.Set("important", "modified")
db.Set("temp", "data")
db.Rollback()  // "important" back to "original", "temp" not saved

// Statistics
stats := db.GetStats()
// stats.RecordCount = 5
// stats.Writes = 10
// stats.Reads = 3
memory := db.MemoryUsage()  // Approximate bytes
```

## 📋 Instructions

1. **Database struct:**
   - `data map[string]Record` - Main storage
   - `indexes map[string]map[string][]string` - Secondary indexes
   - `transactionLog []Operation` - Operations since Begin()
   - `inTransaction bool` - Transaction state
   - Stats counters (embed or separate struct)

2. **Set operation:**
   - Get current timestamp: `time.Now().Unix()`
   - Check if key exists (update vs create)
   - If in transaction, log operation with old value
   - Create/update Record
   - Store in `data` map
   - Increment write counter

3. **Delete operation:**
   - Check if key exists (return error if not)
   - If in transaction, log operation with old value
   - Remove from `data` map
   - Increment delete counter

4. **GetByPrefix:**
   - Iterate over all keys in `data`
   - Use `strings.HasPrefix(key, prefix)` to filter
   - Collect matching keys

5. **Transactions:**
   - Begin: Set `inTransaction = true`, clear `transactionLog`
   - Each operation in transaction: append to `transactionLog`
   - Commit: Set `inTransaction = false`, clear `transactionLog`
   - Rollback: Iterate `transactionLog` in reverse, restore old values

6. **Indexing:**
   - CreateIndex: Add entry to `indexes` map
   - AddToIndex: Append key to `indexes[indexName][indexValue]`
   - QueryIndex: Return `indexes[indexName][indexValue]`

7. **MemoryUsage:**
   - Estimate: `len(key) + len(value)` for each record
   - Add overhead for Record struct (16 bytes for timestamps)
   - Add overhead for indexes

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected test count: ~50-60 tests covering all functionality

## 🤔 Think About

1. **Why use a transaction log?**
   - Enables rollback by reversing operations
   - Maintains ACID properties (Atomicity)
   - Real databases use similar Write-Ahead Logging (WAL)

2. **What are the performance trade-offs?**
   - Indexes: Faster queries but slower writes and more memory
   - Transactions: Overhead of logging operations
   - In-memory: Fast but limited by RAM

3. **How would you make this production-ready?**
   - Persistence (save to disk)
   - Concurrency (mutexes, goroutines)
   - TTL (time-to-live) for records
   - Memory limits and eviction policies
   - Replication and clustering

4. **What's the difference between this and Redis?**
   - Redis: Persistent, networked, supports more data types
   - This: In-memory only, single-process, educational
   - Both: Key-value stores with similar operations

## 💡 Hints

<details>
<summary>Hint 1: Database struct layout</summary>

```go
type Database struct {
    data           map[string]Record
    indexes        map[string]map[string][]string
    transactionLog []Operation
    inTransaction  bool
    reads          int64
    writes         int64
    deletes        int64
}

func NewDatabase() *Database {
    return &Database{
        data:           make(map[string]Record),
        indexes:        make(map[string]map[string][]string),
        transactionLog: []Operation{},
        inTransaction:  false,
    }
}
```
</details>

<details>
<summary>Hint 2: Set with transaction support</summary>

```go
import "time"

func (db *Database) Set(key, value string) error {
    now := time.Now().Unix()

    // If in transaction, log the operation
    if db.inTransaction {
        oldRecord, exists := db.data[key]
        var oldValue *Record
        if exists {
            oldValue = &oldRecord
        }
        db.transactionLog = append(db.transactionLog, Operation{
            OpType:   "set",
            Key:      key,
            OldValue: oldValue,
        })
    }

    // Create or update record
    record := Record{
        Key:       key,
        Value:     value,
        CreatedAt: now,
        UpdatedAt: now,
    }

    // If updating, preserve CreatedAt
    if existing, exists := db.data[key]; exists {
        record.CreatedAt = existing.CreatedAt
    }

    db.data[key] = record
    db.writes++
    return nil
}
```
</details>

<details>
<summary>Hint 3: Transaction rollback</summary>

```go
func (db *Database) Rollback() error {
    if !db.inTransaction {
        return errors.New("no transaction in progress")
    }

    // Process log in reverse order
    for i := len(db.transactionLog) - 1; i >= 0; i-- {
        op := db.transactionLog[i]

        switch op.OpType {
        case "set":
            if op.OldValue == nil {
                // Key didn't exist, delete it
                delete(db.data, op.Key)
            } else {
                // Restore old value
                db.data[op.Key] = *op.OldValue
            }
        case "delete":
            // Restore deleted record
            if op.OldValue != nil {
                db.data[op.Key] = *op.OldValue
            }
        }
    }

    db.inTransaction = false
    db.transactionLog = []Operation{}
    return nil
}
```
</details>

<details>
<summary>Hint 4: GetByPrefix implementation</summary>

```go
import "strings"

func (db *Database) GetByPrefix(prefix string) []string {
    var keys []string
    for key := range db.data {
        if strings.HasPrefix(key, prefix) {
            keys = append(keys, key)
        }
    }
    return keys
}
```
</details>

<details>
<summary>Hint 5: Memory usage estimation</summary>

```go
func (db *Database) MemoryUsage() int64 {
    var total int64

    // Data storage
    for key, record := range db.data {
        total += int64(len(key))           // Key string
        total += int64(len(record.Value))  // Value string
        total += 16                        // Two int64 timestamps
    }

    // Indexes storage
    for indexName, indexMap := range db.indexes {
        total += int64(len(indexName))
        for indexValue, keys := range indexMap {
            total += int64(len(indexValue))
            for _, key := range keys {
                total += int64(len(key))
            }
        }
    }

    // Transaction log
    for _, op := range db.transactionLog {
        total += int64(len(op.OpType))
        total += int64(len(op.Key))
        if op.OldValue != nil {
            total += int64(len(op.OldValue.Value))
        }
    }

    return total
}
```
</details>

## 🎓 What This Teaches

- **System design** - Architecting a complete system with multiple interacting components
- **Map usage** - Using maps for fast key-value lookups (O(1) average case)
- **Slice manipulation** - Transaction logs, collecting results, filtering
- **Struct composition** - Record and Operation types modeling real data
- **State management** - Transaction states, tracking statistics
- **Error handling** - Returning meaningful errors for invalid operations
- **Performance considerations** - Indexing trade-offs, memory estimation
- **Real-world patterns** - CRUD operations, transactions, indexing like production databases
- **Code organization** - Grouping related functionality, clear interfaces
- **Testing complex systems** - Verifying interactions between multiple components

## 🔥 Bonus Challenges

If you finish the main implementation, try these extensions:

1. **TTL (Time-To-Live):** Add expiration to records, automatically delete expired keys
2. **Sorted Sets:** Implement a sorted set data structure with score-based ordering
3. **Batch Operations:** SetMany, DeleteMany for bulk operations
4. **Query Language:** Simple query parser for "GET WHERE prefix = X"
5. **Persistence:** SaveToFile and LoadFromFile using JSON encoding
6. **Concurrency:** Add mutex locks to make it thread-safe
7. **LRU Eviction:** Implement Least Recently Used eviction when memory limit reached

---

**Congratulations!** You've completed the Data Structures module. You now have hands-on experience with:
- Slices, maps, structs, and custom types
- Performance optimization and benchmarking
- Generic collections and type safety
- Graph algorithms (BFS, DFS)
- Building production-like systems

**Next Module:** `03-algorithms` - Sorting, searching, recursion, and algorithmic thinking
