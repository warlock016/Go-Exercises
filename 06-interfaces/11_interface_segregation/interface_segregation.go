package interface_segregation

// TODO(human): Define Creator interface

// TODO(human): Define Reader interface

// TODO(human): Define Updater interface

// TODO(human): Define Deleter interface

// TODO(human): Define Lister interface

// TODO(human): Define MemoryStore struct

// TODO(human): Implement Create() for MemoryStore

// TODO(human): Implement Read() for MemoryStore

// TODO(human): Implement Update() for MemoryStore

// TODO(human): Implement Delete() for MemoryStore

// TODO(human): Implement List() for MemoryStore

// CopyData copies a single record from one store to another
func CopyData(from Reader, to Creator) error {
	// TODO(human): Implement
	return nil
}

// MigrateAll migrates all records from one store to another
func MigrateAll(from Reader, fromList Lister, to Creator) error {
	// TODO(human): Implement
	return nil
}
