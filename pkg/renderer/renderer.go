package renderer

import (
	"log"

	"github.com/aaronvb/logrequest"
)

// Renderer contains the interface which our servers use to render the output.
type Renderer interface {
	// Start is called when we start our server.
	Start()

	// ErrorLogger can be used if we need to access the logger interface.
	// This is useful in certain cases such as the ErrorLog when using the http.Server
	ErrorLogger() *log.Logger

	// Fatal is used when we need to display a message and should always exit the CLI.
	Fatal(error)

	IncomingRequest(logrequest.RequestFields, string)
}
