package device

import (
	"errors"
	"sync"
)

var (
	ErrDeviceNotFound = errors.New("device not found")
	ErrDeviceExists   = errors.New("device already registered")
)

// Registry stores and manages all known devices
type Registry struct {
	mu   sync.RWMutex
	byID map[string]*Device
	byIP map[string]*Device
}

// NewRegistry creates an empty device registry
func NewRegistry() *Registry {
	// TODO(human): Initialize maps
	return nil
}

// Register adds a device to the registry
func (r *Registry) Register(device *Device) error {
	// TODO(human): Add to both maps, check for duplicates
	return nil
}

// Get retrieves a device by ID
func (r *Registry) Get(id string) (*Device, error) {
	// TODO(human): Lookup with RLock
	return nil, nil
}

// GetByIP retrieves a device by IP address
func (r *Registry) GetByIP(ip string) (*Device, error) {
	// TODO(human): Lookup with RLock
	return nil, nil
}

// List returns all registered devices
func (r *Registry) List() []*Device {
	// TODO(human): Return slice of all devices
	return nil
}

// Remove deletes a device from the registry
func (r *Registry) Remove(id string) error {
	// TODO(human): Remove from both maps
	return nil
}
