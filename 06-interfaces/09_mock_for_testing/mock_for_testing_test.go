package mock_for_testing

import (
	"testing"
	"time"
)

func TestRealTimeNow(t *testing.T) {
	rt := &RealTime{}
	now := rt.Now()

	// Should be close to actual time.Now()
	diff := time.Since(now)
	if diff > time.Second {
		t.Errorf("RealTime.Now() is %v in the past, should be current time", diff)
	}
}

func TestMockTimeNow(t *testing.T) {
	fixedTime := time.Date(2024, 1, 1, 10, 30, 0, 0, time.UTC)
	mt := &MockTime{CurrentTime: fixedTime}

	got := mt.Now()
	if got != fixedTime {
		t.Errorf("MockTime.Now() = %v, want %v", got, fixedTime)
	}
}

func TestSchedulerShouldRun(t *testing.T) {
	tests := []struct {
		name string
		hour int
		want bool
	}{
		{"Before business hours - 8 AM", 8, false},
		{"Start of business hours - 9 AM", 9, true},
		{"Mid day - 12 PM", 12, true},
		{"End of business hours - 4 PM", 16, true},
		{"After business hours - 5 PM", 17, false},
		{"After business hours - 6 PM", 18, false},
		{"Night - 11 PM", 23, false},
		{"Early morning - 3 AM", 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTime := &MockTime{
				CurrentTime: time.Date(2024, 1, 1, tt.hour, 0, 0, 0, time.UTC),
			}
			scheduler := &Scheduler{TimeProvider: mockTime}

			got := scheduler.ShouldRun()
			if got != tt.want {
				t.Errorf("ShouldRun() at hour %d = %v, want %v", tt.hour, got, tt.want)
			}
		})
	}
}

func TestSchedulerTimeUntilNextRun(t *testing.T) {
	tests := []struct {
		name     string
		hour     int
		minute   int
		wantHour time.Duration
	}{
		{"8:00 AM - 1 hour until 9 AM", 8, 0, 1 * time.Hour},
		{"8:30 AM - 30 min until 9 AM", 8, 30, 30 * time.Minute},
		{"6:00 PM - 15 hours until 9 AM", 18, 0, 15 * time.Hour},
		{"11:00 PM - 10 hours until 9 AM", 23, 0, 10 * time.Hour},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTime := &MockTime{
				CurrentTime: time.Date(2024, 1, 1, tt.hour, tt.minute, 0, 0, time.UTC),
			}
			scheduler := &Scheduler{TimeProvider: mockTime}

			got := scheduler.TimeUntilNextRun()
			if got != tt.wantHour {
				t.Errorf("TimeUntilNextRun() = %v, want %v", got, tt.wantHour)
			}
		})
	}
}

func TestImplementsTimeProvider(t *testing.T) {
	var _ TimeProvider = &RealTime{}
	var _ TimeProvider = &MockTime{}
	t.Log("✓ Both RealTime and MockTime implement TimeProvider interface")
}

func TestDependencyInjection(t *testing.T) {
	// Test that we can swap implementations
	scheduler := &Scheduler{}

	// Use RealTime in production
	scheduler.TimeProvider = &RealTime{}
	_ = scheduler.ShouldRun()

	// Use MockTime in tests
	scheduler.TimeProvider = &MockTime{
		CurrentTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
	}
	shouldRun := scheduler.ShouldRun()

	if !shouldRun {
		t.Error("Scheduler with 10 AM time should run")
	}
}
