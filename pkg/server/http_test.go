package server

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aaronvb/logrequest"
)

type MockPrinter struct {
	fields logrequest.RequestFields
	params string
}

func (mp *MockPrinter) Fatal(error)              {}
func (mp *MockPrinter) Start()                   {}
func (mp *MockPrinter) ErrorLogger() *log.Logger { return log.New(os.Stderr, "", 0) }
func (mp *MockPrinter) IncomingRequest(fields logrequest.RequestFields, params string) {
	mp.fields = fields
	mp.params = params
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