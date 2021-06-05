package renderer

import (
	"testing"

	"github.com/pterm/pterm"
)

func TestWrite(t *testing.T) {
	pterm.DisableOutput()

	pl := PrinterLog{}
	b := []byte("foobar")
	i, _ := pl.Write(b)

	if i != len(b) {
		t.Errorf("Expected str len %d, got %d", len(b), i)
	}
}
