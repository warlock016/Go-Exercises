# Module: Config

## Purpose

Load and validate application configuration from environment variables and optional `.env` files. This module centralizes all configuration logic, ensuring API keys and settings are available before any API calls are made.

---

## Package Location

```
config/config.go
```

---

## Dependencies

| Package | Purpose |
|---------|---------|
| `errors` (internal) | `ErrUnauthorized` for missing API key |

---

## Public API

### Types

```go
// Config holds all application configuration
type Config struct {
    GeocodeAPIKey string // API key for Nominatim geocoding service
}
```

### Functions

```go
// Load reads configuration from environment variables and optional .env file
// Returns an error if required configuration is missing
func Load() (*Config, error)

// LoadFromEnvFile attempts to load a .env file from the given path
// This is optional - returns nil error if file doesn't exist
func LoadFromEnvFile(path string) error
```

---

## Behavior Specification

### Load()

**Input:** None

**Output:** `(*Config, error)`

**Behavior:**
1. Attempt to load `.env` file from current directory (optional, ignore if missing)
2. Read `GEOCODE_API_KEY` from environment
3. Validate that required keys are present and non-empty
4. Return populated Config struct or error

**Flow Diagram:**
```
Load()
   │
   ├──► LoadFromEnvFile(".env")
   │         │
   │         ├── File exists → Parse and set env vars
   │         └── File missing → Continue (not an error)
   │
   ├──► os.Getenv("GEOCODE_API_KEY")
   │         │
   │         ├── Empty → Return error wrapping ErrUnauthorized
   │         └── Present → Store in Config
   │
   └──► Return &Config{...}, nil
```

**Edge Cases:**

| Case | Behavior |
|------|----------|
| `.env` file missing | Continue without error (optional) |
| `.env` file malformed | Return descriptive error |
| `GEOCODE_API_KEY` empty | Return error: "GEOCODE_API_KEY not set" |
| `GEOCODE_API_KEY` whitespace only | Trim and treat as empty |
| Environment variable set | Use it, even if `.env` also exists |

---

### LoadFromEnvFile(path string)

**Input:** `path string` - Path to .env file

**Output:** `error`

**Behavior:**
1. Check if file exists at path
2. If not exists, return `nil` (not an error)
3. If exists, parse file line by line
4. For each line matching `KEY=VALUE`, call `os.Setenv(KEY, VALUE)`
5. Ignore empty lines and lines starting with `#`
6. Return error only if file exists but cannot be parsed

**.env File Format:**
```
# Comment lines start with #
GEOCODE_API_KEY=your_api_key_here

# Empty lines are ignored
ANOTHER_KEY=another_value
```

**Parsing Rules:**
- Split on first `=` only (values can contain `=`)
- Trim whitespace from key and value
- Handle quoted values: `KEY="value with spaces"` → `value with spaces`
- Skip lines without `=`

---

## Error Messages

| Scenario | Error Message |
|----------|---------------|
| Missing API key | `"configuration error: GEOCODE_API_KEY environment variable is not set"` |
| Malformed .env | `"configuration error: failed to parse .env file: {details}"` |
| File read error | `"configuration error: failed to read .env file: {details}"` |

---

## Environment Variables

| Variable | Required | Description | Default |
|----------|----------|-------------|---------|
| `GEOCODE_API_KEY` | Yes | API key for Nominatim geocoding | (none) |

---

## Useful Packages

| Package | Why Useful |
|---------|------------|
| `os` | `os.Getenv()`, `os.Setenv()`, `os.Open()` |
| `bufio` | `bufio.Scanner` for line-by-line file reading |
| `strings` | `strings.TrimSpace()`, `strings.SplitN()`, `strings.HasPrefix()` |
| `fmt` | Error message formatting |

**Note:** You should implement .env parsing yourself rather than using external packages like `godotenv`. This is a learning exercise!

---

## Test Cases

| Scenario | Setup | Expected Result |
|----------|-------|-----------------|
| Happy path | `GEOCODE_API_KEY=abc123` in env | `Config{GeocodeAPIKey: "abc123"}`, nil error |
| Missing key | No env var set | Error containing "GEOCODE_API_KEY" |
| Empty key | `GEOCODE_API_KEY=""` | Error (empty is not valid) |
| Whitespace key | `GEOCODE_API_KEY="   "` | Error (whitespace-only is not valid) |
| .env file missing | No .env file | Fallback to env vars, no error |
| .env file exists | Valid .env with key | Config populated from file |
| .env with comments | `# comment\nKEY=value` | Comments ignored, KEY parsed |
| .env with quotes | `KEY="value"` | Quotes stripped from value |
| Env overrides .env | Both set, different values | Environment variable takes precedence |

---

## Implementation Hints

1. **Order of precedence:** Environment variables should override `.env` file values. Load `.env` first, then let existing env vars take precedence.

2. **Trimming whitespace:** Always trim keys and values. `KEY = value ` should become `KEY` and `value`.

3. **Handling quotes:** Check if value starts and ends with the same quote character (`"` or `'`), then strip both.

4. **SplitN for parsing:** Use `strings.SplitN(line, "=", 2)` to split on only the FIRST `=`. This handles values like `KEY=a=b=c`.

5. **File existence check:** Use `os.Stat()` to check if file exists before trying to open:
   ```go
   if _, err := os.Stat(path); os.IsNotExist(err) {
       return nil // File doesn't exist, not an error
   }
   ```

6. **Error wrapping:** Wrap the sentinel error for proper classification:
   ```go
   return nil, fmt.Errorf("GEOCODE_API_KEY not set: %w", errors.ErrUnauthorized)
   ```
   This allows `errors.Is(err, ErrUnauthorized)` to work in main.

7. **Testing environment variables:** In tests, use `t.Setenv()` (Go 1.17+) which automatically restores the original value after the test:
   ```go
   func TestLoad_MissingKey(t *testing.T) {
       t.Setenv("GEOCODE_API_KEY", "")
       _, err := Load()
       // ...
   }
   ```

---

## Example Usage

```go
// In main.go
cfg, err := config.Load()
if err != nil {
    fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    os.Exit(1)
}

// Pass config to clients
geocoder := client.NewGeocoder(cfg.GeocodeAPIKey)
```

---

## Future Extensions (Not Required Now)

These are mentioned for awareness but NOT part of the current implementation:

- `WEATHER_TIMEOUT` - Custom timeout duration
- `RETRY_MAX_ATTEMPTS` - Override default retry count
- `LOG_LEVEL` - Debug/Info/Warn/Error logging level
