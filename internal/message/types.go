package message

import "github.com/pion/webrtc/v3"

type Response struct {
	Params *webrtc.SessionDescription `json:"params"`
	Result *webrtc.SessionDescription `json:"result"`
	Method string                     `json:"method"`
	Id     uint64                     `json:"id"`
}

type SendAnswer struct {
	SID    string                     `json:"sid"`
	Answer *webrtc.SessionDescription `json:"answer"`
}

type ResponseCandidate struct {
	Target    int                      `json:"target"`
	Candidate *webrtc.ICECandidateInit `json:"candidate"`
}

type TrickleResponse struct {
	Params ResponseCandidate `json:"params"`
	Method string            `json:"method"`
}
