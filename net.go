package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

// PORT is listenning port
const PORT = ":9669"

var (
	listConnections map[string]net.Conn = make(map[string]net.Conn)
	mutex                               = new(sync.Mutex)
)

func main() {
	for _, addr := range os.Args[1:] {
		if err := connectNode(addr); err != nil {
			log.Println(err)
		}
	}
	serve()
}

func serve() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", PORT)
	if err != nil {
		log.Println(err)
		os.Exit(500) // TODO errors handling
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println(err)
		os.Exit(500) // TODO errors handling
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue // TODO errors handling
		}

		go receiveMessage(conn)
	}
}

func receiveMessage(conn net.Conn) {
	defer conn.Close()
	for {
		mess, _ := bufio.NewReader(conn).ReadString('\n')

		log.Println(string(mess))
		time.Sleep(3 * time.Second)
	}
}

func connectNode(IP string) error {
	log.Printf("Connect to %s\n", IP+PORT)

	tcpAddr, err := net.ResolveTCPAddr("tcp", IP+PORT)
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}

	fmt.Fprintf(conn, "Hello boi! \n")
	return nil
}
