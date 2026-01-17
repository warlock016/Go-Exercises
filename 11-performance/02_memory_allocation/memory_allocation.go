package memory_allocation

// User represents a simple user
type User struct {
	ID   int
	Name string
}

// ConcatStrings joins strings using naive concatenation
func ConcatStrings(strs []string) string {
	// TODO(human): Implement with string concatenation
	return ""
}

// ConcatStringsOptimized joins strings efficiently
func ConcatStringsOptimized(strs []string) string {
	// TODO(human): Implement with strings.Builder
	return ""
}

// ProcessItems doubles each item, returning new slice
func ProcessItems(items []int) []int {
	// TODO(human): Implement creating new slice
	return nil
}

// ProcessItemsInPlace doubles each item in the existing slice
func ProcessItemsInPlace(items []int) []int {
	// TODO(human): Implement modifying in place
	return nil
}

// CreateUsers creates user pointers from names
func CreateUsers(names []string) []*User {
	// TODO(human): Implement returning slice of pointers
	return nil
}

// CreateUsersOptimized creates users with fewer allocations
func CreateUsersOptimized(names []string) []User {
	// TODO(human): Implement returning slice of values
	return nil
}
