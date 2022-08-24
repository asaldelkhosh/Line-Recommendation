package main

import (
	"log"

	"github.com/amirhnajafiz/broadcaster/internal/pion/sfu"
)

// Server
// manages the peer connections.
type Server interface {
	Accept() error
}

func main() {
	// creating connection to ion-sfu server
	conn, err := sfu.Connect("localhost:7000")
	if err != nil {
		log.Fatalf("failed to dial ion-sfu server: %v\n", err)

		return
	}

	// closing connection to our ion-sfu server
	defer conn.Close()

	// starting broadcast server
	s := New(conn)

	// accept peer on server
	if er := s.Accept(); er != nil {
		log.Fatalf("failed to accept on broadcast server: %v\n", er)
	}
}
