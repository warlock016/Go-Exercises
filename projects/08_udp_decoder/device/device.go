package device

import (
	"context"
	"net"
	"sync"
)

// Device represents a Tuya IoT device
type Device struct {
	ID       string
	Name     string
	IP       string
	LocalKey string
	MAC      string

	mu    sync.RWMutex
	conn  net.Conn
	state *State
	seq   uint32
}

// NewDevice creates a new device from configuration
func NewDevice(id, name, ip, localKey, mac string) (*Device, error) {
	// TODO(human): Validate inputs, create device
	return nil, nil
}

// Connect establishes a TCP connection to the device
func (d *Device) Connect(ctx context.Context) error {
	// TODO(human): Dial TCP to port 6668
	return nil
}

// Disconnect closes the device connection
func (d *Device) Disconnect() error {
	// TODO(human): Close connection if open
	return nil
}

// Query sends a DP_QUERY and returns current measurements
func (d *Device) Query(ctx context.Context) (*State, error) {
	// TODO(human): Send query packet, read response, decrypt, parse
	return nil, nil
}

// SendHeartbeat sends a keepalive to the device
func (d *Device) SendHeartbeat(ctx context.Context) error {
	// TODO(human): Send heartbeat command
	return nil
}

// GetState returns the cached device state
func (d *Device) GetState() *State {
	// TODO(human): Return copy of state (thread-safe)
	return nil
}

// UpdateState updates cached state from passive monitoring
func (d *Device) UpdateState(dps map[string]interface{}, source string) {
	// TODO(human): Update state with mutex protection
}

// nextSeq returns the next sequence number
func (d *Device) nextSeq() uint32 {
	// TODO(human): Atomically increment and return sequence
	return 0
}
