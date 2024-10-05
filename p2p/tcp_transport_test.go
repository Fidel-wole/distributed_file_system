package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTcpTransport(t *testing.T) {
	
	tcpOpts := TCPTransportOps{
		ListenAddr: ":4000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder: GOBDecoder{},
	}
	tr := NewTCPTransport(tcpOpts)
	assert.Equal(t, tr.ListenAddr, tcpOpts.ListenAddr)
	
	//Server
	assert.Nil(t, tr.ListenAndAccept())
	select{}
}
