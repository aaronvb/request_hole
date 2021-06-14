package renderer

import (
	"sync"

	"github.com/aaronvb/request_hole/pkg/protocol"
)

// Renderer contains the interface which our servers use to render the output.
type Renderer interface {
	// Start is called when we start our server.
	Start(*sync.WaitGroup, chan protocol.RequestPayload, chan int)
}
