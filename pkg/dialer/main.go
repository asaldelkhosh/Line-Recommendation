package dialer

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

var (
	addr string
)

func MakeConnection() (*websocket.Conn, error) {
	flag.StringVar(&addr, "a", "localhost:7000", "address to use")
	flag.Parse()

	u := url.URL{
		Scheme: "ws",
		Host:   addr,
		Path:   "/ws",
	}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	return c, err
}
