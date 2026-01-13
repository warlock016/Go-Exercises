# Go Naming Conventions Guide

A comprehensive guide to idiomatic Go naming, from variables to packages.

---

## Table of Contents

1. [Core Philosophy](#core-philosophy)
2. [Variable Naming](#variable-naming)
3. [Type-Specific Conventions](#type-specific-conventions)
4. [Function & Method Naming](#function--method-naming)
5. [Package Naming](#package-naming)
6. [Constants & Exported Identifiers](#constants--exported-identifiers)
7. [Common Patterns](#common-patterns)
8. [Anti-Patterns to Avoid](#anti-patterns-to-avoid)
9. [Quick Reference Tables](#quick-reference-tables)

---

## Core Philosophy

### The Golden Rule

**Name length should be proportional to scope size.**

```go
// Small scope (loop) → single letter is fine
for i, v := range items {
    process(v)
}

// Medium scope (function) → short but clear
func process(items []int) {
    results := make([]int, 0, len(items))
    for _, v := range items {
        results = append(results, v*2)
    }
}

// Large scope (package-level, exported) → descriptive
var DefaultTimeout = 30 * time.Second
```

### Go vs Other Languages

| Aspect | Go Style | Java/C# Style |
|--------|----------|---------------|
| Loop counter | `i`, `j`, `k` | `index`, `counter` |
| Error variable | `err` | `error`, `exception` |
| Receiver | `s`, `c`, `r` | `this`, `self` |
| Short-lived | `v`, `x`, `n` | `value`, `item`, `number` |
| Channel | `ch`, `done` | `channel`, `doneChannel` |

**Why?** Go values readability through brevity. Long names in small scopes add noise without adding clarity.

---

## Variable Naming

### By Scope Size

#### Tiny Scope (1-5 lines)
Single letters are preferred:

```go
// Good - clear in context
for i := 0; i < n; i++ { ... }
for k, v := range m { ... }
if err != nil { return err }

// Unnecessary verbosity
for index := 0; index < count; index++ { ... }  // Too long for simple loop
```

#### Small Scope (5-15 lines)
Short, conventional abbreviations:

```go
func fetchUser(id int) (*User, error) {
    ctx := context.Background()
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var u User
    if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
        return nil, err
    }
    return &u, nil
}
```

#### Medium Scope (15-50 lines)
More descriptive, but still concise:

```go
func processOrders(orders []Order) (*Summary, error) {
    validOrders := make([]Order, 0, len(orders))
    failedCount := 0

    for _, order := range orders {
        if err := validateOrder(order); err != nil {
            failedCount++
            continue
        }
        validOrders = append(validOrders, order)
    }

    totalRevenue := calculateRevenue(validOrders)
    // ... more processing
}
```

#### Large Scope (package-level, exported)
Fully descriptive:

```go
// Package-level variables
var (
    DefaultRetryCount    = 3
    MaxConcurrentFetches = 10
    RequestTimeout       = 30 * time.Second
)
```

### Common Short Names

These are universally understood in Go:

| Name | Meaning | Context |
|------|---------|---------|
| `i`, `j`, `k` | Loop indices | `for i := 0; ...` |
| `n` | Count or length | `n := len(items)` |
| `v` | Value in iteration | `for _, v := range ...` |
| `k` | Key in map iteration | `for k, v := range m` |
| `s` | String | `func process(s string)` |
| `b` | Byte slice or bool | `b := []byte(s)` |
| `r` | Reader, rune, or receiver | `for _, r := range s` |
| `w` | Writer | `func Write(w io.Writer)` |
| `t` | Type parameter, time, or test | `func Do[T any](...)` |
| `ok` | Boolean success flag | `v, ok := m[key]` |
| `err` | Error | Always `err`, never `e` or `error` |

---

## Type-Specific Conventions

### Channels

**Pattern:** Suffix with `Ch` or `Chan`, OR use descriptive noun (often plural hints at channel).

```go
// Signal channels (no data, just signals)
done := make(chan struct{})        // "done" is idiomatic for cancellation
quit := make(chan struct{})        // alternative for shutdown
shutdown := make(chan struct{})    // explicit shutdown signal

// Data channels - by what they carry
errCh := make(chan error)          // errors channel
results := make(chan Result)       // results (plural suggests stream)
jobs := make(chan Job)             // job queue
events := make(chan Event)         // event stream

// Directional channels in function signatures
func worker(jobs <-chan Job, results chan<- Result) {
    // jobs is receive-only, results is send-only
}

// When you have multiple related channels
type Pipeline struct {
    inputCh  chan Request
    outputCh chan Response
    errorCh  chan error
    doneCh   chan struct{}
}
```

**Anti-patterns:**
```go
// Too short - what does it carry?
c := make(chan error)

// Too verbose
errorChannel := make(chan error)
theChannelForErrors := make(chan error)

// Misleading
errors := make(chan error)  // Could be confused with []error
```

### WaitGroups

**Always `wg`** - this is universal:

```go
var wg sync.WaitGroup

// If you somehow need multiple (rare):
var fetchWg sync.WaitGroup
var processWg sync.WaitGroup
```

### Mutexes

**Pattern:** `mu` for single mutex, descriptive prefix for multiple:

```go
// Single mutex protecting a struct
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

// Multiple mutexes (each protects different data)
type Cache struct {
    dataMu  sync.RWMutex
    data    map[string][]byte

    statsMu sync.Mutex
    hits    int
    misses  int
}

// Package-level mutex
var (
    configMu sync.RWMutex
    config   Config
)
```

**Placement convention:** Mutex should be declared immediately before the field(s) it protects:

```go
type SafeMap struct {
    mu   sync.RWMutex  // protects the following field
    data map[string]int

    // other unprotected fields can go here
    name string
}
```

### Maps

**Pattern:** `thingsByKey` or descriptive noun:

```go
// Descriptive of the mapping relationship
usersByID := make(map[int]*User)
countByName := make(map[string]int)
configByEnv := make(map[string]Config)

// Simple cases in small scope
m := make(map[string]int)           // OK in tiny scope
results := make(map[string]string)  // OK when context is clear

// Cache patterns
cache := make(map[string][]byte)
seen := make(map[string]bool)       // "seen" implies set semantics
visited := make(map[int]bool)       // graph traversal
```

### Slices

**Pattern:** Plural noun describing contents:

```go
// Good
users := []User{}
items := make([]Item, 0, capacity)
results := make([]Result, len(inputs))
errors := []error{}  // Note: errors (slice) vs errCh (channel)

// Specific roles
pending := []Job{}
completed := []Job{}
failed := []Job{}

// When slice has specific purpose
buffer := make([]byte, 1024)
output := make([]string, 0)
```

### Structs (Local Types)

**Pattern:** Descriptive of what it bundles:

```go
// Result wrappers
type result struct {
    value int
    err   error
}

type fetchResult struct {
    url     string
    content string
    err     error
}

// Indexed results (for order preservation)
type indexedResult struct {
    index int
    value string
}

// Configuration bundles
type options struct {
    timeout    time.Duration
    retries    int
    maxWorkers int
}
```

### Context

**Always `ctx`**:

```go
func DoSomething(ctx context.Context, id int) error {
    // ctx is universal
}

// When deriving contexts
func process(ctx context.Context) {
    childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    // ...
}
```

### HTTP Types

```go
// Request/Response
req, err := http.NewRequest(...)
resp, err := client.Do(req)

// Writers/ResponseWriters
func handler(w http.ResponseWriter, r *http.Request) {
    // w and r are idiomatic
}

// Client
client := &http.Client{Timeout: 10 * time.Second}
```

### Time

```go
// Durations
timeout := 30 * time.Second
interval := 5 * time.Minute
delay := 100 * time.Millisecond

// Time values
now := time.Now()
deadline := now.Add(timeout)
start := time.Now()
elapsed := time.Since(start)

// Tickers and Timers
ticker := time.NewTicker(interval)
timer := time.NewTimer(delay)
```

---

## Function & Method Naming

### General Rules

1. **Use MixedCaps** (camelCase for unexported, PascalCase for exported)
2. **Start with verb** for actions: `Get`, `Set`, `Create`, `Delete`, `Process`, `Handle`
3. **Getters omit "Get"**: `user.Name()` not `user.GetName()`
4. **Boolean functions**: Use `Is`, `Has`, `Can`, `Should` prefixes

```go
// Good
func (u *User) Name() string        // Getter - no "Get" prefix
func (u *User) SetName(n string)    // Setter - "Set" prefix
func (u *User) IsAdmin() bool       // Boolean - "Is" prefix
func (u *User) HasPermission(p Permission) bool
func CreateUser(name string) *User  // Constructor - "Create" or "New"
func NewServer(addr string) *Server // Constructor - "New" prefix

// Avoid
func (u *User) GetName() string     // Don't use Get for getters
func (u *User) Admin() bool         // Unclear it returns bool
```

### Method Receivers

**Use 1-2 letter abbreviation of type name:**

```go
func (s *Server) Start() error { ... }
func (c *Client) Connect() error { ... }
func (u *User) FullName() string { ... }
func (rw *ResponseWriter) Write(b []byte) { ... }  // 2 letters for 2-word type
func (sc *SafeCounter) Increment() { ... }
```

### Constructor Functions

**Pattern:** `New` + type name, or `Create` + noun:

```go
// Returns pointer - use New
func NewServer(addr string) *Server { ... }
func NewClient(config Config) *Client { ... }

// Returns value - can use New or just type name
func NewPoint(x, y int) Point { ... }

// When New is awkward, use Create or Make
func CreateTempFile(dir string) (*os.File, error) { ... }
func MakeBuffer(size int) []byte { ... }
```

---

## Package Naming

### Rules

1. **Lowercase, single word** preferred
2. **No underscores or mixedCaps**
3. **Short but descriptive**
4. **Noun, not verb**

```go
// Good
package http
package json
package user
package auth
package storage

// Avoid
package httpUtil      // Use httputil (no caps)
package user_service  // No underscores
package utilities     // Too generic
package common        // Meaningless
package misc          // Avoid grab-bag packages
```

### Avoid Stutter

Package name + exported identifier should read naturally:

```go
// Good
http.Client           // Not http.HttpClient
json.Marshal          // Not json.JsonMarshal
user.New()            // Not user.NewUser()
auth.Token            // Not auth.AuthToken

// The package provides context
bytes.Buffer          // Not bytes.ByteBuffer
strings.Reader        // Not strings.StringReader
```

---

## Constants & Exported Identifiers

### Constants

```go
// Single constants - descriptive
const MaxRetries = 3
const DefaultTimeout = 30 * time.Second

// Related constants - use const block
const (
    StatusPending   = "pending"
    StatusRunning   = "running"
    StatusCompleted = "completed"
    StatusFailed    = "failed"
)

// Iota enums
const (
    LogLevelDebug = iota
    LogLevelInfo
    LogLevelWarn
    LogLevelError
)

// Private constants - can be shorter
const (
    maxBufSize = 4096
    minWorkers = 1
)
```

### Exported Identifiers

Follow standard Go conventions - the export itself signals importance:

```go
// Types
type Config struct { ... }
type Server struct { ... }
type RequestHandler func(r *Request) Response

// Variables
var DefaultClient = &Client{Timeout: 30 * time.Second}
var ErrNotFound = errors.New("not found")

// Error variables - always start with Err
var ErrTimeout = errors.New("timeout")
var ErrInvalidInput = errors.New("invalid input")

// Error types - always end with Error
type ValidationError struct { ... }
type TimeoutError struct { ... }
```

---

## Common Patterns

### The Comma-OK Idiom

```go
// Map lookup
value, ok := myMap[key]
if !ok {
    // key not found
}

// Type assertion
str, ok := v.(string)
if !ok {
    // v is not a string
}

// Channel receive
value, ok := <-ch
if !ok {
    // channel closed
}
```

### Error Handling

```go
// Always "err"
result, err := doSomething()
if err != nil {
    return err
}

// When checking multiple errors sequentially
if err := step1(); err != nil {
    return fmt.Errorf("step1: %w", err)
}
if err := step2(); err != nil {
    return fmt.Errorf("step2: %w", err)
}
```

### Defer with Named Returns

```go
func process() (result string, err error) {
    // Named returns allow defer to modify them
    defer func() {
        if err != nil {
            err = fmt.Errorf("process failed: %w", err)
        }
    }()
    // ...
}
```

### Builder Pattern

```go
type ServerBuilder struct {
    addr    string
    timeout time.Duration
}

func (b *ServerBuilder) WithAddr(addr string) *ServerBuilder {
    b.addr = addr
    return b
}

func (b *ServerBuilder) WithTimeout(d time.Duration) *ServerBuilder {
    b.timeout = d
    return b
}

func (b *ServerBuilder) Build() *Server {
    return &Server{addr: b.addr, timeout: b.timeout}
}
```

---

## Anti-Patterns to Avoid

### 1. Hungarian Notation

```go
// Bad - type is already in declaration
var strName string
var intCount int
var arrUsers []User
var mapScores map[string]int

// Good
var name string
var count int
var users []User
var scores map[string]int
```

### 2. Overly Verbose Names

```go
// Bad
theListOfAllUsersWhoAreAdmins := getAdminUsers()
numberOfItemsInTheShoppingCart := cart.ItemCount()

// Good
admins := getAdminUsers()
itemCount := cart.ItemCount()
```

### 3. Meaningless Names

```go
// Bad - what do these mean?
data := fetch()
result := process(data)
info := getInfo()
temp := calculate()

// Good - describe what they contain
users := fetchUsers()
report := generateReport(users)
config := loadConfig()
total := calculateTotal()
```

### 4. Abbreviations Only You Understand

```go
// Bad
usrMgr := NewUserManager()
cfgSvc := NewConfigService()
reqHndlr := NewRequestHandler()

// Good
userManager := NewUserManager()
configService := NewConfigService()
handler := NewRequestHandler()  // Or just 'h' in small scope
```

### 5. Single Letters in Large Scope

```go
// Bad - unclear after 50 lines
func processData(d []Data) {
    r := make([]Result, 0)
    for _, x := range d {
        // 30 more lines...
        // what is r? what is x?
    }
}

// Good
func processData(data []Data) {
    results := make([]Result, 0)
    for _, item := range data {
        // Now clear even 50 lines later
    }
}
```

---

## Quick Reference Tables

### Variables by Type

| Type | Short Scope | Long Scope |
|------|-------------|------------|
| `error` | `err` | `err` |
| `context.Context` | `ctx` | `ctx` |
| `sync.WaitGroup` | `wg` | `wg` |
| `sync.Mutex` | `mu` | `dataMu`, `cacheMu` |
| `chan T` | `ch` | `resultsCh`, `jobs` |
| `chan struct{}` | `done` | `shutdownSignal` |
| `map[K]V` | `m` | `usersByID`, `countByName` |
| `[]T` | `items` | `pendingJobs` |
| `*http.Request` | `r`, `req` | `req` |
| `http.ResponseWriter` | `w` | `w` |
| `io.Reader` | `r` | `reader` |
| `io.Writer` | `w` | `writer` |
| `string` | `s` | `name`, `query` |
| `int` (loop) | `i`, `j` | `index` |
| `int` (count) | `n` | `count`, `total` |
| `bool` | `ok`, `b` | `isValid`, `found` |
| `time.Duration` | `d` | `timeout`, `interval` |
| `time.Time` | `t` | `start`, `deadline` |

### Function Naming

| Purpose | Pattern | Example |
|---------|---------|---------|
| Constructor | `New` + Type | `NewServer()` |
| Getter | Just field name | `user.Name()` |
| Setter | `Set` + Field | `user.SetName()` |
| Bool check | `Is`/`Has`/`Can` + Adj | `IsValid()`, `HasAccess()` |
| Conversion | `To` + Type | `ToString()`, `ToJSON()` |
| Parse | `Parse` + Source | `ParseInt()`, `ParseConfig()` |
| Find | `Find` + What | `FindByID()`, `FindAll()` |

### Method Receivers

| Type Name | Receiver |
|-----------|----------|
| `Server` | `s` |
| `Client` | `c` |
| `User` | `u` |
| `Config` | `c` or `cfg` |
| `SafeMap` | `sm` |
| `HTTPClient` | `hc` |
| `ResponseWriter` | `rw` |

---

## Resources

- [Effective Go - Names](https://go.dev/doc/effective_go#names)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Standard Library](https://pkg.go.dev/std) - Best examples of Go naming

---

*Last Updated: 2026-01-12*
