package options_pattern

import (
	"testing"
	"time"
)

func TestNewServerDefaults(t *testing.T) {
	server := NewServer()

	if server.Host() != "localhost" {
		t.Errorf("Default host = %q, want %q", server.Host(), "localhost")
	}

	if server.Port() != 8080 {
		t.Errorf("Default port = %d, want 8080", server.Port())
	}

	if server.Timeout() != 30*time.Second {
		t.Errorf("Default timeout = %v, want 30s", server.Timeout())
	}

	if server.MaxConnections() != 100 {
		t.Errorf("Default maxConn = %d, want 100", server.MaxConnections())
	}
}

func TestWithHost(t *testing.T) {
	tests := []struct {
		name string
		host string
	}{
		{"IPv4", "0.0.0.0"},
		{"IPv6", "::"},
		{"domain", "example.com"},
		{"empty", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewServer(WithHost(tt.host))
			if server.Host() != tt.host {
				t.Errorf("Host() = %q, want %q", server.Host(), tt.host)
			}
		})
	}
}

func TestWithPort(t *testing.T) {
	tests := []struct {
		name string
		port int
	}{
		{"common port", 3000},
		{"http", 80},
		{"https", 443},
		{"custom", 9999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewServer(WithPort(tt.port))
			if server.Port() != tt.port {
				t.Errorf("Port() = %d, want %d", server.Port(), tt.port)
			}
		})
	}
}

func TestWithTimeout(t *testing.T) {
	tests := []struct {
		name    string
		timeout time.Duration
	}{
		{"short", 5 * time.Second},
		{"medium", 30 * time.Second},
		{"long", 2 * time.Minute},
		{"very long", 10 * time.Minute},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewServer(WithTimeout(tt.timeout))
			if server.Timeout() != tt.timeout {
				t.Errorf("Timeout() = %v, want %v", server.Timeout(), tt.timeout)
			}
		})
	}
}

func TestWithMaxConnections(t *testing.T) {
	tests := []struct {
		name    string
		maxConn int
	}{
		{"small", 10},
		{"medium", 100},
		{"large", 1000},
		{"unlimited", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewServer(WithMaxConnections(tt.maxConn))
			if server.MaxConnections() != tt.maxConn {
				t.Errorf("MaxConnections() = %d, want %d", server.MaxConnections(), tt.maxConn)
			}
		})
	}
}

func TestMultipleOptions(t *testing.T) {
	server := NewServer(
		WithHost("0.0.0.0"),
		WithPort(9000),
		WithTimeout(60*time.Second),
		WithMaxConnections(200),
	)

	if server.Host() != "0.0.0.0" {
		t.Errorf("Host() = %q, want %q", server.Host(), "0.0.0.0")
	}

	if server.Port() != 9000 {
		t.Errorf("Port() = %d, want 9000", server.Port())
	}

	if server.Timeout() != 60*time.Second {
		t.Errorf("Timeout() = %v, want 60s", server.Timeout())
	}

	if server.MaxConnections() != 200 {
		t.Errorf("MaxConnections() = %d, want 200", server.MaxConnections())
	}
}

func TestPartialOptions(t *testing.T) {
	// Only set port and maxConn, others should use defaults
	server := NewServer(
		WithPort(3000),
		WithMaxConnections(500),
	)

	if server.Host() != "localhost" {
		t.Errorf("Host() = %q, want %q (default)", server.Host(), "localhost")
	}

	if server.Port() != 3000 {
		t.Errorf("Port() = %d, want 3000", server.Port())
	}

	if server.Timeout() != 30*time.Second {
		t.Errorf("Timeout() = %v, want 30s (default)", server.Timeout())
	}

	if server.MaxConnections() != 500 {
		t.Errorf("MaxConnections() = %d, want 500", server.MaxConnections())
	}
}

func TestOptionsOrder(t *testing.T) {
	// Options should be applied in order, last one wins
	server := NewServer(
		WithPort(8080),
		WithPort(9000),
		WithPort(3000),
	)

	if server.Port() != 3000 {
		t.Errorf("Port() = %d, want 3000 (last option should win)", server.Port())
	}
}

func TestOptionsReuse(t *testing.T) {
	// Options can be stored and reused
	productionOpts := []Option{
		WithHost("0.0.0.0"),
		WithTimeout(120 * time.Second),
		WithMaxConnections(500),
	}

	server1 := NewServer(productionOpts...)
	server2 := NewServer(productionOpts...)

	// Both should have same configuration
	if server1.Host() != server2.Host() {
		t.Error("Reused options should produce same configuration")
	}

	// But they should be different instances
	if server1 == server2 {
		t.Error("NewServer should create new instances")
	}
}

func TestOptionsCombination(t *testing.T) {
	// Test combining option slices
	baseOpts := []Option{
		WithHost("localhost"),
		WithPort(8080),
	}

	devOpts := []Option{
		WithTimeout(10 * time.Second),
	}

	server := NewServer(append(baseOpts, devOpts...)...)

	if server.Host() != "localhost" {
		t.Errorf("Host() = %q, want %q", server.Host(), "localhost")
	}

	if server.Port() != 8080 {
		t.Errorf("Port() = %d, want 8080", server.Port())
	}

	if server.Timeout() != 10*time.Second {
		t.Errorf("Timeout() = %v, want 10s", server.Timeout())
	}
}

func TestServerIndependence(t *testing.T) {
	// Test that servers are independent
	server1 := NewServer(WithPort(3000))
	server2 := NewServer(WithPort(4000))

	if server1.Port() == server2.Port() {
		t.Error("Servers should have independent configurations")
	}

	if server1.Port() != 3000 {
		t.Errorf("Server1 port = %d, want 3000", server1.Port())
	}

	if server2.Port() != 4000 {
		t.Errorf("Server2 port = %d, want 4000", server2.Port())
	}
}
