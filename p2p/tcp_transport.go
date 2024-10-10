package p2p

import (
	"fmt"
	"net"
	"sync"
	"io"
)

// TCPPeer represents the remote node over a TCP established connection
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

type TCPTransportOps struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}
type TCPTransport struct {
	TCPTransportOps
	Listener net.Listener // corrected spelling
	Rpcch chan RPC
	
	//mutex
	mu    sync.RWMutex
	Peers map[net.Addr]Peer
}

// NewTCPTransport creates a new TCPTransport object.
func NewTCPTransport(opts TCPTransportOps) *TCPTransport {
	return &TCPTransport{
		TCPTransportOps: opts,
		Rpcch: make(chan RPC),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.Listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		go t.handleConn(conn)
	}
}

type Temp struct {
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP Handshake error: %s\n", err)
		return
	}

	// Read Loop
	rpc := &RPC{}
	for {
		if err := t.Decoder.Decode(conn, rpc); err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by peer.")
				break // Exit the loop if EOF is encountered
			}
			fmt.Printf("Error decoding message: %s\n", err)
			continue
		}
		rpc.From = conn.RemoteAddr()
		fmt.Printf("message: %+v\n", rpc)
	}
	conn.Close() // Ensure the connection is closed after the loop
}
