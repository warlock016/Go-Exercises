# Exercise 11: Dependency Injection

**Learning Goal:** Master dependency injection using functions for testable code

---

## 📝 Problem Description

Dependency injection passes dependencies as parameters instead of creating them internally. This makes code:
- Testable (inject mocks)
- Flexible (swap implementations)
- Decoupled (no hard dependencies)

You'll implement a user service with injectable dependencies.

---

## 🎯 Function Signatures

```go
type UserStore interface {
    GetUser(id int) (string, error)
    SaveUser(id int, name string) error
}

type Logger interface {
    Log(message string)
}

type UserService struct {
    store  UserStore
    logger Logger
}

func NewUserService(store UserStore, logger Logger) *UserService

func (s *UserService) GetUser(id int) (string, error)

func (s *UserService) CreateUser(id int, name string) error
```

---

## 📖 Examples

```go
// Production code
store := &DatabaseStore{}
logger := &FileLogger{}
service := NewUserService(store, logger)
user, _ := service.GetUser(1)

// Test code with mocks
mockStore := &MockStore{...}
mockLogger := &MockLogger{...}
testService := NewUserService(mockStore, mockLogger)
// Test with full control over dependencies
```

---

## 📋 Instructions

1. Define `UserStore` and `Logger` interfaces
2. Implement `UserService` struct
3. Implement `NewUserService` constructor
4. Implement `GetUser` that uses store and logger
5. Implement `CreateUser` that validates and saves
6. Create mock implementations for testing
7. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Complete Solution</summary>

```go
package dependency_injection

import "errors"

type UserStore interface {
	GetUser(id int) (string, error)
	SaveUser(id int, name string) error
}

type Logger interface {
	Log(message string)
}

type UserService struct {
	store  UserStore
	logger Logger
}

func NewUserService(store UserStore, logger Logger) *UserService {
	return &UserService{
		store:  store,
		logger: logger,
	}
}

func (s *UserService) GetUser(id int) (string, error) {
	s.logger.Log("Getting user")
	return s.store.GetUser(id)
}

func (s *UserService) CreateUser(id int, name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	s.logger.Log("Creating user")
	return s.store.SaveUser(id, name)
}
```

</details>

---

## 🤔 Think About

1. Why use interfaces for dependencies?
2. How does this improve testability?
3. When would you inject functions vs interfaces?
4. What's the trade-off between flexibility and simplicity?

---

## 🎓 What This Teaches

- **Dependency injection**: Passing dependencies rather than creating them
- **Interfaces**: Defining contracts for dependencies
- **Testability**: Injecting mocks for testing
- **Decoupling**: Reducing hardcoded dependencies
- **SOLID principles**: Dependency inversion principle

---

**Tier:** 3 - Integration
**Estimated Time:** 40-50 minutes
