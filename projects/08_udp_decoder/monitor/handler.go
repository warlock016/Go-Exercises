package monitor

import (
	"github.com/google/gopacket"
)

// processPacket handles a captured UDP packet
func (m *Monitor) processPacket(packet gopacket.Packet) {
	// TODO(human): Extract from main.go processPacket()
	// 1. Extract IPv4 and UDP layers
	// 2. Lookup device by source IP
	// 3. Decode and decrypt payload
	// 4. Parse DPs from JSON
	// 5. Send update to channel
}
