package renderer

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/aaronvb/request_hole/pkg/protocol"
	"github.com/pterm/pterm"
)

// Printer is our CLI output that currently uses pterm.
// See https://github.com/pterm/pterm for more info on pterm.
type Printer struct {
	// Spinner is a constant, we set it during the start method and can later stop it when we exit.
	Spinner *pterm.SpinnerPrinter

	// Details will output the headers with the request. Default is false unless the
	// flag is passed.
	Details bool
}

// Start renders the spinner and starts receive incoming requests from the channel.
func (p *Printer) Start(wg *sync.WaitGroup, rp chan protocol.RequestPayload, q chan int, e chan int) {
	defer wg.Done()

	p.startSpinner()

	// Receive incoming requests on RequestPayload channel or
	// exit blocking select if quit is received from protocol
	for {
		select {
		case r := <-rp:
			p.incomingRequest(r)
		case <-q:
			close(rp)
			return
		}
	}
}

// incomingRequest handles the output for incoming requests to the protocol.
func (p *Printer) incomingRequest(r protocol.RequestPayload) {
	p.Spinner.Stop()

	prefix := pterm.Prefix{
		Text:  r.Fields.Method,
		Style: pterm.NewStyle(pterm.BgGray, pterm.FgWhite),
	}

	text := p.incomingRequestText(r)
	pterm.Info.WithPrefix(prefix).Println(text)

	// If the details flag is passed we print headers as well,
	// default is false if no flag is passed.
	if p.Details {
		table := p.incomingRequestHeadersTable(r)
		if table != "" {
			pterm.Printf("%s\n", table)
		}
	}

	p.startSpinner()
}

// incomingRequestText converts the RequestPayload into a printable string.
func (p *Printer) incomingRequestText(r protocol.RequestPayload) string {
	urlWithStyle := ""
	if r.Fields.Url != "" {
		urlWithStyle = pterm.DefaultBasicText.
			WithStyle(pterm.NewStyle(pterm.FgWhite)).Sprintf("%s ", r.Fields.Url)
	}

	paramsWithStyle := pterm.DefaultBasicText.
		WithStyle(pterm.NewStyle(pterm.Fuzzy)).Sprintf(r.Message)

	text := fmt.Sprintf("%s%s", urlWithStyle, paramsWithStyle)
	return text
}

// incomingRequestHeadersTable constructs the headers table string from the RequestPayload.
// This takes the headers map from the request and sorts it alphabetically by key.
func (p *Printer) incomingRequestHeadersTable(r protocol.RequestPayload) string {
	if len(r.Headers) == 0 {
		return ""
	}

	keys := make([]string, 0, len(r.Headers))
	for key := range r.Headers {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var headersFormatted [][]string

	headerRow := []string{"Header", "Value"}
	headersFormatted = append(headersFormatted, headerRow)

	for _, key := range keys {
		value := strings.Join(r.Headers[key], ",")
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
