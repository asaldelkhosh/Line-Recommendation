package engine

import "github.com/pion/mediadevices/pkg/codec/vpx"

func GetParams() vpx.VP8Params {
	vpxP, _ := vpx.NewVP8Params()
	vpxP.BitRate = 500_000

	return vpxP
}
