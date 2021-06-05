package renderer

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aaronvb/logrequest"
)

func TestLoggerStartText(t *testing.T) {
	logger := Logger{Port: 123, Addr: "foo.bar"}
	text := logger.startText()
	expected := fmt.Sprintf("Listening on http://%s:%d", logger.Addr, logger.Port)

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
	text := logger.incomingRequestText(fields, params)
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
