package engine

import (
	"github.com/pion/mediadevices"
	"github.com/pion/mediadevices/pkg/codec/vpx"
	"github.com/pion/webrtc/v3"
)

// GetMediaEngine will create and return the media engine and code selector
func GetMediaEngine() (*webrtc.MediaEngine, *mediadevices.CodecSelector) {
	mediaEngine := webrtc.MediaEngine{}

	// vpx parameters for media engine
	vpxP, _ := vpx.NewVP8Params()
	vpxP.BitRate = 500_000
	// code selector for media engine
	codecSelector := mediadevices.NewCodecSelector(
		mediadevices.WithVideoEncoders(&vpxP),
	)
	codecSelector.Populate(&mediaEngine)

	return &mediaEngine, codecSelector
}
