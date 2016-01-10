package connection

import (
	"github.com/WooDNSword/registrant/config"
	"net"
	"testing"
)

// TestConnect determines whether Connect establishes a successful
// connection with a simple ad-hoc listener.
func TestConnect(t *testing.T) {
	port := "7776"

	func() {
		// Create a listener.
		listener, err := net.Listen("tcp", ":"+port)

		if err != nil {
			// This should probably not happen.
			t.Error(err)
		}

		// Listen on another thread until we can connect.
		go listener.Accept()
	}()

	endpoint := config.Endpoint{
		Address: "localhost",
		Port:    port,
	}

	// Establish a new connection.
	_, err := Connect("tcp", endpoint)

	// Fail if the connection was unsuccessful.
	if err != nil {
		t.Error(err)
	}
}
