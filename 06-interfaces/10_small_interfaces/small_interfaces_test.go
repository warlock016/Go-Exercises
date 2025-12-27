package small_interfaces

import (
	"reflect"
	"testing"
)

func TestFileSize(t *testing.T) {
	file := File{FileName: "test.txt", FileSize: 100}
	if file.Size() != 100 {
		t.Errorf("File.Size() = %d, want 100", file.Size())
	}
}

func TestFileName(t *testing.T) {
	file := File{FileName: "test.txt", FileSize: 100}
	if file.Name() != "test.txt" {
		t.Errorf("File.Name() = %q, want %q", file.Name(), "test.txt")
	}
}

func TestDirectorySize(t *testing.T) {
	dir := Directory{
		DirName: "docs",
		Files: []File{
			{FileName: "a.txt", FileSize: 100},
			{FileName: "b.txt", FileSize: 200},
			{FileName: "c.txt", FileSize: 50},
		},
	}

	want := int64(350)
	if dir.Size() != want {
		t.Errorf("Directory.Size() = %d, want %d", dir.Size(), want)
	}
}

func TestDirectoryName(t *testing.T) {
	dir := Directory{DirName: "docs", Files: []File{}}
	if dir.Name() != "docs" {
		t.Errorf("Directory.Name() = %q, want %q", dir.Name(), "docs")
	}
}

func TestTotalSize(t *testing.T) {
	file1 := File{FileName: "a.txt", FileSize: 100}
	file2 := File{FileName: "b.txt", FileSize: 200}
	dir := Directory{
		DirName: "docs",
		Files:   []File{{FileName: "c.txt", FileSize: 50}},
	}

	items := []Sizer{file1, file2, dir}
	total := TotalSize(items)

	want := int64(350) // 100 + 200 + 50
	if total != want {
		t.Errorf("TotalSize() = %d, want %d", total, want)
	}
}

func TestTotalSizeEmpty(t *testing.T) {
	total := TotalSize([]Sizer{})
	if total != 0 {
		t.Errorf("TotalSize([]) = %d, want 0", total)
	}
}

func TestListNames(t *testing.T) {
	file1 := File{FileName: "a.txt", FileSize: 100}
	file2 := File{FileName: "b.txt", FileSize: 200}
	dir := Directory{DirName: "docs", Files: []File{}}

	items := []Namer{file1, file2, dir}
	names := ListNames(items)

	want := []string{"a.txt", "b.txt", "docs"}
	if !reflect.DeepEqual(names, want) {
		t.Errorf("ListNames() = %v, want %v", names, want)
	}
}

func TestListNamesEmpty(t *testing.T) {
	names := ListNames([]Namer{})
	if names == nil {
		t.Error("ListNames([]) should return empty slice, not nil")
	}
	if len(names) != 0 {
		t.Errorf("ListNames([]) length = %d, want 0", len(names))
	}
}

func TestLargestBySize(t *testing.T) {
	file1 := File{FileName: "small.txt", FileSize: 100}
	file2 := File{FileName: "large.txt", FileSize: 500}
	file3 := File{FileName: "medium.txt", FileSize: 200}

	items := []Sizer{file1, file2, file3}
	largest, ok := LargestBySize(items)

	if !ok {
		t.Fatal("LargestBySize() should return true for non-empty slice")
	}

	if largest.Size() != 500 {
		t.Errorf("LargestBySize() size = %d, want 500", largest.Size())
	}

	// Should be file2
	if f, ok := largest.(File); !ok || f.FileName != "large.txt" {
		t.Errorf("LargestBySize() returned wrong item")
	}
}

func TestLargestBySizeEmpty(t *testing.T) {
	_, ok := LargestBySize([]Sizer{})
	if ok {
		t.Error("LargestBySize([]) should return false for empty slice")
	}
}

func TestImplementsSizer(t *testing.T) {
	var _ Sizer = File{}
	var _ Sizer = Directory{}
	t.Log("✓ Both File and Directory implement Sizer interface")
}

func TestImplementsNamer(t *testing.T) {
	var _ Namer = File{}
	var _ Namer = Directory{}
	t.Log("✓ Both File and Directory implement Namer interface")
}

func TestSmallInterfacesAreFlexible(t *testing.T) {
	// This test demonstrates the flexibility of small interfaces

	file := File{FileName: "test.txt", FileSize: 100}

	// File can be used as Sizer
	var s Sizer = file
	_ = s.Size()

	// File can be used as Namer
	var n Namer = file
	_ = n.Name()

	// File can be used in functions accepting either interface
	TotalSize([]Sizer{file})
	ListNames([]Namer{file})
}
