package data_modeling

// TODO(human): Define Book struct with ID, Title, Author, ISBN, Available fields

// TODO(human): Define Member struct with ID, Name, CheckedOutBooks fields

// TODO(human): Define Library struct with Books and Members maps

// NewLibrary creates and returns a new empty library
func NewLibrary() *Library {
	// TODO(human): Create and initialize Library
	return nil
}

// AddBook adds a book to the library and returns the book's ID
// Book IDs should start at 1 and increment
func AddBook(lib *Library, title, author, isbn string) int {
	// TODO(human): Generate ID, create book, add to library
	return 0
}

// RemoveBook removes a book from the library by ID
// Returns true if removed, false if not found
func RemoveBook(lib *Library, bookID int) bool {
	// TODO(human): Check existence and delete if found
	return false
}

// AddMember adds a member to the library and returns the member's ID
// Member IDs should start at 1 and increment
func AddMember(lib *Library, name string) int {
	// TODO(human): Generate ID, create member, add to library
	return 0
}

// CheckoutBook checks out a book to a member
// Returns true if successful, false if:
// - Book doesn't exist or isn't available
// - Member doesn't exist
func CheckoutBook(lib *Library, memberID, bookID int) bool {
	// TODO(human): Validate book and member, update both records
	return false
}

// ReturnBook processes a book return from a member
// Returns true if successful, false if member doesn't have this book checked out
func ReturnBook(lib *Library, memberID, bookID int) bool {
	// TODO(human): Find book in member's list, remove it, update records
	return false
}

// FindBooksByAuthor returns a slice of all books by the given author
func FindBooksByAuthor(lib *Library, author string) []Book {
	// TODO(human): Filter books by author
	return nil
}

// GetMemberBooks returns a slice of all books currently checked out by a member
// Returns empty slice if member doesn't exist or has no books
func GetMemberBooks(lib *Library, memberID int) []Book {
	// TODO(human): Lookup member's checked out books
	return nil
}

// GetAvailableBooks returns a slice of all available books
func GetAvailableBooks(lib *Library) []Book {
	// TODO(human): Filter books by availability
	return nil
}
