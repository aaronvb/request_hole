package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aaronvb/logrequest"
)

type MockPrinter struct {
	fields  logrequest.RequestFields
	params  string
	headers map[string][]string
}

func (mp *MockPrinter) IncomingRequest(fields logrequest.RequestFields, params string, headers map[string][]string) {
func (mp *MockPrinter) Fatal(error) {}
func (mp *MockPrinter) Start()      {}

func (mp *MockPrinter) IncomingRequest(fields logrequest.RequestFields, params string) {
	mp.fields = fields
	mp.params = params
}

func (mp *MockPrinter) IncomingRequestHeaders(headers map[string][]string) {
	mp.headers = headers
}

func TestResponseCodeFlag(t *testing.T) {
	tests := []int{200, 404, 201, 500}

	renderer := &MockPrinter{}
	for _, respCode := range tests {
		httpServer := Http{ResponseCode: respCode, Output: renderer}
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

func TestLogRequest(t *testing.T) {
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
	renderer := &MockPrinter{}
	httpServer := Http{ResponseCode: 200, Output: renderer}
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

		if renderer.fields.Method != test.method {
			t.Errorf("Expected %s, got %s", test.method, renderer.fields.Method)
		}

		if renderer.fields.Url != test.path {
			t.Errorf("Expected %s, got %s", test.path, renderer.fields.Url)
		}

		if renderer.params != test.expectedParams {
			t.Errorf("Expected %s, got %s", test.expectedParams, renderer.params)
		}
	}
}

func TestLogRequestHeaders(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Foo":          "bar",
		"Bearer":       "hello!",
	}

	renderer := &MockPrinter{}
	httpServer := Http{ResponseCode: 200, Output: renderer}
	srv := httptest.NewServer(httpServer.routes())
	defer srv.Close()

	req, err := http.NewRequest(http.MethodGet, srv.URL, nil)
	if err != nil {
		t.Error(err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

	for key, value := range headers {
		result := renderer.headers[key][0]
		if value != result {
			t.Errorf("Expected %s, got %s", value, result)
		}
	}
}
