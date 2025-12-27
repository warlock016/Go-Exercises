# Tuya Protocol 3.3/3.4 Reference

This document describes the Tuya local protocol versions 3.3 and 3.4, used by smart home devices to communicate over the local network.

---

## Overview

Tuya devices communicate using:
- **TCP port 6668** - For active queries (connect, send command, receive response)
- **UDP port 6666** - Unencrypted broadcasts (protocol 3.1)
- **UDP port 6667** - Encrypted broadcasts (protocol 3.3+)

**CRITICAL: Two different encryption keys are used:**

| Traffic Type | Port | Encryption Key | Contains |
|--------------|------|----------------|----------|
| UDP Discovery | 6667 | `MD5("yGAdlopoPVldABfn")` | Device IPs, IDs, versions |
| TCP Device Comm | 6668 | Device's LocalKey | DPs (measurements) |

UDP broadcasts use a **fixed global key**, NOT the device's LocalKey!

---

## Packet Structure

Every Tuya packet follows this binary format:

```
┌────────────┬────────────┬────────────┬────────────┬────────────┬─────────────────┬────────────┬────────────┐
│  Prefix    │  Sequence  │  Command   │  Length    │  RetCode   │  Payload        │    CRC     │  Suffix    │
│  4 bytes   │  4 bytes   │  4 bytes   │  4 bytes   │  4 bytes   │  variable       │  4 bytes   │  4 bytes   │
│  0x000055AA│  Big-End   │  Big-End   │  Big-End   │  4 bytes   │  Encrypted JSON │  4 bytes   │  0x0000AA55│
└────────────┴────────────┴────────────┴────────────┴────────────┴─────────────────┴────────────┴────────────┘

Total minimum size: 24 bytes (header) + 8 bytes (footer) = 32 bytes (empty payload)
```

### Field Details

| Field | Size | Byte Order | Description |
|-------|------|------------|-------------|
| Prefix | 4 | Big-Endian | Magic number `0x000055AA` |
| Sequence | 4 | Big-Endian | Incrementing counter, starts at 1 |
| Command | 4 | Big-Endian | Command type (see table below) |
| Length | 4 | Big-Endian | Bytes after this field (RetCode + Payload + CRC + Suffix) |
| RetCode | 4 | Big-Endian | Return code (0 = success) |
| Payload | Variable | N/A | Encrypted JSON data |
| CRC | 4 | Big-Endian | CRC32 checksum of packet (optional, often 0) |
| Suffix | 4 | Big-Endian | Magic number `0x0000AA55` |

---

## Command Codes

### TCP Commands (Port 6668)

| Code | Hex | Name | Direction | Description |
|------|-----|------|-----------|-------------|
| 7 | 0x07 | CONTROL | Client → Device | Set data point values |
| 8 | 0x08 | STATUS | Device → Client | Status response |
| 9 | 0x09 | HEARTBEAT | Both | Keep-alive ping/pong |
| 10 | 0x0a | DP_QUERY | Client → Device | Request all data points |
| 13 | 0x0d | DP_QUERY_NEW | Client → Device | Alternative query format |

### UDP Broadcast Commands (Port 6667)

| Code | Hex | Name | Description |
|------|-----|------|-------------|
| 19 | 0x13 | UDP_NEW | Device discovery broadcast (v3.3) |
| 35 | 0x23 | BROADCAST_V34 | Device discovery broadcast (v3.4) |

### Common Usage

**Querying Device Status:**
1. Send command `0x0a` (DP_QUERY) with empty or `{}` payload
2. Receive command `0x08` (STATUS) with data points

**Heartbeat:**
1. Send command `0x09` with empty payload
2. Receive command `0x09` response

---

## Data Points (DPs)

Tuya devices expose functionality through numbered "Data Points". Each DP has:
- A numeric ID (1, 2, 3, etc.)
- A value (boolean, integer, string, or enum)

### Example Payload (Decrypted JSON)

**Status Response:**
```json
{
  "dps": {
    "1": true,
    "2": 25,
    "18": 1200,
    "19": 85,
    "20": 2300
  }
}
```

### Common DP IDs for Smart Plugs

| DP | Type | Description |
|----|------|-------------|
| 1 | bool | Power on/off |
| 2 | int | Countdown timer (seconds) |
| 18 | int | Current (mA) |
| 19 | int | Power (W × 10) |
| 20 | int | Voltage (V × 10) |

*Note: DP meanings vary by device type and manufacturer.*

---

## Encryption

### Protocol 3.3 Encryption

- **Algorithm:** AES-128-ECB
- **Key:** 16-byte LocalKey (from Tuya IoT Platform or device pairing)
- **Padding:** PKCS7 (pad to 16-byte blocks)
- **Mode:** ECB (Electronic Codebook) - each 16-byte block encrypted independently

### Encryption Process (for sending)

```
1. Take JSON payload: {"dps": {"1": true}}
2. Convert to bytes: [123, 34, 100, 112, ...]
3. Add PKCS7 padding to make length multiple of 16
4. Encrypt each 16-byte block with AES-ECB
5. Result is encrypted payload
```

### Decryption Process (for receiving)

```
1. Take encrypted bytes from packet
2. Decrypt each 16-byte block with AES-ECB
3. Remove PKCS7 padding
4. Parse as JSON
```

### PKCS7 Padding

- Pad with N bytes, each with value N
- If data is already multiple of 16, add full block of 16 bytes (value 0x10)

**Example:**
```
Data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13]  (13 bytes)
Pad:  [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 3, 3, 3]  (16 bytes)
                                                   ↑ 3 padding bytes, each = 0x03
```

### Key Derivation

For protocol 3.3, the LocalKey is used directly (no hashing required).

**LocalKey sources:**
1. Tuya IoT Platform → Device Management → Local Key
2. Extracted from mobile app traffic
3. Third-party tools like `tuya-cli`

---

## Packet Examples

### DP_QUERY Request (to send)

```
Bytes (hex):
000055AA                    # Prefix
00000001                    # Sequence = 1
0000000A                    # Command = 10 (DP_QUERY)
00000018                    # Length = 24 (8 + encrypted + 8)
00000000                    # RetCode = 0
[encrypted payload]         # Encrypted "{}" or empty
00000000                    # CRC (often 0)
0000AA55                    # Suffix
```

### STATUS Response (received)

```
Bytes (hex):
000055AA                    # Prefix
00000001                    # Sequence = 1
00000008                    # Command = 8 (STATUS)
00000030                    # Length = 48
00000000                    # RetCode = 0
[32 bytes encrypted]        # Encrypted JSON with dps
00000000                    # CRC
0000AA55                    # Suffix
```

---

## Reading Packets in Go

Use `encoding/binary` with `binary.BigEndian`:

```go
prefix := binary.BigEndian.Uint32(data[0:4])
seq := binary.BigEndian.Uint32(data[4:8])
cmd := binary.BigEndian.Uint32(data[8:12])
length := binary.BigEndian.Uint32(data[12:16])
```

## Writing Packets in Go

```go
buf := make([]byte, 16) // header only
binary.BigEndian.PutUint32(buf[0:4], 0x000055AA)
binary.BigEndian.PutUint32(buf[4:8], sequence)
binary.BigEndian.PutUint32(buf[8:12], command)
binary.BigEndian.PutUint32(buf[12:16], payloadLen+12) // +retcode+crc+suffix
```

---

## TCP Connection Flow

1. **Connect:** `net.Dial("tcp", "device_ip:6668")`
2. **Send DP_QUERY:** Write encoded packet bytes
3. **Read Response:** Read bytes, parse packet, decrypt
4. **Keep-Alive:** Send HEARTBEAT every 10-30 seconds
5. **Disconnect:** `conn.Close()`

### Timeout Recommendations

- Connection timeout: 5 seconds
- Read timeout: 10 seconds
- Heartbeat interval: 15-30 seconds

---

## UDP Discovery Broadcasts

Devices periodically broadcast discovery info on UDP port 6667. These broadcasts tell you:
- Device IP addresses
- Device IDs (gwId)
- Protocol versions
- Whether the device requires encryption

### Discovery Broadcast Encryption

**IMPORTANT:** Discovery broadcasts use a DIFFERENT key than device communication!

```go
// The global UDP discovery key (same for ALL Tuya devices)
udpKey := md5.Sum([]byte("yGAdlopoPVldABfn"))
// Result: 6c1ec8e2bb9bb59ab50b0daf649b410a
```

### Discovery Payload Extraction

For UDP broadcasts, use simple slicing (no version string detection):
```go
// Remove 20-byte header and 8-byte footer
encryptedPayload := data[20 : len(data)-8]
```

### Decrypted Discovery Response

```json
{
  "ip": "192.168.178.21",
  "gwId": "bf51e799667e9b1fbazfkx",
  "active": 2,
  "ablilty": 0,
  "encrypt": true,
  "productKey": "keym9qkuywghyrvs",
  "version": "3.3"
}
```

Protocol 3.4 devices include additional fields:
```json
{
  "ip": "192.168.178.65",
  "gwId": "bf446c6d6fe6f70e16on63",
  "version": "3.4",
  "token": true,
  "wf_cfg": true
}
```

### Capturing Discovery Broadcasts

1. Listen on network interface for UDP port 6667
2. Check for Tuya prefix `0x000055AA`
3. Extract payload: `data[20:len(data)-8]`
4. Decrypt with `MD5("yGAdlopoPVldABfn")` key
5. Parse JSON to get device info

**Note:** Discovery only tells you WHICH devices exist. To get actual measurements (DPs), you must connect via TCP to the device.

---

## Error Handling

| RetCode | Meaning |
|---------|---------|
| 0 | Success |
| 1 | Invalid request |
| 2 | Device busy |
| 3 | Network error |

If decryption produces garbage (not valid JSON), the LocalKey is likely incorrect.

---

## Think About

1. Why does Tuya use ECB mode instead of CBC? (Hint: stateless encryption)
2. How would you validate a packet before attempting decryption?
3. What happens if sequence numbers get out of sync?
4. How would you handle a device that doesn't respond to queries?

---

## References

- [TuyAPI](https://github.com/codetheweb/tuyapi) - Node.js implementation
- [LocalTuya](https://github.com/rospogrigio/localtuya) - Home Assistant integration
- [tinytuya](https://github.com/jasonacox/tinytuya) - Python implementation
