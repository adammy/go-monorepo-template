package httpapi

// Server defines the interface for a server.
type Server interface {
	// Start the Server.
	Start() error
}
