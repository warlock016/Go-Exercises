# Go Patterns for This Project

This guide covers Go patterns you'll use in the Tuya IoT Measurement API.

---

## 1. Thread-Safe Map with sync.RWMutex

When multiple goroutines access shared data, use `sync.RWMutex`:

```go
type Registry struct {
    mu      sync.RWMutex
    devices map[string]*Device
}

// Read operation - multiple readers allowed
func (r *Registry) Get(id string) (*Device, bool) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    device, ok := r.devices[id]
    return device, ok
}

// Write operation - exclusive access
func (r *Registry) Add(device *Device) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.devices[device.ID] = device
}
```

**Key Points:**
- `RLock()` for read-only operations (multiple concurrent readers)
- `Lock()` for write operations (exclusive access)
- Always use `defer` to ensure unlock happens

**Documentation:** https://pkg.go.dev/sync#RWMutex

---

## 2. Context for Cancellation and Timeouts

Use `context.Context` for operations that may need cancellation:

```go
func (d *Device) Query(ctx context.Context) (*State, error) {
    // Create connection with timeout
    dialer := net.Dialer{Timeout: 5 * time.Second}
    conn, err := dialer.DialContext(ctx, "tcp", d.IP+":6668")
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    // Set read deadline from context
    if deadline, ok := ctx.Deadline(); ok {
        conn.SetReadDeadline(deadline)
    }

    // ... send query, read response
}
```

**Creating contexts:**
```go
// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
// Call cancel() when you want to stop
```

**Documentation:** https://pkg.go.dev/context

---

## 3. Goroutine Patterns

### Starting a Background Task

```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Start monitor in background
    go func() {
        if err := monitor.Start(ctx); err != nil {
            log.Printf("monitor error: %v", err)
        }
    }()

    // Start HTTP server (blocking)
    server.ListenAndServe()
}
```

### Graceful Shutdown with Signals

```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())

    // Handle shutdown signals
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        <-sigChan
        log.Println("Shutting down...")
        cancel()
    }()

    // Start server with context
    server := &http.Server{Addr: ":8080"}
    go func() {
        <-ctx.Done()
        server.Shutdown(context.Background())
    }()

    server.ListenAndServe()
}
```

**Documentation:** https://pkg.go.dev/os/signal

---

## 4. Channel Patterns

### Updates Channel for Event Streaming

```go
type Monitor struct {
    updates chan DeviceUpdate
}

func NewMonitor() *Monitor {
    return &Monitor{
        updates: make(chan DeviceUpdate, 100), // buffered channel
    }
}

// Producer: send updates
func (m *Monitor) processPacket(pkt Packet) {
    update := DeviceUpdate{...}
    select {
    case m.updates <- update:
        // sent successfully
    default:
        // channel full, drop update
        log.Println("update channel full")
    }
}

// Consumer: receive updates
func consumeUpdates(updates <-chan DeviceUpdate) {
    for update := range updates {
        // process update
    }
}
```

**Key Points:**
- Use buffered channels to prevent blocking
- Use `select` with `default` for non-blocking sends
- Close channels from the sender side only

---

## 5. HTTP Server Patterns

### Basic ServeMux Routing

```go
func (s *Server) Routes() {
    s.mux.HandleFunc("GET /api/v1/devices", s.handleListDevices)
    s.mux.HandleFunc("GET /api/v1/devices/{id}", s.handleGetDevice)
    s.mux.HandleFunc("GET /api/v1/devices/{id}/status", s.handleGetStatus)
    s.mux.HandleFunc("POST /api/v1/devices/{id}/query", s.handleQuery)
}

func (s *Server) handleListDevices(w http.ResponseWriter, r *http.Request) {
    devices := s.registry.List()
    json.NewEncoder(w).Encode(devices)
}

func (s *Server) handleGetDevice(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")  // Go 1.22+ path parameters
    device, ok := s.registry.Get(id)
    if !ok {
        http.Error(w, "device not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(device)
}
```

**Go 1.22+ Note:** Path parameters like `{id}` work with `http.ServeMux`. Access via `r.PathValue("id")`.

**Documentation:** https://pkg.go.dev/net/http#ServeMux

### JSON Responses

```go
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
    writeJSON(w, status, map[string]string{"error": message})
}
```

---

## 6. Binary Encoding

### Reading Big-Endian Values

```go
import "encoding/binary"

func parseHeader(data []byte) (prefix, seq, cmd, length uint32) {
    prefix = binary.BigEndian.Uint32(data[0:4])
    seq = binary.BigEndian.Uint32(data[4:8])
    cmd = binary.BigEndian.Uint32(data[8:12])
    length = binary.BigEndian.Uint32(data[12:16])
    return
}
```

### Writing Big-Endian Values

```go
func encodeHeader(seq, cmd, length uint32) []byte {
    buf := make([]byte, 16)
    binary.BigEndian.PutUint32(buf[0:4], 0x000055AA)
    binary.BigEndian.PutUint32(buf[4:8], seq)
    binary.BigEndian.PutUint32(buf[8:12], cmd)
    binary.BigEndian.PutUint32(buf[12:16], length)
    return buf
}
```

**Documentation:** https://pkg.go.dev/encoding/binary

---

## 7. Error Handling Patterns

### Sentinel Errors

```go
package protocol

import "errors"

var (
    ErrInvalidPrefix  = errors.New("invalid packet prefix")
    ErrInvalidSuffix  = errors.New("invalid packet suffix")
    ErrPayloadTooShort = errors.New("payload too short")
    ErrDecryptFailed  = errors.New("decryption failed")
)
```

### Error Wrapping

```go
func (d *Device) Connect(ctx context.Context) error {
    conn, err := net.Dial("tcp", d.IP+":6668")
    if err != nil {
        return fmt.Errorf("connecting to %s: %w", d.ID, err)
    }
    d.conn = conn
    return nil
}

// Caller can check:
if errors.Is(err, net.ErrClosed) {
    // handle closed connection
}
```

**Documentation:** https://pkg.go.dev/errors

---

## 8. Interface-Based Design

### Define Behavior, Not Data

```go
// Good: interface defines behavior
type Cipher interface {
    Encrypt(plaintext []byte) ([]byte, error)
    Decrypt(ciphertext []byte) ([]byte, error)
}

// Implementation
type AESCipher struct {
    key []byte
}

func (c *AESCipher) Encrypt(plaintext []byte) ([]byte, error) {
    // implementation
}

func (c *AESCipher) Decrypt(ciphertext []byte) ([]byte, error) {
    // implementation
}
```

**Key Points:**
- Keep interfaces small (1-3 methods)
- Define interfaces where they're used, not where implemented
- Accept interfaces, return structs

---

## 9. Constructor Patterns

```go
func NewDevice(id, ip, localKey string) (*Device, error) {
    if id == "" {
        return nil, errors.New("device ID required")
    }
    if len(localKey) != 16 {
        return nil, errors.New("local key must be 16 bytes")
    }

    return &Device{
        ID:       id,
        IP:       ip,
        LocalKey: localKey,
        state:    &State{},
    }, nil
}
```

---

## 10. Testing Patterns

### Table-Driven Tests

```go
func TestDecrypt(t *testing.T) {
    tests := []struct {
        name       string
        ciphertext []byte
        key        []byte
        want       []byte
        wantErr    bool
    }{
        {
            name:       "valid packet",
            ciphertext: []byte{...},
            key:        []byte("1234567890123456"),
            want:       []byte(`{"dps":{"1":true}}`),
        },
        {
            name:       "invalid key length",
            ciphertext: []byte{...},
            key:        []byte("short"),
            wantErr:    true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Decrypt(tt.ciphertext, tt.key)
            if (err != nil) != tt.wantErr {
                t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !bytes.Equal(got, tt.want) {
                t.Errorf("Decrypt() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

---

## Think About

1. When should you use `RLock` vs `Lock`?
2. Why use buffered channels for the updates channel?
3. What happens if you forget `defer cancel()` after creating a context with timeout?
4. How would you test the Registry without creating real devices?

---

## Quick Reference

| Pattern | Package | Use Case |
|---------|---------|----------|
| RWMutex | sync | Thread-safe shared data |
| Context | context | Timeouts, cancellation |
| Channels | builtin | Communication between goroutines |
| ServeMux | net/http | HTTP routing |
| BigEndian | encoding/binary | Binary protocol parsing |
| Sentinel errors | errors | Typed error checking |
