// Package connection provides utilities for establishing and handling
// connections and other networking artifacts.
package connection

import (
	"net"
)

// Connect establishes a connection to the supplied endpoint over the specified
// protocol and returns a net.Conn object as well as an error object.
func Connect(protocol string, endpoint string) (net.Conn, error) {
	conn, err := net.Dial(protocol, endpoint)
	return conn, err
}
