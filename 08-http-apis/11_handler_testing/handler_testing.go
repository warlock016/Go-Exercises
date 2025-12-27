package handler_testing

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// User represents a user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserService manages users
type UserService struct {
	users  map[int]User
	nextID int
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]User),
		nextID: 1,
	}
}

// ListUsersHandler returns all users
func (s *UserService) ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUserHandler creates a new user
func (s *UserService) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	if input.Email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	user := User{
		ID:    s.nextID,
		Name:  input.Name,
		Email: input.Email,
	}
	s.users[s.nextID] = user
	s.nextID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUserHandler returns a single user by ID
func (s *UserService) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from path /users/{id}
	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 {
		http.Error(w, "user ID required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	user, exists := s.users[id]
	if !exists {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
