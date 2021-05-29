package renderer

import (
	"fmt"
	"testing"

	"github.com/aaronvb/logrequest"
	"github.com/aaronvb/request_hole/pkg/version"
)

func TestStartText(t *testing.T) {
	printer := Printer{Addr: "localhost", Port: 8080}
	result := printer.startText()
	expected := fmt.Sprintf("Request Hole %s\nListening at http://%s:%d", version.BuildVersion, printer.Addr, printer.Port)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestIncomingRequestText(t *testing.T) {
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
