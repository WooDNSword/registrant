// Package connection provides utilities for establishing and handling
// connections and other networking artifacts.
// TODO: Write tests for the package.
package connection

import (
	"github.com/WooDNSword/registrant/config"
	"fmt"
	"net"
)

// Connect establishes a connection to the supplied endpoint over the specified
// protocol and returns a net.Conn object as well as an error object.
func Connect(protocol string, endpoint config.Endpoint) (net.Conn, error) {
	conn, err := net.Dial(protocol, EndpointToString(endpoint))
	return conn, err
}

// TODO: Determine if EndpointToString actually belongs in the connection
// package. Would it be a better fit for the config package?
func EndpointToString(endpoint config.Endpoint) string {
	return endpoint.Address + ":" + endpoint.Port
}

func HandleConnection(connectionHandler func(net.Conn), conn net.Conn) error {
	connectionHandler(conn)
	return conn.Close()
}

func Initiate(endpoint config.Endpoint, connectionHandler func(net.Conn)) {
	// TODO: Move protocol parameter from being hardcoded to being represented
	// in a config file.
	conn, err := Connect("tcp", endpoint)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Connection established with", EndpointToString(endpoint))
	
	HandleConnection(connectionHandler, conn)
}