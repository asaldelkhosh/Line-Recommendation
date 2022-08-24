package engine

import (
	"github.com/pion/mediadevices"
	"github.com/pion/webrtc/v3"
)

// GetMediaEngine will create and return the media engine and code selector
func GetMediaEngine() (*webrtc.MediaEngine, *mediadevices.CodecSelector) {
	mediaEngine := webrtc.MediaEngine{}

	// vpx parameters for media engine
	vpxP := GetParams()

	// code selector for media engine
	codecSelector := mediadevices.NewCodecSelector(
		mediadevices.WithVideoEncoders(&vpxP),
	)
	codecSelector.Populate(&mediaEngine)

	return &mediaEngine, codecSelector
}
