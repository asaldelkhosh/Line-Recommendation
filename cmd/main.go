package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

var addr string

func main() {
	flag.StringVar(&addr, "a", "localhost:7000", "address to use")
	flag.Parse()

	u := url.URL{
		Scheme: "ws",
		Host:   addr,
		Path:   "/ws",
	}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {
			log.Fatal("close fatal:", err)
		}
	}(c)
}