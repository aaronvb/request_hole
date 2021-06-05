package renderer

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/aaronvb/logrequest"
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

	// Contains build info
	BuildInfo map[string]string

	// Determines if header details should be shown with the request
	Details bool
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
		Sprintf(p.BuildInfo["version"])

	text := fmt.Sprintf("%s %s\nListening on http://%s:%d", primary, version, p.Addr, p.Port)

	if p.Details {
		text = fmt.Sprintf("%s\nDetails: %t", text, p.Details)
	}

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
func (p *Printer) IncomingRequest(fields logrequest.RequestFields, params string, headers map[string][]string) {
	p.Spinner.Stop()
	prefix := pterm.Prefix{
		Text:  fields.Method,
		Style: pterm.NewStyle(pterm.BgGray, pterm.FgWhite),
	}

	text := p.incomingRequestText(fields, params)
	pterm.Info.WithPrefix(prefix).Println(text)

	p.startSpinner()
}

// IncomingRequestHeader handles the output for incoming requests headers to the server.
func (p *Printer) IncomingRequestHeaders(headers map[string][]string) {
	p.Spinner.Stop()

	table := p.incomingRequestHeadersTable(headers)
	pterm.Printf("%s\n\n", table)

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

// incomingRequestHeadersTable constructs the headers table string.
// This takes the headers map from the request and sorts it alphabetically by key.
func (p *Printer) incomingRequestHeadersTable(headers map[string][]string) string {
	keys := make([]string, 0, len(headers))
	for key := range headers {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var headersFormatted [][]string

	headerRow := []string{"Header", "Value"}
	headersFormatted = append(headersFormatted, headerRow)

	for _, key := range keys {
		value := strings.Join(headers[key], ",")
		headersRow := []string{key, value}
		headersFormatted = append(headersFormatted, headersRow)
	}

	headersTable, err := pterm.DefaultTable.WithHasHeader().WithData(headersFormatted).Srender()
	if err != nil {
		pterm.Error.WithShowLineNumber(false).Println(err)
	}

	return headersTable
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
