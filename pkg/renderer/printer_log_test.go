package renderer

import (
	"testing"

	"github.com/pterm/pterm"
)

func TestWriteDefaultPrefix(t *testing.T) {
	pterm.DisableOutput()

	pl := PrinterLog{}
	b := []byte("foobar")
	i, _ := pl.Write(b)

	if i != 34 {
		t.Errorf("Expected str len %d, got %d", 34, i)
	}
}

func TestWriteErrorPrefix(t *testing.T) {
	pterm.DisableOutput()

	pl := PrinterLog{Prefix: pterm.Error}
	b := []byte("foobar")
	i, _ := pl.Write(b)

	if i != 38 {
		t.Errorf("Expected str len %d, got %d", 38, i)
	}
}
