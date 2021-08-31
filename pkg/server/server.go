package server

import (
	"fmt"
	"sync"

	"github.com/aaronvb/request_hole/pkg/protocol"
	"github.com/aaronvb/request_hole/pkg/renderer"
	"github.com/pterm/pterm"
)

// Server orchestrates the start and channels for the protocol and renderers.
type Server struct {
	// FlagData contains the data from the CLI flags.
	FlagData

	// Protocol which we receive incoming requests to. The protocol uses a channel
	// to send incoming requests to the renderers.
	Protocol protocol.Protocol

	// Renderers contains a slice of renderer's. Each will run within it's own
	// goroutine.
	Renderers []renderer.Renderer
}

// FlagData contains the data from the CLI flags.
type FlagData struct {
	// Addr is the address the HTTP server will bind to.
	Addr string

	// BuildInfo contains the build information for rh. Set by goreleaser.
	BuildInfo map[string]string

	// Details determines if header details should be shown with the request,
	Details bool

	// LogFile contains the path and filename to the log file which the server
	// will write to if log flag is passed.
	LogFile string

	// Port is the port the HTTP server will run on.
	Port int

	// ResponseCode is the response which our endpoint will return.
	// Default is 200 if no response code is passed.
	ResponseCode int

	// Web determines if we use the web renderer, otherwise defaults to the printer renderer.
	Web bool

	// WebAddress is the address the web UI will bind to.
	WebAddress string

	// WebPort defines which port we host the web renderer at, defaults to 8081.
	WebPort int
}

// Start handles all of the orchestration.
//
// Prints the CLI header which we use to show flags passed to the CLI(ie: port).
//
// Creates a waitgroup for each goroutine.
//
// Creates a channel between the protocol and renderers to handle incoming request
// payloads and exiting due to errors.
//
// Blocks main program until all goroutines are returned. In most cases the user will
// force exit the CLI from the terminal.
func (s *Server) Start() {
	s.printServerInfo()

	if len(s.Renderers) == 0 {
		pterm.Error.WithShowLineNumber(true).Println("No render provided")
		return
	}

	var wg sync.WaitGroup
	var rpChans []chan protocol.RequestPayload
	var rendererQuitChans []chan int
	var rendererErrorChans []chan int

	// Start the renderers
	for _, renderer := range s.Renderers {
		rp := make(chan protocol.RequestPayload)
		q := make(chan int)
		e := make(chan int)

		wg.Add(1)
		go renderer.Start(&wg, rp, q, e)

		rpChans = append(rpChans, rp)
		rendererQuitChans = append(rendererQuitChans, q)
		rendererErrorChans = append(rendererErrorChans, e)
	}

	// Start the server that accepts incoming requests
	go s.Protocol.Start(rpChans, rendererQuitChans, rendererErrorChans)

	wg.Wait()
}

// printServerInfo prints the top header section of the CLI when we start.
// This contains info such as flag options passed and build info.
func (s *Server) printServerInfo() {
	clear()

	text := s.startText()

	pterm.DefaultBox.
		WithBoxStyle(pterm.NewStyle(pterm.FgGray)).
		Printfln(text)
}

func (s *Server) startText() string {
	primary := pterm.DefaultBasicText.
		WithStyle(pterm.NewStyle(pterm.Bold)).
		Sprintf("Request Hole")
	version := pterm.DefaultBasicText.
		WithStyle(pterm.NewStyle(pterm.Fuzzy)).
		Sprintf(s.FlagData.BuildInfo["version"])

	text := fmt.Sprintf("%s %s\nListening on http://%s:%d", primary, version, s.FlagData.Addr, s.FlagData.Port)

	if s.FlagData.Web {
		text = fmt.Sprintf("%s\nWeb running on: http://%s:%d", text, s.FlagData.WebAddress, s.FlagData.WebPort)
	}

	if s.FlagData.Details {
		text = fmt.Sprintf("%s\nDetails: %t", text, s.FlagData.Details)
	}

	if s.FlagData.LogFile != "" {
		text = fmt.Sprintf("%s\nLog: %s", text, s.FlagData.LogFile)
	}

	return text
}

// clear will clear the terminal, called at the start.
func clear() {
	print("\033[H\033[2J")
}
