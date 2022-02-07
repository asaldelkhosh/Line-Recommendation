package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
	"github.com/sourcegraph/jsonrpc2"
	"io"
	"log"
)

type Message struct {
	Connection     *websocket.Conn
	PeerConnection *webrtc.PeerConnection
	ConnectionID   *uint64
}

func (m Message) ReadMessage(done chan struct{}) {
	defer close(done)

	for {
		_, message, err := m.Connection.ReadMessage()
		if err != nil || err == io.EOF {
			log.Fatal("Error reading: ", err)
			return
		}

		fmt.Printf("recv: %s", message)

		var response Response
		_ = json.Unmarshal(message, &response)

		if response.Id == *m.ConnectionID {
			result := *response.Result
			_ = response.Result
			if err := m.PeerConnection.SetRemoteDescription(result); err != nil {
				log.Fatal(err)
			}
		} else if response.Id != 0 && response.Method == "offer" {
			_ = m.PeerConnection.SetRemoteDescription(*response.Params)
			answer, err := m.PeerConnection.CreateAnswer(nil)

			if err != nil {
				log.Fatal(err)
			}

			_ = m.PeerConnection.SetLocalDescription(answer)

			connectionUUID := uuid.New()
			*m.ConnectionID = uint64(connectionUUID.ID())

			offerJSON, err := json.Marshal(&SendAnswer{
				Answer: m.PeerConnection.LocalDescription(),
				SID:    "test room",
			})

			params := (*json.RawMessage)(&offerJSON)

			answerMessage := jsonrpc2.Request{
				Method: "answer",
				Params: params,
				ID: jsonrpc2.ID{
					IsString: false,
					Str:      "",
					Num:      *m.ConnectionID,
				},
			}

			reqBodyBytes := new(bytes.Buffer)
			_ = json.NewEncoder(reqBodyBytes).Encode(answerMessage)

			messageBytes := reqBodyBytes.Bytes()
			_ = m.Connection.WriteMessage(websocket.TextMessage, messageBytes)
		} else if response.Method == "trickle" {
			var trickleResponse TrickleResponse

			if err := json.Unmarshal(message, &trickleResponse); err != nil {
				log.Fatal(err)
			}

			err := m.PeerConnection.AddICECandidate(*trickleResponse.Params.Candidate)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
