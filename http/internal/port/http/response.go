package http

import (
	"github.com/sweetie-pie/line-recommendation/internal/model"
)

type Response struct {
	Routes   []RouteResponse `json:"routes"`
	Searches []RouteResponse `json:"searches"`
}

type RouteResponse struct {
	ID    uint       `json:"id"`
	Start model.Node `json:"start"`
	Stop  model.Node `json:"stop"`
}
