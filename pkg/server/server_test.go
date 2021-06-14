package server

import (
	"fmt"
	"testing"

	"github.com/pterm/pterm"
)

func TestStartText(t *testing.T) {
	pterm.DisableColor()
	flags := FlagData{Addr: "localhost", Port: 8080, BuildInfo: map[string]string{"version": "dev"}}
	server := Server{FlagData: flags}
	result := server.startText()
	expected := fmt.Sprintf("Request Hole %s\nListening on http://%s:%d", "dev", server.FlagData.Addr, server.FlagData.Port)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestStartTextWithDetails(t *testing.T) {
	pterm.DisableColor()
	flags := FlagData{
		Addr:      "localhost",
		Port:      8080,
		BuildInfo: map[string]string{"version": "dev"},
		Details:   true}
	server := Server{FlagData: flags}
	result := server.startText()
	expected := fmt.Sprintf(
		"Request Hole %s\nListening on http://%s:%d\nDetails: %t", "dev",
		server.FlagData.Addr, server.FlagData.Port, server.FlagData.Details)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestStartTextWithLogFile(t *testing.T) {
	pterm.DisableColor()
	flags := FlagData{
		Addr:      "localhost",
		Port:      8080,
		BuildInfo: map[string]string{"version": "dev"},
		LogFile:   "rh.log",
	}
	server := Server{FlagData: flags}
	result := server.startText()
	expected := fmt.Sprintf(
		"Request Hole %s\nListening on http://%s:%d\nLog: %s", "dev",
		server.FlagData.Addr, server.FlagData.Port, server.FlagData.LogFile)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
