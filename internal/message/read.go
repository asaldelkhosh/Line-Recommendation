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

var (
	peerConnection *webrtc.PeerConnection
	connectionID   uint64
)

func ReadMessage(connection *websocket.Conn, done chan struct{}) {
	defer close(done)

	for {
		_, message, err := connection.ReadMessage()
		if err != nil || err == io.EOF {
			log.Fatal("Error reading: ", err)
			return
		}

		fmt.Printf("recv: %s", message)

		var response Response
		_ = json.Unmarshal(message, &response)

		if response.Id == connectionID {
			result := *response.Result
			_ = response.Result
			if err := peerConnection.SetRemoteDescription(result); err != nil {
				log.Fatal(err)
			}
		} else if response.Id != 0 && response.Method == "offer" {
			_ = peerConnection.SetRemoteDescription(*response.Params)
			answer, err := peerConnection.CreateAnswer(nil)

			if err != nil {
				log.Fatal(err)
			}

			_ = peerConnection.SetLocalDescription(answer)

			connectionUUID := uuid.New()
			connectionID = uint64(connectionUUID.ID())

			offerJSON, err := json.Marshal(&SendAnswer{
				Answer: peerConnection.LocalDescription(),
				SID:    "test room",
			})

			params := (*json.RawMessage)(&offerJSON)

			answerMessage := jsonrpc2.Request{
				Method: "answer",
				Params: params,
				ID: jsonrpc2.ID{
					IsString: false,
					Str:      "",
					Num:      connectionID,
				},
			}

			reqBodyBytes := new(bytes.Buffer)
			_ = json.NewEncoder(reqBodyBytes).Encode(answerMessage)

			messageBytes := reqBodyBytes.Bytes()
			_ = connection.WriteMessage(websocket.TextMessage, messageBytes)
		} else if response.Method == "trickle" {
			var trickleResponse TrickleResponse

			if err := json.Unmarshal(message, &trickleResponse); err != nil {
				log.Fatal(err)
			}

			err := peerConnection.AddICECandidate(*trickleResponse.Params.Candidate)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
