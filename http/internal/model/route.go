package model

type Route struct {
	Base
	StartID uint `json:"start"`
	StopID  uint `json:"stop"`
}
