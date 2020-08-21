package main

import (
	"log"

	"github.com/eigenfire/ngenii/network"
)

func main() {
	listener, err := network.CreateTLSListener(":8443", "certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Fatal error: %s", err)
		}

		go network.HandleConnection(conn)
	}
}
