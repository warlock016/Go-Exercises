package monitor

import (
	"context"

	"github.com/warlock16/udp_decoder/device"
)

// Monitor captures UDP packets from Tuya devices
type Monitor struct {
	iface    string
	registry *device.Registry
	updates  chan<- Update
}

// Update represents a device state change from passive monitoring
type Update struct {
	DeviceID string
	IP       string
	DPs      map[string]interface{}
}

// NewMonitor creates a new packet monitor
func NewMonitor(iface string, registry *device.Registry, updates chan<- Update) *Monitor {
	// TODO(human): Store dependencies
	return nil
}

// Start begins packet capture (blocks until context cancelled)
func (m *Monitor) Start(ctx context.Context) error {
	// TODO(human): Extract from main.go startCapture()
	// Add context cancellation support
	return nil
}
