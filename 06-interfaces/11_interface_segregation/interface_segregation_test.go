package interface_segregation

import (
	"errors"
	"sort"
	"testing"
)

func TestMemoryStoreCreate(t *testing.T) {
	store := &MemoryStore{data: make(map[string]string)}

	err := store.Create("user1", "Alice")
	if err != nil {
		t.Errorf("Create() error = %v", err)
	}

	// Verify created
	val, _ := store.Read("user1")
	if val != "Alice" {
		t.Errorf("After Create, Read() = %q, want %q", val, "Alice")
	}
}

func TestMemoryStoreRead(t *testing.T) {
	store := &MemoryStore{data: map[string]string{"user1": "Alice"}}

	val, err := store.Read("user1")
	if err != nil {
		t.Errorf("Read() error = %v", err)
	}
	if val != "Alice" {
		t.Errorf("Read() = %q, want %q", val, "Alice")
	}

	// Read non-existent
	_, err = store.Read("nonexistent")
	if err == nil {
		t.Error("Read(nonexistent) should return error")
	}
}

func TestMemoryStoreUpdate(t *testing.T) {
	store := &MemoryStore{data: map[string]string{"user1": "Alice"}}

	err := store.Update("user1", "Alice Updated")
	if err != nil {
		t.Errorf("Update() error = %v", err)
	}

	val, _ := store.Read("user1")
	if val != "Alice Updated" {
		t.Errorf("After Update, Read() = %q, want %q", val, "Alice Updated")
	}

	// Update non-existent should error
	err = store.Update("nonexistent", "value")
	if err == nil {
		t.Error("Update(nonexistent) should return error")
	}
}

func TestMemoryStoreDelete(t *testing.T) {
	store := &MemoryStore{data: map[string]string{"user1": "Alice"}}

	err := store.Delete("user1")
	if err != nil {
		t.Errorf("Delete() error = %v", err)
	}

	// Should not exist anymore
	_, err = store.Read("user1")
	if err == nil {
		t.Error("After Delete, Read() should return error")
	}
}

func TestMemoryStoreList(t *testing.T) {
	store := &MemoryStore{data: map[string]string{
		"user1": "Alice",
		"user2": "Bob",
		"user3": "Charlie",
	}}

	list := store.List()
	sort.Strings(list) // Sort for consistent comparison

	want := []string{"user1", "user2", "user3"}
	if len(list) != len(want) {
		t.Errorf("List() length = %d, want %d", len(list), len(want))
	}

	for i, id := range want {
		if list[i] != id {
			t.Errorf("List()[%d] = %q, want %q", i, list[i], id)
		}
	}
}

func TestCopyData(t *testing.T) {
	source := &MemoryStore{data: map[string]string{"key1": "value1"}}
	dest := &MemoryStore{data: make(map[string]string)}

	// Note: CopyData signature expects (from Reader, to Creator)
	// For this test, we need a specific key to copy
	// This assumes CopyData copies a specific key - adjust based on implementation
	err := CopyData(source, dest, "key1")
	if err != nil {
		t.Errorf("CopyData() error = %v", err)
	}

	if _, ok := dest.data["key1"]; !ok {
		t.Errorf("failed to copy %s to destination", "key1")
	}
}

func TestMigrateAll(t *testing.T) {
	source := &MemoryStore{data: map[string]string{
		"user1": "Alice",
		"user2": "Bob",
		"user3": "Charlie",
	}}
	dest := &MemoryStore{data: make(map[string]string)}

	err := MigrateAll(source, source, dest)
	if err != nil {
		t.Errorf("MigrateAll() error = %v", err)
	}

	// Verify all data copied
	for id, expectedVal := range source.data {
		val, err := dest.Read(id)
		if err != nil {
			t.Errorf("After MigrateAll, Read(%q) error = %v", id, err)
		}
		if val != expectedVal {
			t.Errorf("After MigrateAll, Read(%q) = %q, want %q", id, val, expectedVal)
		}
	}
}

func TestInterfaceSegregation(t *testing.T) {
	store := &MemoryStore{data: make(map[string]string)}

	// MemoryStore implements all interfaces
	var _ Creator = store
	var _ Reader = store
	var _ Updater = store
	var _ Deleter = store
	var _ Lister = store

	t.Log("✓ MemoryStore implements all CRUD interfaces")
}

// Helper types for TestFunctionsAcceptMinimalInterfaces
// These demonstrate interface segregation - each implements only what it needs

type OnlyReader struct {
	data map[string]string
}

func (r *OnlyReader) Read(id string) (string, error) {
	val, ok := r.data[id]
	if !ok {
		return "", errors.New("not found")
	}
	return val, nil
}

type OnlyLister struct {
	keys []string
}

func (l *OnlyLister) List() []string {
	return l.keys
}

type OnlyCreator struct {
	data map[string]string
}

func (c *OnlyCreator) Create(id string, value string) error {
	c.data[id] = value
	return nil
}

func TestFunctionsAcceptMinimalInterfaces(t *testing.T) {
	// This test demonstrates that functions accept minimal interfaces
	// These types don't implement full CRUD, but work with segregated functions
	reader := &OnlyReader{data: map[string]string{"key1": "value1"}}
	lister := &OnlyLister{keys: []string{"key1"}}
	creator := &OnlyCreator{data: make(map[string]string)}

	// MigrateAll works with segregated interfaces
	err := MigrateAll(reader, lister, creator)
	if err != nil {
		t.Errorf("MigrateAll with segregated interfaces error = %v", err)
	}

	t.Log("✓ Functions work with types implementing only the needed interfaces")
}
