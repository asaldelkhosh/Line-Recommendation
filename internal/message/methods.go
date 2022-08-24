package message

import (
	"bytes"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
)

func (m *Message) onAccept(response Response) error {
	if err := m.PeerConnection.SetRemoteDescription(*response.Result); err != nil {
		return err
	}

	return nil
}

func (m *Message) onOffer(response Response) error {
	_ = m.PeerConnection.SetRemoteDescription(*response.Params)

	answer, err := m.PeerConnection.CreateAnswer(nil)
	if err != nil {
		return err
	}

	_ = m.PeerConnection.SetLocalDescription(answer)

	*m.ConnectionID = uint64(uuid.New().ID())

	offerJSON, _ := json.Marshal(&SendAnswer{
		Answer: m.PeerConnection.LocalDescription(),
		SID:    "test room",
	})

	answerMessage := jsonrpc2.Request{
		Method: "answer",
		Params: (*json.RawMessage)(&offerJSON),
		ID: jsonrpc2.ID{
			IsString: false,
			Str:      "",
			Num:      *m.ConnectionID,
		},
	}

	reqBodyBytes := new(bytes.Buffer)
	if er := json.NewEncoder(reqBodyBytes).Encode(answerMessage); er != nil {
		return er
	}

	return m.Connection.WriteMessage(websocket.TextMessage, reqBodyBytes.Bytes())
}

func (m *Message) onTrickle(message []byte) error {
	var trickleResponse TrickleResponse

	if err := json.Unmarshal(message, &trickleResponse); err != nil {
		return err
	}

	if err := m.PeerConnection.AddICECandidate(*trickleResponse.Params.Candidate); err != nil {
		return err
	}

	return nil
}
