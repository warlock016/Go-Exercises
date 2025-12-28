package mock_for_testing

import "time"

// TODO(human): Define TimeProvider interface
type TimeProvider interface {
	Now() time.Time
}

// TODO(human): Define RealTime struct
type RealTime struct{}

// TODO(human): Define MockTime struct
type MockTime struct {
	CurrentTime time.Time
}

// TODO(human): Define Scheduler struct
type Scheduler struct {
	TimeProvider TimeProvider
}

// TODO(human): Implement Now() for RealTime
func (t *RealTime) Now() time.Time {
	return time.Now()
}

// TODO(human): Implement Now() for MockTime
func (t *MockTime) Now() time.Time {
	return t.CurrentTime
}

// ShouldRun returns true if current time is during business hours (9 AM - 5 PM)
func (s *Scheduler) ShouldRun() bool {
	// TODO(human): Implement
	now := s.TimeProvider.Now().Hour()
	if now >= 9 && now < 17 {
		return true
	}
	return false
}

// TimeUntilNextRun returns duration until next business hour (9 AM)
func (s *Scheduler) TimeUntilNextRun() time.Duration {
	// TODO(human): Implement
	var next time.Time
	switch {
	case s.TimeProvider.Now().Hour() < 9:
		next = time.Date(s.TimeProvider.Now().Year(), time.Month(s.TimeProvider.Now().Month()), s.TimeProvider.Now().Day(), 9, 0, 0, 0, time.UTC)
	case s.TimeProvider.Now().Hour() >= 17:
		next = time.Date(s.TimeProvider.Now().Year(), time.Month(s.TimeProvider.Now().Month()), s.TimeProvider.Now().Day()+1, 9, 0, 0, 0, time.UTC)
	default:
		next = s.TimeProvider.Now()
	}
	return next.Sub(s.TimeProvider.Now())
}
