package protocol

// Encode creates a wire-format packet for sending to a device
func Encode(cmd uint32, payload []byte, key []byte, seq uint32) ([]byte, error) {
	// TODO(human): Build packet with header, encrypted payload, footer
	return nil, nil
}

// Decode parses a wire-format packet and decrypts the payload
func Decode(data []byte, key []byte) (*Packet, error) {
	// TODO(human): Extract from main.go parseTuyaPacket()
	return nil, nil
}
