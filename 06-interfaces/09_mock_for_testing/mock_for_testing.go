package mock_for_testing

// TODO(human): Define TimeProvider interface

// TODO(human): Define RealTime struct

// TODO(human): Define MockTime struct

// TODO(human): Define Scheduler struct

// TODO(human): Implement Now() for RealTime

// TODO(human): Implement Now() for MockTime

// ShouldRun returns true if current time is during business hours (9 AM - 5 PM)
func (s *Scheduler) ShouldRun() bool {
	// TODO(human): Implement
	return false
}

// TimeUntilNextRun returns duration until next business hour (9 AM)
func (s *Scheduler) TimeUntilNextRun() time.Duration {
	// TODO(human): Implement
	return 0
}
