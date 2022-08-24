package main

import (
	"log"

	"github.com/amirhnajafiz/broadcaster/internal/cmd/serve"
	"github.com/amirhnajafiz/broadcaster/internal/pion/sfu"
)

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
	serve.Start(conn)
}
