# Modbus Simulator - Project Specification

A Modbus TCP server simulator for IIoT development and testing.

**Estimated Time:** 8-12 hours
**Difficulty:** Advanced
**Prerequisites:** Projects 0-2 completed, Module 09 (Concurrency)

---

## ⚠️ Status: DEFERRED

This project requires advanced concepts you haven't learned yet:
- **Concurrency** (goroutines, channels, sync primitives)
- **Binary protocols** (encoding/binary package)
- **TCP networking** (net.Listener, net.Conn)
- **Interface design** (for different register types)

**When to start:** After completing Module 09 (Concurrency) - approximately Week 12+

---

## Learning Goals

This project will practice:
- **TCP server** - Raw socket programming with `net` package
- **Binary protocols** - Reading/writing binary data with `encoding/binary`
- **Concurrency** - Handling multiple simultaneous connections
- **Protocol implementation** - Following Modbus TCP specification
- **State management** - Simulating device registers
- **Testing** - Testing binary protocol handlers

---

## What is Modbus?

Modbus is a communication protocol widely used in industrial automation (SCADA, PLCs, sensors). It's simple, well-documented, and perfect for learning binary protocols.

### Modbus Concepts

| Term | Description |
|------|-------------|
| **Coils** | Single-bit read/write values (on/off, true/false) |
| **Discrete Inputs** | Single-bit read-only values |
| **Holding Registers** | 16-bit read/write values |
| **Input Registers** | 16-bit read-only values |
| **Unit ID** | Device address (1-247) |
| **Function Code** | Operation type (read coils, write register, etc.) |

### Common Function Codes

| Code | Name | Description |
|------|------|-------------|
| 0x01 | Read Coils | Read 1-2000 coils |
| 0x02 | Read Discrete Inputs | Read 1-2000 discrete inputs |
| 0x03 | Read Holding Registers | Read 1-125 registers |
| 0x04 | Read Input Registers | Read 1-125 registers |
| 0x05 | Write Single Coil | Write one coil |
| 0x06 | Write Single Register | Write one register |
| 0x0F | Write Multiple Coils | Write 1-1968 coils |
| 0x10 | Write Multiple Registers | Write 1-123 registers |

---

## Requirements

### Functional Requirements

1. **Listen for Modbus TCP connections** on configurable port (default 502)
2. **Simulate device registers**:
   - Holding registers (addresses 0-999)
   - Input registers (addresses 0-999)
   - Coils (addresses 0-999)
   - Discrete inputs (addresses 0-999)
3. **Support function codes**: 0x01, 0x02, 0x03, 0x04, 0x05, 0x06
4. **Handle multiple clients** simultaneously
5. **Simulate dynamic values** (optional: temperature, counters, etc.)
6. **Log all transactions** for debugging

### Command-Line Interface

```bash
# Start simulator with defaults
modbussim

# Custom port
modbussim -port 5020

# Pre-load register values from file
modbussim -config registers.json

# Verbose logging
modbussim -verbose
```

### Non-Functional Requirements

- Handle malformed requests gracefully (return exception response)
- Support at least 10 concurrent connections
- Response time < 50ms for standard operations
- Clean shutdown on SIGINT (Ctrl+C)

---

## Modbus TCP Frame Format

### Request Frame (from client)
```
| MBAP Header (7 bytes)      | PDU (variable)           |
|----------------------------|--------------------------|
| Transaction ID | Protocol  | Length | Unit | FC | Data|
| 2 bytes        | 2 bytes   | 2 bytes| 1 b  | 1b | ... |
```

### MBAP Header Fields
- **Transaction ID**: Client-assigned ID, echo back in response
- **Protocol ID**: Always 0x0000 for Modbus
- **Length**: Bytes following (Unit ID + PDU)
- **Unit ID**: Device address (usually 1)

### Example: Read Holding Registers Request
```
Transaction ID:  0x00 0x01
Protocol ID:     0x00 0x00
Length:          0x00 0x06
Unit ID:         0x01
Function Code:   0x03
Start Address:   0x00 0x00  (register 0)
Quantity:        0x00 0x0A  (10 registers)
```

### Example: Read Holding Registers Response
```
Transaction ID:  0x00 0x01
Protocol ID:     0x00 0x00
Length:          0x00 0x17  (23 bytes follow)
Unit ID:         0x01
Function Code:   0x03
Byte Count:      0x14       (20 bytes of data)
Data:            [20 bytes of register values]
```

---

## Suggested File Structure

```
03_modbus_simulator/
├── main.go              # Entry point: flag parsing, server setup
├── server.go            # TCP server: Accept connections, dispatch handlers
├── handler.go           # Protocol handler: Parse requests, build responses
├── registers.go         # Register storage: Coils, holding registers, etc.
├── protocol.go          # Modbus protocol: Frame parsing, function codes
├── protocol_test.go     # Tests for protocol parsing
├── registers_test.go    # Tests for register operations
└── README.md            # This spec
```

---

## Useful Packages

| Package | Purpose | Documentation |
|---------|---------|---------------|
| `net` | TCP server/client | https://pkg.go.dev/net |
| `encoding/binary` | Binary encoding (big-endian) | https://pkg.go.dev/encoding/binary |
| `sync` | Mutex for thread-safe registers | https://pkg.go.dev/sync |
| `context` | Cancellation, timeouts | https://pkg.go.dev/context |
| `io` | Reading exact bytes | https://pkg.go.dev/io |
| `log` | Transaction logging | https://pkg.go.dev/log |

---

## Key Concepts to Learn

### 1. TCP Server

```go
listener, err := net.Listen("tcp", ":502")
if err != nil {
    log.Fatal(err)
}
defer listener.Close()

for {
    conn, err := listener.Accept()
    if err != nil {
        log.Println("Accept error:", err)
        continue
    }
    go handleConnection(conn)  // Handle in goroutine
}
```

### 2. Reading Binary Data

```go
import "encoding/binary"

// Read fixed-size header
header := make([]byte, 7)
if _, err := io.ReadFull(conn, header); err != nil {
    return err
}

// Parse big-endian values
transactionID := binary.BigEndian.Uint16(header[0:2])
protocolID := binary.BigEndian.Uint16(header[2:4])
length := binary.BigEndian.Uint16(header[4:6])
unitID := header[6]
```

### 3. Writing Binary Data

```go
response := make([]byte, 0, 256)

// Build header
var header [7]byte
binary.BigEndian.PutUint16(header[0:2], transactionID)
binary.BigEndian.PutUint16(header[2:4], 0) // Protocol ID
binary.BigEndian.PutUint16(header[4:6], uint16(len(pdu)+1))
header[6] = unitID

response = append(response, header[:]...)
response = append(response, pdu...)

conn.Write(response)
```

### 4. Thread-Safe Register Storage

```go
type Registers struct {
    mu       sync.RWMutex
    holding  [1000]uint16
    input    [1000]uint16
    coils    [1000]bool
    discrete [1000]bool
}

func (r *Registers) ReadHolding(addr, count uint16) []uint16 {
    r.mu.RLock()
    defer r.mu.RUnlock()
    return r.holding[addr : addr+count]
}
```

### 5. Handling Multiple Clients

```go
func handleConnection(conn net.Conn) {
    defer conn.Close()

    for {
        request, err := readRequest(conn)
        if err != nil {
            if err != io.EOF {
                log.Println("Read error:", err)
            }
            return
        }

        response := processRequest(request)
        if _, err := conn.Write(response); err != nil {
            log.Println("Write error:", err)
            return
        }
    }
}
```

---

## Hints (Read Only If Stuck)

<details>
<summary>Hint 1: Modbus exception responses</summary>

When a request is invalid, return an exception:
```go
// Exception response: FC | 0x80, then error code
// Error codes:
// 0x01 = Illegal function
// 0x02 = Illegal data address
// 0x03 = Illegal data value

func buildException(fc byte, errorCode byte) []byte {
    return []byte{fc | 0x80, errorCode}
}
```
</details>

<details>
<summary>Hint 2: Read registers implementation</summary>

```go
func handleReadHoldingRegisters(pdu []byte, regs *Registers) []byte {
    startAddr := binary.BigEndian.Uint16(pdu[0:2])
    quantity := binary.BigEndian.Uint16(pdu[2:4])

    // Validate
    if startAddr+quantity > 1000 {
        return buildException(0x03, 0x02)
    }

    // Build response
    byteCount := quantity * 2
    response := make([]byte, 1+int(byteCount))
    response[0] = byte(byteCount)

    for i := uint16(0); i < quantity; i++ {
        val := regs.ReadHolding(startAddr + i)
        binary.BigEndian.PutUint16(response[1+i*2:], val)
    }

    return response
}
```
</details>

<details>
<summary>Hint 3: Simulating dynamic values</summary>

```go
func (r *Registers) SimulateSensors() {
    go func() {
        ticker := time.NewTicker(time.Second)
        for range ticker.C {
            r.mu.Lock()
            // Temperature (register 0): 20-30°C with noise
            r.holding[0] = uint16(2500 + rand.Intn(500)) // 25.00°C in centi-degrees

            // Counter (register 1): increment
            r.holding[1]++
            r.mu.Unlock()
        }
    }()
}
```
</details>

---

## Success Criteria

- [ ] Server listens on TCP port 502 (or configurable)
- [ ] Handles Read Holding Registers (0x03)
- [ ] Handles Read Input Registers (0x04)
- [ ] Handles Write Single Register (0x06)
- [ ] Handles multiple simultaneous clients
- [ ] Returns exception for invalid addresses
- [ ] Returns exception for invalid function codes
- [ ] Transaction IDs echo correctly
- [ ] Tests exist for protocol parsing
- [ ] Tests exist for register operations
- [ ] Clean shutdown on SIGINT

---

## Testing Tools

### Using `modbus-cli` (if available)
```bash
# Install: go install github.com/simonvetter/modbus-cli@latest
modbus-cli -host localhost:502 read holding 0 10
modbus-cli -host localhost:502 write holding 0 12345
```

### Using Python (pymodbus)
```python
from pymodbus.client import ModbusTcpClient

client = ModbusTcpClient('localhost', port=502)
client.connect()

# Read holding registers
result = client.read_holding_registers(0, 10)
print(result.registers)

# Write single register
client.write_register(0, 12345)
```

---

## Resources

- [Modbus TCP Specification](https://modbus.org/specs.php)
- [Simply Modbus](https://www.simplymodbus.ca/) - Good visual explanations
- [Modbus Protocol Reference](https://www.modbustools.com/modbus.html)

---

## Getting Started (When Ready)

1. First, complete Module 09 (Concurrency)
2. Read the Modbus TCP specification (at least function codes 0x03, 0x06)
3. Start with `protocol.go` - implement frame parsing with tests
4. Add `registers.go` - simple in-memory storage
5. Build `server.go` - basic TCP accept loop
6. Implement `handler.go` - connect parsing to storage
7. Add error handling and logging

This is a challenging project - take your time and test each piece thoroughly!
