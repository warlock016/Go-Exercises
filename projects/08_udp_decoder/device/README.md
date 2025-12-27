# Device Package

Manages device connections, state caching, and the device registry.

## Learning Goals

- Use `sync.RWMutex` for thread-safe shared data
- Manage TCP connections with proper lifecycle
- Design registry pattern with dual indexing (by ID and by IP)
- Cache state with timestamps

## Files to Create

| File | Purpose |
|------|---------|
| `device.go` | Device struct, TCP connection methods |
| `registry.go` | Thread-safe device storage |
| `state.go` | Measurement state caching |

## Key Concepts

### Registry Pattern

Store devices in a map with mutex protection:
- Multiple goroutines will read device state (API handlers)
- One goroutine will write updates (monitor)
- Need to lookup by both device ID (API) and IP address (monitor)

### Connection Lifecycle

1. **Connect:** Dial TCP to device:6668
2. **Query:** Send DP_QUERY command, read response
3. **Heartbeat:** Keep connection alive (every 15-30 seconds)
4. **Disconnect:** Clean close when done

### State Caching

Store the most recent device measurements:
- DPs map (data points)
- Timestamp of last update
- Source (passive monitoring vs active query)

## Hints

### Basic Hint
Look at `docs/GO_PATTERNS.md` for the RWMutex pattern.

### Intermediate Hint
For dual indexing, maintain two maps:
```go
type Registry struct {
    mu      sync.RWMutex
    byID    map[string]*Device
    byIP    map[string]*Device
}
```

### TCP Connection Hint
Use `net.DialTimeout()` for connection with timeout:
```go
conn, err := net.DialTimeout("tcp", ip+":6668", 5*time.Second)
```

## Testing

Mock the TCP connection for unit tests. Test the registry thread-safety with concurrent goroutines.

```bash
go test -race ./device/
```

## Think About

1. What happens if two goroutines try to update the same device's state?
2. How would you handle a device that goes offline?
3. Should Connect() be idempotent (safe to call multiple times)?
