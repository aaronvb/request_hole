package server

import (
	"fmt"
	"testing"

	"github.com/pterm/pterm"
)

func TestStartText(t *testing.T) {
	pterm.DisableColor()
	flags := FlagData{Addr: "localhost", Port: 8080, BuildInfo: map[string]string{"version": "dev"}, Protocol: "http"}
	server := Server{FlagData: flags}
	result := server.startText()
	expected := fmt.Sprintf("Request Hole %s\nListening on %s://%s:%d", "dev", server.FlagData.Protocol, server.FlagData.Addr, server.FlagData.Port)

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
		Details:   true,
		Protocol:  "http"}
	server := Server{FlagData: flags}
	result := server.startText()
	expected := fmt.Sprintf(
		"Request Hole %s\nListening on %s://%s:%d\nDetails: %t", "dev",
		server.FlagData.Protocol, server.FlagData.Addr, server.FlagData.Port, server.FlagData.Details)

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
		Protocol:  "ws",
	}
	server := Server{FlagData: flags}
	result := server.startText()
	expected := fmt.Sprintf(
		"Request Hole %s\nListening on %s://%s:%d\nLog: %s", "dev",
		server.FlagData.Protocol, server.FlagData.Addr, server.FlagData.Port, server.FlagData.LogFile)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestStartTextWithWebUIDefault(t *testing.T) {
	pterm.DisableColor()
	flags := FlagData{
		Addr:      "localhost",
		Port:      8080,
		BuildInfo: map[string]string{"version": "dev"},
		Web:       true,
		Protocol:  "ws",
	}
	server := Server{FlagData: flags}
	result := server.startText()
	expected := fmt.Sprintf(
		"Request Hole %s\nListening on %s://%s:%d\nWeb running on: http://%s:%d", "dev",
		server.FlagData.Protocol, server.FlagData.Addr, server.FlagData.Port, server.FlagData.WebAddress, server.FlagData.WebPort)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestStartTextWithWebUICustomFlags(t *testing.T) {
	pterm.DisableColor()
	flags := FlagData{
		Addr:       "localhost",
		Protocol:   "ws",
		Port:       8080,
		BuildInfo:  map[string]string{"version": "dev"},
		Web:        true,
		WebAddress: "0.0.0.0",
		WebPort:    8082,
	}
	server := Server{FlagData: flags}
	result := server.startText()
	expected := fmt.Sprintf(
		"Request Hole %s\nListening on %s://%s:%d\nWeb running on: http://%s:%d", "dev",
		server.FlagData.Protocol, server.FlagData.Addr, server.FlagData.Port, server.FlagData.WebAddress, server.FlagData.WebPort)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
