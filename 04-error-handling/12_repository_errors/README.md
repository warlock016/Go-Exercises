# Exercise 12: Repository Errors

## 🎯 Learning Goal
Learn data access layer error handling patterns: wrapping database errors, mapping to domain errors, and handling common database scenarios.

## 📝 Problem Description

The repository pattern requires translating low-level database errors into domain-level errors that business logic can understand.

## 🔧 Function Signatures

```go
// User represents a domain user
type User struct {
    ID    int
    Name  string
    Email string
}

// UserRepository handles user data access
type UserRepository struct {
    users map[int]User
}

// FindByID returns user by ID or ErrNotFound
func (r *UserRepository) FindByID(id int) (User, error)

// Create creates a user or returns ErrAlreadyExists
func (r *UserRepository) Create(user User) error

// Update updates a user or returns ErrNotFound
func (r *UserRepository) Update(user User) error

// Delete deletes a user or returns ErrNotFound
func (r *UserRepository) Delete(id int) error
```

## 🎓 What This Teaches

- **Repository pattern** - Data access abstraction
- **Error translation** - Database errors to domain errors
- **CRUD error handling** - Standard data operation errors

---

**Next Exercise:** `13_error_middleware` - HTTP panic recovery middleware
