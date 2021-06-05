package renderer

import (
	"github.com/aaronvb/logrequest"
)

// Renderer contains the interface which our servers use to render the output.
type Renderer interface {
	// Start is called when we start our server.
	Start()

	// Fatal is used when we need to display a message and should always exit the CLI.
	Fatal(error)

	// IncomingRequest is called when we receive an incoming request to the server.
	IncomingRequest(logrequest.RequestFields, string)

	// IncomingRequestHeaders is called when the details flag is passed and we want to
	// render the headers.
	IncomingRequestHeaders(map[string][]string)
}
