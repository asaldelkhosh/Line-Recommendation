package model

type Search struct {
	Base
	StartID uint `json:"start"`
	StopID  uint `json:"stop"`
}
