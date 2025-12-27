package api

import "net/http"

// handleListDevices returns all registered devices
func (s *Server) handleListDevices(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Get devices from registry, return as JSON
}

// handleGetDevice returns info for a single device
func (s *Server) handleGetDevice(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Get device by ID from path, return as JSON
}

// handleGetStatus returns cached state for a device
func (s *Server) handleGetStatus(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Get device state, return as JSON
}

// handleQuery forces an active query to a device
func (s *Server) handleQuery(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Call device.Query(), return result
}

// handleHealth returns server health status
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Return health info
}
