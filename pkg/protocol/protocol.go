package protocol

import (
	"sync"

	"github.com/aaronvb/logrequest"
)

// Protocol is the interface for the servers that accept incoming requests.
// Incoming requests are then sent to the renderers through the RequestPayload channel.
// If a protocol closes(ie: from and error), we use the second channel which is used to
// send an int(1 signals quit).
type Protocol interface {
	Start(*sync.WaitGroup, []chan RequestPayload, []chan int)
}

// RequestPayload is the request payload we receive from an incoming request that we use with
// the renderers.
type RequestPayload struct {
	Fields  logrequest.RequestFields `json:"fields"`
	Headers map[string][]string      `json:"headers"`
	Params  string                   `json:"params"`
}