package renderer

import (
	"fmt"
	"testing"

	"github.com/aaronvb/logrequest"
	"github.com/aaronvb/request_hole/pkg/protocol"
	"github.com/pterm/pterm"
)

func TestIncomingRequestText(t *testing.T) {
	pterm.DisableColor()
	printer := Printer{}
	fields := logrequest.RequestFields{
		Method: "GET",
		Url:    "/foobar",
	}
	params := "{\"foo\" => \"bar\"}"
	rp := protocol.RequestPayload{Fields: fields, Message: params}
	result := printer.incomingRequestText(rp)
	expected := fmt.Sprintf("%s %s", fields.Url, params)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestIncomingRequestHeadersTables(t *testing.T) {
	pterm.DisableColor()
	printer := Printer{}
	headers := map[string][]string{
		"hello": {"world", "foobar"},
		"foo":   {"bar"},
	}
	rp := protocol.RequestPayload{Headers: headers}
	result := printer.incomingRequestHeadersTable(rp)

	headersForTable := [][]string{}
	headersForTable = append(headersForTable, []string{"Header", "Value"})
	headersForTable = append(headersForTable, []string{"foo", "bar"})
	headersForTable = append(headersForTable, []string{"hello", "world,foobar"})

	expected, err := pterm.DefaultTable.WithHasHeader().WithData(headersForTable).Srender()
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
