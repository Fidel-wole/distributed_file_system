package main

import (
	"log"

	"github.com/Fidel-wole/distributed_file_system/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOps{
		ListenAddr: ":4000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("failed to listen and accept: %v", err)
	}
	select {}
}
