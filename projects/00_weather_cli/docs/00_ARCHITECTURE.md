# Weather CLI Architecture Overview

## Purpose

This document describes the high-level architecture for a production-quality weather CLI application. The design emphasizes separation of concerns, testability, and robust error handling using patterns learned in the error-handling module.

---

## System Diagram

```
┌─────────────────────────────────────────────────────────────────────────────────┐
│                              WEATHER CLI SYSTEM                                 │
├─────────────────────────────────────────────────────────────────────────────────┤
│                                                                                 │
│  ┌──────────────────────────────────────────────────────────────────────────┐   │
│  │                              main.go                                     │   │
│  │  • Parse CLI flags                                                       │   │
│  │  • Orchestrate module calls                                              │   │
│  │  • Handle errors → user-friendly messages                                │   │
│  │  • Exit with appropriate codes                                           │   │
│  └──────────────────────────────────────────────────────────────────────────┘   │
│         │                    │                    │                    │        │
│         ▼                    ▼                    ▼                    ▼        │
│   ┌────────────┐      ┌────────────┐      ┌────────────┐      ┌────────────┐    │
│   │   config   │      │ validation │      │   client   │      │ formatter  │    │
│   │            │      │            │      │            │      │            │    │
│   │ Load env   │      │ Validate   │      │ Weather +  │      │ Table/JSON │    │
│   │ API keys   │      │ all inputs │      │ Geocoder   │      │ output     │    │
│   └────────────┘      └────────────┘      └────────────┘      └────────────┘    │
│         │                    │                    │                    │        │
│         │                    │                    ▼                    │        │
│         │                    │           ┌────────────┐                │        │
│         │                    │           │   retry    │                │        │
│         │                    │           │            │                │        │
│         │                    │           │ Backoff    │                │        │
│         │                    │           │ logic      │                │        │
│         │                    │           └────────────┘                │        │
│         │                    │                    │                    │        │
│         ▼                    ▼                    ▼                    ▼        │
│  ┌──────────────────────────────────────────────────────────────────────────┐   │
│  │                              errors                                      │   │
│  │  • Sentinel errors (ErrNotFound, ErrRateLimited, ...)                    │   │
│  │  • Custom types (WeatherAPIError, GeocodeError, ValidationErrors)        │   │
│  └──────────────────────────────────────────────────────────────────────────┘   │
│                                      │                                          │
│  ┌──────────────────────────────────────────────────────────────────────────┐   │
│  │                              types                                       │   │
│  │  • WeatherData, Location, HourlyData                                     │   │
│  │  • Shared data structures across packages                                │   │
│  └──────────────────────────────────────────────────────────────────────────┘   │
│                                                                                 │
└─────────────────────────────────────────────────────────────────────────────────┘
```

---

## Package Dependency Graph

```
                            main
                             │
         ┌───────────────────┼───────────────────┐
         │                   │                   │
         ▼                   ▼                   ▼
      config            validation           formatter
         │                   │                   │
         │                   │                   │
         ▼                   ▼                   ▼
      errors              errors              types
                             │                   │
                             ▼                   │
                          types ◄────────────────┘
                             │
         ┌───────────────────┘
         │
         ▼
      client ──────► retry
         │              │
         ▼              ▼
      errors         errors
         │
         ▼
      types
```

### Dependency Rules

1. **errors** - Has NO dependencies (leaf package)
2. **types** - Has NO dependencies (leaf package)
3. **config** - Depends on: `errors`
4. **validation** - Depends on: `errors`, `types`
5. **retry** - Depends on: `errors`
6. **client** - Depends on: `errors`, `types`, `retry`
7. **formatter** - Depends on: `types`
8. **main** - Depends on: ALL packages (orchestrator)

---

## Data Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           REQUEST FLOW                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   CLI Arguments                                                             │
│        │                                                                    │
│        ▼                                                                    │
│   ┌─────────────┐                                                           │
│   │ Flag Parse  │  --lat, --long, --start, --end, --format                  │
│   └─────────────┘                                                           │
│        │                                                                    │
│        ▼                                                                    │
│   ┌─────────────┐                                                           │
│   │ Config Load │  Load GEOCODE_API_KEY from env/.env                       │
│   └─────────────┘                                                           │
│        │                                                                    │
│        ▼                                                                    │
│   ┌─────────────┐                                                           │
│   │ Validation  │  Validate all inputs, collect ALL errors                  │
│   └─────────────┘                                                           │
│        │                                                                    │
│        ├──────────────────────────────────┐                                 │
│        ▼                                  ▼                                 │
│   ┌─────────────┐                   ┌─────────────┐                         │
│   │ Weather API │                   │ Geocoder API│                         │
│   │ (Open-Meteo)│                   │ (Nominatim) │                         │
│   └─────────────┘                   └─────────────┘                         │
│        │                                  │                                 │
│        │    ┌─────────────────────────────┘                                 │
│        ▼    ▼                                                               │
│   ┌─────────────┐                                                           │
│   │  Formatter  │  Combine weather + location → output                      │
│   └─────────────┘                                                           │
│        │                                                                    │
│        ▼                                                                    │
│   ┌─────────────┐                                                           │
│   │   stdout    │  Table or JSON format                                     │
│   └─────────────┘                                                           │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Error Propagation Path

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                          ERROR FLOW                                         │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   Layer 4: API Clients (weather.go, geocoder.go)                            │
│   ├── Network timeout         → ErrTimeout (retryable)                      │
│   ├── HTTP 404                → ErrNotFound (not retryable)                 │
│   ├── HTTP 429                → ErrRateLimited (retryable)                  │
│   ├── HTTP 5xx                → WeatherAPIError/GeocodeError (retryable)    │
│   └── JSON parse failure      → wrapped error (not retryable)               │
│        │                                                                    │
│        ▼                                                                    │
│   Layer 3: Retry Logic (retry.go)                                           │
│   ├── Classifies error as retryable/non-retryable                           │
│   ├── Retries with exponential backoff if retryable                         │
│   └── Returns final error after max attempts                                │
│        │                                                                    │
│        ▼                                                                    │
│   Layer 2: Validation (validation.go)                                       │
│   ├── Invalid lat/long        → ValidationErrors                            │
│   ├── Invalid date format     → ValidationErrors                            │
│   └── Invalid output format   → ValidationErrors                            │
│        │                                                                    │
│        ▼                                                                    │
│   Layer 1: Main Orchestration (main.go)                                     │
│   ├── errors.Is(err, ErrNotFound)     → "Location not found..."             │
│   ├── errors.Is(err, ErrRateLimited)  → "Too many requests..."              │
│   ├── errors.As(err, &ValidationErrors{}) → Show all validation errors      │
│   └── Default                         → "An error occurred: ..."            │
│        │                                                                    │
│        ▼                                                                    │
│   Output: stderr + os.Exit(1)                                               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Package Structure

```
projects/00_weather_cli/
│
├── main.go                      # Entry point - orchestration only
├── go.mod                       # Module: github.com/username/weather-cli
├── go.sum
│
├── errors/                      # Package: weather-cli/errors
│   └── errors.go                # Sentinel errors + custom error types
│
├── config/                      # Package: weather-cli/config
│   └── config.go                # Configuration loading and validation
│
├── validation/                  # Package: weather-cli/validation
│   └── validation.go            # Input validation with aggregated errors
│
├── client/                      # Package: weather-cli/client
│   ├── weather.go               # Open-Meteo API client
│   ├── geocoder.go              # Nominatim API client
│   └── retry.go                 # Shared retry logic with backoff
│
├── formatter/                   # Package: weather-cli/formatter
│   └── formatter.go             # Table and JSON output formatting
│
├── types/                       # Package: weather-cli/types
│   └── types.go                 # Shared data structures
│
└── docs/                        # Architecture documentation (this folder)
    ├── 00_ARCHITECTURE.md
    ├── 01_ERRORS.md
    └── ...
```

---

## Design Principles Applied

### 1. Separation of Concerns
Each package has a single responsibility:
- `errors` - Error definitions only
- `config` - Configuration loading only
- `validation` - Input validation only
- `client` - API communication only
- `formatter` - Output formatting only
- `types` - Data structure definitions only
- `main` - Orchestration only

### 2. Dependency Inversion
- Client functions accept `context.Context` for timeout/cancellation
- Retry logic is generic and reusable
- Formatter accepts typed data, not raw API responses

### 3. Error Handling Patterns
- **Sentinel errors** for well-known conditions (ErrNotFound, ErrTimeout)
- **Custom error types** for rich context (WeatherAPIError, ValidationErrors)
- **Error wrapping** for stack traces and context preservation
- **errors.Is/As** for error inspection in main

### 4. Fail Fast, Fail Informatively
- Validate ALL inputs before making API calls
- Aggregate validation errors (show all at once, not one at a time)
- User-friendly error messages in main, technical details in logs

---

## Implementation Order

Recommended order for implementing modules:

1. **types/** - Define data structures first (no dependencies)
2. **errors/** - Define error types (no dependencies)
3. **config/** - Load configuration (depends on errors)
4. **validation/** - Input validation (depends on errors, types)
5. **client/retry.go** - Retry logic (depends on errors)
6. **client/weather.go** - Weather API (depends on errors, types, retry)
7. **client/geocoder.go** - Geocoder API (depends on errors, types, retry)
8. **formatter/** - Output formatting (depends on types)
9. **main.go** - Orchestration (depends on all)

---

## Success Criteria

After implementation, the CLI should:

1. ✅ Accept `--lat`, `--long`, `--start`, `--end`, `--format` flags
2. ✅ Validate all inputs and show ALL errors at once
3. ✅ Fetch weather data from Open-Meteo API
4. ✅ Fetch location data from Nominatim API
5. ✅ Retry transient failures (timeout, 429, 5xx) with backoff
6. ✅ Output combined data in table or JSON format
7. ✅ Print errors to stderr with user-friendly messages
8. ✅ Exit with code 0 on success, 1 on error
