package serve

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

func Start(conn *websocket.Conn) {
	// media engine and code selector
	mediaEngine, codecSelector := engine.GetMediaEngine()

	h := handler.Handler{
		Conn:         conn,
		CodeSelector: codecSelector,
	}

	api := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine))
	peerConnection, err := api.NewPeerConnection(getPeerDefaultConfigs())
	if err != nil {
		panic(err)
	}

	h.Handle(peerConnection)
}
