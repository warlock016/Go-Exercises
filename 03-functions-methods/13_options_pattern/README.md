# Exercise 13: Options Pattern

**Learning Goal:** Master the functional options pattern for flexible configuration

---

## 📝 Problem Description

The options pattern uses functional options to configure objects, avoiding:
- Long parameter lists
- Breaking changes when adding config
- Config struct exports

Pattern popularized by Dave Cheney and Rob Pike. Used extensively in production Go code (grpc, zap logger, etc.).

---

## 🎯 Function Signatures

```go
type Server struct {
    host    string
    port    int
    timeout time.Duration
    maxConn int
}

type Option func(*Server)

func NewServer(opts ...Option) *Server

func WithHost(host string) Option
func WithPort(port int) Option
func WithTimeout(timeout time.Duration) Option
func WithMaxConnections(maxConn int) Option
```

---

## 📖 Examples

```go
// Default server
server := NewServer()
// host="localhost", port=8080, timeout=30s, maxConn=100

// Custom configuration
server := NewServer(
    WithHost("0.0.0.0"),
    WithPort(9000),
    WithTimeout(60 * time.Second),
)

// Partial configuration (others use defaults)
server := NewServer(
    WithPort(3000),
    WithMaxConnections(200),
)

// Options can be stored and reused
productionOpts := []Option{
    WithHost("0.0.0.0"),
    WithTimeout(120 * time.Second),
    WithMaxConnections(500),
}
server := NewServer(productionOpts...)
```

---

## 📋 Instructions

1. Define `Server` struct with private fields
2. Define `Option` function type
3. Implement `NewServer` with sensible defaults
4. Apply all provided options
5. Implement option functions
6. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Basic Concept</summary>

Options modify a struct:
```go
type Option func(*Server)

func WithPort(port int) Option {
    return func(s *Server) {
        s.port = port
    }
}
```

</details>

<details>
<summary>Complete Solution</summary>

```go
package options_pattern

import "time"

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

type Option func(*Server)

func NewServer(opts ...Option) *Server {
	server := &Server{
		host:    "localhost",
		port:    8080,
		timeout: 30 * time.Second,
		maxConn: 100,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

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

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConnections(maxConn int) Option {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

func (s *Server) Host() string           { return s.host }
func (s *Server) Port() int              { return s.port }
func (s *Server) Timeout() time.Duration { return s.timeout }
func (s *Server) MaxConnections() int    { return s.maxConn }
```

</details>

---

## 🤔 Think About

1. Why use private fields with public getters?
2. What are the benefits over a config struct?
3. How does this prevent breaking changes?
4. When should you NOT use this pattern?

---

## 🎓 What This Teaches

- **Options pattern**: Functional configuration
- **Variadic options**: Flexible parameter lists
- **Backward compatibility**: Adding options without breaking changes
- **Encapsulation**: Private fields with controlled access
- **API design**: Creating flexible, maintainable APIs

---

**Tier:** 4 - Mastery
**Estimated Time:** 45-55 minutes
