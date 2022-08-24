package main

import (
	"github.com/amirhnajafiz/broadcaster/internal/handler"
	"github.com/amirhnajafiz/broadcaster/internal/pion/engine"
	"github.com/gorilla/websocket"
	_ "github.com/pion/mediadevices/pkg/driver/camera"
	_ "github.com/pion/mediadevices/pkg/driver/microphone"
	"github.com/pion/webrtc/v3"
)

const (
	stunServer = "stun:stun.l.google.com:19302"
)

// getPeerDefaultConfigs returns a peer configurations.
func getPeerDefaultConfigs() webrtc.Configuration {
	return webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{stunServer},
			},
		},
		SDPSemantics: webrtc.SDPSemanticsUnifiedPlanWithFallback,
	}
}

// server manages to make a webrtc server.
type server struct {
	API     *webrtc.API
	Handler handler.Handler
}

// New server.
func New(conn *websocket.Conn) Server {
	// creating a server
	s := &server{}

	// creating a media-engine and a code-selector
	mediaEngine, codecSelector := engine.GetMediaEngine()

	// creating a handler, to handle peer tracks and requests
	s.Handler = handler.Handler{
		Conn:         conn,
		CodeSelector: codecSelector,
	}

	// webrtc api
	s.API = webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine))

	return s
}

// Accept new peers.
func (s *server) Accept() error {
	peerConnection, err := s.API.NewPeerConnection(getPeerDefaultConfigs())
	if err != nil {
		return err
	}

	// handle peer
	s.Handler.Handle(peerConnection)

	return nil
}
