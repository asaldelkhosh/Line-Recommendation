package media

import (
	"fmt"
	"github.com/pion/mediadevices"
	"github.com/pion/mediadevices/pkg/frame"
	"github.com/pion/mediadevices/pkg/prop"
	"github.com/pion/webrtc/v3"
)

type Media struct {
	CodeSelector   *mediadevices.CodecSelector
	PeerConnection *webrtc.PeerConnection
}

func (m Media) InitUserMedia() {
	fmt.Println(mediadevices.EnumerateDevices())

	s, err := mediadevices.GetUserMedia(mediadevices.MediaStreamConstraints{
		Video: func(constraints *mediadevices.MediaTrackConstraints) {
			constraints.FrameFormat = prop.FrameFormat(frame.FormatYUY2)
			constraints.Width = prop.Int(640)
			constraints.Height = prop.Int(480)
		},
		Codec: m.CodeSelector,
	})

	if err != nil {
		panic(err)
	}

	for _, track := range s.GetTracks() {
		track.OnEnded(func(err error) {
			fmt.Printf("Track (ID: %s) ended with error: %v\n", track.ID(), err)
		})
		_, err = m.PeerConnection.AddTransceiverFromTrack(track,
			webrtc.RTPTransceiverInit{
				Direction: webrtc.RTPTransceiverDirectionSendonly,
			},
		)

		if err != nil {
			panic(err)
		}
	}
}
