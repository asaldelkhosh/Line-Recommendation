package engine

import (
	"github.com/pion/mediadevices"
	"github.com/pion/mediadevices/pkg/codec/vpx"
	"github.com/pion/webrtc/v3"
)

func GetMediaEngine() (*webrtc.MediaEngine, *mediadevices.CodecSelector) {
	mediaEngine := webrtc.MediaEngine{}

	vpxP, _ := vpx.NewVP8Params()
	vpxP.BitRate = 500_000
	codecSelector := mediadevices.NewCodecSelector(
		mediadevices.WithVideoEncoders(&vpxP),
	)
	codecSelector.Populate(&mediaEngine)

	return &mediaEngine, codecSelector
}
