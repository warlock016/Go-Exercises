# Protocol Package

Handles Tuya protocol 3.3 encoding, decoding, and encryption.

## Learning Goals

- Parse binary protocols using `encoding/binary`
- Implement AES-ECB encryption/decryption
- Handle PKCS7 padding correctly
- Design clean interfaces for crypto operations

## Files to Create

| File | Purpose |
|------|---------|
| `protocol.go` | Constants and command codes |
| `crypto.go` | AES-ECB encryption/decryption, key handling |
| `packet.go` | Packet encoding and decoding |

## What to Extract from main.go

Your existing `main.go` already has working implementations. Extract and refactor:

| Function | Lines | New Location | Notes |
|----------|-------|--------------|-------|
| `deriveKeyMD5()` | 371-375 | `crypto.go` | Keep as-is |
| `tryDecrypt()` | 377-404 | `crypto.go` | Rename to `DecryptECB()` |
| `removePKCS7Padding()` | 406-431 | `crypto.go` | Keep as-is |
| `parseTuyaPacket()` | 196-278 | `packet.go` | Rename to `Decode()` |
| `decodePayload()` | 280-369 | `packet.go` | Key selection logic |

**Also add the UDP broadcast key constant:**
```go
const UDPKeyString = "yGAdlopoPVldABfn"

// UDPKey is MD5(UDPKeyString) - used for discovery broadcasts
var UDPKey = deriveKeyMD5(UDPKeyString)
```

## New Code Needed

1. **Encryption** - Reverse of your decryption logic (same AES-ECB, just encrypt instead of decrypt)
2. **Packet Encoding** - Build outbound packets with proper header/footer for TCP queries
3. **PKCS7 Padding Addition** - Add padding before encryption

## Interface Design

Consider defining interfaces for testability:

```go
type Cipher interface {
    Encrypt(plaintext []byte) ([]byte, error)
    Decrypt(ciphertext []byte) ([]byte, error)
}

type Codec interface {
    Encode(cmd uint32, payload []byte, seq uint32) ([]byte, error)
    Decode(data []byte) (*Packet, error)
}
```

## Hints

### Basic Hint
Look at your existing `tryDecrypt()` function. Encryption is the same process but using `cipher.Encrypt()` instead of `cipher.Decrypt()`.

### Intermediate Hint
For packet encoding:
1. Calculate total length (header + retcode + encrypted_payload + crc + suffix)
2. Build header with `binary.BigEndian.PutUint32()`
3. Append encrypted payload
4. Append footer

### Complete Solution Pattern
See `docs/TUYA_PROTOCOL.md` for exact byte offsets and format.

## Testing

Create test vectors by capturing real packets with your existing code, then use those to verify your encode/decode round-trips correctly.

```bash
go test -v ./protocol/
```

## Think About

1. Why does `parseTuyaPacket()` need to detect version strings?
2. What happens if you try to decrypt with the wrong key?
3. How would you handle packets that don't have proper padding?
