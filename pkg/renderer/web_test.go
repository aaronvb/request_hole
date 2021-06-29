package renderer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/aaronvb/logrequest"
	"github.com/aaronvb/request_hole/pkg/protocol"
)

func TestIncomingRequest(t *testing.T) {
	rp := protocol.RequestPayload{Fields: logrequest.RequestFields{Url: "/foo"}}
	webServer := Web{}

	if len(webServer.requests) != 0 {
		t.Errorf("Expected %d, got %d", 0, len(webServer.requests))
	}

	webServer.incomingRequest(rp)

	if len(webServer.requests) != 1 {
		t.Errorf("Expected %d, got %d", 0, len(webServer.requests))
	}
}

// Handlers

// GET /requests
func TestRequestHandler(t *testing.T) {
	testTable := []struct {
		req protocol.RequestPayload
	}{
		{protocol.RequestPayload{Fields: logrequest.RequestFields{Url: "/foo"}}},
		{protocol.RequestPayload{Fields: logrequest.RequestFields{Url: "/bar"}}},
		{protocol.RequestPayload{Fields: logrequest.RequestFields{Method: http.MethodGet}}},
		{protocol.RequestPayload{Params: "{\"foo\" => \"bar\"}"}},
	}

	webServer := Web{requests: make([]*protocol.RequestPayload, 0)}
	srv := httptest.NewServer(webServer.routes())

	defer srv.Close()

	for i, test := range testTable {
		webServer.incomingRequest(test.req)

		req, err := http.NewRequest(http.MethodGet, srv.URL+"/requests", nil)
		if err != nil {
			t.Error(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		rps := []protocol.RequestPayload{}
		err = json.Unmarshal(body, &rps)
		if err != nil {
			t.Error(err)
		}

		if len(rps) != i+1 {
			t.Errorf("Expected %d, got %d", i+1, len(rps))
		}

		if reflect.DeepEqual(rps[i], test.req) != true {
			t.Errorf("Expected %v, got %v", test.req, rps[i])
		}

		resp.Body.Close()
	}
}
