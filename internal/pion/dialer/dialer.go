package dialer

import (
	"flag"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

var (
	addr string
)

// MakeConnection will create a websocket from the following url
// and will return the connection.
func MakeConnection() (*websocket.Conn, error) {
	// creating the address with a flag
	flag.StringVar(&addr, "a", "localhost:7000", "address to use")
	flag.Parse()

	// url that we use to create our websocket dialer
	u := url.URL{
		Scheme: "ws",
		Host:   addr,
		Path:   "/ws",
	}
	log.Printf("connecting to %s", u.String())

	// creating the websocket dialer with dial method
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	return c, err
}
