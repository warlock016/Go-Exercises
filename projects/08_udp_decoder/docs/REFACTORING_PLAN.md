# Refactoring Plan: main.go to Packages

This document maps the current `main.go` (432 lines) to the target package structure.

---

## Current main.go Structure

| Lines | Function/Type | Purpose |
|-------|---------------|---------|
| 23-32 | `Packet` struct | Packet structure (unused, needs revision) |
| 36-87 | `main()` | CLI, config loading, orchestration |
| 89-95 | `DeviceInfo` struct | Device credentials |
| 97-112 | `listInterfaces()` | List network interfaces |
| 114-145 | `startCapture()` | Open interface, packet loop |
| 147-194 | `processPacket()` | Extract layers, call parser |
| 196-278 | `parseTuyaPacket()` | Parse header, extract payload |
| 280-369 | `decodePayload()` | Key selection, decryption attempts |
| 371-375 | `deriveKeyMD5()` | MD5 hash for key derivation |
| 377-404 | `tryDecrypt()` | AES-ECB decryption |
| 406-431 | `removePKCS7Padding()` | PKCS7 unpadding |

---

## Target Package Mapping

### 1. `protocol/` Package

**Purpose:** Binary protocol encoding/decoding, encryption

**Files to create:**

#### `protocol/constants.go`
```
From main.go:
- Prefix = 0x000055AA (line 197)
- Suffix = 0x0000AA55 (line 198)
- Add: Command codes (DP_QUERY = 0x0a, STATUS = 0x08, etc.)
- Add: UDP_KEY = "yGAdlopoPVldABfn"
```

#### `protocol/crypto.go`
```
From main.go:
- deriveKeyMD5() (lines 371-375)
- tryDecrypt() (lines 377-404) → rename to DecryptECB()
- removePKCS7Padding() (lines 406-431)
- Add: EncryptECB() (reverse of decrypt)
- Add: addPKCS7Padding()
```

#### `protocol/packet.go`
```
From main.go:
- parseTuyaPacket() (lines 196-278) → rename to Decode()
- Packet struct (lines 23-32) → revise fields
- Add: Encode() for building outbound packets
```

**Interface:**
```go
type Codec interface {
    Encode(cmd uint32, payload []byte, key []byte, seq uint32) ([]byte, error)
    Decode(data []byte, key []byte, isUDPBroadcast bool) (*Packet, error)
}

type Cipher interface {
    Encrypt(plaintext, key []byte) ([]byte, error)
    Decrypt(ciphertext, key []byte) ([]byte, error)
}
```

---

### 2. `config/` Package

**Purpose:** Load and validate configuration from .env

**Files to create:**

#### `config/config.go`
```
From main.go:
- Environment parsing logic (lines 43-57)
- DeviceInfo struct (lines 89-95) → expand to DeviceConfig

New functionality:
- Parse DEVICE_N_* pattern for multiple devices
- Validate required fields (ID, LOCAL_KEY, IP)
- Return Config struct with APIPort, MonitorInterface, []DeviceConfig
```

**Types:**
```go
type Config struct {
    APIPort          int
    MonitorInterface string
    Devices          []DeviceConfig
}

type DeviceConfig struct {
    ID       string
    Name     string
    IP       string
    LocalKey string
    MAC      string
    Version  string  // detected from discovery
}
```

---

### 3. `device/` Package

**Purpose:** Device management, state caching, TCP connection

**Files to create:**

#### `device/device.go`
```
From main.go:
- DeviceInfo struct (lines 89-95) → expand to Device

New functionality:
- Connect(ctx) - TCP dial to port 6668
- Query(ctx) - Send DP_QUERY, receive response
- Disconnect()
- GetState() / UpdateState()
```

#### `device/registry.go`
```
New functionality:
- Thread-safe map[string]*Device (by ID)
- Secondary index by IP (for correlating UDP packets)
- Register(), Get(), GetByIP(), List(), Remove()
```

#### `device/state.go`
```
New functionality:
- DPs map[string]interface{}
- LastUpdated time.Time
- Source string ("passive" or "active")
```

---

### 4. `monitor/` Package

**Purpose:** UDP packet capture and processing

**Files to create:**

#### `monitor/monitor.go`
```
From main.go:
- startCapture() (lines 114-145)
- Context support for graceful shutdown

New functionality:
- NewMonitor(iface string, registry *device.Registry)
- Start(ctx context.Context) error
- Updates channel for notifying of new data
```

#### `monitor/handler.go`
```
From main.go:
- processPacket() (lines 147-194)
- Integration with device.Registry.GetByIP()
- Call protocol.Decode() and protocol.Decrypt()
```

**Key changes:**
- Remove `device DeviceInfo` parameter
- Lookup device by source IP from Registry
- Only process port 6667 broadcasts
- Update device state via Registry

---

### 5. `api/` Package

**Purpose:** REST API endpoints

**Files to create:**

#### `api/server.go`
```
New functionality:
- HTTP server with http.ServeMux
- Routes configuration
- Graceful shutdown
```

#### `api/handlers.go`
```
New functionality:
- handleListDevices()
- handleGetDevice()
- handleGetStatus()
- handleQuery()
- handleHealth()
```

#### `api/responses.go`
```
New functionality:
- JSON response types
- writeJSON(), writeError() helpers
```

---

### 6. Updated `main.go`

**Purpose:** CLI, orchestration, signal handling

**Keep:**
- Command-line flag parsing (lines 37-41, 59-66)
- listInterfaces() (lines 97-112) - optional, could move to monitor/

**New:**
```go
func main() {
    // 1. Parse flags
    // 2. Load config
    cfg, err := config.Load(".env")

    // 3. Create registry, register devices
    registry := device.NewRegistry()
    for _, d := range cfg.Devices {
        device, _ := device.NewDevice(d)
        registry.Register(device)
    }

    // 4. Start monitor in goroutine
    mon := monitor.NewMonitor(cfg.MonitorInterface, registry)
    go mon.Start(ctx)

    // 5. Start API server (blocking)
    srv := api.NewServer(cfg.APIPort, registry)
    srv.ListenAndServe()

    // 6. Handle signals for graceful shutdown
}
```

---

## Refactoring Order

**Phase 1: Protocol (no dependencies)**
1. Create `protocol/constants.go`
2. Move crypto functions to `protocol/crypto.go`
3. Move packet parsing to `protocol/packet.go`
4. Add encryption and encoding functions
5. Write tests

**Phase 2: Config (depends on nothing)**
1. Create `config/config.go`
2. Implement multi-device parsing
3. Write tests

**Phase 3: Device (depends on protocol)**
1. Create `device/device.go` with struct
2. Create `device/registry.go`
3. Create `device/state.go`
4. Write tests

**Phase 4: Monitor (depends on device, protocol)**
1. Create `monitor/monitor.go`
2. Create `monitor/handler.go`
3. Integrate with Registry
4. Write tests

**Phase 5: API (depends on device)**
1. Create `api/responses.go`
2. Create `api/handlers.go`
3. Create `api/server.go`
4. Write tests

**Phase 6: Integration**
1. Update `main.go` to use packages
2. Add signal handling
3. Integration tests

---

## Key Decisions

### 1. UDP Key Handling

The UDP discovery key should be in `protocol/constants.go`:
```go
const UDPKeyString = "yGAdlopoPVldABfn"

var UDPKey = func() []byte {
    hash := md5.Sum([]byte(UDPKeyString))
    return hash[:]
}()
```

### 2. Packet Struct Revision

Current struct has `Payload uint64` which is wrong. Revise to:
```go
type Packet struct {
    Prefix   uint32
    Sequence uint32
    Command  uint32
    Length   uint32
    RetCode  uint32
    Payload  []byte  // decrypted JSON
    CRC      uint32
    Suffix   uint32
}
```

### 3. Error Handling

Create sentinel errors in each package:
```go
// protocol/errors.go
var (
    ErrInvalidPrefix  = errors.New("invalid packet prefix")
    ErrDecryptFailed  = errors.New("decryption failed")
)

// device/errors.go
var (
    ErrDeviceNotFound = errors.New("device not found")
    ErrNotConnected   = errors.New("device not connected")
)
```

---

## File Count Summary

| Package | Files | Lines (estimate) |
|---------|-------|------------------|
| protocol/ | 4 | ~200 |
| config/ | 2 | ~100 |
| device/ | 4 | ~250 |
| monitor/ | 3 | ~150 |
| api/ | 4 | ~200 |
| main.go | 1 | ~100 |
| **Total** | **18** | **~1000** |

This represents a healthy expansion from 432 lines to ~1000 lines, with proper separation of concerns and testability.
