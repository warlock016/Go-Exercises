package api

import (
	"encoding/json"
	"net/http"
	"time"
)

// DeviceInfo is the API representation of a device
type DeviceInfo struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	IP       string     `json:"ip"`
	MAC      string     `json:"mac,omitempty"`
	Online   bool       `json:"online"`
	LastSeen *time.Time `json:"last_seen,omitempty"`
}

// StateResponse is the API representation of device state
type StateResponse struct {
	DeviceID    string                 `json:"device_id"`
	DPs         map[string]interface{} `json:"dps"`
	LastUpdated time.Time              `json:"last_updated"`
	Source      string                 `json:"source"`
}

// ErrorResponse represents an error
type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code,omitempty"`
}

// HealthResponse represents server health
type HealthResponse struct {
	Status      string `json:"status"`
	DeviceCount int    `json:"device_count"`
}

// writeJSON writes a JSON response
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	// TODO(human): Set content-type, write status, encode JSON
}

// writeError writes an error response
func writeError(w http.ResponseWriter, status int, message string) {
	// TODO(human): Use writeJSON with ErrorResponse
}

// Helper for common error responses
func notFound(w http.ResponseWriter, message string) {
	writeError(w, http.StatusNotFound, message)
}

func badRequest(w http.ResponseWriter, message string) {
	writeError(w, http.StatusBadRequest, message)
}

func internalError(w http.ResponseWriter, message string) {
	writeError(w, http.StatusInternalServerError, message)
}

// ToDeviceInfo converts a device to API representation
func ToDeviceInfo(d interface{}) DeviceInfo {
	// TODO(human): Convert device.Device to DeviceInfo
	return DeviceInfo{}
}
