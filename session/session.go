// Package session provides utilities for dealing with sessions with registrar
// connections.
package session

import (
	"net"
)

func Handler(conn net.Conn) {
	// TODO: Develop message format struct type.
	// TODO: Send IDENT message.
	// TODO: Delete the placeholder "Greetings, ..." message.
	conn.Write([]byte("Greetings, registrar!"))
}