package trafficlight

// TrafficLight represents a simple traffic light controller
type TrafficLight struct {
	// TODO(human): Add fields to track:
	// - Current state (RED, GREEN, or YELLOW)
	// - Number of complete cycles (how many times returned to RED from YELLOW)
	state  string
	cycles int
}

// NewTrafficLight creates a new traffic light starting in RED state
func NewTrafficLight() *TrafficLight {
	// TODO(human): Initialize the traffic light
	// Hint: Start in "RED" state with 0 cycles

	return &TrafficLight{
		state:  "RED",
		cycles: 0,
	}
}

// Next advances the traffic light to the next state
// State transitions: RED → GREEN → YELLOW → RED
func (t *TrafficLight) Next() {

	switch t.state {
	case "RED":
		t.state = "GREEN"
	case "GREEN":
		t.state = "YELLOW"
	case "YELLOW":
		t.state = "RED"
		t.cycles++
	}
	// TODO(human): Implement state transition
	//
	// Pattern: Use a switch statement on current state
	// - RED → GREEN
	// - GREEN → YELLOW
	// - YELLOW → RED (and increment cycles)
	//
	// Think: When should you increment the cycle count?
	// Answer: When transitioning from YELLOW back to RED
}

// Current returns the current state of the traffic light
func (t *TrafficLight) Current() string {
	// TODO(human): Return the current state
	return t.state
}

// Cycles returns the number of complete cycles (RED→GREEN→YELLOW→RED)
func (t *TrafficLight) Cycles() int {
	// TODO(human): Return the cycle count
	return t.cycles
}
