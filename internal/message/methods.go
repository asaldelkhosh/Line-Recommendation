package message

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
)

func (m *Message) onAccept(response Response) {
	result := *response.Result
	_ = response.Result
	if err := m.PeerConnection.SetRemoteDescription(result); err != nil {
		log.Fatal(err)
	}
}

func (m *Message) onOffer(response Response) {
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
}

func (m *Message) onTrickle(message []byte) {
	var trickleResponse TrickleResponse

	if err := json.Unmarshal(message, &trickleResponse); err != nil {
		log.Fatal(err)
	}

	err := m.PeerConnection.AddICECandidate(*trickleResponse.Params.Candidate)

	if err != nil {
		log.Fatal(err)
	}
}
