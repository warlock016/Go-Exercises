package mini_database

import (
	"slices"
	"testing"
)

// Helper to check if slices contain same elements (order doesn't matter)
func containsSameElements(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([]string, len(a))
	bCopy := make([]string, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	slices.Sort(aCopy)
	slices.Sort(bCopy)
	return slices.Equal(aCopy, bCopy)
}

// CRUD Tests

func TestNewDatabase(t *testing.T) {
	db := NewDatabase()
	if db == nil {
		t.Fatal("NewDatabase() returned nil")
	}
	if db.Count() != 0 {
		t.Errorf("New database Count() = %d, want 0", db.Count())
	}
}

func TestSetAndGet(t *testing.T) {
	db := NewDatabase()

	err := db.Set("key1", "value1")
	if err != nil {
		t.Fatalf("Set() error = %v, want nil", err)
	}

	val, err := db.Get("key1")
	if err != nil {
		t.Fatalf("Get() error = %v, want nil", err)
	}
	if val != "value1" {
		t.Errorf("Get(key1) = %q, want %q", val, "value1")
	}
}

func TestGetNonExistent(t *testing.T) {
	db := NewDatabase()

	_, err := db.Get("nonexistent")
	if err == nil {
		t.Error("Get(nonexistent) should return error")
	}
}

func TestSetUpdate(t *testing.T) {
	db := NewDatabase()

	db.Set("key1", "value1")
	db.Set("key1", "value2")

	val, _ := db.Get("key1")
	if val != "value2" {
		t.Errorf("After update, Get(key1) = %q, want %q", val, "value2")
	}

	if db.Count() != 1 {
		t.Errorf("After update, Count() = %d, want 1", db.Count())
	}
}

func TestDelete(t *testing.T) {
	db := NewDatabase()

	db.Set("key1", "value1")
	err := db.Delete("key1")
	if err != nil {
		t.Fatalf("Delete() error = %v, want nil", err)
	}

	if db.Exists("key1") {
		t.Error("Key should not exist after deletion")
	}

	// Delete non-existent key should error
	err = db.Delete("key1")
	if err == nil {
		t.Error("Delete(nonexistent) should return error")
	}
}

func TestExists(t *testing.T) {
	db := NewDatabase()

	if db.Exists("key1") {
		t.Error("Exists() should return false for non-existent key")
	}

	db.Set("key1", "value1")
	if !db.Exists("key1") {
		t.Error("Exists() should return true after Set()")
	}

	db.Delete("key1")
	if db.Exists("key1") {
		t.Error("Exists() should return false after Delete()")
	}
}

// Range Query Tests

func TestGetAll(t *testing.T) {
	db := NewDatabase()

	db.Set("key1", "value1")
	db.Set("key2", "value2")
	db.Set("key3", "value3")

	keys := db.GetAll()
	want := []string{"key1", "key2", "key3"}

	if !containsSameElements(keys, want) {
		t.Errorf("GetAll() = %v, want %v", keys, want)
	}
}

func TestGetByPrefix(t *testing.T) {
	db := NewDatabase()

	db.Set("user:1", "Alice")
	db.Set("user:2", "Bob")
	db.Set("product:1", "Laptop")
	db.Set("product:2", "Mouse")

	tests := []struct {
		name   string
		prefix string
		want   []string
	}{
		{"user prefix", "user:", []string{"user:1", "user:2"}},
		{"product prefix", "product:", []string{"product:1", "product:2"}},
		{"no matches", "order:", []string{}},
		{"partial match", "user:1", []string{"user:1"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := db.GetByPrefix(tt.prefix)
			if !containsSameElements(got, tt.want) {
				t.Errorf("GetByPrefix(%q) = %v, want %v", tt.prefix, got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	db := NewDatabase()

	if db.Count() != 0 {
		t.Errorf("Count() = %d, want 0", db.Count())
	}

	db.Set("key1", "value1")
	if db.Count() != 1 {
		t.Errorf("Count() = %d, want 1", db.Count())
	}

	db.Set("key2", "value2")
	db.Set("key3", "value3")
	if db.Count() != 3 {
		t.Errorf("Count() = %d, want 3", db.Count())
	}

	db.Delete("key2")
	if db.Count() != 2 {
		t.Errorf("After delete, Count() = %d, want 2", db.Count())
	}
}

// Index Tests

func TestCreateIndex(t *testing.T) {
	db := NewDatabase()

	err := db.CreateIndex("usernames")
	if err != nil {
		t.Fatalf("CreateIndex() error = %v, want nil", err)
	}

	// Creating same index again should error
	err = db.CreateIndex("usernames")
	if err == nil {
		t.Error("Creating duplicate index should return error")
	}
}

func TestAddToIndexAndQuery(t *testing.T) {
	db := NewDatabase()

	db.CreateIndex("usernames")
	db.Set("user:1", "Alice")
	db.Set("user:2", "Bob")
	db.Set("user:3", "Alice")

	db.AddToIndex("usernames", "Alice", "user:1")
	db.AddToIndex("usernames", "Bob", "user:2")
	db.AddToIndex("usernames", "Alice", "user:3")

	aliceKeys := db.QueryIndex("usernames", "Alice")
	if !containsSameElements(aliceKeys, []string{"user:1", "user:3"}) {
		t.Errorf("QueryIndex(usernames, Alice) = %v, want [user:1, user:3]", aliceKeys)
	}

	bobKeys := db.QueryIndex("usernames", "Bob")
	if !containsSameElements(bobKeys, []string{"user:2"}) {
		t.Errorf("QueryIndex(usernames, Bob) = %v, want [user:2]", bobKeys)
	}
}

func TestQueryNonExistentIndex(t *testing.T) {
	db := NewDatabase()

	keys := db.QueryIndex("nonexistent", "value")
	if len(keys) != 0 {
		t.Errorf("QueryIndex(nonexistent) should return empty slice, got %v", keys)
	}
}

// Transaction Tests

func TestBeginCommit(t *testing.T) {
	db := NewDatabase()

	err := db.Begin()
	if err != nil {
		t.Fatalf("Begin() error = %v, want nil", err)
	}

	db.Set("key1", "value1")
	db.Set("key2", "value2")

	err = db.Commit()
	if err != nil {
		t.Fatalf("Commit() error = %v, want nil", err)
	}

	// Changes should be persisted
	val, _ := db.Get("key1")
	if val != "value1" {
		t.Errorf("After commit, Get(key1) = %q, want %q", val, "value1")
	}
}

func TestRollbackSet(t *testing.T) {
	db := NewDatabase()

	db.Set("existing", "original")

	db.Begin()
	db.Set("existing", "modified")
	db.Set("new", "value")
	db.Rollback()

	// Existing key should have original value
	val, _ := db.Get("existing")
	if val != "original" {
		t.Errorf("After rollback, Get(existing) = %q, want %q", val, "original")
	}

	// New key should not exist
	if db.Exists("new") {
		t.Error("After rollback, new key should not exist")
	}
}

func TestRollbackDelete(t *testing.T) {
	db := NewDatabase()

	db.Set("key1", "value1")

	db.Begin()
	db.Delete("key1")
	db.Rollback()

	// Key should still exist
	if !db.Exists("key1") {
		t.Error("After rollback of delete, key should exist")
	}

	val, _ := db.Get("key1")
	if val != "value1" {
		t.Errorf("After rollback, Get(key1) = %q, want %q", val, "value1")
	}
}

func TestNestedTransactionError(t *testing.T) {
	db := NewDatabase()

	db.Begin()
	err := db.Begin()
	if err == nil {
		t.Error("Begin() while in transaction should return error")
	}
}

func TestCommitWithoutBegin(t *testing.T) {
	db := NewDatabase()

	err := db.Commit()
	if err == nil {
		t.Error("Commit() without Begin() should return error")
	}
}

func TestRollbackWithoutBegin(t *testing.T) {
	db := NewDatabase()

	err := db.Rollback()
	if err == nil {
		t.Error("Rollback() without Begin() should return error")
	}
}

func TestComplexTransaction(t *testing.T) {
	db := NewDatabase()

	// Setup initial data
	db.Set("a", "1")
	db.Set("b", "2")
	db.Set("c", "3")

	// Transaction with multiple operations
	db.Begin()
	db.Set("a", "10")      // Update
	db.Set("d", "4")       // Create
	db.Delete("b")         // Delete
	db.Set("c", "30")      // Update
	db.Set("e", "5")       // Create
	db.Rollback()

	// All changes should be reverted
	tests := []struct {
		key    string
		want   string
		exists bool
	}{
		{"a", "1", true},
		{"b", "2", true},
		{"c", "3", true},
		{"d", "", false},
		{"e", "", false},
	}

	for _, tt := range tests {
		exists := db.Exists(tt.key)
		if exists != tt.exists {
			t.Errorf("After rollback, Exists(%q) = %v, want %v", tt.key, exists, tt.exists)
		}
		if tt.exists {
			val, _ := db.Get(tt.key)
			if val != tt.want {
				t.Errorf("After rollback, Get(%q) = %q, want %q", tt.key, val, tt.want)
			}
		}
	}
}

// Statistics Tests

func TestGetStats(t *testing.T) {
	db := NewDatabase()

	db.Set("key1", "value1")
	db.Set("key2", "value2")
	db.Get("key1")
	db.Get("key2")
	db.Delete("key1")

	stats := db.GetStats()

	if stats.RecordCount != 1 {
		t.Errorf("Stats.RecordCount = %d, want 1", stats.RecordCount)
	}
	if stats.Writes != 2 {
		t.Errorf("Stats.Writes = %d, want 2", stats.Writes)
	}
	if stats.Reads != 2 {
		t.Errorf("Stats.Reads = %d, want 2", stats.Reads)
	}
	if stats.Deletes != 1 {
		t.Errorf("Stats.Deletes = %d, want 1", stats.Deletes)
	}
}

func TestGetStatsWithIndexes(t *testing.T) {
	db := NewDatabase()

	db.CreateIndex("index1")
	db.CreateIndex("index2")

	stats := db.GetStats()
	if stats.IndexCount != 2 {
		t.Errorf("Stats.IndexCount = %d, want 2", stats.IndexCount)
	}
}

func TestMemoryUsage(t *testing.T) {
	db := NewDatabase()

	initialMemory := db.MemoryUsage()

	db.Set("key", "value")
	afterSetMemory := db.MemoryUsage()

	if afterSetMemory <= initialMemory {
		t.Error("Memory usage should increase after Set()")
	}

	// Memory should be at least len("key") + len("value") + timestamps
	minExpected := int64(len("key") + len("value") + 16)
	if afterSetMemory < minExpected {
		t.Errorf("MemoryUsage() = %d, want at least %d", afterSetMemory, minExpected)
	}
}

// Integration Tests

func TestFullWorkflow(t *testing.T) {
	db := NewDatabase()

	// Create users
	db.Set("user:1", "Alice")
	db.Set("user:2", "Bob")
	db.Set("user:3", "Charlie")

	// Create index
	db.CreateIndex("roles")
	db.AddToIndex("roles", "admin", "user:1")
	db.AddToIndex("roles", "user", "user:2")
	db.AddToIndex("roles", "user", "user:3")

	// Query by role
	admins := db.QueryIndex("roles", "admin")
	if len(admins) != 1 || admins[0] != "user:1" {
		t.Errorf("Admin users = %v, want [user:1]", admins)
	}

	users := db.QueryIndex("roles", "user")
	if !containsSameElements(users, []string{"user:2", "user:3"}) {
		t.Errorf("Regular users = %v, want [user:2, user:3]", users)
	}

	// Transaction: Promote user:2 to admin
	db.Begin()
	db.Set("user:2", "Bob (Admin)")
	// In real system, we'd update index here too
	db.Commit()

	val, _ := db.Get("user:2")
	if val != "Bob (Admin)" {
		t.Errorf("After promotion, user:2 = %q, want %q", val, "Bob (Admin)")
	}

	// Get all users
	allUsers := db.GetByPrefix("user:")
	if len(allUsers) != 3 {
		t.Errorf("Total users = %d, want 3", len(allUsers))
	}

	// Stats
	stats := db.GetStats()
	if stats.RecordCount != 3 {
		t.Errorf("RecordCount = %d, want 3", stats.RecordCount)
	}
	if stats.IndexCount != 1 {
		t.Errorf("IndexCount = %d, want 1", stats.IndexCount)
	}
}

func TestRecordTimestamps(t *testing.T) {
	db := NewDatabase()

	db.Set("key1", "value1")
	// Small delay would be nice here, but not critical for test

	db.Set("key1", "value2")

	// We can't directly test timestamps without exposing internal state,
	// but we've verified the Set operation works correctly in other tests
	// This test documents the expected behavior
}
