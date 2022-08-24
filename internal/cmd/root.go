package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/amirhnajafiz/broadcaster/internal/message"
	"github.com/amirhnajafiz/broadcaster/internal/pion/dialer"
	"github.com/amirhnajafiz/broadcaster/internal/pion/engine"
	"github.com/amirhnajafiz/broadcaster/internal/pion/media"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	_ "github.com/pion/mediadevices/pkg/driver/camera"
	_ "github.com/pion/mediadevices/pkg/driver/microphone"
	"github.com/pion/webrtc/v3"
	"github.com/sourcegraph/jsonrpc2"
)

var (
	peerConnection *webrtc.PeerConnection
	connectionID   uint64
)

type SendOffer struct {
	SID   string                     `json:"sid"`
	Offer *webrtc.SessionDescription `json:"offer"`
}

type Candidate struct {
	Target    int                  `json:"target"`
	Candidate *webrtc.ICECandidate `json:"candidate"`
}

// GetConfigs will create the webrtc server configurations
func GetConfigs() webrtc.Configuration {
	return webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
		SDPSemantics: webrtc.SDPSemanticsUnifiedPlanWithFallback,
	}
}

func main() {
	// creating the websocket connection
	c, err := dialer.MakeConnection()
	if err != nil {
		log.Fatal("dial:", err)
	}

	// closing connection when we are done
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {
			log.Fatal("close fatal:", err)
		}
	}(c)

	// webrtc configuration
	config := GetConfigs()

	// media engine and code selector
	mediaEngine, codecSelector := engine.GetMediaEngine()

	api := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine))
	peerConnection, err = api.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// create a message
	msg := message.Message{
		Connection:     c,
		PeerConnection: peerConnection,
		ConnectionID:   &connectionID,
	}
	done := make(chan struct{})

	// use our message package to create a reader
	go msg.ReadMessage(done)

	// user media devices
	med := media.Media{
		CodeSelector:   codecSelector,
		PeerConnection: peerConnection,
	}
	med.InitUserMedia()

	// WebRTC offer
	offer, err := peerConnection.CreateOffer(nil)

	// Remote Session description
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		panic(err)
	}

	// Handling OnICECandidate event
	peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			candidateJSON, err := json.Marshal(&Candidate{
				Candidate: candidate,
				Target:    0,
			})

			params := (*json.RawMessage)(&candidateJSON)

			if err != nil {
				log.Fatal(err)
			}

			m := &jsonrpc2.Request{
				Method: "trickle",
				Params: params,
			}

			reqBodyBytes := new(bytes.Buffer)
			_ = json.NewEncoder(reqBodyBytes).Encode(m)

			messageBytes := reqBodyBytes.Bytes()
			_ = c.WriteMessage(websocket.TextMessage, messageBytes)
		}
	})

	peerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed to %s \n", state.String())
	})

	offerJson, err := json.Marshal(&SendOffer{
		Offer: peerConnection.LocalDescription(),
		SID:   "test room",
	})

	params := (*json.RawMessage)(&offerJson)

	connectionUUID := uuid.New()
	connectionID = uint64(connectionUUID.ID())

	offerMessage := &jsonrpc2.Request{
		Method: "join",
		Params: params,
		ID: jsonrpc2.ID{
			IsString: false,
			Str:      "",
			Num:      connectionID,
		},
	}

	reqBodyBytes := new(bytes.Buffer)
	_ = json.NewEncoder(reqBodyBytes).Encode(offerMessage)

	messageBytes := reqBodyBytes.Bytes()
	_ = c.WriteMessage(websocket.TextMessage, messageBytes)

	<-done
}
