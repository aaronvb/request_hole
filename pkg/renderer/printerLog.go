package renderer

import "github.com/pterm/pterm"

// printerLog is our interface to Logger and accepts a pterm prefix
type printerLog struct {
	prefix pterm.PrefixPrinter
}

// Write will be used by the function calling our printerLog.
// If no prefix is passed, we default to Info
func (pl printerLog) Write(b []byte) (n int, err error) {
	if pl.prefix.Prefix.Text == "" {
		pl.prefix = pterm.Info
	}
	pl.prefix.WithShowLineNumber(false).Println(string(b))
	return len(b), nil
}
