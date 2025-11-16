# Exercise 11: Data Modeling

## 🎯 Learning Goal
Design a complete real-world system using structs, maps, and slices. Learn to model entities, relationships, and operations for a library management system.

## 📝 Problem Description

Data modeling is the process of designing how your program represents real-world entities and their relationships. Good data modeling:

- Makes code intuitive and easy to understand
- Enables efficient operations (O(1) lookups vs O(n) searches)
- Prevents invalid states through proper structure
- Scales as requirements grow

In this exercise, you'll build a **Library System** that manages books, members, and checkout operations. This teaches you how to:

- Define structs for entities (Book, Member, Library)
- Choose appropriate data structures (map for O(1) lookup, slice for ordering)
- Implement business logic (checkout, return, search)
- Maintain data consistency (availability tracking, member limits)

This mirrors real-world backend systems where you design databases, APIs, and domain models.

## 🔧 Type Definitions and Function Signatures

Define these types in `data_modeling.go`:

```go
// Book represents a book in the library
type Book struct {
	ID        int
	Title     string
	Author    string
	ISBN      string
	Available bool
}

// Member represents a library member
type Member struct {
	ID              int
	Name            string
	CheckedOutBooks []int // Book IDs
}

// Library represents the entire library system
type Library struct {
	Books   map[int]Book      // Book ID -> Book
	Members map[int]Member    // Member ID -> Member
}
```

Implement these functions:

```go
// NewLibrary creates and returns a new empty library
func NewLibrary() *Library

// AddBook adds a book to the library and returns the book's ID
// Book IDs should start at 1 and increment
func AddBook(lib *Library, title, author, isbn string) int

// RemoveBook removes a book from the library by ID
// Returns true if removed, false if not found
func RemoveBook(lib *Library, bookID int) bool

// AddMember adds a member to the library and returns the member's ID
// Member IDs should start at 1 and increment
func AddMember(lib *Library, name string) int

// CheckoutBook checks out a book to a member
// Returns true if successful, false if:
// - Book doesn't exist or isn't available
// - Member doesn't exist
func CheckoutBook(lib *Library, memberID, bookID int) bool

// ReturnBook processes a book return from a member
// Returns true if successful, false if:
// - Member doesn't have this book checked out
func ReturnBook(lib *Library, memberID, bookID int) bool

// FindBooksByAuthor returns a slice of all books by the given author
func FindBooksByAuthor(lib *Library, author string) []Book

// GetMemberBooks returns a slice of all books currently checked out by a member
// Returns empty slice if member doesn't exist or has no books
func GetMemberBooks(lib *Library, memberID int) []Book

// GetAvailableBooks returns a slice of all available books
func GetAvailableBooks(lib *Library) []Book
```

## 💡 Examples

```go
// Create library
lib := NewLibrary()

// Add books
id1 := AddBook(lib, "1984", "George Orwell", "978-0451524935")
id2 := AddBook(lib, "Animal Farm", "George Orwell", "978-0451526342")
id3 := AddBook(lib, "To Kill a Mockingbird", "Harper Lee", "978-0061120084")

// Add members
member1 := AddMember(lib, "Alice")
member2 := AddMember(lib, "Bob")

// Checkout books
success := CheckoutBook(lib, member1, id1)  // true
success = CheckoutBook(lib, member1, id1)   // false (already checked out)

// Get member's books
books := GetMemberBooks(lib, member1)  // [Book{ID:1, Title:"1984", ...}]

// Find by author
orwell := FindBooksByAuthor(lib, "George Orwell")  // [Book{ID:1, ...}, Book{ID:2, ...}]

// Return book
success = ReturnBook(lib, member1, id1)  // true

// Get available books
available := GetAvailableBooks(lib)  // Now includes id1 again
```

## 📋 Instructions

1. **NewLibrary:** Return `&Library{Books: make(map[int]Book), Members: make(map[int]Member)}`
2. **AddBook:** Generate ID (hint: use `len(lib.Books) + 1`), create Book struct, add to map
3. **RemoveBook:** Check if exists with comma-ok idiom, use `delete(lib.Books, bookID)`
4. **AddMember:** Similar to AddBook but for Members
5. **CheckoutBook:** Validate book exists and is available, validate member exists, update book's Available, append to member's CheckedOutBooks
6. **ReturnBook:** Find book in member's slice, remove it, set book's Available to true
7. **FindBooksByAuthor:** Loop through Books map, collect matching books into slice
8. **GetMemberBooks:** Get member's CheckedOutBooks IDs, look up each book in Books map
9. **GetAvailableBooks:** Loop through Books, collect where Available == true

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected test count: ~40-45 tests across all functions

## 🤔 Think About

1. **Why use `map[int]Book` instead of `[]Book`?**
   - O(1) lookup by ID vs O(n) search
   - No need to maintain order for most operations

2. **Why store Book IDs in Member.CheckedOutBooks instead of full Book structs?**
   - Avoids data duplication
   - Single source of truth (Books map)
   - Prevents inconsistency if book details change

3. **What happens if you modify a Book after getting it from the map?**
   - Maps store values, not pointers (unless you use `map[int]*Book`)
   - Modifications to retrieved Book don't affect the map
   - Must reassign to map to persist changes: `lib.Books[id] = modifiedBook`

4. **How would you prevent members from checking out too many books?**
   - Add MaxBooks constant, check `len(member.CheckedOutBooks) < MaxBooks` before checkout

## 💡 Hints

<details>
<summary>Hint 1: Generating IDs</summary>

Simple ID generation:

```go
// For books
newID := len(lib.Books) + 1

// Create and add book
book := Book{
    ID:        newID,
    Title:     title,
    Author:    author,
    ISBN:      isbn,
    Available: true,
}
lib.Books[newID] = book
return newID
```

**Note:** In production, use UUID or database auto-increment for IDs.
</details>

<details>
<summary>Hint 2: Checking out books</summary>

Checkout requires multiple checks and updates:

```go
// 1. Check if book exists and is available
book, exists := lib.Books[bookID]
if !exists || !book.Available {
    return false
}

// 2. Check if member exists
member, exists := lib.Members[memberID]
if !exists {
    return false
}

// 3. Update book availability
book.Available = false
lib.Books[bookID] = book  // Important: reassign to map!

// 4. Add to member's checked out books
member.CheckedOutBooks = append(member.CheckedOutBooks, bookID)
lib.Members[memberID] = member  // Important: reassign to map!

return true
```
</details>

<details>
<summary>Hint 3: Returning books</summary>

Return involves finding and removing from slice:

```go
member, exists := lib.Members[memberID]
if !exists {
    return false
}

// Find and remove book from member's slice
found := false
for i, id := range member.CheckedOutBooks {
    if id == bookID {
        // Remove by slicing
        member.CheckedOutBooks = append(member.CheckedOutBooks[:i], member.CheckedOutBooks[i+1:]...)
        found = true
        break
    }
}

if !found {
    return false
}

// Update book availability
book := lib.Books[bookID]
book.Available = true
lib.Books[bookID] = book

// Update member
lib.Members[memberID] = member
return true
```
</details>

<details>
<summary>Hint 4: Filtering and searching</summary>

Pattern for collecting matching items:

```go
func FindBooksByAuthor(lib *Library, author string) []Book {
    var result []Book
    for _, book := range lib.Books {
        if book.Author == author {
            result = append(result, book)
        }
    }
    return result
}
```
</details>

<details>
<summary>Full Solution</summary>

See solution file for complete implementation. Try implementing yourself first!

Key patterns:
- Use maps for O(1) entity lookup
- Store IDs to reference entities (avoid duplication)
- Always reassign structs to maps after modification
- Validate all inputs (existence checks, availability checks)
- Use slices for filtering/searching results
</details>

## 🎓 What This Teaches

- **Entity design** - Modeling real-world objects with appropriate fields
- **Relationship modeling** - Using IDs to reference related entities
- **Data structure selection** - Maps for lookup, slices for lists
- **Business logic implementation** - Checkout/return with validation
- **State management** - Tracking availability, member books
- **Search and filter patterns** - Finding entities by criteria
- **Map value semantics** - Understanding that maps store values, requiring reassignment
- **Real-world system design** - Patterns used in backend services, databases, APIs

---

**Next Exercise:** `12_custom_types` - Define custom types with methods (preview of Module 03)
