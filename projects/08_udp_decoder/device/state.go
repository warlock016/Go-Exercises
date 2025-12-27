package device

import "time"

// State represents the current known state of a device
type State struct {
	DPs         map[string]interface{} // Data points: {"1": true, "2": 25}
	LastUpdated time.Time
	Source      string // "passive" or "active"
}

// NewState creates an empty state
func NewState() *State {
	// TODO(human): Initialize with empty DPs map
	return nil
}

// Update merges new DPs into the state
func (s *State) Update(dps map[string]interface{}, source string) {
	// TODO(human): Merge DPs, update timestamp and source
}

// Age returns how long since last update
func (s *State) Age() time.Duration {
	// TODO(human): Calculate duration since LastUpdated
	return 0
}

// Copy returns a deep copy of the state (for thread-safe reads)
func (s *State) Copy() *State {
	// TODO(human): Create copy with cloned DPs map
	return nil
}
