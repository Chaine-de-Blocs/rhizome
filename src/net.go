package src

import (
	"log"
	"net"
	"os"
	"rhizome/src/protocol"
	"sync"
	"time"
)

// DefaultPort is the default port when not set during runtime
const DefaultPort = ":9669"

// TODO lock when working with nodeList
var (
	nodeList = make([]*Node, 0)
	mutex    = new(sync.Mutex)
)

// Listen for incoming messages
func Listen() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", DefaultPort)
	if err != nil {
		log.Println(err)
		os.Exit(500) // TODO errors handling
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println(err)
		os.Exit(500) // TODO errors handling
	}

	log.Printf("Listenning on port %s\n", DefaultPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue // TODO errors handling
		}

		// Time out after 500ms (TODO handles timeout behavior)
		conn.SetDeadline(time.Now().Add(time.Millisecond * 500))

		n := Node{
			Socket: conn,
		}

		go n.ProcessMessages()
	}
}

// OpenConnection joins a bootstrap peer by IPAddress (includes PORT)
func OpenConnection(IPAddress string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", IPAddress)
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}

	node := Node{
		Socket: conn,
	}

	msg := protocol.Message{
		Command: protocol.CmdJoin,
		Payload: nil,
	}

	if err := node.SendMessage(msg); err != nil {
		return err
	}

	go node.ProcessMessages()

	return nil
}

// ConnectToNode initiates a connection with a remote node
func ConnectToNode(node Node) (*Node, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", node.GetIPAddress())
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	n := &Node{
		Socket: conn,
	}

	return n, nil
}

// IsListedNode checks if the node is in the list
func IsListedNode(node *Node) bool {
	for _, n := range nodeList {
		if n.GetIPAddress() == node.GetIPAddress() {
			return true
		}
	}
	return false
}
