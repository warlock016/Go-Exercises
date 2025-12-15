# Home Data Miner - Architecture Guide

## Overview

This document describes the architectural patterns and data flow for building a Go API client that interacts with VictoriaMetrics' Prometheus-compatible API to fetch home automation metrics.

---

## System Context

```
┌─────────────────┐       HTTP/REST        ┌──────────────────────┐
│                 │ ───────────────────────▶│                      │
│  Home Data      │   GET + Query Params   │   VictoriaMetrics    │
│  Miner (Go)     │ ◀─────────────────────── │   :8428              │
│                 │      JSON Response      │                      │
└─────────────────┘                         └──────────────────────┘
                                                      │
                                                      │ Scrapes
                                                      ▼
                                            ┌──────────────────────┐
                                            │   Home Assistant     │
                                            │   Sensors/Entities   │
                                            └──────────────────────┘
```

### Key Components

| Component | Role | Port |
|-----------|------|------|
| Home Data Miner | Your Go application - fetches and processes metrics | N/A |
| VictoriaMetrics | Time-series database with Prometheus API | 8428 |
| Home Assistant | Source of sensor data (power, temp, etc.) | 8123 |

---

## Data Flow

### Request Lifecycle

```
1. Build Request
   ┌──────────────────────────────────────────────────────────┐
   │  Config (BaseURL, Timeout)                               │
   │       │                                                  │
   │       ▼                                                  │
   │  Construct URL + Query Parameters                        │
   │       │                                                  │
   │       ▼                                                  │
   │  Create http.Request (GET)                               │
   │       │                                                  │
   │       ▼                                                  │
   │  Set Headers (if needed)                                 │
   └──────────────────────────────────────────────────────────┘
                          │
                          ▼
2. Execute Request
   ┌──────────────────────────────────────────────────────────┐
   │  http.Client.Do(request)                                 │
   │       │                                                  │
   │       ▼                                                  │
   │  Check Response Status Code                              │
   │       │                                                  │
   │       ├── 2xx ──▶ Continue                               │
   │       └── !2xx ─▶ Return HTTP Error                      │
   └──────────────────────────────────────────────────────────┘
                          │
                          ▼
3. Process Response
   ┌──────────────────────────────────────────────────────────┐
   │  Read Response Body                                      │
   │       │                                                  │
   │       ▼                                                  │
   │  Unmarshal JSON into APIResponse                         │
   │       │                                                  │
   │       ▼                                                  │
   │  Check status == "success"                               │
   │       │                                                  │
   │       ├── success ──▶ Unmarshal Data field               │
   │       └── error ────▶ Return API Error                   │
   └──────────────────────────────────────────────────────────┘
```

### Error Handling Layers

```
┌─────────────────────────────────────────────────────────────┐
│ Layer 1: Network Errors                                     │
│   - Connection refused                                      │
│   - Timeout                                                 │
│   - DNS resolution failed                                   │
│   Handle: Return error from client.Do()                     │
├─────────────────────────────────────────────────────────────┤
│ Layer 2: HTTP Errors                                        │
│   - 4xx Client errors (bad request, not found)              │
│   - 5xx Server errors                                       │
│   Handle: Check resp.StatusCode before reading body         │
├─────────────────────────────────────────────────────────────┤
│ Layer 3: API Errors                                         │
│   - status: "error" in JSON response                        │
│   - Invalid PromQL syntax                                   │
│   Handle: Check APIResponse.Status field                    │
├─────────────────────────────────────────────────────────────┤
│ Layer 4: Parsing Errors                                     │
│   - Malformed JSON                                          │
│   - Unexpected data structure                               │
│   Handle: Return error from json.Unmarshal()                │
└─────────────────────────────────────────────────────────────┘
```

---

## Component Design

### Package Structure

```
06_home_data_miner/
├── main.go              # Entry point, CLI handling
├── .env                 # Configuration (gitignored)
│
├── client/              # API client package
│   ├── client.go        # Client struct, NewClient, HTTP transport config
│   ├── types.go         # Response types: VMEnvelope, Metric, InstantMetric, etc.
│   ├── queries.go       # Query and QueryRange request builders
│   ├── series.go        # Labels, LabelValues, Series discovery methods
│   └── errors.go        # Custom error types for API/HTTP/parse errors
│
├── docs/                # Documentation
│   ├── ARCHITECTURE.md  # This file
│   ├── API_REFERENCE.md # VictoriaMetrics API docs
│   └── PATTERNS.md      # Go patterns guide
│
├── cache/               # Cached API response samples (for development)
│
└── TODO.md              # Implementation checklist
```

### File Responsibilities

| File | Responsibility | Key Types/Functions |
|------|---------------|---------------------|
| `client.go` | HTTP client wrapper, connection pooling, transport config | `Client`, `NewClient()`, `Do()` |
| `types.go` | API response marshaling, data structures | `VMEnvelope`, `QueryResponse`, `Metric`, `InstantMetric`, `RangeMetric` |
| `queries.go` | Query/QueryRange URL construction, PromQL building | `RequestQuery()`, `RequestQueryRange()`, `BuildSelector()` |
| `series.go` | Discovery endpoint builders | `RequestLabels()`, `RequestLabelValues()`, `RequestSeries()` |
| `errors.go` | Typed errors for better handling | `APIError`, `HTTPError`, `ParseError` |

### Client Structure Concept

```
┌─────────────────────────────────────────┐
│              Client                      │
├─────────────────────────────────────────┤
│ - httpClient  *http.Client              │
│ - baseURL     string                    │
│ - timeout     time.Duration             │
├─────────────────────────────────────────┤
│ + NewClient(config) *Client             │
│ + Query(promQL) (Result, error)         │
│ + QueryRange(...) (Result, error)       │
│ + GetSeries(...) ([]Series, error)      │
│ + GetLabels() ([]string, error)         │
│ + GetLabelValues(name) ([]string, error)│
│ - doRequest(endpoint, params) ([], err) │
└─────────────────────────────────────────┘
         │
         │ uses
         ▼
┌─────────────────────────────────────────┐
│           http.Client                    │
│  (standard library, handles connection  │
│   pooling, timeouts, redirects)         │
└─────────────────────────────────────────┘
```

### Why This Design?

1. **Single Client Instance**: Reuses TCP connections (connection pooling)
2. **Encapsulated Configuration**: Base URL, timeout stored once
3. **Private Helper Method**: `doRequest` handles common HTTP logic
4. **Public API Methods**: Each endpoint becomes a clean method call

---

## Configuration Management

### Environment-Based Configuration

```
┌──────────────┐      Load       ┌──────────────┐      Create      ┌──────────────┐
│    .env      │ ──────────────▶ │   Config     │ ───────────────▶ │   Client     │
│    file      │                 │   struct     │                  │   struct     │
└──────────────┘                 └──────────────┘                  └──────────────┘

.env contents:
  VICTORIA_METRICS_URL=http://homeassistant.local:8428
  REQUEST_TIMEOUT=10s
```

### Configuration Validation Flow

```
Load from .env
     │
     ▼
Parse into Config struct
     │
     ▼
Validate required fields ──── Missing? ──▶ Return error with field name
     │
     │ OK
     ▼
Validate URL format ──────── Invalid? ──▶ Return error with details
     │
     │ OK
     ▼
Set defaults for optional fields
     │
     ▼
Return validated Config
```

---

## Dependency Relationships

```
main.go
   │
   ├──▶ client/client.go (imports)
   │         │
   │         ├──▶ client/types.go
   │         │
   │         └──▶ net/http, net/url, encoding/json (stdlib)
   │
   └──▶ github.com/hashicorp/go-envparse (external)
```

### Import Hierarchy Rule

```
Standard library imports
    ↓
External dependencies (github.com/...)
    ↓
Internal packages (./client)
```

---

## Concurrency Considerations

### Safe Concurrent Usage

```
┌─────────────────────────────────────────────────────────────┐
│  http.Client is SAFE for concurrent use                     │
│                                                             │
│  Multiple goroutines can call client.Do() simultaneously    │
│  Connection pooling is handled automatically                │
└─────────────────────────────────────────────────────────────┘

Goroutine 1 ─────┐
                 │
Goroutine 2 ─────┼────▶  Shared http.Client  ────▶  Server
                 │
Goroutine 3 ─────┘
```

### Future Extension: Parallel Queries

```
When fetching multiple series:

Sequential (slow):          Parallel (fast):
┌─────┐                     ┌─────┐ ┌─────┐ ┌─────┐
│ Q1  │                     │ Q1  │ │ Q2  │ │ Q3  │
└──┬──┘                     └──┬──┘ └──┬──┘ └──┬──┘
   │                           │      │      │
┌──▼──┐                        │      │      │
│ Q2  │                        ▼      ▼      ▼
└──┬──┘                     ┌─────────────────────┐
   │                        │   Collect Results   │
┌──▼──┐                     └─────────────────────┘
│ Q3  │
└─────┘

Time: 3x          vs        Time: 1x (roughly)
```

---

## Testing Strategy

### Test Pyramid for API Client

```
                    ┌─────────────┐
                   ╱  Integration  ╲
                  ╱   (real API)    ╲
                 ╱   Few, slow       ╲
                ├─────────────────────┤
               ╱     HTTP Tests        ╲
              ╱   (httptest server)     ╲
             ╱   Medium count, fast      ╲
            ├───────────────────────────────┤
           ╱         Unit Tests              ╲
          ╱   (URL building, parsing)         ╲
         ╱   Many, very fast                   ╲
        └───────────────────────────────────────┘
```

### What to Test at Each Level

| Level | What to Test | Example |
|-------|--------------|---------|
| Unit | URL construction | `params.Encode()` produces correct string |
| Unit | Response parsing | JSON → struct works for all response types |
| HTTP | Request formation | Correct endpoint, headers, params |
| HTTP | Error handling | 4xx/5xx responses handled correctly |
| Integration | Real queries | Actual VictoriaMetrics responds |

---

## Evolution Path

### Phase 1: Working Prototype (Current)
```
main.go does everything
  - Load config
  - Build request
  - Execute
  - Print response
```

### Phase 2: Extracted Client
```
main.go ──▶ client.Query()
              │
              └── All HTTP logic encapsulated
```

### Phase 3: Multiple Endpoints
```
main.go ──┬──▶ client.Query()
          ├──▶ client.QueryRange()
          ├──▶ client.GetSeries()
          └──▶ client.GetLabels()
```

### Phase 4: Domain-Specific Methods
```
main.go ──┬──▶ client.GetCurrentPower("total_power")
          ├──▶ client.GetPowerHistory("total_power", "7d")
          └──▶ client.DiscoverEntities()
```

---

## Key Decisions to Make

As you implement, you'll need to decide:

1. **Error Wrapping Strategy**: Use `fmt.Errorf("context: %w", err)` or custom error types?

2. **Response Type Flexibility**: Use `json.RawMessage` for dynamic data or separate structs per endpoint?

3. **Time Handling**: Accept `time.Time` and convert, or accept strings like "now", "-7d"?

4. **Logging**: Silent client or configurable logging for debugging?

5. **Retry Logic**: Automatic retries for transient failures or leave to caller?

These are design decisions with no single "right" answer - consider your use case.
