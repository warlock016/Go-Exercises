package mini_database

import (
	"errors"
	"strings"
	"time"
)

// TODO(human): Define Record struct with Key, Value, CreatedAt, UpdatedAt fields
type Record struct {
	Key       string
	Value     string
	CreatedAt int64
	UpdatedAt int64
}

// TODO(human): Define Operation struct with OpType, Key, OldRecord fields
type Operation struct {
	OpType    string
	Key       string
	OldRecord *Record
}

// TODO(human): Define Stats struct with RecordCount, Reads, Writes, Deletes, IndexCount fields
type Stats struct {
	RecordCount int
	Reads       int64
	Writes      int64
	Deletes     int64
	IndexCount  int
}

// TODO(human): Define Database type with necessary fields
type Database struct {
	Data           map[string]Record
	Indexes        map[string]map[string][]string
	transactionLog []Operation
	inTransaction  bool
	stats          *Stats
}

// NewDatabase creates a new empty database
func NewDatabase() *Database {
	// TODO(human): Implement
	return &Database{
		Data:           map[string]Record{},
		Indexes:        map[string]map[string][]string{},
		transactionLog: []Operation{},
		inTransaction:  false,
		stats:          &Stats{},
	}
}

// CRUD Operations

// Set stores a key-value pair (creates or updates)
func (db *Database) Set(key, value string) error {
	// TODO(human): Implement
	newRec := Record{
		Key:   key,
		Value: value,
	}

	if db.inTransaction {
		newOp := Operation{
			Key:       key,
			OldRecord: &Record{},
		}

		if _, exists := db.Data[key]; exists { // update
			*newOp.OldRecord = db.Data[key]
			newOp.OpType = "update"
		} else { // create
			newOp.OpType = "create"
		}

		db.transactionLog = append(db.transactionLog, newOp)
	}

	if _, exists := db.Data[key]; exists {
		newRec.UpdatedAt = time.Now().Unix()
	} else {
		newRec.CreatedAt = time.Now().Unix()
		db.stats.RecordCount += 1
	}

	db.Data[key] = newRec
	db.stats.Writes += 1

	return nil
}

// Get retrieves a value by key (returns error if not found)
func (db *Database) Get(key string) (string, error) {
	// TODO(human): Implement
	if db.Exists(key) {
		db.stats.Reads += 1
		return db.Data[key].Value, nil
	}
	return "", errors.New("key not found")
}

// Delete removes a key-value pair (returns error if not found)
func (db *Database) Delete(key string) error {
	// TODO(human): Implement
	if _, exists := db.Data[key]; !exists {
		return errors.New("key not found")
	}

	if db.inTransaction {
		newOp := Operation{
			OpType: "delete",
			Key:    key,
			OldRecord: &Record{
				Key:       db.Data[key].Key,
				Value:     db.Data[key].Value,
				CreatedAt: db.Data[key].CreatedAt,
				UpdatedAt: db.Data[key].UpdatedAt,
			},
		}

		*newOp.OldRecord = db.Data[key]

		db.transactionLog = append(db.transactionLog, newOp)
	}

	delete(db.Data, key)
	db.stats.Deletes += 1
	db.stats.RecordCount -= 1

	return nil
}

// Exists checks if a key exists in the database
func (db *Database) Exists(key string) bool {
	// TODO(human): Implement
	// db.stats.Reads += 1
	if _, exists := db.Data[key]; exists {
		return true
	}
	return false
}

// Range Queries

// GetAll returns all keys in the database (order not guaranteed)
func (db *Database) GetAll() []string {
	// TODO(human): Implement
	if db.Data == nil {
		return nil
	}
	result := make([]string, 0, len(db.Data))

	for _, v := range db.Data {
		result = append(result, v.Key)
	}
	return result
}

// GetByPrefix returns all keys that start with the given prefix
func (db *Database) GetByPrefix(prefix string) []string {
	// TODO(human): Implement

	if db.Data == nil { // empty or inexistent DB
		return nil
	}

	result := []string{}

	for k := range db.Data {
		if strings.Contains(k, prefix) {
			result = append(result, k)
		}
	}

	return result
}

// Count returns the number of records in the database
func (db *Database) Count() int {
	// TODO(human): Implement
	return len(db.Data)
}

// Indexing

// CreateIndex creates a secondary index with the given name
func (db *Database) CreateIndex(indexName string) error {
	// TODO(human): Implement
	if _, exists := db.Indexes[indexName]; exists {
		return errors.New("index already exists")
	}

	newIdx := map[string][]string{}
	db.Indexes[indexName] = newIdx
	db.stats.Writes += 1
	db.stats.IndexCount += 1

	return nil
}

// AddToIndex adds a record to an index with the given index value
func (db *Database) AddToIndex(indexName, indexValue, key string) error {
	// TODO(human): Implement

	if _, exists := db.Data[key]; !exists {
		return errors.New("non-existent key")
	}

	newKey := db.Data[key].Key

	if _, exists := db.Indexes[indexName]; !exists {
		db.Indexes[indexName] = map[string][]string{}
	} else if _, exists := db.Indexes[indexName][indexValue]; !exists {
		db.Indexes[indexName][indexValue] = make([]string, 0)
	}

	db.Indexes[indexName][indexValue] = append(db.Indexes[indexName][indexValue], newKey)
	db.stats.Writes += 1
	return nil
}

// QueryIndex returns all keys that have the given value in the index
func (db *Database) QueryIndex(indexName, indexValue string) []string {
	// TODO(human): Implement
	if _, exists := db.Indexes[indexName]; !exists {
		return nil
	} else if _, exists := db.Indexes[indexName][indexValue]; !exists {
		return nil
	}

	result := []string{}
	result = append(result, db.Indexes[indexName][indexValue]...)
	db.stats.Reads += 1

	return result
}

// Transactions

// Begin starts a new transaction
func (db *Database) Begin() error {
	// TODO(human): Implement
	if db.inTransaction {
		return errors.New("transaction already in progress")
	}

	db.transactionLog = []Operation{}
	db.inTransaction = true

	return nil
}

// Commit commits the current transaction (clears transaction log)
func (db *Database) Commit() error {
	// TODO(human): Implement
	if !db.inTransaction {
		return errors.New("no transaction in progress")
	}
	db.inTransaction = false
	db.transactionLog = []Operation{}
	return nil
}

// Rollback rolls back all operations since Begin()
func (db *Database) Rollback() error {
	// TODO(human): Implement
	if !db.inTransaction {
		return errors.New("no transaction in progress")
	}
	// Roll back transactions

	for i := len(db.transactionLog) - 1; i >= 0; i-- {

		switch db.transactionLog[i].OpType {
		case "delete":
			oldKey := db.transactionLog[i].Key
			OldRecord := db.transactionLog[i].OldRecord
			db.Data[oldKey] = *OldRecord
		case "create":
			oldKey := db.transactionLog[i].Key
			delete(db.Data, oldKey)
		case "update":
			oldKey := db.transactionLog[i].Key
			OldRecord := db.transactionLog[i].OldRecord
			db.Data[oldKey] = *OldRecord
		default:
			return errors.New("unknown action: not CUD")
		}
		// db.transactionLog[i].
	}

	db.inTransaction = false
	db.transactionLog = []Operation{}
	return nil
}

// Statistics

// GetStats returns database statistics
func (db *Database) GetStats() Stats {
	// TODO(human): Implement
	if db.stats == nil {
		return Stats{}
	}
	return *db.stats
}

// MemoryUsage estimates the database memory usage in bytes (approximate)
func (db *Database) MemoryUsage() int64 {
	// TODO(human): Implement
	if db == nil {
		return 0
	}

	var result int64

	for k, v := range db.Data {
		result += int64(len(k))
		result += int64(len(v.Key))
		result += int64(len(v.Value))
		result += 8 * 2 // v.createdAt + v.updatedAt
	}

	return result
}
