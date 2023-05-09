package model

type Search struct {
	Base
	Start Node `json:"start"`
	Stop  Node `json:"stop"`
}
