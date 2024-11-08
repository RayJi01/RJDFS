package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.Mutex
	peers map[net.Listener]Peer
}

// NewTCPTransport: Create a new instances of TCPTransport handler
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

// ListenAndAccept: Listen on the listenAddr defined in TCPTransport, and start the accept loop to keep accepting new connections
func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	//start the accept loop and keep accepting connection
	go t.startAcceptLoop()
	return nil
}

// startAcceptLoop: Keep looping and accept new connection to the system
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP Error: %v\n", err)
		}

		go t.handleConn(conn)
	}
}

// handleConn: do things on the accepted connections
func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Printf("Incoming TCP Connection: %v\n", conn)
}