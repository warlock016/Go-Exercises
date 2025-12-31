package mutex_shared_state

import (
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		c := NewSafeCounter()
		if c == nil {
			t.Fatal("NewSafeCounter() returned nil")
		}

		if v := c.Value(); v != 0 {
			t.Errorf("Initial value = %d, want 0", v)
		}

		c.Inc()
		if v := c.Value(); v != 1 {
			t.Errorf("After Inc() = %d, want 1", v)
		}

		c.Inc()
		c.Inc()
		if v := c.Value(); v != 3 {
			t.Errorf("After 3 Inc() = %d, want 3", v)
		}

		c.Dec()
		if v := c.Value(); v != 2 {
			t.Errorf("After Dec() = %d, want 2", v)
		}
	})

	t.Run("concurrent increments", func(t *testing.T) {
		c := NewSafeCounter()
		var wg sync.WaitGroup

		n := 1000
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.Inc()
			}()
		}
		wg.Wait()

		if v := c.Value(); v != n {
			t.Errorf("After %d concurrent Inc() = %d, want %d", n, v, n)
		}
	})

	t.Run("concurrent mixed operations", func(t *testing.T) {
		c := NewSafeCounter()
		var wg sync.WaitGroup

		// 500 increments, 300 decrements = 200
		for i := 0; i < 500; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.Inc()
			}()
		}
		for i := 0; i < 300; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.Dec()
			}()
		}
		wg.Wait()

		if v := c.Value(); v != 200 {
			t.Errorf("After mixed operations = %d, want 200", v)
		}
	})
}

func TestSafeMap(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		m := NewSafeMap()
		if m == nil {
			t.Fatal("NewSafeMap() returned nil")
		}

		if l := m.Len(); l != 0 {
			t.Errorf("Initial Len() = %d, want 0", l)
		}

		m.Set("a", 1)
		m.Set("b", 2)

		if l := m.Len(); l != 2 {
			t.Errorf("Len() after 2 Set() = %d, want 2", l)
		}

		v, ok := m.Get("a")
		if !ok || v != 1 {
			t.Errorf("Get(a) = %d, %v, want 1, true", v, ok)
		}

		v, ok = m.Get("missing")
		if ok {
			t.Errorf("Get(missing) ok = %v, want false", ok)
		}

		m.Delete("a")
		v, ok = m.Get("a")
		if ok {
			t.Errorf("Get(a) after Delete ok = %v, want false", ok)
		}
	})

	t.Run("concurrent writes", func(t *testing.T) {
		m := NewSafeMap()
		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				key := string(rune('a' + i%26))
				m.Set(key, i)
			}(i)
		}
		wg.Wait()

		// Should have at most 26 keys (a-z)
		if l := m.Len(); l > 26 {
			t.Errorf("Len() = %d, want <= 26", l)
		}
	})

	t.Run("concurrent read and write", func(t *testing.T) {
		m := NewSafeMap()
		m.Set("key", 0)

		var wg sync.WaitGroup
		done := make(chan struct{})

		// Writer
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				m.Set("key", i)
			}
			close(done)
		}()

		// Readers
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case <-done:
						return
					default:
						m.Get("key") // Should not panic
					}
				}
			}()
		}

		wg.Wait()
	})
}

func TestSafeCache(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		c := NewSafeCache()
		if c == nil {
			t.Fatal("NewSafeCache() returned nil")
		}

		v, ok := c.Get("missing")
		if ok {
			t.Errorf("Get(missing) ok = %v, want false", ok)
		}

		c.Set("key", "value")
		v, ok = c.Get("key")
		if !ok || v != "value" {
			t.Errorf("Get(key) = %q, %v, want 'value', true", v, ok)
		}
	})

	t.Run("GetOrSet new key", func(t *testing.T) {
		c := NewSafeCache()

		v := c.GetOrSet("new", "value")
		if v != "value" {
			t.Errorf("GetOrSet(new, value) = %q, want 'value'", v)
		}

		// Verify it was set
		v, ok := c.Get("new")
		if !ok || v != "value" {
			t.Errorf("Get(new) = %q, %v, want 'value', true", v, ok)
		}
	})

	t.Run("GetOrSet existing key", func(t *testing.T) {
		c := NewSafeCache()
		c.Set("existing", "original")

		v := c.GetOrSet("existing", "new")
		if v != "original" {
			t.Errorf("GetOrSet(existing, new) = %q, want 'original'", v)
		}
	})

	t.Run("concurrent reads", func(t *testing.T) {
		c := NewSafeCache()
		c.Set("key", "value")

		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < 100; j++ {
					v, ok := c.Get("key")
					if !ok || v != "value" {
						t.Errorf("Concurrent Get() = %q, %v", v, ok)
					}
				}
			}()
		}
		wg.Wait()
	})

	t.Run("concurrent GetOrSet", func(t *testing.T) {
		c := NewSafeCache()

		var wg sync.WaitGroup
		results := make(chan string, 100)

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				v := c.GetOrSet("race", "value")
				results <- v
			}(i)
		}
		wg.Wait()
		close(results)

		// All results should be "value"
		for v := range results {
			if v != "value" {
				t.Errorf("Concurrent GetOrSet returned %q, want 'value'", v)
			}
		}
	})
}

// Benchmarks to demonstrate RWMutex benefits
func BenchmarkSafeCounter(b *testing.B) {
	c := NewSafeCounter()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Inc()
		}
	})
}

func BenchmarkSafeMapWrite(b *testing.B) {
	m := NewSafeMap()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m.Set("key", i)
			i++
		}
	})
}

func BenchmarkSafeCacheRead(b *testing.B) {
	c := NewSafeCache()
	c.Set("key", "value")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get("key")
		}
	})
}
