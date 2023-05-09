package http

import (
	"github.com/sweetie-pie/line-recommendation/internal/model"
)

type Response struct {
	Routes   []RouteResponse  `json:"routes"`
	Searches []SearchResponse `json:"searches"`
}

type RouteResponse struct {
	ID    uint       `json:"id"`
	Start model.Node `json:"start"`
	Stop  model.Node `json:"stop"`
}

type SearchResponse struct {
	X1 int64 `json:"x1"`
	Y1 int64 `json:"y1"`
	X2 int64 `json:"x2"`
	Y2 int64 `json:"y2"`
}
