package message

import (
	"encoding/json"
	"io"
	"log"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
)

// Message
// manages to read data from peers.
type Message struct {
	Connection     *websocket.Conn
	PeerConnection *webrtc.PeerConnection
	ConnectionID   *uint64
}

func (m *Message) ReadMessage(done chan struct{}) {
	// closing the channel
	defer close(done)

	for {
		// read messages
		_, message, err := m.Connection.ReadMessage()
		if err != nil || err == io.EOF {
			log.Fatalf("failed in reading message: %v\n", err)

			return
		}

		// creating an empty response
		var response Response

		// parse messages from json object
		_ = json.Unmarshal(message, &response)

		// check for response id
		if response.Id == *m.ConnectionID {
			handleError(m.onAccept(response)) // accept
		} else if response.Id != 0 && response.Method == "offer" {
			handleError(m.onOffer(response)) // offer
		} else if response.Method == "trickle" {
			handleError(m.onTrickle(message)) // trickle
		}
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("failed in message reading: %v\n", err)
	}
}
