package protocol

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
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
		params         map[string]string
		expectedParams string
		headerKey      string
		headerValue    string
	}{
		{
			http.MethodGet,
			"/foo",
			"",
			nil,
			"",
			"Foo",
			"Bar",
		},
		{
			http.MethodPost,
			"/foo/bar",
			"{\"foo\": \"bar\"}",
			nil,
			"{\"foo\" => \"bar\"}",
			"Content-Type",
			"application/json",
		},
		{
			http.MethodPost,
			"/form/data",
			"",
			map[string]string{"aloha": "friday"},
			"{\"aloha\" => \"friday\"}",
			"Content-Type",
			"multipart/form-data",
		},
		{
			http.MethodDelete,
			"/foo/1",
			"",
			nil,
			"",
			"Bearer",
			"hello!",
		},
		{
			http.MethodGet,
			"/foo/bar?hello=world",
			"",
			nil,
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
		r, _ := regexp.Compile(`multipart\/form-data`)
		matched := r.MatchString(test.headerValue)
		if test.headerKey == "Content-Type" && matched {
			b := &bytes.Buffer{}
			writer := multipart.NewWriter(b)
			for k := range test.params {
				fw, err := writer.CreateFormField(k)
				if err != nil {
					t.Errorf("Error POST to httptest server")
				}

				_, err = io.Copy(fw, strings.NewReader(test.params[k]))
				if err != nil {
					t.Errorf("Error POST to httptest server")
				}
			}
			writer.Close()
			req, err := http.NewRequest(test.method, srv.URL+test.path, bytes.NewReader(b.Bytes()))

			if err != nil {
				t.Error(err)
			}

			req.Header.Set("Content-Type", writer.FormDataContentType())
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Request failed with response code: %d", resp.StatusCode)
			}
			resp.Body.Close()
		} else {
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
		}

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

		headerValue := rp.Headers[test.headerKey][0]
		r, _ = regexp.Compile(test.headerValue)
		matched = r.MatchString(headerValue)
		if !matched {
			t.Errorf("Expected %s, got %s", test.headerValue, headerValue)
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
	errorChans := []chan int{e1, e2}

	q1 := make(chan int, 1)
	q2 := make(chan int, 1)
	quitChans := []chan int{q1, q2}

	e1 <- 1
	httpServer := Http{}
	httpServer.Start(make([]chan RequestPayload, 0), quitChans, errorChans)

	expectedQ1 := <-q1
	expectedQ2 := <-q2

	if expectedQ1 != 1 || expectedQ2 != 1 {
		t.Error("Expected channel to receive quit signal")
	}
}
