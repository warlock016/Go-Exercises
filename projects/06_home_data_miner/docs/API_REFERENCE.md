# VictoriaMetrics API Reference

This document covers the VictoriaMetrics Prometheus-compatible API endpoints relevant to the Home Data Miner project.

---

## Base Configuration

| Property | Value |
|----------|-------|
| Base URL | `http://homeassistant.local:8428` |
| Protocol | HTTP (no TLS in local setup) |
| Authentication | None required (local network) |
| Content-Type | Responses are `application/json` |

---

## Endpoints Overview

### Query Endpoints (Core)

| Endpoint | Method | Purpose | Priority |
|----------|--------|---------|----------|
| `/api/v1/query` | GET | Instant query (current values) | ✅ Implemented |
| `/api/v1/query_range` | GET | Range query (historical data) | ✅ Implemented |
| `/api/v1/series` | GET | Find series matching a selector | 🔶 High |
| `/api/v1/labels` | GET | List all label names | 🔶 High |
| `/api/v1/label/<name>/values` | GET | List values for a label | 🔶 High |

### Export Endpoints (Data Extraction)

| Endpoint | Method | Purpose | Priority |
|----------|--------|---------|----------|
| `/api/v1/export` | GET | Export data as JSON lines | 🔷 Medium |
| `/api/v1/export/csv` | GET | Export data as CSV | 🔷 Medium |
| `/api/v1/export/native` | GET | Export raw binary format | ⚪ Low |

### Status & Metadata Endpoints

| Endpoint | Method | Purpose | Priority |
|----------|--------|---------|----------|
| `/api/v1/series/count` | GET | Total number of series | 🔷 Medium |
| `/api/v1/status/tsdb` | GET | Time series database stats | ⚪ Low |
| `/api/v1/status/active_queries` | GET | Currently running queries | ⚪ Low |
| `/api/v1/status/top_queries` | GET | Most frequent/slow queries | ⚪ Low |

### Delete Endpoints (Destructive)

| Endpoint | Method | Purpose | Priority |
|----------|--------|---------|----------|
| `/api/v1/admin/tsdb/delete_series` | POST | Delete time series | ⚠️ Careful |

**Source:** [VictoriaMetrics API Examples](https://docs.victoriametrics.com/victoriametrics/url-examples/)

---

## Endpoint Details

### 1. Instant Query - `/api/v1/query`

Evaluates a PromQL expression at a single point in time.

**Request:**
```
GET /api/v1/query?query=<promql>&time=<timestamp>
```

**Parameters:**

| Parameter | Required | Description | Example |
|-----------|----------|-------------|---------|
| `query` | Yes | PromQL expression | `W_value{entity_id="total_power"}` |
| `time` | No | Evaluation timestamp (default: now) | `1765645456` or `2024-01-15T10:00:00Z` |

**Response Structure:**
```json
{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "__name__": "W_value",
          "db": "home_assistant",
          "domain": "sensor",
          "entity_id": "total_power"
        },
        "value": [1765645456, "55"]
      }
    ]
  },
  "stats": {
    "seriesFetched": "1",
    "executionTimeMsec": 3
  }
}
```

**Key Points:**
- `value` is an array: `[unix_timestamp, "string_value"]`
- The value is always a **string**, even for numbers
- `resultType` is typically `"vector"` for instant queries

---

### 2. Range Query - `/api/v1/query_range`

Evaluates a PromQL expression over a time range.

**Request:**
```
GET /api/v1/query_range?query=<promql>&start=<time>&end=<time>&step=<duration>
```

**Parameters:**

| Parameter | Required | Description | Example |
|-----------|----------|-------------|---------|
| `query` | Yes | PromQL expression | `W_value{entity_id="total_power"}` |
| `start` | Yes | Start timestamp | `1765600000` or `-1h` |
| `end` | Yes | End timestamp | `1765645456` or `now` |
| `step` | Yes | Resolution step | `60s`, `5m`, `1h` |

**Response Structure:**
```json
{
  "status": "success",
  "data": {
    "resultType": "matrix",
    "result": [
      {
        "metric": {
          "__name__": "W_value",
          "entity_id": "total_power"
        },
        "values": [
          [1765600000, "52.3"],
          [1765600060, "53.1"],
          [1765600120, "51.8"]
        ]
      }
    ]
  }
}
```

**Key Points:**
- `resultType` is `"matrix"` for range queries
- `values` is an array of `[timestamp, "value"]` pairs
- `step` determines data point frequency

---

### 3. Series Discovery - `/api/v1/series`

Find time series that match certain label selectors.

**Request:**
```
GET /api/v1/series?match[]=<selector>&start=<time>&end=<time>
```

**Parameters:**

| Parameter | Required | Description | Example |
|-----------|----------|-------------|---------|
| `match[]` | Yes (1+) | Series selector(s) | `W_value{db="home_assistant"}` |
| `start` | No | Start time for search | `-7d` |
| `end` | No | End time for search | `now` |

**Note:** `match[]` can be specified multiple times to match multiple selectors.

**Response Structure:**
```json
{
  "status": "success",
  "data": [
    {
      "__name__": "W_value",
      "db": "home_assistant",
      "domain": "sensor",
      "entity_id": "total_power"
    },
    {
      "__name__": "W_value",
      "db": "home_assistant",
      "domain": "sensor",
      "entity_id": "kitchen_power"
    }
  ]
}
```

**Key Points:**
- Returns label sets, not values
- Useful for discovering what metrics exist
- Each item in `data` is a map of label name → value

---

### 4. Label Names - `/api/v1/labels`

Returns a list of all label names in the database.

**Request:**
```
GET /api/v1/labels?start=<time>&end=<time>
```

**Parameters:**

| Parameter | Required | Description |
|-----------|----------|-------------|
| `start` | No | Start time for search |
| `end` | No | End time for search |

**Response Structure:**
```json
{
  "status": "success",
  "data": [
    "__name__",
    "db",
    "domain",
    "entity_id",
    "friendly_name"
  ]
}
```

---

### 5. Label Values - `/api/v1/label/<name>/values`

Returns all values for a specific label name.

**Request:**
```
GET /api/v1/label/entity_id/values?start=<time>&end=<time>
```

**Parameters:**

| Parameter | Required | Description |
|-----------|----------|-------------|
| `start` | No | Start time for search |
| `end` | No | End time for search |

**Response Structure:**
```json
{
  "status": "success",
  "data": [
    "total_power",
    "kitchen_power",
    "living_room_temp",
    "bedroom_humidity"
  ]
}
```

---

### 6. Export JSON - `/api/v1/export`

Exports raw time series data as JSON lines (one JSON object per line).

**Request:**
```
GET /api/v1/export?match[]=<selector>&start=<time>&end=<time>
```

**Parameters:**

| Parameter | Required | Description | Example |
|-----------|----------|-------------|---------|
| `match[]` | Yes | Series selector(s) | `W_value{entity_id="total_power"}` |
| `start` | No | Start timestamp | `-1d` |
| `end` | No | End timestamp | `now` |
| `max_rows_per_line` | No | Limit values per JSON line | `1000` |

**Response Structure (JSON lines, not array):**
```json
{"metric":{"__name__":"W_value","entity_id":"total_power"},"values":[52.3,53.1,51.8],"timestamps":[1765600000,1765600060,1765600120]}
{"metric":{"__name__":"W_value","entity_id":"kitchen_power"},"values":[10.5,11.2],"timestamps":[1765600000,1765600060]}
```

**Key Points:**
- Response is **not** wrapped in `{"status":"success","data":...}`
- Each line is a separate JSON object (JSON Lines format)
- Good for streaming large datasets
- Values and timestamps are parallel arrays

**curl Example:**
```bash
curl -G 'http://homeassistant.local:8428/api/v1/export' \
  --data-urlencode 'match[]=W_value{db="home_assistant"}' \
  -d 'start=-1h' -d 'end=now'
```

---

### 7. Export CSV - `/api/v1/export/csv`

Exports data in CSV format with customizable columns.

**Request:**
```
GET /api/v1/export/csv?match[]=<selector>&format=<columns>&start=<time>&end=<time>
```

**Parameters:**

| Parameter | Required | Description | Example |
|-----------|----------|-------------|---------|
| `match[]` | Yes | Series selector(s) | `W_value{entity_id="total_power"}` |
| `format` | No | Column format | `__name__,__value__,__timestamp__:unix_s` |
| `start` | No | Start timestamp | `-1d` |
| `end` | No | End timestamp | `now` |

**Format Placeholders:**

| Placeholder | Description |
|-------------|-------------|
| `__name__` | Metric name |
| `__value__` | Sample value |
| `__timestamp__:unix_s` | Unix timestamp in seconds |
| `__timestamp__:unix_ms` | Unix timestamp in milliseconds |
| `__timestamp__:rfc3339` | RFC3339 formatted time |
| `<label_name>` | Value of specific label (e.g., `entity_id`) |

**Response Example:**
```csv
__name__,entity_id,__value__,__timestamp__:unix_s
W_value,total_power,52.3,1765600000
W_value,total_power,53.1,1765600060
```

**curl Example:**
```bash
curl -G 'http://homeassistant.local:8428/api/v1/export/csv' \
  --data-urlencode 'match[]=W_value{db="home_assistant"}' \
  -d 'format=__name__,entity_id,__value__,__timestamp__:rfc3339' \
  -d 'start=-1h'
```

---

### 8. Series Count - `/api/v1/series/count`

Returns the total number of time series in the database.

**Request:**
```
GET /api/v1/series/count
```

**Parameters:** None required.

**Response Structure:**
```json
{
  "status": "success",
  "data": [
    12345
  ]
}
```

**Use Case:** Quick health check or capacity monitoring.

---

### 9. TSDB Status - `/api/v1/status/tsdb`

Returns statistics about the time series database.

**Request:**
```
GET /api/v1/status/tsdb
```

**Response Structure:**
```json
{
  "status": "success",
  "data": {
    "totalSeries": 12345,
    "totalLabelValuePairs": 67890,
    "seriesCountByMetricName": [
      {"name": "W_value", "count": 5},
      {"name": "V_value", "count": 3}
    ],
    "seriesCountByLabelValuePair": [
      {"name": "db=home_assistant", "count": 10}
    ],
    "labelValueCountByLabelName": [
      {"name": "entity_id", "count": 15}
    ]
  }
}
```

**Use Case:** Understanding database structure, finding most common metrics/labels.

---

### 10. Top Queries - `/api/v1/status/top_queries`

Lists the most frequently executed or slowest queries.

**Request:**
```
GET /api/v1/status/top_queries?topN=<count>&maxLifetime=<duration>
```

**Parameters:**

| Parameter | Required | Description | Default |
|-----------|----------|-------------|---------|
| `topN` | No | Number of queries to return | `10` |
| `maxLifetime` | No | Time window to consider | `10m` |

**Response Structure:**
```json
{
  "topByCount": [
    {"query": "W_value{entity_id=\"total_power\"}", "count": 150}
  ],
  "topByAvgDuration": [
    {"query": "sum(W_value)", "avgDurationSeconds": 0.5}
  ],
  "topBySumDuration": [
    {"query": "W_value", "sumDurationSeconds": 12.5}
  ]
}
```

**Use Case:** Performance monitoring, identifying expensive queries.

---

## PromQL Quick Reference

### Basic Selectors

| Pattern | Meaning | Example |
|---------|---------|---------|
| `metric_name` | All series with this name | `W_value` |
| `{label="value"}` | Filter by label | `{entity_id="total_power"}` |
| `metric{label="value"}` | Combined | `W_value{entity_id="total_power"}` |

### Label Matchers

| Operator | Meaning | Example |
|----------|---------|---------|
| `=` | Exact match | `entity_id="total_power"` |
| `!=` | Not equal | `domain!="binary_sensor"` |
| `=~` | Regex match | `entity_id=~".*power.*"` |
| `!~` | Regex not match | `entity_id!~"test.*"` |

### Multiple Labels

```
W_value{db="home_assistant", domain="sensor", entity_id="total_power"}
```

All conditions must match (AND logic).

---

## Time Formats

VictoriaMetrics accepts multiple time formats:

| Format | Example | Description |
|--------|---------|-------------|
| Unix timestamp | `1765645456` | Seconds since epoch |
| RFC3339 | `2024-01-15T10:30:00Z` | ISO format with timezone |
| Relative | `-7d` | 7 days ago |
| Relative | `-1h` | 1 hour ago |
| Relative | `now` | Current time |

### Duration Suffixes

| Suffix | Meaning |
|--------|---------|
| `s` | Seconds |
| `m` | Minutes |
| `h` | Hours |
| `d` | Days |
| `w` | Weeks |
| `y` | Years |

---

## Error Responses

### API Error Response

```json
{
  "status": "error",
  "errorType": "bad_data",
  "error": "invalid parameter 'query': parse error at char 5: unexpected character"
}
```

### Common Error Types

| errorType | Meaning |
|-----------|---------|
| `bad_data` | Invalid query syntax or parameters |
| `timeout` | Query execution timeout |
| `canceled` | Query was canceled |
| `execution` | Error during query execution |

---

## Metrics in Your Home Assistant Setup

Based on your curl commands, your setup has:

### Known Metrics

| Metric Name | Labels | Description |
|-------------|--------|-------------|
| `W_value` | `db`, `domain`, `entity_id` | Power consumption in Watts |

### Known Label Values

| Label | Known Values |
|-------|--------------|
| `db` | `home_assistant` |
| `domain` | `sensor` |
| `entity_id` | `total_power` |

### Discovery Queries

To discover more:
```
# Find all metric names
GET /api/v1/label/__name__/values

# Find all entity IDs
GET /api/v1/label/entity_id/values

# Find all series with power data
GET /api/v1/series?match[]={db="home_assistant"}
```

---

## Rate Limits and Best Practices

### Query Efficiency

1. **Be specific with selectors** - More labels = faster query
   ```
   # Slower (scans all)
   W_value

   # Faster (indexed lookup)
   W_value{entity_id="total_power"}
   ```

2. **Limit time ranges** - Smaller ranges = less data processed
   ```
   # Avoid unbounded
   GET /api/v1/series?match[]={...}

   # Better
   GET /api/v1/series?match[]={...}&start=-7d&end=now
   ```

3. **Use appropriate step** - For range queries, match step to your needs
   ```
   # 1 week of data with 1-minute steps = 10,080 points (excessive)
   # 1 week of data with 1-hour steps = 168 points (reasonable)
   ```

### Connection Management

- Reuse HTTP client (connection pooling)
- Set reasonable timeouts (10-30 seconds typical)
- Consider retries for transient failures
