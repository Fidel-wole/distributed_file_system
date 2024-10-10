package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTcpTransport(t *testing.T) {
	
	tcpOpts := TCPTransportOps{
		ListenAddr: "3000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder: DefaultDecoder{},
	}
	tr := NewTCPTransport(tcpOpts)
	assert.Equal(t, tr.ListenAddr, tcpOpts.ListenAddr)
	
	//Server
	assert.Nil(t, tr.ListenAndAccept()) // Assert that there is no error
	
	// Consider replacing the blocking select with a timeout or signal
	//select{} // This will block indefinitely
	// You might want to use a different approach here
}
