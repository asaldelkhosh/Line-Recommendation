package engine

import "github.com/pion/mediadevices/pkg/codec/vpx"

const (
	// vpx bitrate
	bitRate = 500_000
)

// getVPXPParams returns a VP8 params with
// default bitrate.
func getVPXPParams() *vpx.VP8Params {
	vpxP, _ := vpx.NewVP8Params()
	vpxP.BitRate = bitRate

	return &vpxP
}
