# Polyglot Hexagonal Web Service for PV KPI (PR) Calculation on Raspberry Pi

This document captures a practical design approach for a **web-based KPI calculator** (PR and related metrics) that ingests **Meteocontrol CSV exports**, processes them primarily with **Python/Pandas**, and serves results via a **Go-based web service** with a **JavaScript UI**, deployable on a **Raspberry Pi** for **~2–25 concurrent users**.

---

## 1) Goal and Context

### Problem
- Users upload CSV files containing PV telemetry time-series data (exported from Meteocontrol API).
- The system processes these files and returns:
  - An **Excel report** (XLSX)
  - A **short KPI summary** displayed in the Web UI

### Constraints
- Runs on a **Raspberry Pi** on a **local network**
- **2–25** concurrent users (web traffic)
- Compute workload (Pandas) is the main resource bottleneck
- Need clean interaction between parts written in **Go + Python + JS**

---

## 2) Hexagonal Architecture in a Polyglot System

Hexagonal architecture (Ports & Adapters) is about **boundaries** and **contracts**, not only folders.

### Core Idea
- Define the system around an **application core** that expresses:
  - KPI calculation use-cases
  - Data processing policies (timezone, resampling, gap handling, etc.)
- Everything external (Web UI, storage, compute engine, report writers) is an **adapter**.

### In a multi-language project
- Languages can live in **separate adapters/services**
- Sustainability comes from a strict, versioned **contract** (schemas + semantics).

---

## 3) Architecture Overview

### Components

1) **UI (JavaScript)**
- Upload CSV files
- Display job status / progress
- Display summary KPIs & warnings
- Download Excel results

2) **Go “App Server” (System Hub)**
- Handles multiple concurrent users efficiently
- Web API endpoints
- Upload handling + storage
- Job orchestration / queueing
- Persistence for job metadata (SQLite)
- Calls Python compute adapter via a stable contract

3) **Python “Compute Service” (Pandas Engine)**
- Parses and normalizes telemetry data
- Aligns time series / resamples
- Computes KPIs (PR, yields, completeness, warnings)
- Writes result report (XLSX) and summary JSON

---

## 4) How Go and Python Should Communicate (IPC Options)

### Option A — Separate Python compute service (Recommended)
- Go runs web server + orchestration
- Python runs a compute API (HTTP or gRPC)
- Shared data via mounted directory/volume

**Pros**
- Clean hexagonal boundary
- Easy concurrency control in Go
- Observability/logging is easier than subprocess
- Python can be moved to another host later if needed

**Cons**
- Requires stable API contract (recommended anyway)

---

### Option B — Python subprocess pool controlled by Go
- Go spawns Python processes and communicates via stdin/stdout or files

**Pros**
- Potentially simpler deployment (one “product” feel)
- No internal networking required

**Cons**
- Harder to manage reliably (timeouts, logs, partial failures)
- Contract can become implicit if not enforced carefully

---

### Option C — Embedding/FFI
- Go embeds Python or uses C-FFI

**Not recommended**
- Complex on ARM/RPi
- Hard debugging
- Dependency/version management issues

---

## 5) Ports and Adapters (Concrete Definition)

### Go-side Ports (Interfaces)
Define ports in the Go “application core”:

#### Upload / Storage
- `UploadStoragePort`
  - `SaveUpload(jobId, file) -> path`
  - `Open(jobId) -> stream/path`

#### Job tracking
- `JobStorePort`
  - `CreateJob(userId, metadata)`
  - `SetStatus(jobId, queued|running|done|failed + progress)`
  - `SaveResultRefs(jobId, resultPaths)`

#### KPI computation
- `KpiComputePort`
  - `Compute(jobId, inputRef, options) -> (summary, resultRef)`

#### Report generation (optional)
- `ReportPort`
  - `WriteExcel(resultData) -> path`

### Adapters (Implementations)
- Web upload adapter (HTTP handlers)
- Local filesystem adapter for uploads/results
- SQLite adapter for job store
- Python compute client adapter (HTTP/gRPC)
- UI adapter (static assets + API)

---

## 6) Contract Design (Critical for Polyglot Systems)

A stable contract is the glue that makes multi-language sustainable.

### Recommended Contract Style: File references + JSON summary
Instead of shipping huge time series over HTTP:
- Go stores uploads under a shared directory
- Go calls Python with a jobId + input folder reference
- Python reads input files, writes outputs to results folder

#### Request fields (example)
- `contractVersion`
- `jobId`
- `inputDir`
- `outputDir`
- `options`
  - timezone policy
  - resampling interval (e.g., 5min/15min)
  - column mapping profile
  - facility metadata (Pnom, etc.)

#### Response fields (example)
- `jobId`
- `status`
- `summaryPath` (JSON)
- `resultPath` (XLSX)
- `warnings` (optional inline)

### Version everything
- `contractVersion: 1`
- Column mapping profiles (what columns are required/optional)
- Unit conventions
- Timestamp conventions (UTC vs local)

---

## 7) End-to-End Workflow

1) User uploads CSV(s) in the UI  
2) Go creates a `jobId`
3) Go stores files: `/data/uploads/<jobId>/...`
4) Go stores job metadata in SQLite and enqueues job
5) Worker picks job (concurrency-limited)
6) Go calls Python compute service:
   - `POST /compute { jobId, inputDir, outputDir, options }`
7) Python:
   - parses & normalizes CSV(s)
   - aligns/resamples time series
   - computes KPIs and data quality metrics
   - writes:
     - `/data/results/<jobId>/result.xlsx`
     - `/data/results/<jobId>/summary.json`
8) Go marks job as done and stores result refs
9) UI polls `/jobs/<jobId>` for status & summary
10) User downloads XLSX report

---

## 8) Concurrency on Raspberry Pi

You have two different concurrency concerns:

### A) Web concurrency (2–25 users)
- Go handles this easily (HTTP requests, status polling, downloads)

### B) Compute concurrency (Pandas CPU/RAM)
- This is the bottleneck
- Limit compute jobs:
  - Pi 4: often **1–2** jobs concurrently (depends on RAM and file sizes)
  - Pi 5: possibly **2–4**, depending on memory and CPU load
- Queue additional jobs and show queue position in UI

#### Implementation suggestion
- Go uses a worker pool / buffered channel to limit compute concurrency
- Python service can run 1–2 workers; Go orchestrates queueing centrally

---

## 9) Excel Report Generation: Python vs Go

### Generate Excel in Python (Recommended)
- DataFrames are already in Python
- Libraries: `xlsxwriter` / `openpyxl`
- Easy multi-sheet output:
  - Raw / Cleaned data
  - KPI summary
  - Daily aggregation
  - Warnings / Data quality

### Generate Excel in Go (Possible but usually not worth it)
- Library: `excelize`
- Requires marshaling data out of Python (more friction)
- Better to keep reporting close to compute pipeline

---

## 10) Domain Pitfalls to Bake into the Compute Rules

These are typical trouble spots for PV telemetry and should become explicit policies:

### Timezone and DST
- Choose a canonical internal timeline (usually **UTC**)
- Require an explicit plant timezone in options
- Flag DST anomalies, missing or duplicated timestamps

### Time alignment
Define alignment/resampling policy:
- common interval: 1/5/15 minutes
- interpolation policy: none vs forward fill vs linear
- gap thresholds: when to ignore or flag periods

### Data completeness scoring
- KPI confidence should reflect coverage quality
- PR computed on partial data must be flagged

### Curtailment / setpoint / remote control periods
- Optionally detect/label periods where PR is not meaningful
- Separate “technical PR” from “curtailed PR” depending on rules

### Data validation and normalization
- delimiter (comma/semicolon)
- decimal comma vs dot
- column naming variations
- unit normalization (W vs kW, Wh vs kWh)

---

## 11) Deployment on Raspberry Pi

### Recommended: Docker Compose
Services:
- `app-go` (Go server)
- `compute-py` (Python compute API)
Shared volume:
- `/data` mounted into both containers

Benefits:
- Reproducible
- Isolated dependencies
- Easy to move to another host later

### Alternative: systemd services (no Docker)
- `kpi-app.service` (Go binary)
- `kpi-compute.service` (Python venv + uvicorn/gunicorn)

---

## 12) Suggested Mono-Repo Structure (Hexagonal-Friendly)

/go-app
/cmd/server # main
/internal/core # application services + ports
/internal/adapters/http # REST handlers
/internal/adapters/storage # filesystem
/internal/adapters/jobstore# sqlite
/internal/adapters/computeclient # calls python

/py-compute
/core # parse -> normalize -> align -> compute -> report
/adapters/api # FastAPI endpoints or gRPC server
/schemas # contract (mirrors /contracts ideally)

/ui

React/Vue/Svelte or simple HTML

/contracts

JSON schema or protobuf definitions
versioned docs: assumptions, mapping rules, units, time policies

/deploy
docker-compose.yml or systemd units


The `/contracts` directory is the glue that keeps a polyglot system maintainable.

---

## 13) Practical “V1” Build Plan (Avoid Over-Architecture)

1) Define **contract v1**:
   - request options, input/output paths, summary fields
2) Implement Go server:
   - upload endpoint
   - job queue + status endpoints
   - results download endpoint
   - UI serving
3) Implement Python compute:
   - `/compute` endpoint
   - read CSVs from inputDir
   - write `result.xlsx` + `summary.json` to outputDir
4) Add robustness:
   - timeouts and cancellation
   - jobId-based structured logging
   - strict input validation (required columns & units)
5) Add “profiles”:
   - different calculation/mapping recipes per export format / site template

---

## 14) Key Recommendation Summary

- Use **Option A**: Go app server + Python compute service
- Communicate using:
  - **Shared storage** for large data (CSV/XLSX)
  - **Versioned JSON contract** for options and summaries
- Let Go handle:
  - multi-user web concurrency
  - job orchestration and queueing
  - persistence and serving results
- Let Python handle:
  - Pandas-heavy processing
  - KPI computation policies
  - Excel report generation
- Limit compute concurrency on Raspberry Pi and queue jobs

---
