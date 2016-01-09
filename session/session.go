// Package session provides utilities for dealing with sessions with registrar
// connections.
package session

import (
	"github.com/WooDNSword/registrant/connection"
	"net"
)

func Handler(conn net.Conn) {
	// TODO: Develop message format struct type.
	// TODO: Send IDENT message.
	identMsg := connection.Message{
		Type: "IDENT",
		Content: []string{"registrant"},
	}
	
	identMsg.Send(conn)
}
