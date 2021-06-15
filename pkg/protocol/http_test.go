package protocol

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseCodeFlag(t *testing.T) {
	tests := []int{200, 404, 201, 500}

	for _, respCode := range tests {
		httpServer := Http{ResponseCode: respCode}
		srv := httptest.NewServer(httpServer.routes())
		req, err := http.NewRequest(http.MethodGet, srv.URL+"/", nil)
		if err != nil {
			t.Error(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != respCode {
			t.Errorf("Expected %d, got %d", respCode, resp.StatusCode)
		}
		resp.Body.Close()
		srv.Close()
	}
}

func TestLogRequestOneRenderer(t *testing.T) {
	testTable := []struct {
		method         string
		path           string
		body           string
		expectedParams string
		headerKey      string
		headerValue    string
	}{
		{
			http.MethodGet,
			"/foo",
			"",
			"",
			"Foo",
			"Bar",
		},
		{
			http.MethodPost,
			"/foo/bar",
			"{\"foo\": \"bar\"}",
			"{\"foo\" => \"bar\"}",
			"Content-Type",
			"application/json",
		},
		{
			http.MethodDelete,
			"/foo/1",
			"",
			"",
			"Bearer",
			"hello!",
		},
		{
			http.MethodGet,
			"/foo/bar?hello=world",
			"",
			"{\"hello\" => \"world\"}",
			"Content-Type",
			"application/json!",
		},
	}

	rpChannel := make(chan RequestPayload, len(testTable))
	httpServer := Http{ResponseCode: 200, rendererChannels: []chan RequestPayload{rpChannel}}
	srv := httptest.NewServer(httpServer.routes())
	defer srv.Close()

	for _, test := range testTable {
		b := bytes.NewBuffer([]byte(test.body))
		req, err := http.NewRequest(test.method, srv.URL+test.path, b)

		if test.headerKey != "" {
			req.Header.Set(test.headerKey, test.headerValue)
		}

		if err != nil {
			t.Error(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}
		resp.Body.Close()

		rp := <-rpChannel

		if rp.Fields.Method != test.method {
			t.Errorf("Expected %s, got %s", test.method, rp.Fields.Method)
		}

		if rp.Fields.Url != test.path {
			t.Errorf("Expected %s, got %s", test.path, rp.Fields.Url)
		}

		if rp.Params != test.expectedParams {
			t.Errorf("Expected %s, got %s", test.expectedParams, rp.Params)
		}

		expectedHeaderValue := rp.Headers[test.headerKey][0]
		if expectedHeaderValue != test.headerValue {
			t.Errorf("Expected %s, got %s", expectedHeaderValue, test.headerValue)
		}
	}
}

func TestLogRequestManyRenderers(t *testing.T) {
	testTable := []struct {
		method         string
		path           string
		Body           string
		expectedParams string
	}{
		{http.MethodGet, "/foo", "", ""},
		{http.MethodPost, "/foo/bar", "{\"foo\": \"bar\"}", "{\"foo\" => \"bar\"}"},
		{http.MethodDelete, "/foo/1", "", ""},
		{http.MethodGet, "/foo/bar?hello=world", "", "{\"hello\" => \"world\"}"},
	}

	rpChannelA := make(chan RequestPayload, len(testTable))
	rpChannelB := make(chan RequestPayload, len(testTable))
	httpServer := Http{
		ResponseCode:     200,
		rendererChannels: []chan RequestPayload{rpChannelA, rpChannelB}}
	srv := httptest.NewServer(httpServer.routes())
	defer srv.Close()

	for _, test := range testTable {
		b := bytes.NewBuffer([]byte(test.Body))
		req, err := http.NewRequest(test.method, srv.URL+test.path, b)

		if test.Body != "" {
			req.Header.Set("Content-Type", "application/json")
		}

		if err != nil {
			t.Error(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}
		resp.Body.Close()

		rpA := <-rpChannelA
		rpB := <-rpChannelB

		if rpA.Fields.Method != test.method {
			t.Errorf("Expected %s, got %s", test.method, rpA.Fields.Method)
		}

		if rpA.Fields.Url != test.path {
			t.Errorf("Expected %s, got %s", test.path, rpA.Fields.Url)
		}

		if rpA.Params != test.expectedParams {
			t.Errorf("Expected %s, got %s", test.expectedParams, rpA.Params)
		}

		if rpB.Fields.Method != test.method {
			t.Errorf("Expected %s, got %s", test.method, rpB.Fields.Method)
		}

		if rpB.Fields.Url != test.path {
			t.Errorf("Expected %s, got %s", test.path, rpB.Fields.Url)
		}

		if rpB.Params != test.expectedParams {
			t.Errorf("Expected %s, got %s", test.expectedParams, rpB.Params)
		}
	}
}

func TestQuitRenderers(t *testing.T) {
	q1 := make(chan int, 1)
	q2 := make(chan int, 1)
	chans := []chan int{q1, q2}

	httpServer := Http{rendererQuitChannels: chans}
	httpServer.quitRenderers()
	expectedQ1 := <-q1
	expectedQ2 := <-q2

	if expectedQ1 != 1 || expectedQ2 != 1 {
		t.Error("Expected channel to receive quit signal")
	}
}

func TestErrorFromRenderer(t *testing.T) {
	e1 := make(chan int, 1)
	e2 := make(chan int, 1)
	chans := []chan int{e1, e2}

	e1 <- 1
	err := <-merge(chans)

	if err != 1 {
		t.Error("Expect	channel to receive error signal")
	}
}
