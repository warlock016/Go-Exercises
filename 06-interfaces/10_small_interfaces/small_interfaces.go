package small_interfaces

// TODO(human): Define Sizer interface
type Sizer interface {
	Size() int64
}

// TODO(human): Define Namer interface
type Namer interface {
	Name() string
}

type SizedNamed interface {
	Namer
	Sizer
}

// TODO(human): Define File struct
type File struct {
	FileName string
	FileSize int64 // int64??
}

// TODO(human): Define Directory struct
type Directory struct {
	DirName string
	Files   []File
}

// TODO(human): Implement Size() for File
func (f File) Size() int64 {
	return f.FileSize
}

// TODO(human): Implement Name() for File
func (f File) Name() string {
	return f.FileName
}

// TODO(human): Implement Size() for Directory (sum of all files)
func (d Directory) Size() int64 {

	var total int64
	for _, f := range d.Files {
		total += f.FileSize
	}
	return total
}

// TODO(human): Implement Name() for Directory
func (d Directory) Name() string {
	return d.DirName
}

// TotalSize returns the sum of all item sizes
func TotalSize(items []Sizer) int64 {
	// TODO(human): Implement
	var total int64

	for _, i := range items {
		total += i.Size()
	}
	return total
}

// ListNames returns all item names
func ListNames(items []Namer) []string {
	// TODO(human): Implement

	result := make([]string, 0, len(items))

	for _, i := range items {
		result = append(result, i.Name())
	}

	return result
}

// LargestBySize returns the largest item by size
func LargestBySize(items []Sizer) (Sizer, bool) {
	// TODO(human): Implement
	maxSize := int64(-1)
	var largest Sizer
	for _, i := range items {
		if i.Size() > maxSize {
			maxSize = i.Size()
			largest = i
		}
	}

	if maxSize < 0 {
		return nil, false
	}
	return largest, true
}
