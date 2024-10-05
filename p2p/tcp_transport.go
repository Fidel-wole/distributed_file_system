package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP establsished connection
type TCPPeer struct {
	// conn is the underlying connection to the peer
	conn net.Conn

	//if we dial a conn => outbount == true
	//if we accept and receive a conn => outbount == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	ListenAddress string
	Listner       net.Listener
	//mutex
	mu    sync.RWMutex
	Peers map[net.Addr]Peer
}

// NewTCPTransport creates a new TCPTransport object.
func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		ListenAddress: listenAddress,
		Peers:         make(map[net.Addr]Peer),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.Listner, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listner.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("New incoming connection %+v\n", peer)
}
