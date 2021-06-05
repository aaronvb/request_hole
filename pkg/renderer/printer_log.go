package renderer

import "github.com/pterm/pterm"

// printerLog is our interface to Logger and accepts a pterm prefix
type PrinterLog struct {
	Prefix pterm.PrefixPrinter
}

// Write will be used by the function calling our printerLog.
// If no prefix is passed, we default to Info
func (pl *PrinterLog) Write(b []byte) (n int, err error) {
	if pl.Prefix.Prefix.Text == "" {
		pl.Prefix = pterm.Info
	}

	str := pl.Prefix.WithShowLineNumber(false).Sprint(string(b))
	pterm.Println(str)

	return len(str), nil
}
