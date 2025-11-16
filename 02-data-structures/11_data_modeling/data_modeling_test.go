package data_modeling

import (
	"testing"
)

func TestNewLibrary(t *testing.T) {
	lib := NewLibrary()

	if lib == nil {
		t.Fatal("NewLibrary() returned nil")
	}

	if lib.Books == nil {
		t.Error("NewLibrary() Books map is nil")
	}

	if lib.Members == nil {
		t.Error("NewLibrary() Members map is nil")
	}

	if len(lib.Books) != 0 {
		t.Errorf("NewLibrary() Books length = %d, want 0", len(lib.Books))
	}

	if len(lib.Members) != 0 {
		t.Errorf("NewLibrary() Members length = %d, want 0", len(lib.Members))
	}
}

func TestAddBook(t *testing.T) {
	lib := NewLibrary()

	id1 := AddBook(lib, "1984", "George Orwell", "978-0451524935")
	if id1 != 1 {
		t.Errorf("AddBook() first ID = %d, want 1", id1)
	}

	id2 := AddBook(lib, "Animal Farm", "George Orwell", "978-0451526342")
	if id2 != 2 {
		t.Errorf("AddBook() second ID = %d, want 2", id2)
	}

	if len(lib.Books) != 2 {
		t.Errorf("AddBook() library has %d books, want 2", len(lib.Books))
	}

	book1, exists := lib.Books[id1]
	if !exists {
		t.Fatal("AddBook() book 1 not found in library")
	}

	if book1.Title != "1984" {
		t.Errorf("AddBook() book1.Title = %q, want %q", book1.Title, "1984")
	}

	if book1.Author != "George Orwell" {
		t.Errorf("AddBook() book1.Author = %q, want %q", book1.Author, "George Orwell")
	}

	if book1.ISBN != "978-0451524935" {
		t.Errorf("AddBook() book1.ISBN = %q, want %q", book1.ISBN, "978-0451524935")
	}

	if !book1.Available {
		t.Error("AddBook() book1.Available = false, want true")
	}
}

func TestRemoveBook(t *testing.T) {
	lib := NewLibrary()
	id1 := AddBook(lib, "1984", "George Orwell", "978-0451524935")
	id2 := AddBook(lib, "Animal Farm", "George Orwell", "978-0451526342")

	removed := RemoveBook(lib, id1)
	if !removed {
		t.Error("RemoveBook() returned false, want true")
	}

	if len(lib.Books) != 1 {
		t.Errorf("RemoveBook() library has %d books, want 1", len(lib.Books))
	}

	_, exists := lib.Books[id1]
	if exists {
		t.Error("RemoveBook() book still exists after removal")
	}

	_, exists = lib.Books[id2]
	if !exists {
		t.Error("RemoveBook() removed wrong book")
	}

	removed = RemoveBook(lib, 999)
	if removed {
		t.Error("RemoveBook() non-existent book returned true, want false")
	}
}

func TestAddMember(t *testing.T) {
	lib := NewLibrary()

	id1 := AddMember(lib, "Alice")
	if id1 != 1 {
		t.Errorf("AddMember() first ID = %d, want 1", id1)
	}

	id2 := AddMember(lib, "Bob")
	if id2 != 2 {
		t.Errorf("AddMember() second ID = %d, want 2", id2)
	}

	if len(lib.Members) != 2 {
		t.Errorf("AddMember() library has %d members, want 2", len(lib.Members))
	}

	member1, exists := lib.Members[id1]
	if !exists {
		t.Fatal("AddMember() member 1 not found in library")
	}

	if member1.Name != "Alice" {
		t.Errorf("AddMember() member1.Name = %q, want %q", member1.Name, "Alice")
	}

	if member1.CheckedOutBooks == nil {
		t.Error("AddMember() member1.CheckedOutBooks is nil, want empty slice")
	}

	if len(member1.CheckedOutBooks) != 0 {
		t.Errorf("AddMember() member1.CheckedOutBooks length = %d, want 0", len(member1.CheckedOutBooks))
	}
}

func TestCheckoutBook(t *testing.T) {
	lib := NewLibrary()
	bookID := AddBook(lib, "1984", "George Orwell", "978-0451524935")
	memberID := AddMember(lib, "Alice")

	// Successful checkout
	success := CheckoutBook(lib, memberID, bookID)
	if !success {
		t.Error("CheckoutBook() valid checkout returned false, want true")
	}

	book, _ := lib.Books[bookID]
	if book.Available {
		t.Error("CheckoutBook() book.Available = true, want false")
	}

	member, _ := lib.Members[memberID]
	if len(member.CheckedOutBooks) != 1 {
		t.Fatalf("CheckoutBook() member has %d books, want 1", len(member.CheckedOutBooks))
	}

	if member.CheckedOutBooks[0] != bookID {
		t.Errorf("CheckoutBook() member.CheckedOutBooks[0] = %d, want %d", member.CheckedOutBooks[0], bookID)
	}

	// Try to checkout already checked out book
	success = CheckoutBook(lib, memberID, bookID)
	if success {
		t.Error("CheckoutBook() unavailable book returned true, want false")
	}

	// Try to checkout non-existent book
	success = CheckoutBook(lib, memberID, 999)
	if success {
		t.Error("CheckoutBook() non-existent book returned true, want false")
	}

	// Try to checkout with non-existent member
	success = CheckoutBook(lib, 999, bookID)
	if success {
		t.Error("CheckoutBook() non-existent member returned true, want false")
	}
}

func TestReturnBook(t *testing.T) {
	lib := NewLibrary()
	bookID := AddBook(lib, "1984", "George Orwell", "978-0451524935")
	memberID := AddMember(lib, "Alice")

	// Checkout first
	CheckoutBook(lib, memberID, bookID)

	// Return book
	success := ReturnBook(lib, memberID, bookID)
	if !success {
		t.Error("ReturnBook() valid return returned false, want true")
	}

	book, _ := lib.Books[bookID]
	if !book.Available {
		t.Error("ReturnBook() book.Available = false, want true")
	}

	member, _ := lib.Members[memberID]
	if len(member.CheckedOutBooks) != 0 {
		t.Errorf("ReturnBook() member has %d books, want 0", len(member.CheckedOutBooks))
	}

	// Try to return book not checked out
	success = ReturnBook(lib, memberID, bookID)
	if success {
		t.Error("ReturnBook() book not checked out returned true, want false")
	}

	// Try to return with non-existent member
	success = ReturnBook(lib, 999, bookID)
	if success {
		t.Error("ReturnBook() non-existent member returned true, want false")
	}
}

func TestMultipleCheckouts(t *testing.T) {
	lib := NewLibrary()
	book1 := AddBook(lib, "1984", "George Orwell", "978-0451524935")
	book2 := AddBook(lib, "Animal Farm", "George Orwell", "978-0451526342")
	book3 := AddBook(lib, "To Kill a Mockingbird", "Harper Lee", "978-0061120084")
	memberID := AddMember(lib, "Alice")

	CheckoutBook(lib, memberID, book1)
	CheckoutBook(lib, memberID, book2)
	CheckoutBook(lib, memberID, book3)

	member, _ := lib.Members[memberID]
	if len(member.CheckedOutBooks) != 3 {
		t.Errorf("Multiple checkouts: member has %d books, want 3", len(member.CheckedOutBooks))
	}

	// Return middle book
	ReturnBook(lib, memberID, book2)

	member, _ = lib.Members[memberID]
	if len(member.CheckedOutBooks) != 2 {
		t.Errorf("After return: member has %d books, want 2", len(member.CheckedOutBooks))
	}

	// Check correct books remain
	hasBook1 := false
	hasBook3 := false
	for _, id := range member.CheckedOutBooks {
		if id == book1 {
			hasBook1 = true
		}
		if id == book3 {
			hasBook3 = true
		}
		if id == book2 {
			t.Error("After return: member still has book2")
		}
	}

	if !hasBook1 || !hasBook3 {
		t.Error("After return: member missing book1 or book3")
	}
}

func TestFindBooksByAuthor(t *testing.T) {
	lib := NewLibrary()
	AddBook(lib, "1984", "George Orwell", "978-0451524935")
	AddBook(lib, "Animal Farm", "George Orwell", "978-0451526342")
	AddBook(lib, "To Kill a Mockingbird", "Harper Lee", "978-0061120084")

	tests := []struct {
		name       string
		author     string
		wantCount  int
		wantTitles []string
	}{
		{"Orwell books", "George Orwell", 2, []string{"1984", "Animal Farm"}},
		{"Lee books", "Harper Lee", 1, []string{"To Kill a Mockingbird"}},
		{"Non-existent author", "J.K. Rowling", 0, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			books := FindBooksByAuthor(lib, tt.author)

			if len(books) != tt.wantCount {
				t.Errorf("FindBooksByAuthor(%q) returned %d books, want %d", tt.author, len(books), tt.wantCount)
			}

			for _, wantTitle := range tt.wantTitles {
				found := false
				for _, book := range books {
					if book.Title == wantTitle {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("FindBooksByAuthor(%q) missing book %q", tt.author, wantTitle)
				}
			}
		})
	}
}

func TestGetMemberBooks(t *testing.T) {
	lib := NewLibrary()
	book1 := AddBook(lib, "1984", "George Orwell", "978-0451524935")
	book2 := AddBook(lib, "Animal Farm", "George Orwell", "978-0451526342")
	memberID := AddMember(lib, "Alice")

	// No books checked out
	books := GetMemberBooks(lib, memberID)
	if len(books) != 0 {
		t.Errorf("GetMemberBooks() with no checkouts = %d books, want 0", len(books))
	}

	// Checkout books
	CheckoutBook(lib, memberID, book1)
	CheckoutBook(lib, memberID, book2)

	books = GetMemberBooks(lib, memberID)
	if len(books) != 2 {
		t.Errorf("GetMemberBooks() = %d books, want 2", len(books))
	}

	// Check titles
	titles := make(map[string]bool)
	for _, book := range books {
		titles[book.Title] = true
	}

	if !titles["1984"] || !titles["Animal Farm"] {
		t.Error("GetMemberBooks() missing expected books")
	}

	// Non-existent member
	books = GetMemberBooks(lib, 999)
	if len(books) != 0 {
		t.Errorf("GetMemberBooks() non-existent member = %d books, want 0", len(books))
	}
}

func TestGetAvailableBooks(t *testing.T) {
	lib := NewLibrary()
	book1 := AddBook(lib, "1984", "George Orwell", "978-0451524935")
	_ = AddBook(lib, "Animal Farm", "George Orwell", "978-0451526342")
	_ = AddBook(lib, "To Kill a Mockingbird", "Harper Lee", "978-0061120084")
	memberID := AddMember(lib, "Alice")

	// All books available
	books := GetAvailableBooks(lib)
	if len(books) != 3 {
		t.Errorf("GetAvailableBooks() initially = %d books, want 3", len(books))
	}

	// Checkout one book
	CheckoutBook(lib, memberID, book1)

	books = GetAvailableBooks(lib)
	if len(books) != 2 {
		t.Errorf("GetAvailableBooks() after checkout = %d books, want 2", len(books))
	}

	for _, book := range books {
		if book.ID == book1 {
			t.Error("GetAvailableBooks() includes checked out book")
		}
	}

	// Return book
	ReturnBook(lib, memberID, book1)

	books = GetAvailableBooks(lib)
	if len(books) != 3 {
		t.Errorf("GetAvailableBooks() after return = %d books, want 3", len(books))
	}
}
