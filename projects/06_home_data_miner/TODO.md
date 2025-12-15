# Home Data Miner - Implementation TODO

## Documentation Reference

| Document | Purpose | Location |
|----------|---------|----------|
| Architecture Guide | System design, data flow, component structure | [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) |
| API Reference | VictoriaMetrics endpoints, parameters, responses | [docs/API_REFERENCE.md](docs/API_REFERENCE.md) |
| Go Patterns | HTTP client patterns, code examples, idioms | [docs/PATTERNS.md](docs/PATTERNS.md) |

---

## Phase 0: Package Structure ✅ COMPLETE

### 0.1 Create client/ Package Directory
- [x] Create `client/` directory
- [x] Create `client/client.go` (empty placeholder)
- [x] Create `client/types.go` (empty placeholder)
- [x] Create `client/queries.go` (empty placeholder)
- [x] Create `client/series.go` (empty placeholder)
- [x] Create `client/errors.go` (empty placeholder)

### Current File Structure

```
06_home_data_miner/
├── main.go              # Entry point, orchestration (~130 lines)
├── .env                 # Configuration
├── client/              # API client package
│   ├── client.go        # HAClient, NewHAClient, LoadEnv (~57 lines)
│   ├── types.go         # Response types + Labels, Status, Series (~50 lines)
│   ├── queries.go       # HAQuery, RequestQuery, RequestQueryRange (~122 lines)
│   ├── series.go        # RequestStatus, RequestLabels, etc. (~63 lines)
│   └── errors.go        # Custom error types (pending)
├── docs/
│   ├── ARCHITECTURE.md
│   ├── API_REFERENCE.md
│   └── PATTERNS.md
├── cache/               # Sample API responses
└── TODO.md
```

---

## Phase 1: Foundation ✅ COMPLETE

### 1.1 Fix Basic Request Construction
- [x] Update URL building to match working curl pattern
- [x] Add query parameters using `url.Values`
- [x] Remove unnecessary `Content-Type` header for GET requests
- [x] Test basic `/api/v1/query` endpoint works

### 1.2 Create Client Configuration
- [x] `HAClient` struct exists with `http.Client` and `baseURL`
- [x] `NewClient(baseURL, timeout)` constructor implemented
- [x] Create `client/` package structure (Phase 0)

### 1.3 URL Helper Functions
- [x] `RequestQuery()` - builds instant query URLs
- [x] `RequestQueryRange()` - builds range query URLs with proper param separation

---

## Phase 1.5: Code Migration ✅ COMPLETE

Move code from `main.go` to `client/` package files.

### 1.5.1 Populate `client/types.go` ✅
- [x] Move `VMEnvelope` struct
- [x] Move `QueryResponse` struct
- [x] Move `InstantResponse`, `InstantMetric` structs
- [x] Move `RangeResponse`, `RangeMetric` structs
- [x] Move `Metric` struct
- [x] Add `Labels`, `Status`, `Series` types for discovery endpoints

### 1.5.2 Populate `client/client.go` ✅
- [x] `HAClient` struct with exported `Client` and `BaseURL` fields
- [x] `NewHAClient()` constructor with transport config
- [x] `LoadEnv()` helper for configuration

### 1.5.3 Populate `client/queries.go` ✅
- [x] Move `HAQuery` struct
- [x] Move `RequestQuery()` method
- [x] Move `RequestQueryRange()` method
- [ ] Extract `BuildSelector()` helper (DRY the PromQL construction) - optional

### 1.5.4 Populate `client/series.go` ✅
- [x] Move `RequestStatus()` method
- [x] Move `RequestLabels()` method
- [x] Move `RequestLabelValue()` method
- [x] Move `RequestSeriesCount()` method

### 1.5.5 Populate `client/errors.go` 🔶 Pending
> **New code - implement when needed**

- [ ] Define `APIError` type (for status != "success")
- [ ] Define `HTTPError` type (for non-2xx responses)
- [ ] Define `ParseError` type (for JSON unmarshal failures)
- [ ] Implement `Error()` method on each type

### 1.5.6 Update `main.go` ✅
- [x] Add import for `client` package
- [x] Update type references to use `client.` prefix
- [x] Remove moved code
- [x] Verify compilation and runtime

---

## Phase 2: API Methods - Complete Implementation

### 2.1 Query Endpoints (Core) ✅ Partially Complete
> **Docs:** See [API_REFERENCE.md#endpoint-details](docs/API_REFERENCE.md#endpoint-details)

| Function | Endpoint | Status | Docs Section |
|----------|----------|--------|--------------|
| `Query()` | `/api/v1/query` | ✅ Done | [#1](docs/API_REFERENCE.md#1-instant-query---apiv1query) |
| `QueryRange()` | `/api/v1/query_range` | ✅ Done | [#2](docs/API_REFERENCE.md#2-range-query---apiv1query_range) |
| `GetSeries()` | `/api/v1/series` | 🔶 TODO | [#3](docs/API_REFERENCE.md#3-series-discovery---apiv1series) |
| `GetLabels()` | `/api/v1/labels` | ✅ Done | [#4](docs/API_REFERENCE.md#4-label-names---apiv1labels) |
| `GetLabelValues()` | `/api/v1/label/<name>/values` | ✅ Done | [#5](docs/API_REFERENCE.md#5-label-values---apiv1labelnamevalues) |

### 2.2 Export Endpoints (Data Extraction)
> **Docs:** See [API_REFERENCE.md#6-export-json](docs/API_REFERENCE.md#6-export-json---apiv1export)

| Function | Endpoint | Status | Docs Section |
|----------|----------|--------|--------------|
| `Export()` | `/api/v1/export` | 🔶 TODO | [#6](docs/API_REFERENCE.md#6-export-json---apiv1export) |
| `ExportCSV()` | `/api/v1/export/csv` | 🔶 TODO | [#7](docs/API_REFERENCE.md#7-export-csv---apiv1exportcsv) |

**Note:** Export endpoints return JSON Lines or CSV, not standard API response wrapper.

### 2.3 Status & Metadata Endpoints
> **Docs:** See [API_REFERENCE.md#8-series-count](docs/API_REFERENCE.md#8-series-count---apiv1seriescount)

| Function | Endpoint | Status | Docs Section |
|----------|----------|--------|--------------|
| `GetSeriesCount()` | `/api/v1/series/count` | ✅ Done | [#8](docs/API_REFERENCE.md#8-series-count---apiv1seriescount) |
| `GetTSDBStatus()` | `/api/v1/status/tsdb` | ✅ Done | [#9](docs/API_REFERENCE.md#9-tsdb-status---apiv1statustsdb) |
| `GetTopQueries()` | `/api/v1/status/top_queries` | 🔷 Optional | [#10](docs/API_REFERENCE.md#10-top-queries---apiv1statustop_queries) |

---

## Phase 3: Response Types
> **Docs:** See [PATTERNS.md#4-json-parsing-patterns](docs/PATTERNS.md#4-json-parsing-patterns) for unmarshaling strategies
> **Docs:** See [API_REFERENCE.md](docs/API_REFERENCE.md) for exact JSON response structures

### 3.1 Define Response Structs ✅ MOSTLY COMPLETE

- [x] `VMEnvelope` wrapper with `json.RawMessage` for dynamic `Data` field
- [x] `QueryResponse` with `resultType` discriminator
- [x] `InstantResponse` and `InstantMetric` for `/api/v1/query`
- [x] `RangeResponse` and `RangeMetric` for `/api/v1/query_range`
- [x] Handle `[timestamp, "value"]` array with `[2]any`
- [x] `Labels` type alias for label discovery
- [ ] Fix `Status` struct - missing fields (see statusCache.json for full structure)
- [ ] Fix `Series.RequestCount` → `RequestsCount` (JSON field is plural)

### 3.2 Response Parsing Logic ✅ COMPLETE

- [x] Two-step unmarshaling (envelope → specific type)
- [x] Switch on `resultType` for vector vs matrix
- [x] Error field checking

---

## Phase 4: Convenience & Discovery
> **Docs:** See [ARCHITECTURE.md#evolution-path](docs/ARCHITECTURE.md#evolution-path) for design progression

### 4.1 High-Level Discovery
- [ ] `DiscoverAllSeries() ([]Series, error)` - fetch all available metrics
- [ ] `GetEntityIDs() ([]string, error)` - list all entity_id values
- [ ] `GetDomains() ([]string, error)` - list all domain values

### 4.2 Power Monitoring Specific
- [ ] `GetCurrentPower(entityID string) (float64, error)`
- [ ] `GetPowerHistory(entityID string, duration string) ([]DataPoint, error)`

---

## Phase 5: Testing
> **Docs:** See [PATTERNS.md#8-testing-http-clients](docs/PATTERNS.md#8-testing-http-clients) for httptest patterns
> **Docs:** See [ARCHITECTURE.md#testing-strategy](docs/ARCHITECTURE.md#testing-strategy) for test pyramid

### 5.1 Unit Tests
- [ ] Test URL construction produces correct encoded strings
- [ ] Test JSON parsing for each response type
- [ ] Test `BuildSelector()` helper

### 5.2 HTTP Tests (with httptest)
- [ ] Test successful responses return correct data
- [ ] Test 4xx/5xx status codes return appropriate errors
- [ ] Test API error responses (status: "error") are handled

---

## Next Steps

**Immediate:** Fix `Status` and `Series` types in `client/types.go`

1. Fix `Series.RequestCount` → `RequestsCount` (JSON uses plural)
2. Complete `Status` struct with all fields from statusCache.json:
   - `seriesCountByLabelName`
   - `seriesCountByFocusLabelValue`
   - `seriesCountByLabelValuePair`
   - `labelValueCountByLabelName`
3. Consider splitting into `MetricSeries` (with request tracking) and `NameValuePair` (simple)

**After Types:** Implement remaining discovery functions (`GetSeries`)

---

## Sources

- [VictoriaMetrics API Examples](https://docs.victoriametrics.com/victoriametrics/url-examples/)
- [VictoriaMetrics Single-Node Documentation](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/)
