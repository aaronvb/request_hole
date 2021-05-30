package renderer

import (
	"fmt"
	"testing"

	"github.com/aaronvb/logrequest"
	"github.com/pterm/pterm"
)

func TestStartText(t *testing.T) {
	pterm.DisableColor()
	printer := Printer{Addr: "localhost", Port: 8080, BuildInfo: map[string]string{"version": "dev"}}
	result := printer.startText()
	expected := fmt.Sprintf("Request Hole %s\nListening on http://%s:%d", "dev", printer.Addr, printer.Port)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestIncomingRequestText(t *testing.T) {
	pterm.DisableColor()
	printer := Printer{}
	fields := logrequest.RequestFields{
		Method: "GET",
		Url:    "/foobar",
	}
	params := "{\"foo\" => \"bar\"}"
	result := printer.incomingRequestText(fields, params)
	expected := fmt.Sprintf("%s %s", fields.Url, params)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
