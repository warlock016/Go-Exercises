# Exercise 10: Small Interfaces

**Concept:** Single-method interfaces and the power of small abstractions
**Difficulty:** Medium
**Estimated Time:** 55 minutes

## Learning Goal

Learn Go's philosophy: "The bigger the interface, the weaker the abstraction." Small, focused interfaces are more powerful and flexible.

## The Problem

Large interfaces are rigid:

```go
// Bad - large, inflexible interface
type FileSystem interface {
    Open() error
    Close() error
    Read() []byte
    Write([]byte) error
    GetSize() int64
    GetName() string
    GetModTime() time.Time
}
```

Small interfaces compose better:

```go
// Good - small, focused interfaces
type Sizer interface {
    Size() int64
}

type Namer interface {
    Name() string
}

// Compose when needed
type SizedNamed interface {
    Sizer
    Namer
}
```

## Your Task

Design small interfaces for file system objects.

### Interfaces

1. **Sizer** - things with a size
   - `Size() int64`

2. **Namer** - things with a name
   - `Name() string`

### Types

1. **File** - has name and size
2. **Directory** - has name and contains files

### Functions

1. **TotalSize(items []Sizer) int64** - sum of all sizes
2. **ListNames(items []Namer) []string** - collect all names
3. **LargestBySize(items []Sizer) (Sizer, bool)** - find largest item

## Function Signatures

```go
type Sizer interface {
    Size() int64
}

type Namer interface {
    Name() string
}

type File struct {
    FileName string
    FileSize int64
}

type Directory struct {
    DirName string
    Files   []File
}

func (f File) Size() int64
func (f File) Name() string
func (d Directory) Size() int64
func (d Directory) Name() string

func TotalSize(items []Sizer) int64
func ListNames(items []Namer) []string
func LargestBySize(items []Sizer) (Sizer, bool)
```

## What This Teaches

- Single-method interfaces
- Interface composition
- Accept minimal interfaces
- Small, focused abstractions

---

**Next up:** Exercise 11 - Interface Segregation
