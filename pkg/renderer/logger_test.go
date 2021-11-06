package renderer

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aaronvb/logrequest"
	"github.com/aaronvb/request_hole/pkg/protocol"
)

func TestLoggerStartText(t *testing.T) {
	logger := Logger{Addr: "localhost", Port: 1234, Protocol: "ws"}
	text := logger.startText()
	expected := fmt.Sprintf("Listening on %s://%s:%d", logger.Protocol, logger.Addr, logger.Port)

	if text != expected {
		t.Errorf("Expected %s, got %s", expected, text)
	}
}

func TestLoggerIncomingRequest(t *testing.T) {
	logger := Logger{}
	fields := logrequest.RequestFields{
		Method: "GET",
		Url:    "/foobar",
	}
	params := "{\"foo\" => \"bar\"}"
	rp := protocol.RequestPayload{Fields: fields, Message: params}
	text := logger.incomingRequestText(rp)
	expected := fmt.Sprintf("%s %s %s", fields.Method, fields.Url, params)

	if text != expected {
		t.Errorf("Expected %s, got %s", expected, text)
	}
}

func TestIncomingRequestHeadersText(t *testing.T) {
	logger := Logger{}
	headers := map[string][]string{
		"hello": {"world", "foobar"},
		"foo":   {"bar"},
	}
	exepectedHeaders := map[string]string{
		"hello": "world,foobar",
		"foo":   "bar",
	}
	expectedSortedKeys := []string{
		"foo", "hello",
	}

	headersWithJoinedValues, keys := logger.incomingRequestHeaders(headers)

	if !reflect.DeepEqual(headersWithJoinedValues, exepectedHeaders) {
		t.Errorf("Expected %s, got %s", exepectedHeaders, headersWithJoinedValues)
	}

	if !reflect.DeepEqual(keys, expectedSortedKeys) {
		t.Errorf("Expected %s, got %s", expectedSortedKeys, keys)
	}
}
