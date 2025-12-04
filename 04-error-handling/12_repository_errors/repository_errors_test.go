package repository_errors

import (
	"testing"
)

func TestUserRepository_FindByID(t *testing.T) {
	repo := NewUserRepository()
	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	repo.Create(user)

	found, err := repo.FindByID(1)
	if err != nil {
		t.Errorf("FindByID() error = %v", err)
	}
	if found.Name != "Alice" {
		t.Errorf("FindByID() Name = %q, want %q", found.Name, "Alice")
	}

	_, err = repo.FindByID(999)
	if err == nil {
		t.Error("FindByID() with invalid ID should return error")
	}
}

func TestUserRepository_Create(t *testing.T) {
	repo := NewUserRepository()
	user := User{ID: 1, Name: "Bob", Email: "bob@example.com"}

	err := repo.Create(user)
	if err != nil {
		t.Errorf("Create() error = %v", err)
	}

	err = repo.Create(user)
	if err == nil {
		t.Error("Create() duplicate should return error")
	}
}

func TestUserRepository_Update(t *testing.T) {
	repo := NewUserRepository()
	user := User{ID: 1, Name: "Charlie", Email: "charlie@example.com"}
	repo.Create(user)

	user.Name = "Charles"
	err := repo.Update(user)
	if err != nil {
		t.Errorf("Update() error = %v", err)
	}

	err = repo.Update(User{ID: 999})
	if err == nil {
		t.Error("Update() non-existent user should return error")
	}
}

func TestUserRepository_Delete(t *testing.T) {
	repo := NewUserRepository()
	user := User{ID: 1, Name: "Dave", Email: "dave@example.com"}
	repo.Create(user)

	err := repo.Delete(1)
	if err != nil {
		t.Errorf("Delete() error = %v", err)
	}

	err = repo.Delete(1)
	if err == nil {
		t.Error("Delete() non-existent user should return error")
	}
}
