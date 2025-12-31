package handler_testing

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func setupUserService(t *testing.T) *UserService {

	t.Helper()

	users := map[int]User{
		1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
		2: {ID: 2, Name: "Bob", Email: "bob@example.com"},
		3: {ID: 3, Name: "Charlie", Email: "charlie@example.com"},
	}

	return &UserService{
		users:  users,
		nextID: 4,
	}
}

func TestListUsersHandler(t *testing.T) {

	tests := []struct {
		name       string
		setupUsers map[int]User
		wantStatus int
		wantCount  int
	}{
		{
			name: "returns all users",
			setupUsers: map[int]User{
				1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
				2: {ID: 2, Name: "Bob", Email: "bob@example.com"},
			},
			wantStatus: http.StatusOK,
			wantCount:  2,
		},
		{
			name:       "empty list returns empty array",
			setupUsers: map[int]User{},
			wantStatus: http.StatusOK,
			wantCount:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			userService := UserService{
				users:  tt.setupUsers,
				nextID: 3,
			}

			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			w := httptest.NewRecorder()

			userService.ListUsersHandler(w, req)

			if val := w.Header().Get("Content-Type"); val != "application/json" {
				t.Errorf("\"Content-Type\" header missing, want: \"application/json\"")
			}

			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}

			decoded := []User{}
			err := json.NewDecoder(w.Body).Decode(&decoded)
			if err != nil {
				t.Errorf("failed to decode response: %+v", decoded)
			}

			if w.Header().Get("Content-Type") != "application/json" {
				t.Errorf("incorrect response header, got: %s, want: \"%s\"", w.Header().Get("Content-Type"), "application/json")
			}

			if len(decoded) != tt.wantCount {
				t.Errorf("unexpected user count: %d, want: %d", len(decoded), len(userService.users))
			}

			for _, u := range decoded {
				if u.ID != tt.setupUsers[u.ID].ID || u.Name != tt.setupUsers[u.ID].Name || u.Email != tt.setupUsers[u.ID].Email {
					t.Errorf("invalid user entry")
				}
			}
		})
	}
}

func TestCreateUserHandler(t *testing.T) {

	tests := []struct {
		name     string
		newUser  User
		rawBody  string
		wantCode int
	}{
		{
			name: "valid user",
			newUser: User{
				Name:  "Alice",
				Email: "alice@example.com",
			},
			wantCode: http.StatusCreated,
		},
		{
			name: "empty name",
			newUser: User{
				Name:  "",
				Email: "bob@example.com",
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "empty email",
			newUser: User{
				Name:  "Charlie",
				Email: "",
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "malformed body",
			newUser:  User{ /*invalid user*/ },
			rawBody:  "{invalid",
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := User{}
			var body io.Reader
			var userByte []byte

			if tt.rawBody != "" {
				body = strings.NewReader(tt.rawBody)
			} else {
				userByte, _ = json.Marshal(tt.newUser)
				body = bytes.NewReader(userByte)
			}

			req := httptest.NewRequest(http.MethodPost, "/users", body)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			service := NewUserService()
			service.CreateUserHandler(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("failed to create user, Status: %d, want: %d", w.Code, tt.wantCode)
			}

			if tt.wantCode != http.StatusCreated {
				return
			}

			if w.Header().Get("Content-Type") != "application/json" {
				t.Errorf("incorrect response header, got: %s, want: \"%s\"", w.Header().Get("Content-Type"), "application/json")
			}

			if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
				t.Errorf("failed to decode response: %v", err)
			}

			if result.Name != tt.newUser.Name || result.Email != tt.newUser.Email {
				t.Errorf("created incorrect user: got %+v, want: %+v", result, tt.newUser)
			}
		})
	}
}

func TestGetUserHandler(t *testing.T) {

	tests := []struct {
		name       string
		userId     string
		wantCode   int
		refService *UserService
	}{
		{name: "valid user", userId: "1", wantCode: http.StatusOK, refService: setupUserService(t)},
		{name: "unavailable user", userId: "4", wantCode: http.StatusNotFound, refService: setupUserService(t)},
		{name: "incorrect userId", userId: "abc", wantCode: http.StatusBadRequest, refService: setupUserService(t)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			path := "/users/" + tt.userId

			req := httptest.NewRequest(http.MethodGet, path, nil)
			// req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			tt.refService.GetUserHandler(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("invalid status: got %d, want: %d", w.Code, tt.wantCode)
			}

			if tt.wantCode != http.StatusOK {
				return
			}

			if w.Header().Get("Content-Type") != "application/json" {
				t.Errorf("invalid \"Content-Type\" header: got %s, want: \"%s\"", w.Header().Get("Content-Type"), "application/json")
			}

			resUser := User{}
			err := json.NewDecoder(w.Body).Decode(&resUser)
			if err != nil {
				t.Errorf("failed to decode response: %v", err)
			}

			id, err := strconv.ParseInt(tt.userId, 10, 64)
			if err != nil {
				t.Errorf("failed to convert test UserID: %v", err)
			}
			if resUser.ID != tt.refService.users[int(id)].ID || resUser.Name != tt.refService.users[int(id)].Name || resUser.Email != tt.refService.users[int(id)].Email {
				t.Errorf("response returned incorrect user")
			}
		})
	}
}

// TODO(human): Write comprehensive tests for UserService handlers
//
// Test cases to implement:
//
// 1. TestListUsersHandler
//    - Empty list returns empty array
//    - List with users returns all users
//    - Content-Type is application/json
//
// 2. TestCreateUserHandler
//    - Valid user creation returns 201
//    - Created user has correct data
//    - Missing name returns 400
//    - Missing email returns 400
//    - Malformed JSON returns 400
//    - Content-Type is application/json
//
// 3. TestGetUserHandler
//    - Existing user returns 200 with user data
//    - Non-existent user returns 404
//    - Invalid ID format returns 400
//    - Content-Type is application/json
//
// Use table-driven tests with httptest.NewRequest and httptest.NewRecorder
// Example structure:
//
// func TestCreateUserHandler(t *testing.T) {
//     tests := []struct {
//         name       string
//         body       string
//         wantStatus int
//         wantUser   *User
//     }{
//         // test cases here
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             // test implementation
//         })
//     }
// }
