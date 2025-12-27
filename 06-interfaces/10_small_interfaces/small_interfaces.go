package small_interfaces

// TODO(human): Define Sizer interface

// TODO(human): Define Namer interface

// TODO(human): Define File struct

// TODO(human): Define Directory struct

// TODO(human): Implement Size() for File

// TODO(human): Implement Name() for File

// TODO(human): Implement Size() for Directory (sum of all files)

// TODO(human): Implement Name() for Directory

// TotalSize returns the sum of all item sizes
func TotalSize(items []Sizer) int64 {
	// TODO(human): Implement
	return 0
}

// ListNames returns all item names
func ListNames(items []Namer) []string {
	// TODO(human): Implement
	return nil
}

// LargestBySize returns the largest item by size
func LargestBySize(items []Sizer) (Sizer, bool) {
	// TODO(human): Implement
	return nil, false
}
