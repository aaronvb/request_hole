package graph

import (
	"sync"

	"github.com/aaronvb/request_hole/graph/model"
	"github.com/aaronvb/request_hole/pkg/protocol"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RequestPayloads        *[]*protocol.RequestPayload
	RequestPayloadObserver *map[string]chan *protocol.RequestPayload
	Info                   *model.ServerInfo
	mu                     sync.Mutex
}
