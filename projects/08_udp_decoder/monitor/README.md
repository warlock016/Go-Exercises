# Monitor Package

Captures UDP packets and correlates them to registered devices.

## Learning Goals

- Use gopacket for network packet capture
- Extract relevant layers (IPv4, UDP)
- Correlate packets to devices by IP address
- Run continuous capture in a goroutine with context cancellation

## Files to Create

| File | Purpose |
|------|---------|
| `monitor.go` | Packet capture setup and main loop |
| `handler.go` | Process individual packets |

## What to Extract from main.go

Your existing code has the capture logic. Extract and refactor:

| Function | Lines | Purpose |
|----------|-------|---------|
| `startCapture()` | 112-143 | Opens interface, sets filter, processes packets |
| `processPacket()` | 145-190 | Extracts layers, decodes payload |

## Key Changes from Original

1. **Registry Integration:** Lookup device by source IP
2. **State Updates:** Update device state instead of just printing
3. **Context Support:** Stop capture when context cancelled
4. **Channel Output:** Send updates to a channel for consumers

## Hints

### Basic Hint
The gopacket capture loop can be stopped by closing the handle:
```go
go func() {
    <-ctx.Done()
    handle.Close()
}()
```

### Intermediate Hint
For correlating packets to devices:
```go
device, err := registry.GetByIP(ip.SrcIP.String())
if err != nil {
    // Unknown device, ignore or log
    continue
}
```

### Channel Pattern
Use buffered channel with non-blocking send:
```go
select {
case m.updates <- update:
    // sent
default:
    // channel full, drop
}
```

## Testing

Mock the packet source for unit tests. Test that updates are correctly sent to the channel.

## Think About

1. What happens if a packet comes from an unknown IP?
2. How do you prevent the updates channel from blocking?
3. Should you capture promiscuously or just your interface's traffic?
