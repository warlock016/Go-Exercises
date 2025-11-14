package trafficlight

import "testing"

func TestNewTrafficLight(t *testing.T) {
	light := NewTrafficLight()
	if light.Current() != "RED" {
		t.Errorf("New light should start in RED, got %s", light.Current())
	}
	if light.Cycles() != 0 {
		t.Errorf("New light should have 0 cycles, got %d", light.Cycles())
	}
}

func TestStateTransitions(t *testing.T) {
	light := NewTrafficLight()

	// RED → GREEN
	light.Next()
	if light.Current() != "GREEN" {
		t.Errorf("After Next() from RED, expected GREEN, got %s", light.Current())
	}

	// GREEN → YELLOW
	light.Next()
	if light.Current() != "YELLOW" {
		t.Errorf("After Next() from GREEN, expected YELLOW, got %s", light.Current())
	}

	// YELLOW → RED (should increment cycle)
	light.Next()
	if light.Current() != "RED" {
		t.Errorf("After Next() from YELLOW, expected RED, got %s", light.Current())
	}
	if light.Cycles() != 1 {
		t.Errorf("After completing one cycle, expected Cycles()=1, got %d", light.Cycles())
	}
}

func TestMultipleCycles(t *testing.T) {
	light := NewTrafficLight()

	// Complete 3 full cycles
	for i := 0; i < 9; i++ { // 3 cycles × 3 transitions
		light.Next()
	}

	if light.Current() != "RED" {
		t.Errorf("After 9 transitions, should be RED, got %s", light.Current())
	}
	if light.Cycles() != 3 {
		t.Errorf("After 9 transitions, should have 3 cycles, got %d", light.Cycles())
	}
}
