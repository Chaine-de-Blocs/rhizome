package src

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	"./protocol"
	"strings"
)

// Node is a peer node
type Node struct {
	Socket net.Conn
}

// ProcessMessages processes all incoming messages
func (node *Node) ProcessMessages() {
	defer node.Socket.Close()

	dec := gob.NewDecoder(node.Socket)

	var msg protocol.Message
	if err := dec.Decode(&msg); err != nil {
		log.Println(err)
		return
	}

	log.Printf("Message received : %v\n", msg)

	switch msg.Command {
	case protocol.CmdJoin:
		peerAddr := make([]string, 0)
		for _, n := range nodeList {
			peerAddr = append(peerAddr, n.GetIPAddress())
		}

		sendMsg := protocol.Message{
			Command: protocol.CmdAddresses,
			Payload: peerAddr,
		}

		node.SendMessage(sendMsg)

		if !IsListedNode(node) {
			nodeList = append(nodeList, node)
		}
	case protocol.CmdAddresses:
		// TODO Add addresses to list
		log.Printf("Addresses : %v\n", msg)
	}
}

// SendMessage sends a message to a node
func (node *Node) SendMessage(msg protocol.Message) error {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	if err := enc.Encode(msg); err != nil {
		return err
	}

	_, err := node.Socket.Write(buffer.Bytes())

	return err
}

// GetIPAddress returns the IP address of the node pointing on its default port
func (node *Node) GetIPAddress() string {
	return strings.Split(node.Socket.RemoteAddr().String(), ":")[0] + DefaultPort
}
