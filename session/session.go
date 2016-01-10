// Package session provides utilities for dealing with sessions with registrar
// connections.
package session

import (
	"github.com/WooDNSword/registrant/connection"
	"net"
)

func Handler(conn net.Conn) {
	identMsg := connection.Message{
		Type: "IDENT",
		Content: []string{"registrant"},
	}
	
	identMsg.Send(conn)
	
	// TODO: Receive and respond to post-IDENT STATUS message.
}
