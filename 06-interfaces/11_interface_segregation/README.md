# Exercise 11: Interface Segregation

**Concept:** Interface Segregation Principle - accept minimal interfaces
**Difficulty:** Medium-Hard
**Estimated Time:** 60 minutes

## Learning Goal

Learn the Interface Segregation Principle: functions should accept the smallest interface they need, not large kitchen-sink interfaces.

## The Problem

Functions that accept large interfaces are inflexible:

```go
// Bad - requires all operations even if only need one
func CopyData(db DataStore) { ... }  // Needs only Read(), but DataStore has 10 methods

// Good - accept only what you need
func CopyData(r Reader) { ... }      // Only needs Read()
```

## Your Task

Design a data store with separate small interfaces for different operations.

### Interfaces

1. **Creator** - `Create(id string, value string) error`
2. **Reader** - `Read(id string) (string, error)`
3. **Updater** - `Update(id string, value string) error`
4. **Deleter** - `Delete(id string) error`
5. **Lister** - `List() []string`

### Type

**MemoryStore** - implements all five interfaces

### Functions

1. **CopyData(from Reader, to Creator) error** - copies all data
2. **MigrateAll(from Reader, fromList Lister, to Creator) error** - migrates all records

## Function Signatures

```go
type Creator interface {
    Create(id string, value string) error
}

type Reader interface {
    Read(id string) (string, error)
}

type Updater interface {
    Update(id string, value string) error
}

type Deleter interface {
    Delete(id string) error
}

type Lister interface {
    List() []string
}

type MemoryStore struct {
    data map[string]string
}

func (m *MemoryStore) Create(id string, value string) error
func (m *MemoryStore) Read(id string) (string, error)
func (m *MemoryStore) Update(id string, value string) error
func (m *MemoryStore) Delete(id string) error
func (m *MemoryStore) List() []string

func CopyData(from Reader, to Creator) error
func MigrateAll(from Reader, fromList Lister, to Creator) error
```

## What This Teaches

- Interface Segregation Principle (ISP)
- Accept minimal interfaces
- SOLID principles in Go
- Flexible function design

---

**Next up:** Exercise 12 - Polymorphic Collection
