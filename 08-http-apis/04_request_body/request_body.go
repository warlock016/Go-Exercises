package request_body

import (
	"encoding/json"
	"net/http"
)

// CreateUserRequest represents the input for creating a user
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUserResponse represents the output after creating a user
type CreateUserResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// CreateUserHandler accepts JSON body and creates a user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	w.Header().Set("Content-Type", "application/json")

	req := CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid JSON"))
		return
	}

	if req.Email == "" && req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("name and email are required"))
		return
	} else if req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("name is required"))
		return
	} else if req.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("email is required"))
		return
	}

	resp := CreateUserResponse{
		ID:      42,
		Message: "User created successfully",
		Name:    req.Name,
		Email:   req.Email,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
