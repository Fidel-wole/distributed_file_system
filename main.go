package main

import (
	"log"

	"github.com/Fidel-wole/distributed_file_system/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":4000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("failed to listen and accept: %v", err)
	}
	select {}
}
