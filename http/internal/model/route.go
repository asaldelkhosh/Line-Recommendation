package model

type Route struct {
	Base
	Start Node `json:"start"`
	Stop  Node `json:"stop"`
}
