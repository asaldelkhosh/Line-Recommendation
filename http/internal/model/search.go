package model

type Search struct {
	Base
	StartX int64 `json:"startX"`
	StartY int64 `json:"startY"`
	StopX  int64 `json:"stopX"`
	StopY  int64 `json:"stopY"`
}
