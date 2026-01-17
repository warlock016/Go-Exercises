package mutex_shared_state

import "sync"

// SafeCounter is a concurrent-safe counter
type SafeCounter struct {
	// TODO(human): Define fields
	mu      sync.Mutex
	counter int
}

// NewSafeCounter creates a new SafeCounter
func NewSafeCounter() *SafeCounter {
	// TODO(human): Implement
	return &SafeCounter{}
}

// Inc increments the counter by 1
func (c *SafeCounter) Inc() {
	// TODO(human): Implement
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

// Dec decrements the counter by 1
func (c *SafeCounter) Dec() {
	// TODO(human): Implement
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter--
}

// Value returns the current counter value
func (c *SafeCounter) Value() int {
	// TODO(human): Implement
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

// SafeMap is a concurrent-safe string->int map
type SafeMap struct {
	// TODO(human): Define fields
	mu sync.Mutex
	m  map[string]int
}

// NewSafeMap creates a new SafeMap
func NewSafeMap() *SafeMap {
	// TODO(human): Implement
	return &SafeMap{
		m: make(map[string]int),
	}
}

// Set stores a key-value pair
func (m *SafeMap) Set(key string, value int) {
	// TODO(human): Implement
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = value
}

// Get retrieves a value by key
func (m *SafeMap) Get(key string) (int, bool) {
	// TODO(human): Implement
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.m[key]; !ok {
		return 0, false
	}
	return m.m[key], true
}

// Delete removes a key
func (m *SafeMap) Delete(key string) {
	// TODO(human): Implement
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.m, key)
}

// Len returns the number of entries
func (m *SafeMap) Len() int {
	// TODO(human): Implement
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.m)
	// return 0
}

// SafeCache uses RWMutex for read-heavy workloads
type SafeCache struct {
	// TODO(human): Define fields
	mu sync.RWMutex
	m  map[string]string
}

// NewSafeCache creates a new SafeCache
func NewSafeCache() *SafeCache {
	// TODO(human): Implement
	return &SafeCache{
		m: make(map[string]string),
	}
}

// Get retrieves a value (read lock)
func (c *SafeCache) Get(key string) (string, bool) {
	// TODO(human): Implement
	c.mu.RLock()
	defer c.mu.RUnlock()
	if _, ok := c.m[key]; !ok {
		return "", false
	}
	return c.m[key], true
	// return "", false
}

// Set stores a value (write lock)
func (c *SafeCache) Set(key string, value string) {
	// TODO(human): Implement
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

// GetOrSet returns existing value or sets and returns new value
func (c *SafeCache) GetOrSet(key string, value string) string {
	// TODO(human): Implement - careful with lock upgrade!

	c.mu.RLock()
	if v, ok := c.m[key]; ok {
		c.mu.RUnlock()
		return v
	}
	c.mu.RUnlock()

	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.m[key]; ok {
		return v
	}
	c.m[key] = value
	return value
}
