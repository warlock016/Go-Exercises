# CSV Processor - Architecture Documentation

## Overview

A modular, extensible CSV processor designed for weather data transformation. Built with provider-agnostic architecture to support multiple data sources (WeatherCloud initially, expandable to others).

### Design Philosophy

- **Modularity**: Single responsibility per package
- **Configurability**: Provider settings externalized to YAML/JSON
- **Extensibility**: Interfaces enable new providers/outputs without code changes
- **Error Transparency**: Collect all errors with positional context, report comprehensively

---

## Package Structure

```
01_csv_processor/
├── main.go                          # CLI entry point, orchestration
├── ARCHITECTURE.md                  # This document
├── config/                          # Configuration management
│   ├── config.go                    # Config loading & validation
│   └── providers/                   # Provider-specific profiles
│       └── weathercloud.yaml
├── errors/                          # Error types & accumulator
│   └── errors.go
├── validator/                       # File I/O, encoding, structural checks
│   └── validator.go
├── parser/                          # Type conversion (datetime, float)
│   └── parser.go
├── formatter/                       # Output shaping (JSON, CSV, table)
│   └── formatter.go
├── output/                          # Sink implementations (file, stdout)
│   └── output.go
├── types/                           # Shared data structures
│   └── types.go
└── testdata/                        # Test fixtures
```

---

## Pipeline Architecture

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   config    │───►│  validator  │───►│   parser    │───►│  formatter  │───►│   output    │
│  (Stage 0)  │    │  (Stage 1)  │    │  (Stage 2)  │    │  (Stage 3)  │    │  (Stage 4)  │
└─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘
      │                  │                  │                  │                  │
      ▼                  ▼                  ▼                  ▼                  ▼
   *Config            *RawData         *ParsedData     *FormattedOutput      bytes
```

### Data Flow

| Stage | Input | Output | Responsibility |
|-------|-------|--------|----------------|
| 0: Config | CLI args, YAML file | `*Config` | Load provider profile, validate settings |
| 1: Validator | File path, `*Config` | `*RawData` | Open file, decode encoding, read rows, check structure |
| 2: Parser | `*RawData`, `*Config` | `*ParsedData` | Convert strings to typed values (datetime, float64) |
| 3: Formatter | `*ParsedData`, format string | `*FormattedOutput` | Shape data as JSON, CSV, or table |
| 4: Output | `*FormattedOutput`, sink config | Written bytes | Write to file or stdout |

---

## Package Responsibilities

### `config/`
Loads and validates provider configuration from YAML/JSON files. Defines `ProviderConfig` struct with encoding, delimiter, datetime formats, and column schemas. Entry point: `Load(path string) (*Config, error)`.

### `errors/`
Defines error types for accumulating processing errors with positional context (line, column, value). Provides `ProcessingErrors` with `Add()`, `HasFatalErrors()`, `HasWarnings()`, and `Summary()` methods. Supports stage-aware error classification.

### `validator/`
Handles file I/O and structural validation. Opens files with correct encoding (UTF-8, UTF-16LE), configures CSV reader (delimiter, comment char), reads all rows into memory, and validates row structure (field count consistency). Returns `*RawData` containing headers and raw string rows.

### `parser/`
Converts raw string data to typed values. Parses datetime column using configured format(s), converts numeric columns to float64, handles empty values and format variations (comma thousands separators). Returns `*ParsedData` with typed columns.

### `formatter/`
Shapes parsed data for output. Supports multiple formats: JSON (structured), CSV (re-serialized), table (human-readable). Returns `*FormattedOutput` containing the formatted string.

### `output/`
Writes formatted content to configured sink. Supports file output (with path) and stdout. Implements `OutputWriter` interface for future extensibility (HTTP, S3, etc.).

### `types/`
Shared data structures passed between stages: `RawData`, `ParsedData`, `FormattedOutput`. Defines the contracts between pipeline stages.

---

## Core Types

### Config Types

```go
// config/config.go

type Config struct {
    Provider   ProviderConfig
    InputPath  string          // CSV file to process
    OutputPath string          // Output destination ("" = stdout)
    Format     string          // "json", "csv", "table"
}

type ProviderConfig struct {
    Name               string         `yaml:"name"`
    Encoding           string         `yaml:"encoding"`            // "UTF-8", "UTF-16LE"
    Delimiter          string         `yaml:"delimiter"`           // ";" or ","
    Comment            string         `yaml:"comment"`             // "#" or ""
    SkipRows           int            `yaml:"skip_rows"`           // Rows to skip before header
    ThousandsSeparator string         `yaml:"thousands_separator"` // "," for "1,234.56"
    SchemaMode         string         `yaml:"schema_mode"`         // "dynamic" or "strict"
    Datetime           DatetimeConfig `yaml:"datetime"`
    TypeInference      TypeInference  `yaml:"type_inference"`
    Columns            []ColumnConfig `yaml:"columns"`             // Optional in dynamic mode
}

type DatetimeConfig struct {
    Detection string   `yaml:"detection"` // "name_pattern", "index", "first_column"
    Patterns  []string `yaml:"patterns"`  // Column name patterns for datetime
    Index     int      `yaml:"index"`     // Used when detection: index
    Formats   []string `yaml:"formats"`   // Go time layouts (tried in order)
}

type TypeInference struct {
    Default       string   `yaml:"default"`        // Default type for columns
    StringColumns []string `yaml:"string_columns"` // Patterns to keep as string
}

type ColumnConfig struct {
    Name          string   `yaml:"name"`
    Type          string   `yaml:"type"`           // "datetime", "float64", "string"
    Required      bool     `yaml:"required"`       // Only enforced in strict mode
    Aliases       []string `yaml:"aliases"`        // Alternative column names
    AvailableFrom string   `yaml:"available_from"` // Documentation (e.g., "2025-03")
}
```

### Pipeline Data Types

```go
// types/types.go

// RawData - Output of validator (Stage 1), input to parser (Stage 2)
type RawData struct {
    Headers []string   // Column names from first row
    Rows    [][]string // Raw string values (excluding header)
    Source  string     // File path for error reporting
}

// ParsedData - Output of parser (Stage 2), input to formatter (Stage 3)
type ParsedData struct {
    Headers    []string             // Column names (preserved order)
    Timestamps []time.Time          // Parsed datetime values
    Columns    map[string][]float64 // Column name → numeric values
    RowCount   int                  // Successfully parsed rows
}

// FormattedOutput - Output of formatter (Stage 3), input to output (Stage 4)
type FormattedOutput struct {
    Content  string // Formatted string content
    Format   string // "json", "csv", "table"
    RowCount int    // Rows included in output
}
```

### Error Types

```go
// errors/errors.go

type FieldError struct {
    Stage   string // "config", "validation", "parsing", "formatting", "output"
    Message string
    Line    int    // 1-indexed line number (0 = N/A)
    Column  int    // 0-indexed column position
    Value   string // The problematic value (truncated for display)
}

type ProcessingErrors struct {
    Errors   []FieldError // Fatal or significant errors
    Warnings []FieldError // Non-fatal issues (processing continued)
}

// Methods
func (e *ProcessingErrors) Add(err FieldError)
func (e *ProcessingErrors) AddWarning(err FieldError)
func (e *ProcessingErrors) HasFatalErrors() bool
func (e *ProcessingErrors) HasWarnings() bool
func (e *ProcessingErrors) Summary() string  // "3 errors, 12 warnings"
func (e *ProcessingErrors) Error() string    // Full formatted output
```

---

## Stage Gates

Each stage has criteria that determine whether processing continues or stops.

| Stage | Continue If | Stop If (Fatal) |
|-------|-------------|-----------------|
| 0: Config | Config loaded successfully | File missing, YAML parse error, required fields empty |
| 1: Validator | At least 1 valid row | File unreadable, 0 rows after filtering |
| 2: Parser | At least 1 parseable row | 0 rows with valid datetime AND valid data |
| 3: Formatter | Format recognized | Unknown format string |
| 4: Output | Write successful | Permission denied, disk full, etc. |

### Non-Fatal Errors (Warnings)

These are collected but do not stop processing:

- **Malformed rows**: Wrong field count (skipped, warning logged)
- **Unparseable datetime**: Invalid format (row skipped or null timestamp)
- **Unparseable float**: Non-numeric string (value becomes 0 or NaN)
- **Empty required field**: Missing value in required column (warning)

### Error Aggregation Pattern

```go
func (v *Validator) Validate(path string, cfg *Config) (*RawData, *ProcessingErrors) {
    errs := &ProcessingErrors{}

    // ... processing ...

    if len(row) != expectedCols {
        errs.AddWarning(FieldError{
            Stage:   "validation",
            Message: fmt.Sprintf("wrong field count: got %d, want %d", len(row), expectedCols),
            Line:    lineNum,
            Column:  0,
        })
        continue // Skip malformed row, don't stop
    }

    // ... more processing ...

    // Stage gate check
    if len(rawData.Rows) == 0 {
        errs.Add(FieldError{
            Stage:   "validation",
            Message: "no valid rows after filtering",
            Line:    0,
            Column:  0,
        })
    }

    return rawData, errs
}
```

---

## Configuration Schema

### Schema Strategy: Dynamic vs Strict

WeatherCloud and similar providers exhibit **schema drift** - columns appear/disappear over time (e.g., UV index sensor added March 2025). The configuration supports two modes:

| Mode | Behavior | Use Case |
|------|----------|----------|
| `dynamic` (default) | Discover columns from file headers; validate types only | Production - handles schema drift gracefully |
| `strict` | Require exact column match from config | Testing, validation, known-stable schemas |

### Provider Configuration File (YAML)

```yaml
# config/providers/weathercloud.yaml

name: weathercloud
encoding: UTF-16LE              # File encoding (UTF-8, UTF-16LE, UTF-16BE)
delimiter: ";"                  # Field separator
comment: "#"                    # Comment line prefix (empty = none)
skip_rows: 0                    # Rows to skip before header
thousands_separator: ","        # For numeric parsing (e.g., "1,234.56")

# Schema handling
schema_mode: dynamic            # "dynamic" or "strict"

# Datetime detection (applies to both modes)
datetime:
  detection: name_pattern       # "name_pattern", "index", or "first_column"
  patterns:                     # Column name patterns that indicate datetime
    - "date"
    - "time"
    - "timestamp"
  index: 0                      # Used when detection: index
  formats:                      # Go time layouts (tried in order)
    - "02/01/2006 15:04:05"     # DD/MM/YYYY HH:MM:SS
    - "2006-01-02 15:04:05"     # YYYY-MM-DD HH:MM:SS

# Type inference rules (for dynamic mode)
type_inference:
  default: float64              # Default type for non-datetime columns
  string_columns:               # Column name patterns to keep as string
    - "name"
    - "description"
    - "notes"

# Column definitions (for strict mode, optional hints for dynamic mode)
columns:
  - name: timestamp             # Match by name (flexible) or index (strict)
    type: datetime
    required: true              # Only enforced in strict mode

  - name: temperature_indoor
    type: float64
    aliases:                    # Alternative column names (schema drift support)
      - "temp_indoor"
      - "indoor_temp"

  - name: uv_index              # New sensor - may not exist in older files
    type: float64
    required: false
    available_from: "2025-03"   # Documentation only (not enforced)
```

### Dynamic Mode Behavior

In `dynamic` mode, the parser:

1. **Reads headers from file** - Column names come from the CSV, not config
2. **Detects datetime column** - Uses `datetime.detection` strategy
3. **Infers types** - All non-datetime columns default to `float64` (or configured default)
4. **Applies aliases** - If a column definition has `aliases`, matches any of them
5. **Ignores missing columns** - Columns in config but not in file are skipped (warning)
6. **Accepts extra columns** - Columns in file but not in config are included

### Strict Mode Behavior

In `strict` mode, the parser:

1. **Validates exact schema match** - All `required: true` columns must exist
2. **Uses column index** - Position matters, not just name
3. **Rejects extra columns** - Unexpected columns trigger error
4. **Enforces types** - Type mismatches are fatal

### Supported Values

| Field | Options |
|-------|---------|
| `encoding` | `UTF-8`, `UTF-16LE`, `UTF-16BE` |
| `delimiter` | Any single character: `;`, `,`, `\t` |
| `schema_mode` | `dynamic`, `strict` |
| `type` | `datetime`, `float64`, `string` |
| `datetime.detection` | `name_pattern`, `index`, `first_column` |

---

## CLI Interface

```bash
csvprocessor [flags]

Flags:
  -input string       Input CSV file path (required)
  -provider string    Provider config name (default: "weathercloud")
  -config string      Path to provider config directory (default: "./config/providers")
  -output string      Output file path (default: stdout)
  -format string      Output format: json, csv, table (default: "json")
  -verbose            Show warnings and processing stats
  -help               Show help
```

### Usage Examples

```bash
# Basic usage - JSON to stdout
csvprocessor -input ./data/weather.csv

# CSV output to file
csvprocessor -input ./data/weather.csv -output ./output/clean.csv -format csv

# Table format with verbose warnings
csvprocessor -input ./data/weather.csv -format table -verbose

# Custom provider config
csvprocessor -input ./data/other.csv -provider myconfig -config ./my-configs/
```

---

## Extension Guide

### Adding a New Provider

1. Create config file in `config/providers/`:
   ```yaml
   # config/providers/newprovider.yaml
   name: newprovider
   encoding: UTF-8
   delimiter: ","
   # ... rest of config
   ```

2. Use with CLI:
   ```bash
   csvprocessor -input data.csv -provider newprovider
   ```

No code changes required.

### Adding a New Output Format

1. Implement format function in `formatter/formatter.go`:
   ```go
   func formatXML(data *ParsedData) (string, error) {
       // ... implementation
   }
   ```

2. Register in format switch:
   ```go
   case "xml":
       return formatXML(data)
   ```

### Adding a New Output Sink

1. Implement `OutputWriter` interface:
   ```go
   type HTTPWriter struct {
       Endpoint string
       Client   *http.Client
   }

   func (w *HTTPWriter) Write(output *FormattedOutput) error {
       // POST to endpoint
   }
   ```

2. Register in output factory based on config.

---

## Error Messages Reference

### Config Stage (Stage 0)
- `"config file not found: {path}"` - Provider YAML missing
- `"invalid YAML: {parse_error}"` - Malformed config file
- `"missing required field: {field}"` - Required config field empty

### Validation Stage (Stage 1)
- `"file not found: {path}"` - Input CSV missing
- `"encoding error: {details}"` - Failed to decode file
- `"wrong field count: got {n}, want {m}"` - Row has incorrect columns (warning)
- `"no valid rows after filtering"` - All rows malformed (fatal)

### Parsing Stage (Stage 2)
- `"invalid datetime: {value}"` - Timestamp parse failed (warning)
- `"invalid float: {value}"` - Numeric parse failed (warning)
- `"no parseable rows"` - All rows failed conversion (fatal)

### Formatting Stage (Stage 3)
- `"unsupported format: {format}"` - Unknown format string (fatal)

### Output Stage (Stage 4)
- `"write error: {details}"` - File write failed (fatal)

---

## Processing Modes (Current & Future)

### Current: Single File Mode

The initial implementation processes one CSV file at a time:

```bash
csvprocessor -input ./data/november.csv -output ./output/november.json
```

**Characteristics:**
- Single input path → single output
- Schema discovered from file headers
- Errors/warnings scoped to one file

### Future: Folder Mode (Deferred)

A future enhancement will support processing multiple files sequentially:

```bash
csvprocessor -input ./data/ -output ./output/ -mode folder
```

**Planned Characteristics:**
- Process all matching files in directory
- Each file parsed independently (schema may vary)
- Aggregated error report across all files
- Options for output: merged single file vs. individual files

**Design Considerations for Folder Mode:**
- **Schema union**: Combine columns from all files (superset schema)
- **Schema intersection**: Only columns present in ALL files
- **Per-file output**: Each input file → corresponding output file
- **Merged output**: All files → single combined output

**Implementation will require:**
- File discovery (glob patterns, extension filtering)
- Progress reporting for multi-file operations
- Partial failure handling (continue on single file error?)
- Memory management for large folder processing

This is documented for future reference but **not in scope for initial implementation**.

---

## Implementation Order

Recommended implementation sequence:

1. **errors/** - Foundation for all error handling
2. **types/** - Shared data structures
3. **config/** - Configuration loading
4. **validator/** - File reading and structural validation
5. **parser/** - Type conversion
6. **formatter/** - Output shaping
7. **output/** - Sink implementations
8. **main.go** - CLI orchestration

Each package can be implemented and tested independently before integration.
