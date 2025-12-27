package protocol

// Tuya Protocol 3.3 Constants

const (
	Prefix uint32 = 0x000055AA
	Suffix uint32 = 0x0000AA55

	HeaderSize = 16 // prefix + seq + cmd + length
	FooterSize = 8  // crc + suffix

	// TODO(human): Define command codes (see docs/TUYA_PROTOCOL.md)
)

// Packet represents a decoded Tuya protocol packet
type Packet struct {
	// TODO(human): Define fields based on protocol structure
}
