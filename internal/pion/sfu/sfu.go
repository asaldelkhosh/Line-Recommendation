package sfu

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

// Connect to ION-SFU server.
func Connect(addr string) (*websocket.Conn, error) {
	// server url
	u := url.URL{
		Scheme: "ws",
		Host:   addr,
		Path:   "/ws",
	}

	log.Printf("connecting to ion-sfu on: %s\n", u.String())

	// dial to our server
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	return c, err
}
