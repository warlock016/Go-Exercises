package interface_segregation

import "errors"

// TODO(human): Define Creator interface
type Creator interface {
	Create(string, string) error
}

// TODO(human): Define Reader interface
type Reader interface {
	Read(string) (string, error)
}

// TODO(human): Define Updater interface
type Updater interface {
	Update(string, string) error
}

// TODO(human): Define Deleter interface
type Deleter interface {
	Delete(string) error
}

// TODO(human): Define Lister interface
type Lister interface {
	List() []string
}

// TODO(human): Define MemoryStore struct
type MemoryStore struct {
	data map[string]string
}

// TODO(human): Implement Create() for MemoryStore
func (m *MemoryStore) Create(id, value string) error {

	if _, ok := m.data[id]; ok {
		return errors.New("id already exists")
	}
	m.data[id] = value
	return nil
}

// TODO(human): Implement Read() for MemoryStore
func (m *MemoryStore) Read(id string) (string, error) {

	if value, ok := m.data[id]; ok {
		return value, nil
	}

	return "", errors.New("rd: id not found")
}

// TODO(human): Implement Update() for MemoryStore
func (m *MemoryStore) Update(id, value string) error {

	if _, ok := m.data[id]; !ok {
		return errors.New("upd: id not found")
	}

	m.data[id] = value
	return nil
}

// TODO(human): Implement Delete() for MemoryStore
func (m *MemoryStore) Delete(id string) error {

	if _, ok := m.data[id]; !ok {
		return errors.New("del: id not found")
	}
	delete(m.data, id)
	return nil
}

// TODO(human): Implement List() for MemoryStore
func (m *MemoryStore) List() []string {
	result := make([]string, 0, len(m.data))

	for k := range m.data {
		result = append(result, k)
	}
	return result
}

// CopyData copies a single record from one store to another
func CopyData(from Reader, to Creator, id string) error {
	// TODO(human): Implement

	res, err := from.Read(id)
	if err != nil {
		return err
	}
	err = to.Create(id, res)

	return err
}

// MigrateAll migrates all records from one store to another
func MigrateAll(from Reader, fromList Lister, to Creator) error {
	// TODO(human): Implement

	for _, i := range fromList.List() {
		res, err := from.Read(i)
		if err != nil {
			return err
		}
		err = to.Create(i, res)
		if err != nil {
			return err
		}
	}

	return nil
}
