package renderer

import (
	"fmt"
	"log"
	"os"

	"github.com/aaronvb/logrequest"
	"github.com/aaronvb/request_hole/pkg/version"
	"github.com/pterm/pterm"
)

// Printer is our CLI output that currently uses pterm.
// See https://github.com/pterm/pterm for more info on pterm.
type Printer struct {
	// Spinner is a constant, we set it during the start method and can later stop it when we exit.
	Spinner *pterm.SpinnerPrinter

	// Fields for startText
	Port int
	Addr string
}

// Start renders the initial header and the spinner. The Spinner should be consistent during
// all requests, unless we explicitly tell it to stop.
func (p *Printer) Start() {
	// Clear the terminal
	clear()

	text := p.startText()
	pterm.DefaultBox.
		WithBoxStyle(pterm.NewStyle(pterm.FgGray)).
		Printfln(text)

	p.startSpinner()
}

func (p *Printer) startText() string {
	primary := pterm.DefaultBasicText.
		WithStyle(pterm.NewStyle(pterm.Bold)).
		Sprintf("Request Hole")
	version := pterm.DefaultBasicText.
		WithStyle(pterm.NewStyle(pterm.Fuzzy)).
		Sprintf(version.BuildVersion)

	text := fmt.Sprintf("%s %s\nListening on http://%s:%d", primary, version, p.Addr, p.Port)
	return text
}

// ErrorLogger will create a printerLog which interfaces with Logger.
func (p *Printer) ErrorLogger() *log.Logger {
	errorLog := log.New(&printerLog{prefix: pterm.Error}, "", p.Port)
	return errorLog
}

// Fatal will use the Error prefix to render the error and then exit the CLI.
func (p *Printer) Fatal(err error) {
	p.Spinner.Stop()
	pterm.Error.WithShowLineNumber(false).Println(err)
	os.Exit(1)
}

// IncomingRequest handles the output for incoming requests to the server.
func (p *Printer) IncomingRequest(fields logrequest.RequestFields, params string) {
	p.Spinner.Stop()
	prefix := pterm.Prefix{
		Text:  fields.Method,
		Style: pterm.NewStyle(pterm.BgGray, pterm.FgWhite),
	}

	text := p.incomingRequestText(fields, params)
	pterm.Info.WithPrefix(prefix).Println(text)
	p.startSpinner()
}

func (p *Printer) incomingRequestText(fields logrequest.RequestFields, params string) string {
	urlWithStyle := pterm.DefaultBasicText.
		WithStyle(pterm.NewStyle(pterm.FgWhite)).Sprintf(fields.Url)
	paramsWithStyle := pterm.DefaultBasicText.
		WithStyle(pterm.NewStyle(pterm.Fuzzy)).Sprintf(params)

	text := fmt.Sprintf("%s %s", urlWithStyle, paramsWithStyle)
	return text
}

// Create the spinner which will be displayed at the bottom.
func (p *Printer) startSpinner() {
	listeningText := pterm.DefaultBasicText.
		WithStyle(pterm.NewStyle(pterm.Fuzzy)).
		Sprint("waiting for incoming requests")

	spinner, _ := pterm.DefaultSpinner.
		WithRemoveWhenDone(true).
		WithSequence("⠋",
			"⠙",
			"⠚",
			"⠞",
			"⠖",
			"⠦",
			"⠴",
			"⠲",
			"⠳",
			"⠓").Start(listeningText)

	p.Spinner = spinner
}

// clear will clear the terminal, called at the start.
func clear() {
	print("\033[H\033[2J")
}
