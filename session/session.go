// Package session provides utilities for dealing with sessions with registrar
// connections.
package session

import (
	"net"
)

func Handler(conn net.Conn) {
	conn.Write([]byte("Greetings, registrar!"))
}