# Tuya IoT Measurement API

A Go project to fetch raw measurements from paired Tuya IoT devices and expose them via a REST API.

## Learning Objectives

By completing this project, you will learn:

1. **Protocol Engineering** - Parse binary network protocols, handle encryption/decryption
2. **Modular Go Architecture** - Organize code into focused packages with clear interfaces
3. **Concurrent Programming** - Use goroutines, channels, and sync primitives safely
4. **Network Programming** - TCP client connections, UDP packet capture
5. **REST API Design** - Build HTTP endpoints with Go's standard library

---

## Architecture Overview

```
┌─────────────────┐
│   REST Client   │
└────────┬────────┘
         │ HTTP :8080
┌────────▼────────┐
│   api.Server    │ ◄── GET /devices, POST /query
└────────┬────────┘
         │
┌────────▼────────┐     ┌─────────────────┐     ┌─────────────────┐
│ device.Registry │◄────│  device.Device  │────►│ protocol.Codec  │
│ (thread-safe)   │     │ Connect()/Query()│    │ Encode/Decode   │
└────────┬────────┘     └────────┬────────┘     └─────────────────┘
         │                       │
         │              TCP ─────┴───── Port 6668
         │              (Active Queries)
         │
┌────────▼────────┐
│ monitor.Monitor │ ◄── UDP Broadcast capture (passive)
└─────────────────┘
```

### Two Modes of Operation

| Mode | How It Works | Use Case |
|------|--------------|----------|
| **Passive** | Capture UDP broadcasts from devices | Real-time monitoring without polling |
| **Active** | TCP connect to device, send query, get response | On-demand status fetch |

---

## Project Structure

```
08_udp_decoder/
├── main.go                 # Entry point - orchestrates all components
├── .env                    # Device credentials
├── docs/
│   ├── TUYA_PROTOCOL.md    # Protocol 3.3 specification
│   └── GO_PATTERNS.md      # Relevant Go patterns
├── protocol/               # Binary protocol handling
│   ├── README.md
│   ├── protocol.go         # Constants, command codes
│   ├── packet.go           # Encode/decode packets
│   └── crypto.go           # AES-ECB encryption
├── device/                 # Device management
│   ├── README.md
│   ├── device.go           # Device struct, TCP methods
│   ├── registry.go         # Thread-safe storage
│   └── state.go            # Measurement caching
├── monitor/                # Passive UDP capture
│   ├── README.md
│   ├── monitor.go          # Packet capture loop
│   └── handler.go          # Process captured packets
├── api/                    # REST endpoints
│   ├── README.md
│   ├── server.go           # HTTP server
│   ├── handlers.go         # Request handlers
│   └── responses.go        # JSON types
└── config/                 # Configuration
    ├── README.md
    └── config.go           # Parse .env file
```

---

## Implementation Phases

### Phase 1: Protocol Layer
**Goal:** Extract crypto and packet parsing from existing `main.go` into reusable package.

**You already have working code for:**
- AES-ECB decryption (`tryDecrypt`) - lines 377-404
- MD5 key derivation (`deriveKeyMD5`) - lines 371-375
- PKCS7 padding removal (`removePKCS7Padding`) - lines 406-431
- Packet parsing (`parseTuyaPacket`) - lines 196-278
- Key selection logic (`decodePayload`) - lines 280-369

**New code needed:**
- AES-ECB encryption (reverse of decryption)
- PKCS7 padding addition
- Packet encoding (for sending TCP queries)

**Files:** `protocol/protocol.go`, `protocol/crypto.go`, `protocol/packet.go`

---

### Phase 2: Configuration
**Goal:** Support multiple devices in `.env` file.

**Current working format:**
```env
SMART_PLUG_KEY="=j(Brn(YJuA4n'1O"  # 16-byte local key
SMART_PLUG_IP=192.168.178.21
DEVICE_ID="bf51e799667e9b1fbazfkx"
MAC_ADDRESS=FC:67:1F:CA:89:A0
```

**Target format (multiple devices):**
```env
API_PORT=8080
MONITOR_INTERFACE=en0

DEVICE_1_NAME=SmartMeter
DEVICE_1_ID=bf446c6d6fe6f70e16on63
DEVICE_1_LOCAL_KEY=)+cBrK`YqEy}wlX=
DEVICE_1_IP=192.168.178.65

DEVICE_2_NAME=SmartPlug
DEVICE_2_ID=bf51e799667e9b1fbazfkx
DEVICE_2_LOCAL_KEY==j(Brn(YJuA4n'1O
DEVICE_2_IP=192.168.178.21

DEVICE_3_NAME=SmartBulb
DEVICE_3_ID=bf753f9daac0f59b01h56j
DEVICE_3_LOCAL_KEY=Cno)4MbemamDUw<y
DEVICE_3_IP=192.168.178.110
```

**Files:** `config/config.go`

---

### Phase 3: Device Management
**Goal:** Store device info with thread-safe access.

**Key concepts:**
- Registry pattern (map with mutex)
- State caching with timestamps
- Dual indexing (by ID and by IP)

**Files:** `device/device.go`, `device/registry.go`, `device/state.go`

---

### Phase 4: Refactor Monitor
**Goal:** Extract UDP capture into standalone package.

**Current code to move from main.go:**
- `startCapture()` (lines 114-145)
- `processPacket()` (lines 147-194)
- `parseTuyaPacket()` (lines 196-278) - UDP broadcast handling

**New integration:**
- Correlate packets to devices via Registry (use GetByIP)
- Update device state when packets received
- Distinguish UDP discovery (port 6667) from other traffic

**Files:** `monitor/monitor.go`, `monitor/handler.go`

---

### Phase 5: Active TCP Querying
**Goal:** Connect to devices and request status on-demand.

**New capabilities:**
- `Connect(ctx)` - Dial TCP to port 6668
- `Query(ctx)` - Send DP_QUERY command, await response
- `SendHeartbeat(ctx)` - Keep connection alive

**Protocol details:**
- Command `0x0a` = DP_QUERY (request all data points)
- Command `0x08` = STATUS (device response)
- Sequence numbers must increment

**Files:** Extend `device/device.go`

---

### Phase 6: REST API
**Goal:** Expose device data via HTTP.

**Endpoints:**
| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/v1/devices` | List all devices |
| GET | `/api/v1/devices/:id` | Get device info |
| GET | `/api/v1/devices/:id/status` | Get cached measurements |
| POST | `/api/v1/devices/:id/query` | Force active query |
| GET | `/api/v1/health` | Health check |

**Files:** `api/server.go`, `api/handlers.go`, `api/responses.go`

---

### Phase 7: Main Integration
**Goal:** Wire everything together with graceful shutdown.

**Components to orchestrate:**
1. Load config
2. Create device registry
3. Start passive monitor (goroutine)
4. Start HTTP server (blocking)
5. Handle SIGINT/SIGTERM for clean shutdown

---

## Current Status

**Working (as of 2024-12-19):**
- [x] UDP broadcast capture on port 6667
- [x] Discovery packet decryption with global UDP key
- [x] Device detection (3 devices found: SmartPlug v3.3, SmartBulb v3.3, SmartMeter v3.4)

**Not Yet Implemented:**
- [ ] TCP connection to devices (port 6668)
- [ ] DP_QUERY command sending
- [ ] Device status polling
- [ ] REST API endpoints
- [ ] Multi-device configuration

## Success Criteria

- [ ] `go build ./...` succeeds with no errors
- [ ] `go test ./...` passes
- [ ] Can list devices via `curl localhost:8080/api/v1/devices`
- [ ] Can get cached status via `curl localhost:8080/api/v1/devices/{id}/status`
- [ ] Can force query via `curl -X POST localhost:8080/api/v1/devices/{id}/query`
- [ ] Passive monitor updates device state when UDP packets captured
- [ ] Clean package separation (no circular imports)

---

## Getting Started

1. Read `docs/TUYA_PROTOCOL.md` to understand the protocol
2. Read `docs/GO_PATTERNS.md` for relevant Go idioms
3. Start with Phase 1 - extract protocol code from `main.go`
4. Follow package READMEs for guidance on each component

---

## Testing Commands

```bash
# Build all packages
go build ./...

# Run all tests
go test ./...

# Run specific package tests
go test -v ./protocol/

# Test with race detector
go test -race ./...

# Run the server
go run . -i en0
```

---

## Reference

- [Go net/http documentation](https://pkg.go.dev/net/http)
- [Go sync package](https://pkg.go.dev/sync)
- [gopacket documentation](https://pkg.go.dev/github.com/google/gopacket)
- [encoding/binary](https://pkg.go.dev/encoding/binary)
- [crypto/aes](https://pkg.go.dev/crypto/aes)
