package engine

import (
	"github.com/pion/mediadevices"
	"github.com/pion/webrtc/v3"
)

// GetMediaEngine will create and return the media-engine and code-selector.
func GetMediaEngine() (*webrtc.MediaEngine, *mediadevices.CodecSelector) {
	// webrtc media-engine
	mediaEngine := webrtc.MediaEngine{}

	// code selector for media engine
	codecSelector := mediadevices.NewCodecSelector(
		mediadevices.WithVideoEncoders(getVPXPParams()),
	)
	codecSelector.Populate(&mediaEngine)

	return &mediaEngine, codecSelector
}
