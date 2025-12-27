package config

// Config holds all application configuration
type Config struct {
	APIPort          int
	MonitorInterface string
	Devices          []DeviceConfig
}

// DeviceConfig holds configuration for a single device
type DeviceConfig struct {
	ID       string
	Name     string
	IP       string
	LocalKey string
	MAC      string
}

// Load reads configuration from the specified .env file
func Load(path string) (*Config, error) {
	// TODO(human): Open file, parse env, extract devices
	return nil, nil
}

// parseDevices extracts device configurations from env map
func parseDevices(env map[string]string) ([]DeviceConfig, error) {
	// TODO(human): Use regex to find DEVICE_N_* entries, group by N
	return nil, nil
}

// validateDevice checks that required fields are present
func validateDevice(num string, entry map[string]string) error {
	// TODO(human): Check ID, LOCAL_KEY, IP are present
	return nil
}
