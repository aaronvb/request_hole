package protocol

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWsLogRequestOneRenderer(t *testing.T) {
	testTable := []struct {
		method         string
		path           string
		expectedParams string
		headerKey      string
		headerValue    string
	}{
		{
			http.MethodGet,
			"/foo",
			"",
			"Foo",
			"Bar",
		},
		{
			http.MethodGet,
			"/foo?hello=world",
			"{\"hello\" => \"world\"}",
			"",
			"",
		},
	}

	rpChannel := make(chan RequestPayload, len(testTable))
	wsServer := Ws{rendererChannels: []chan RequestPayload{rpChannel}}
	srv := httptest.NewServer(wsServer.routes())
	defer srv.Close()

	for _, test := range testTable {
		wsUrl := strings.Replace(srv.URL, "http", "ws", 1)

		header := http.Header{test.headerKey: {test.headerValue}}
		wsReq, _, err := websocket.DefaultDialer.Dial(wsUrl+test.path, header)
		if err != nil {
			t.Fatalf("%v", err)
		}

		defer wsReq.Close()

		rp := <-rpChannel

		if rp.Fields.Method != test.method {
			t.Errorf("Expected %s, got %s", test.method, rp.Fields.Method)
		}

		if rp.Fields.Url != test.path {
			t.Errorf("Expected %s, got %s", test.path, rp.Fields.Url)
		}

		if rp.Message != test.expectedParams {
			t.Errorf("Expected %s, got %s", test.expectedParams, rp.Message)
		}

		if rp.Headers[test.headerKey] != nil {
			if rp.Headers[test.headerKey][0] != test.headerValue {
				t.Errorf("Expected %s, got %s", rp.Headers[test.headerKey][0], test.headerValue)
			}
		}
	}
}

func TestWsLogRequestManyRenderers(t *testing.T) {
	testTable := []struct {
		method         string
		path           string
		expectedParams string
	}{
		{http.MethodGet, "/foo", ""},
		{http.MethodGet, "/foo/bar?hello=world", "{\"hello\" => \"world\"}"},
	}

	rpChannelA := make(chan RequestPayload, len(testTable))
	rpChannelB := make(chan RequestPayload, len(testTable))
	wsServer := Ws{rendererChannels: []chan RequestPayload{rpChannelA, rpChannelB}}
	srv := httptest.NewServer(wsServer.routes())
	defer srv.Close()

	for _, test := range testTable {
		wsUrl := strings.Replace(srv.URL, "http", "ws", 1)

		wsReq, _, err := websocket.DefaultDialer.Dial(wsUrl+test.path, nil)
		if err != nil {
			t.Fatalf("%v", err)
		}

		defer wsReq.Close()

		rpA := <-rpChannelA
		rpB := <-rpChannelB

		if rpA.Fields.Method != test.method {
			t.Errorf("Expected %s, got %s", test.method, rpA.Fields.Method)
		}

		if rpA.Fields.Url != test.path {
			t.Errorf("Expected %s, got %s", test.path, rpA.Fields.Url)
		}

		if rpA.Message != test.expectedParams {
			t.Errorf("Expected %s, got %s", test.expectedParams, rpA.Message)
		}

		if rpB.Fields.Method != test.method {
			t.Errorf("Expected %s, got %s", test.method, rpB.Fields.Method)
		}

		if rpB.Fields.Url != test.path {
			t.Errorf("Expected %s, got %s", test.path, rpB.Fields.Url)
		}

		if rpB.Message != test.expectedParams {
			t.Errorf("Expected %s, got %s", test.expectedParams, rpB.Message)
		}
	}
}

func TestWsQuitRenderers(t *testing.T) {
	q1 := make(chan int, 1)
	q2 := make(chan int, 1)
	chans := []chan int{q1, q2}

	wsServer := Ws{rendererQuitChannels: chans}
	wsServer.quitRenderers()
	expectedQ1 := <-q1
	expectedQ2 := <-q2

	if expectedQ1 != 1 || expectedQ2 != 1 {
		t.Error("Expected channel to receive quit signal")
	}
}

func TestWsErrorFromRenderer(t *testing.T) {
	e1 := make(chan int, 1)
	e2 := make(chan int, 1)
	errorChans := []chan int{e1, e2}

	q1 := make(chan int, 1)
	q2 := make(chan int, 1)
	quitChans := []chan int{q1, q2}

	e1 <- 1
	wsServer := Ws{}
	wsServer.Start(make([]chan RequestPayload, 0), quitChans, errorChans)

	expectedQ1 := <-q1
	expectedQ2 := <-q2

	if expectedQ1 != 1 || expectedQ2 != 1 {
		t.Error("Expected channel to receive quit signal")
	}
}

func TestWsLogMessage(t *testing.T) {
	testTable := []struct {
		method  string
		message string
	}{
		{"RECEIVE", "Hello"},
		{"RECEIVE", "World"},
		{"RECEIVE", "ping"},
		{"RECEIVE", "fizz"},
		{"RECEIVE", "buzz"},
	}

	rpChannel := make(chan RequestPayload, len(testTable)+1)
	wsServer := Ws{rendererChannels: []chan RequestPayload{rpChannel}}
	srv := httptest.NewServer(wsServer.routes())
	defer srv.Close()

	wsUrl := strings.Replace(srv.URL, "http", "ws", 1)
	wsReq, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	defer wsReq.Close()

	// Initial handshake request
	wsRequest := <-rpChannel
	if wsRequest.Fields.Method != "GET" {
		t.Errorf("Expected %s, got %s", "GET", wsRequest.Fields.Method)
	}

	// Test each message
	for _, test := range testTable {
		if err := wsReq.WriteMessage(websocket.TextMessage, []byte(test.message)); err != nil {
			t.Fatalf("%v", err)
		}

		wsMessage := <-rpChannel

		if wsMessage.Fields.Method != test.method {
			t.Errorf("Expected %s, got %s", test.method, wsMessage.Fields.Method)
		}

		if wsMessage.Message != test.message {
			t.Errorf("Expected %s, got %s", test.message, wsMessage.Message)
		}
	}
}
