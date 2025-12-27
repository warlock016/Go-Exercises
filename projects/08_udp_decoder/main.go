package main

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/hashicorp/go-envparse"
)

type Packet struct {
	Prefix     uint32 // 4 bytes should be 0x000055AA
	Sequence   uint32 // 4 bytes BE
	Command    uint32 // 4 bytes BE
	PayloadLen uint32 // 4 bytes BE
	ReturnCode uint32 // 4 bytes
	Payload    uint64 // variable length, encrypted
	CRC        uint32 // 4 bytes
	Suffix     uint32 //4 bytes should be 0x0000AA55
}

// Use encoding/binary with binary.BigEndian.Uint32() to read multi-byte integers.

func main() {
	// Command line flags
	iface := flag.String("i", "", "Network interface to capture on")
	target := flag.String("target", "", "Target device IP to filter (optional)")
	listIfaces := flag.Bool("list", false, "List available network interfaces")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to fetch $CWD: %v", err)
	}
	fullPath := path.Join(cwd, ".env")
	env, err := os.Open(fullPath)
	if err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}
	defer env.Close()

	res, err := envparse.Parse(env)
	if err != nil {
		log.Fatalf("failed to parse .env: %v", err)
	}

	if *listIfaces {
		listInterfaces()
		return
	}

	if *iface == "" {
		log.Fatal("Interface required. Use -list to see available interfaces, then -i <interface>")
	}

	fmt.Printf("Starting UDP sniffer on interface: %s\n", *iface)
	if *target != "" {
		fmt.Printf("Filtering for device: %s\n", *target)
	}
	fmt.Println("Press Ctrl+C to stop")

	// Collect all device identifiers for key derivation attempts
	deviceInfo := DeviceInfo{
		LocalKey: strings.Trim(res["SMART_PLUG_KEY"], "\""),
		DeviceID: strings.Trim(res["DEVICE_ID"], "\""),
		MAC:      strings.Trim(res["MAC_ADDRESS"], "\""),
		IP:       strings.Trim(res["SMART_PLUG_IP"], "\""),
	}

	fmt.Printf("DEBUG: LocalKey=%q (len=%d), DeviceID=%q\n", deviceInfo.LocalKey, len(deviceInfo.LocalKey), deviceInfo.DeviceID)

	if err := startCapture(*iface, *target, deviceInfo); err != nil {
		log.Fatalf("Capture failed: %v", err)
	}
}

// DeviceInfo holds device identifiers for key derivation
type DeviceInfo struct {
	LocalKey string
	DeviceID string
	MAC      string
	IP       string
}

func listInterfaces() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf("Error finding devices: %v", err)
	}

	fmt.Println("Available network interfaces:")
	fmt.Println(strings.Repeat("-", 60))
	for _, device := range devices {
		fmt.Printf("Name: %s\n", device.Name)
		for _, addr := range device.Addresses {
			fmt.Printf("  IP: %s\n", addr.IP)
		}
		fmt.Println()
	}
}

func startCapture(iface, targetIP string, device DeviceInfo) error {
	// Open the interface for capture
	handle, err := pcap.OpenLive(
		iface,
		1600, // snapshot length
		true, // promiscuous mode
		pcap.BlockForever,
	)
	if err != nil {
		return fmt.Errorf("opening interface: %w", err)
	}
	defer handle.Close()

	// Build BPF filter
	filter := "udp"
	if targetIP != "" {
		filter = fmt.Sprintf("udp and host %s", targetIP)
	}
	if err := handle.SetBPFFilter(filter); err != nil {
		return fmt.Errorf("setting BPF filter: %w", err)
	}
	fmt.Printf("BPF filter: %s\n\n", filter)

	// Create packet source and process packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		processPacket(packet, device)
	}

	return nil
}

func processPacket(packet gopacket.Packet, device DeviceInfo) {
	// Extract IPv4 layer
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return
	}
	ip, _ := ipLayer.(*layers.IPv4)

	// Extract UDP layer
	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer == nil {
		return
	}
	udp, _ := udpLayer.(*layers.UDP)

	// Get timestamp
	timestamp := packet.Metadata().Timestamp
	if timestamp.IsZero() {
		timestamp = time.Now()
	}

	// Display packet info
	fmt.Printf("[%s] %s:%d -> %s:%d (%d bytes)\n",
		timestamp.Format("15:04:05.000"),
		ip.SrcIP, udp.SrcPort,
		ip.DstIP, udp.DstPort,
		len(udp.Payload),
	)

	// Display payload
	if len(udp.Payload) > 0 {
		// Check if this is a UDP broadcast (port 6667)
		isUDPBroadcast := udp.DstPort == 6667 || udp.SrcPort == 6667

		encr, err := parseTuyaPacket(udp.Payload, isUDPBroadcast)
		if err != nil {
			fmt.Printf("failed to parse Tuya packet: %v", err)
			// return
		} else {
			decr, err := decodePayload(encr, device, isUDPBroadcast)
			if err != nil {
				fmt.Printf("failed to decrypt payload: %v", err)
			} else {
				fmt.Println(string(decr))
			}
		}
	}
}

func parseTuyaPacket(data []byte, isUDPBroadcast bool) ([]byte, error) {
	const (
		Prefix = 0x000055AA
		Suffix = 0x0000AA55
	)

	if len(data) < 28 { // Minimum: header + some payload + footer
		return nil, errors.New("packet too short")
	}

	prefix := binary.BigEndian.Uint32(data[:4])
	if prefix != Prefix {
		return nil, errors.New("invalid prefix")
	}

	suffix := binary.BigEndian.Uint32(data[len(data)-4:])
	if suffix != Suffix {
		return nil, errors.New("invalid suffix")
	}

	sequence := binary.BigEndian.Uint32(data[4:8])
	command := binary.BigEndian.Uint32(data[8:12])
	payloadLen := binary.BigEndian.Uint32(data[12:16])
	retCode := binary.BigEndian.Uint32(data[16:20])

	fmt.Printf("seq: %d, cmd: 0x%02X, dataLen: %d, retCode: %d\n", sequence, command, payloadLen, retCode)

	// For UDP broadcasts, use simple extraction: data[20:-8]
	// This matches the tinytuya/localtuya approach
	if isUDPBroadcast {
		encryptedEnd := len(data) - 8
		payload := data[20:encryptedEnd]
		fmt.Printf("DEBUG: UDP broadcast payload range: [20:%d], size: %d bytes\n", encryptedEnd, len(payload))
		return payload, nil
	}

	// For TCP/other packets, use the more complex version detection
	fmt.Printf("DEBUG: bytes 16-35: %s\n", hex.EncodeToString(data[16:min(36, len(data))]))
	fmt.Printf("DEBUG: bytes as ASCII: %q\n", string(data[16:min(36, len(data))]))

	// Check if there's a "3.x" version string after retCode
	payloadStart := 20

	// Look for version indicator
	versionStr := ""
	if len(data) > 23 && data[20] == '3' && data[21] == '.' {
		versionStr = string(data[20:23])
		payloadStart = 23
		if len(data) > 35 && data[23] >= '0' && data[23] <= '9' {
			payloadStart = 35
		}
	}

	if versionStr != "" {
		fmt.Printf("DEBUG: detected version: %s\n", versionStr)
	}

	encryptedEnd := len(data) - 8
	encryptedLen := encryptedEnd - payloadStart

	// Ensure multiple of 16
	if encryptedLen%16 != 0 {
		for offset := payloadStart; offset < payloadStart+16 && offset < encryptedEnd-16; offset++ {
			testLen := encryptedEnd - offset
			if testLen%16 == 0 && testLen > 0 {
				fmt.Printf("DEBUG: adjusted payload start from %d to %d (len: %d)\n", payloadStart, offset, testLen)
				payloadStart = offset
				encryptedLen = testLen
				break
			}
		}
	}

	fmt.Printf("DEBUG: payload range: [%d:%d], size: %d bytes (%d blocks)\n",
		payloadStart, encryptedEnd, encryptedLen, encryptedLen/16)

	if encryptedLen <= 0 || encryptedLen%16 != 0 {
		return nil, fmt.Errorf("invalid encrypted payload size: %d", encryptedLen)
	}

	payload := data[payloadStart:encryptedEnd]
	return payload, nil
}

func decodePayload(payload []byte, device DeviceInfo, isUDPBroadcast bool) ([]byte, error) {
	// Build list of keys to try from device info
	keysToTry := []struct {
		name string
		key  []byte
	}{}

	if isUDPBroadcast {
		// UDP broadcasts use MD5 of the fixed UDP key
		keysToTry = append(keysToTry,
			struct {
				name string
				key  []byte
			}{"UDP broadcast key (MD5)", deriveKeyMD5("yGAdlopoPVldABfn")},
		)
	}

	// Device-specific keys (for TCP or if UDP key fails)
	keysToTry = append(keysToTry,
		struct {
			name string
			key  []byte
		}{"raw localKey", []byte(device.LocalKey)},
		struct {
			name string
			key  []byte
		}{"MD5(localKey)", deriveKeyMD5(device.LocalKey)},
	)

	// Add device ID based keys if available
	if device.DeviceID != "" {
		keysToTry = append(keysToTry,
			struct {
				name string
				key  []byte
			}{"MD5(deviceID)", deriveKeyMD5(device.DeviceID)},
			struct {
				name string
				key  []byte
			}{"deviceID[:16]", []byte(device.DeviceID)[:min(16, len(device.DeviceID))]},
			struct {
				name string
				key  []byte
			}{"MD5(deviceID+localKey)", deriveKeyMD5(device.DeviceID + device.LocalKey)},
			struct {
				name string
				key  []byte
			}{"MD5(localKey+deviceID)", deriveKeyMD5(device.LocalKey + device.DeviceID)},
		)
	}

	// Add MAC based keys if available
	if device.MAC != "" {
		macClean := strings.ReplaceAll(device.MAC, ":", "")
		keysToTry = append(keysToTry,
			struct {
				name string
				key  []byte
			}{"MD5(MAC)", deriveKeyMD5(macClean)},
			struct {
				name string
				key  []byte
			}{"MD5(MAC+localKey)", deriveKeyMD5(macClean + device.LocalKey)},
		)
	}

	for _, k := range keysToTry {
		if len(k.key) != 16 {
			fmt.Printf("DEBUG: skipping %s - wrong length %d\n", k.name, len(k.key))
			continue
		}

		fmt.Printf("DEBUG: trying %s: %x\n", k.name, k.key)

		result, err := tryDecrypt(payload, k.key)
		if err != nil {
			fmt.Printf("DEBUG: %s failed: %v\n", k.name, err)
			continue
		}

		// Check if result looks like JSON
		if len(result) > 0 && (result[0] == '{' || result[0] == '[') {
			fmt.Printf("DEBUG: %s succeeded!\n", k.name)
			return result, nil
		}
		fmt.Printf("DEBUG: %s decrypted but not JSON (first byte: 0x%02x)\n", k.name, result[0])
	}

	return nil, errors.New("all decryption attempts failed")
}

// deriveKeyMD5 creates an AES key by taking MD5 hash of the localKey (Tuya v3.3+)
func deriveKeyMD5(localKey string) []byte {
	hash := md5.Sum([]byte(localKey))
	return hash[:]
}

func tryDecrypt(payload []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	if len(payload)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("invalid payload size: %d is not multiple of %d", len(payload), block.BlockSize())
	}

	result := make([]byte, len(payload))

	// AES-ECB decryption: decrypt each 16-byte block independently
	for i := 0; i < len(payload); i += 16 {
		block.Decrypt(result[i:i+16], payload[i:i+16])
	}

	// Debug: show first bytes
	fmt.Printf("DEBUG: first 16 bytes: %s\n", hex.EncodeToString(result[:16]))

	// Remove PKCS7 padding
	result, err = removePKCS7Padding(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// removePKCS7Padding removes PKCS7 padding from decrypted data
func removePKCS7Padding(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("empty data")
	}

	// Last byte indicates padding length
	paddingLen := int(data[len(data)-1])

	if paddingLen == 0 || paddingLen > 16 {
		return nil, fmt.Errorf("invalid padding length: %d", paddingLen)
	}

	if paddingLen > len(data) {
		return nil, errors.New("padding length exceeds data length")
	}

	// Verify all padding bytes have the same value
	for i := len(data) - paddingLen; i < len(data); i++ {
		if data[i] != byte(paddingLen) {
			return nil, errors.New("invalid PKCS7 padding")
		}
	}

	return data[:len(data)-paddingLen], nil
}
