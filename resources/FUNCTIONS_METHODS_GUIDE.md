# Functions & Methods Guide

A comprehensive reference for Go functions and methods patterns.

---

## Function Basics

### Function Declaration

```go
func functionName(param1 type1, param2 type2) returnType {
    // body
}

// Multiple parameters of same type
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // "naked" return
}
```

### Variadic Functions

```go
// Accept any number of arguments
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// Calling variadic functions
sum(1, 2, 3)           // Direct arguments
sum(numbers...)        // Unpacking a slice
```

---

## Methods

### Value Receivers

```go
type Point struct {
    X, Y float64
}

// Value receiver - operates on a COPY
func (p Point) Distance() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Cannot modify the original
func (p Point) Scale(factor float64) Point {
    return Point{p.X * factor, p.Y * factor}
}
```

### Pointer Receivers

```go
// Pointer receiver - can modify the original
func (p *Point) Move(dx, dy float64) {
    p.X += dx
    p.Y += dy
}

// Use pointer receivers when:
// 1. You need to modify the receiver
// 2. The struct is large (avoid copying)
// 3. Consistency (if any method needs pointer, use for all)
```

### Method Sets

```go
type Counter struct {
    count int
}

func (c *Counter) Increment() { c.count++ }
func (c Counter) Value() int  { return c.count }

// Value types can call both value and pointer methods
// Pointer types can call both value and pointer methods
// Interface satisfaction depends on method set
```

---

## Closures

### Basic Closure

```go
func counter() func() int {
    count := 0
    return func() int {
        count++  // Captures count from outer scope
        return count
    }
}

c := counter()
c()  // 1
c()  // 2
c()  // 3
```

### Closure with Parameters

```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

double := multiplier(2)
triple := multiplier(3)
double(5)  // 10
triple(5)  // 15
```

---

## Higher-Order Functions

### Functions as Parameters

```go
// Function type
type Predicate func(int) bool

func filter(nums []int, pred Predicate) []int {
    result := make([]int, 0)
    for _, n := range nums {
        if pred(n) {
            result = append(result, n)
        }
    }
    return result
}

// Usage
evens := filter(nums, func(n int) bool {
    return n%2 == 0
})
```

### Common Patterns

```go
// Map: transform each element
func Map[T, U any](items []T, fn func(T) U) []U {
    result := make([]U, len(items))
    for i, item := range items {
        result[i] = fn(item)
    }
    return result
}

// Filter: keep matching elements
func Filter[T any](items []T, pred func(T) bool) []T {
    result := make([]T, 0)
    for _, item := range items {
        if pred(item) {
            result = append(result, item)
        }
    }
    return result
}

// Reduce: accumulate to single value
func Reduce[T, U any](items []T, initial U, fn func(U, T) U) U {
    result := initial
    for _, item := range items {
        result = fn(result, item)
    }
    return result
}
```

---

## Function Factories

```go
// Factory that creates configured functions
func makeValidator(minLen, maxLen int) func(string) bool {
    return func(s string) bool {
        return len(s) >= minLen && len(s) <= maxLen
    }
}

usernameValidator := makeValidator(3, 20)
passwordValidator := makeValidator(8, 100)
```

---

## HTTP Handler Functions

### Basic Handler

```go
// Handler function signature
func handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}

// Register handler
http.HandleFunc("/", handler)
```

### HandlerFunc Type

```go
// http.HandlerFunc is a type that implements http.Handler
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}

// This allows functions to be used as handlers
http.Handle("/", http.HandlerFunc(myHandler))
```

---

## Middleware Pattern

### Basic Middleware

```go
// Middleware wraps a handler with additional behavior
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)  // Call the next handler
    })
}

// Usage
handler := loggingMiddleware(http.HandlerFunc(myHandler))
```

### Chaining Middleware

```go
// Chain multiple middleware
func chain(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
    for i := len(middleware) - 1; i >= 0; i-- {
        h = middleware[i](h)
    }
    return h
}

// Usage
handler := chain(
    finalHandler,
    loggingMiddleware,
    authMiddleware,
    timingMiddleware,
)
```

---

## Options Pattern

### Functional Options

```go
type Server struct {
    host    string
    port    int
    timeout time.Duration
}

type Option func(*Server)

func WithHost(host string) Option {
    return func(s *Server) {
        s.host = host
    }
}

func WithPort(port int) Option {
    return func(s *Server) {
        s.port = port
    }
}

func WithTimeout(d time.Duration) Option {
    return func(s *Server) {
        s.timeout = d
    }
}

func NewServer(opts ...Option) *Server {
    s := &Server{
        host:    "localhost",
        port:    8080,
        timeout: 30 * time.Second,
    }
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// Usage
server := NewServer(
    WithHost("0.0.0.0"),
    WithPort(9000),
    WithTimeout(60 * time.Second),
)
```

---

## Method Chaining (Builder Pattern)

```go
type QueryBuilder struct {
    table   string
    columns []string
    where   []string
}

func NewQuery() *QueryBuilder {
    return &QueryBuilder{}
}

func (q *QueryBuilder) Select(cols ...string) *QueryBuilder {
    q.columns = cols
    return q  // Return self for chaining
}

func (q *QueryBuilder) From(table string) *QueryBuilder {
    q.table = table
    return q
}

func (q *QueryBuilder) Where(condition string) *QueryBuilder {
    q.where = append(q.where, condition)
    return q
}

func (q *QueryBuilder) Build() string {
    // Build SQL string
    return fmt.Sprintf("SELECT %s FROM %s WHERE %s",
        strings.Join(q.columns, ", "),
        q.table,
        strings.Join(q.where, " AND "))
}

// Usage - fluent API
query := NewQuery().
    Select("id", "name", "email").
    From("users").
    Where("active = true").
    Where("age > 18").
    Build()
```

---

## Dependency Injection

```go
// Define interface for dependencies
type UserRepository interface {
    FindByID(id int) (*User, error)
}

type EmailSender interface {
    Send(to, subject, body string) error
}

// Service with injected dependencies
type UserService struct {
    repo   UserRepository
    mailer EmailSender
}

func NewUserService(repo UserRepository, mailer EmailSender) *UserService {
    return &UserService{repo: repo, mailer: mailer}
}

// Now you can inject mocks for testing
```

---

## Common Patterns Summary

| Pattern | Use Case |
|---------|----------|
| Variadic | Unknown number of same-type arguments |
| Closures | Capture state, create factories |
| Higher-order | Transform, filter, reduce collections |
| Middleware | Add cross-cutting concerns (logging, auth) |
| Options | Flexible configuration with defaults |
| Builder | Complex object construction with fluent API |
| DI | Testable code, loose coupling |

---

## Resources

- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
- [Go Blog - First Class Functions](https://go.dev/blog/functions-codewalk)
- [Go by Example - Closures](https://gobyexample.com/closures)
