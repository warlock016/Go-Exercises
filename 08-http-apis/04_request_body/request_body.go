package request_body

import "net/http"

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
}
