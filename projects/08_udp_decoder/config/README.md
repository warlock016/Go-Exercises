# Config Package

Parses configuration from `.env` file, supporting multiple devices.

## Learning Goals

- Parse environment files with `go-envparse`
- Use regex to extract numbered device entries
- Validate configuration completeness
- Design configuration structs

## Files to Create

| File | Purpose |
|------|---------|
| `config.go` | Configuration parsing and validation |

## .env Format

```env
# Server settings
API_PORT=8080
MONITOR_INTERFACE=en0

# Device 1
DEVICE_1_NAME=SmartMeter
DEVICE_1_ID=bf446c6d6fe6f70e16on63
DEVICE_1_LOCAL_KEY=)+cBrK`YqEy}wlX=
DEVICE_1_IP=192.168.1.100
DEVICE_1_MAC=AA:BB:CC:DD:EE:FF

# Device 2
DEVICE_2_NAME=SmartPlug
DEVICE_2_ID=bf51e799667e9b1fbazfkx
DEVICE_2_LOCAL_KEY==j(Brn(YJuA4n'1O
DEVICE_2_IP=192.168.1.101
```

## Required Fields per Device

| Field | Required | Description |
|-------|----------|-------------|
| `DEVICE_N_ID` | Yes | Device ID from Tuya |
| `DEVICE_N_LOCAL_KEY` | Yes | 16-byte encryption key |
| `DEVICE_N_IP` | Yes | Device IP address |
| `DEVICE_N_NAME` | No | Friendly name |
| `DEVICE_N_MAC` | No | MAC address |

## Parsing Strategy

1. Load all key-value pairs from `.env`
2. Find all `DEVICE_N_*` entries using regex
3. Group by device number (N)
4. Validate each device has required fields
5. Return configuration struct

## Hints

### Basic Hint
Use regex to match device entries:
```go
pattern := regexp.MustCompile(`^DEVICE_(\d+)_(\w+)$`)
```

### Intermediate Hint
Group entries by device number:
```go
deviceEntries := make(map[string]map[string]string)
for key, value := range envMap {
    matches := pattern.FindStringSubmatch(key)
    if matches != nil {
        num, field := matches[1], matches[2]
        if deviceEntries[num] == nil {
            deviceEntries[num] = make(map[string]string)
        }
        deviceEntries[num][field] = value
    }
}
```

### Validation Pattern
Check required fields exist:
```go
required := []string{"ID", "LOCAL_KEY", "IP"}
for _, field := range required {
    if entry[field] == "" {
        return fmt.Errorf("device %s missing %s", num, field)
    }
}
```

## Testing

Create test `.env` files with valid and invalid configurations. Test that missing required fields return errors.

```bash
go test -v ./config/
```

## Think About

1. How do you handle a `.env` file that doesn't exist?
2. Should you strip quotes from values?
3. What if a device has ID but no IP?
