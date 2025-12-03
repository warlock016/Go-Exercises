# Simple Data Logger - Project Specification

An HTTP server that ingests, stores, and queries time-series data points.

**Estimated Time:** 6-8 hours
**Difficulty:** Intermediate-Advanced
**Prerequisites:** Project 0 & 1 completed, Module 03 (Functions/Methods), Module 06 (Interfaces basics)

---

## Learning Goals

This project practices:
- **HTTP server** - Creating endpoints with `net/http` (NEW - you've only used client)
- **JSON request handling** - Parsing JSON request bodies (NEW)
- **In-memory storage** - Designing data structures for queries
- **Time handling** - Working with `time.Time` and time ranges
- **Goroutines** - Background periodic export (gentle concurrency intro)
- **RESTful API design** - Resource-based endpoints

---

## Requirements

### Functional Requirements

1. **Ingest data points** via HTTP POST
2. **Query data points** by time range via HTTP GET
3. **List all sensors** that have reported data
4. **Export data** to JSON file periodically (background task)
5. **Health check** endpoint for monitoring

### HTTP API

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/data` | Ingest a new data point |
| GET | `/api/data` | Query data points (with time range) |
| GET | `/api/sensors` | List all known sensors |
| GET | `/api/export` | Trigger manual export to file |
| GET | `/health` | Health check (returns 200 OK) |

### Data Point Structure

```json
{
  "sensor_id": "temp-sensor-01",
  "value": 23.5,
  "unit": "celsius",
  "timestamp": "2025-01-15T14:30:00Z"
}
```

If `timestamp` is omitted, server should use current time.

### Query Parameters for GET /api/data

| Parameter | Description | Example |
|-----------|-------------|---------|
| `sensor_id` | Filter by sensor | `?sensor_id=temp-sensor-01` |
| `start` | Start time (RFC3339) | `?start=2025-01-15T00:00:00Z` |
| `end` | End time (RFC3339) | `?end=2025-01-15T23:59:59Z` |
| `limit` | Max results | `?limit=100` |

### Non-Functional Requirements

- Server should start on configurable port (flag or env var)
- Data stored in memory (lost on restart - that's OK)
- Periodic export runs every N minutes (configurable)
- Graceful error responses with appropriate HTTP status codes
- Request logging to stdout

---

## Example Interactions

### Ingest Data Point
```bash
$ curl -X POST http://localhost:8080/api/data \
  -H "Content-Type: application/json" \
  -d '{"sensor_id": "temp-01", "value": 23.5, "unit": "celsius"}'

{"status": "ok", "id": "dp-12345"}
```

### Query Data
```bash
$ curl "http://localhost:8080/api/data?sensor_id=temp-01&limit=5"

{
  "count": 2,
  "data": [
    {"sensor_id": "temp-01", "value": 23.5, "unit": "celsius", "timestamp": "2025-01-15T14:30:00Z"},
    {"sensor_id": "temp-01", "value": 24.1, "unit": "celsius", "timestamp": "2025-01-15T14:35:00Z"}
  ]
}
```

### List Sensors
```bash
$ curl http://localhost:8080/api/sensors

{
  "sensors": ["temp-01", "humidity-01", "pressure-01"],
  "count": 3
}
```

### Health Check
```bash
$ curl http://localhost:8080/health

{"status": "ok", "uptime": "2h15m30s", "data_points": 1542}
```

### Error Responses
```bash
$ curl -X POST http://localhost:8080/api/data \
  -d 'invalid json'

HTTP/1.1 400 Bad Request
{"error": "invalid JSON: unexpected character 'i'"}

$ curl "http://localhost:8080/api/data?start=invalid-date"

HTTP/1.1 400 Bad Request
{"error": "invalid start time format, expected RFC3339"}
```

---

## Suggested File Structure

```
02_data_logger/
├── main.go           # Entry point: flag parsing, server setup
├── handlers.go       # HTTP handlers: IngestHandler, QueryHandler, etc.
├── storage.go        # In-memory storage: Store struct with methods
├── exporter.go       # Periodic export: ExportToFile, background goroutine
├── models.go         # Data structures: DataPoint, QueryParams, etc.
├── storage_test.go   # Tests for storage operations
└── README.md         # This spec
```

---

## Useful Packages

| Package | Purpose | Documentation |
|---------|---------|---------------|
| `net/http` | HTTP server | https://pkg.go.dev/net/http |
| `encoding/json` | JSON encoding/decoding | https://pkg.go.dev/encoding/json |
| `time` | Time parsing, formatting, durations | https://pkg.go.dev/time |
| `sync` | Mutex for thread-safe storage | https://pkg.go.dev/sync |
| `os` | File operations, environment vars | https://pkg.go.dev/os |
| `log` | Request logging | https://pkg.go.dev/log |
| `flag` | Command-line flags | https://pkg.go.dev/flag |

---

## Key Concepts to Learn

### 1. Basic HTTP Server

```go
http.HandleFunc("/health", healthHandler)
log.Println("Server starting on :8080")
log.Fatal(http.ListenAndServe(":8080", nil))
```

### 2. HTTP Handler Function

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
```

### 3. Reading JSON Request Body

```go
func ingestHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var dp DataPoint
    if err := json.NewDecoder(r.Body).Decode(&dp); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Process dp...
}
```

### 4. Query Parameters

```go
func queryHandler(w http.ResponseWriter, r *http.Request) {
    sensorID := r.URL.Query().Get("sensor_id")  // Returns "" if not present
    startStr := r.URL.Query().Get("start")

    if startStr != "" {
        start, err := time.Parse(time.RFC3339, startStr)
        // ...
    }
}
```

### 5. Time Parsing

```go
// RFC3339 is the standard format for APIs
t, err := time.Parse(time.RFC3339, "2025-01-15T14:30:00Z")

// Format time for output
formatted := t.Format(time.RFC3339)

// Current time
now := time.Now()
```

### 6. Background Goroutine (Periodic Task)

```go
func startPeriodicExport(interval time.Duration) {
    go func() {
        ticker := time.NewTicker(interval)
        for range ticker.C {
            exportToFile()
        }
    }()
}
```

### 7. Thread-Safe Storage

```go
type Store struct {
    mu     sync.RWMutex
    points []DataPoint
}

func (s *Store) Add(dp DataPoint) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.points = append(s.points, dp)
}

func (s *Store) Query(...) []DataPoint {
    s.mu.RLock()
    defer s.mu.RUnlock()
    // Read operations...
}
```

---

## Hints (Read Only If Stuck)

<details>
<summary>Hint 1: Routing different methods to same path</summary>

```go
func dataHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        queryData(w, r)
    case http.MethodPost:
        ingestData(w, r)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}
```
</details>

<details>
<summary>Hint 2: Global vs dependency injection for storage</summary>

Simple approach (global):
```go
var store = NewStore()
```

Better approach (dependency injection):
```go
type Server struct {
    store *Store
}
func (s *Server) ingestHandler(w http.ResponseWriter, r *http.Request) {
    // Use s.store...
}
```
</details>

<details>
<summary>Hint 3: Filtering data by time range</summary>

```go
func (s *Store) Query(sensorID string, start, end time.Time) []DataPoint {
    var results []DataPoint
    for _, dp := range s.points {
        if sensorID != "" && dp.SensorID != sensorID {
            continue
        }
        if !start.IsZero() && dp.Timestamp.Before(start) {
            continue
        }
        if !end.IsZero() && dp.Timestamp.After(end) {
            continue
        }
        results = append(results, dp)
    }
    return results
}
```
</details>

<details>
<summary>Hint 4: JSON response helper</summary>

```go
func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func jsonError(w http.ResponseWriter, status int, message string) {
    jsonResponse(w, status, map[string]string{"error": message})
}
```
</details>

<details>
<summary>Hint 5: Graceful shutdown (advanced)</summary>

```go
import "os/signal"

// Catch interrupt signal
c := make(chan os.Signal, 1)
signal.Notify(c, os.Interrupt)
<-c

// Perform cleanup (export remaining data, etc.)
log.Println("Shutting down...")
```
</details>

---

## Success Criteria

- [ ] Server starts on configurable port (`-port 8080`)
- [ ] POST `/api/data` ingests data points
- [ ] GET `/api/data` returns all data points
- [ ] GET `/api/data?sensor_id=X` filters by sensor
- [ ] GET `/api/data?start=X&end=Y` filters by time range
- [ ] GET `/api/sensors` lists unique sensors
- [ ] GET `/health` returns status and uptime
- [ ] Invalid JSON returns 400 with error message
- [ ] Invalid query params return 400 with error message
- [ ] Request logging shows method, path, status
- [ ] Tests exist for storage operations
- [ ] Periodic export runs in background (optional)

---

## Stretch Goals (Optional)

- [ ] Add authentication (API key in header)
- [ ] Add rate limiting
- [ ] Persist data to file on shutdown, reload on startup
- [ ] Add WebSocket endpoint for real-time data streaming
- [ ] Add `/api/stats` endpoint with aggregations (min, max, avg per sensor)
- [ ] Add Prometheus metrics endpoint (`/metrics`)

---

## Testing the Server

### Using curl
```bash
# Start server
go run . -port 8080

# In another terminal:
curl http://localhost:8080/health
curl -X POST http://localhost:8080/api/data -d '{"sensor_id":"test","value":42}'
curl http://localhost:8080/api/data
```

### Using a test script
Create `test.sh`:
```bash
#!/bin/bash
BASE="http://localhost:8080"

echo "=== Health Check ==="
curl -s $BASE/health | jq

echo "=== Ingest Data ==="
curl -s -X POST $BASE/api/data \
  -H "Content-Type: application/json" \
  -d '{"sensor_id":"temp-01","value":23.5,"unit":"celsius"}' | jq

echo "=== Query Data ==="
curl -s "$BASE/api/data" | jq
```

---

## Getting Started

1. Start with `main.go` - get a basic server running with `/health` endpoint
2. Add `models.go` - define your data structures
3. Implement `storage.go` - in-memory storage with Add and Query
4. Add `handlers.go` - implement POST and GET for `/api/data`
5. Add periodic export in `exporter.go`
6. Write tests for storage operations

Ask for help if you're stuck on a specific concept!
