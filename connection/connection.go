// Package connection provides utilities for establishing and handling
// connections and other networking artifacts.
// TODO: Write tests for the package.
package connection

import (
	"github.com/WooDNSword/registrant/config"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

// TODO: Move Message struct type to shared repository (exists in registrar
// codebase as well).
// TODO: Document Message struct type.
type Message struct {
	Type    string
	Content []string
}

// TODO: Move Message.ByteString() to shared repository (exists in registrar
// codebase as well).
// TODO: Document Message.ByteString().
func (msg Message) ByteString() []byte {
	msgJson, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	
	return msgJson
}

// TODO: Move Message.Send() to shared repository (exists in registrar codebase
// as well).
// TODO: Document Message.Send().
func (msg Message) Send(conn net.Conn) (int, error) {
	return conn.Write(msg.ByteString())
}

// TODO: Move Recv() to shared repository (exists in registrar codebase as
// well).
// TODO: Document Recv().
func Recv(reader io.Reader) Message {
	var msg Message
	
	dec := json.NewDecoder(reader)
	err := dec.Decode(&msg)
	if err != nil {
		panic(err)
	}
	
	return msg
}

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

// TODO: Document ToMessage().
// This function should probably not be used very often. Currently deprecated.
// Read() is better suited to the context this function would probably be used
// in.
func ToMessage(byteString []byte) Message {
	var msg Message
	
	err := json.Unmarshal(byteString, &msg)
	if err != nil {
		panic(err)
	}
	
	return msg
}
