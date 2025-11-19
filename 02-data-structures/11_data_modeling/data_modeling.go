package data_modeling

import "slices"

// TODO(human): Define Book struct with ID, Title, Author, ISBN, Available fields
type Book struct {
	ID        int
	Title     string
	Author    string
	ISBN      string
	Available bool
}

// TODO(human): Define Member struct with ID, Name, CheckedOutBooks fields
type Member struct {
	ID              int
	Name            string
	CheckedOutBooks []int
}

// TODO(human): Define Library struct with Books and Members maps
type Library struct {
	Books        map[int]Book
	Members      map[int]Member
	nextBookId   int
	nextMemberId int
}

// NewLibrary creates and returns a new empty library
func NewLibrary() *Library {
	// TODO(human): Create and initialize Library
	return &Library{
		Books:        make(map[int]Book),
		Members:      make(map[int]Member),
		nextBookId:   1,
		nextMemberId: 1,
	}
}

// AddBook adds a book to the library and returns the book's ID
// Book IDs should start at 1 and increment
func AddBook(lib *Library, title, author, isbn string) int {
	// TODO(human): Generate ID, create book, add to library

	id := lib.nextBookId
	lib.nextBookId++

	lib.Books[id] = Book{
		ID:        id,
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: true,
	}

	return id
}

// RemoveBook removes a book from the library by ID
// Returns true if removed, false if not found
func RemoveBook(lib *Library, bookID int) bool {
	// TODO(human): Check existence and delete if found

	if _, exists := lib.Books[bookID]; exists {
		delete(lib.Books, bookID)
		return true
	}

	return false
}

// AddMember adds a member to the library and returns the member's ID
// Member IDs should start at 1 and increment
func AddMember(lib *Library, name string) int {
	// TODO(human): Generate ID, create member, add to library
	id := lib.nextMemberId
	lib.nextMemberId++

	lib.Members[id] = Member{
		Name:            name,
		CheckedOutBooks: make([]int, 0),
	}

	return id
}

// CheckoutBook checks out a book to a member
// Returns true if successful, false if:
// - Book doesn't exist or isn't available
// - Member doesn't exist
func CheckoutBook(lib *Library, memberID, bookID int) bool {
	// TODO(human): Validate book and member, update both records
	member, exists := lib.Members[memberID]
	book, expected := lib.Books[bookID]

	if exists && expected && book.Available {
		book.Available = false
		member.CheckedOutBooks = append(member.CheckedOutBooks, bookID)
		lib.Members[memberID] = member
		lib.Books[bookID] = book
		return true
	}

	return false
}

// ReturnBook processes a book return from a member
// Returns true if successful, false if member doesn't have this book checked out
func ReturnBook(lib *Library, memberID, bookID int) bool {
	// TODO(human): Find book in member's list, remove it, update records

	member, exists := lib.Members[memberID]
	book, expected := lib.Books[bookID]

	if exists && expected {
		checkedOutBooks := member.CheckedOutBooks
		borrowedBooks := []int{}

		if !slices.Contains(checkedOutBooks, bookID) {
			return false
		}

		for i := range checkedOutBooks {
			if checkedOutBooks[i] == bookID {
				borrowedBooks = append(borrowedBooks, checkedOutBooks[:i]...)
				borrowedBooks = append(borrowedBooks, checkedOutBooks[i+1:]...)
				member.CheckedOutBooks = borrowedBooks // borrowed books updated
				book.Available = true                  // book returned
				break
			}
		}

		lib.Members[memberID] = member
		lib.Books[bookID] = book
		return true
	}

	return false
}

// FindBooksByAuthor returns a slice of all books by the given author
func FindBooksByAuthor(lib *Library, author string) []Book {
	// TODO(human): Filter books by author

	result := []Book{}

	for _, v := range lib.Books {
		if v.Author == author {
			result = append(result, v)
		}
	}

	return result
}

// GetMemberBooks returns a slice of all books currently checked out by a member
// Returns empty slice if member doesn't exist or has no books
func GetMemberBooks(lib *Library, memberID int) []Book {
	// TODO(human): Lookup member's checked out books

	result := []Book{}

	if _, ok := lib.Members[memberID]; !ok { // member does not exist
		return result
	}

	books := lib.Members[memberID].CheckedOutBooks

	if len(books) == 0 {
		return result
	}

	for _, v := range books {
		if _, ok := lib.Books[v]; ok {
			result = append(result, lib.Books[v])
		}
	}

	return result
}

// GetAvailableBooks returns a slice of all available books
func GetAvailableBooks(lib *Library) []Book {
	// TODO(human): Filter books by availability

	result := []Book{}

	for _, v := range lib.Books {
		if v.Available {
			result = append(result, v)
		}
	}

	return result
}
