# API Package

HTTP REST API for accessing device information and measurements.

## Learning Goals

- Build REST endpoints with Go's standard `net/http`
- Use `http.ServeMux` for routing (Go 1.22+ features)
- Return JSON responses
- Handle errors consistently

## Files to Create

| File | Purpose |
|------|---------|
| `server.go` | HTTP server setup and routing |
| `handlers.go` | Request handler functions |
| `responses.go` | JSON response types |

## Endpoints

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/api/v1/devices` | `handleListDevices` | List all devices |
| GET | `/api/v1/devices/{id}` | `handleGetDevice` | Get device info |
| GET | `/api/v1/devices/{id}/status` | `handleGetStatus` | Get cached state |
| POST | `/api/v1/devices/{id}/query` | `handleQuery` | Force active query |
| GET | `/api/v1/health` | `handleHealth` | Health check |

## Go 1.22+ Routing

Use pattern-based routing with path parameters:

```go
mux.HandleFunc("GET /api/v1/devices/{id}", handler)
```

Access path parameters:
```go
id := r.PathValue("id")
```

## Hints

### Basic Hint
See `docs/GO_PATTERNS.md` for JSON response helpers.

### Routing Hint
```go
func (s *Server) Routes() {
    s.mux.HandleFunc("GET /api/v1/devices", s.handleListDevices)
    s.mux.HandleFunc("GET /api/v1/devices/{id}", s.handleGetDevice)
    // ... more routes
}
```

### Handler Pattern
Each handler should:
1. Extract parameters from request
2. Call registry/device methods
3. Return JSON response or error

## Testing

Use `httptest.NewRecorder()` and `httptest.NewRequest()` to test handlers without a real server.

```bash
go test -v ./api/
```

## Think About

1. How should you handle a request for a device that doesn't exist?
2. What status code should `POST /query` return on success?
3. Should you add middleware for logging or CORS?
