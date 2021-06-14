package protocol

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/aaronvb/logparams"
	"github.com/aaronvb/logrequest"
	"github.com/gorilla/mux"
	"github.com/pterm/pterm"
)

// Http is the protocol for accepting http requests.
type Http struct {
	// Addr is the address the HTTP server will bind to.
	Addr string

	// Port is the port the HTTP server will run on.
	Port int

	// ResponseCode is the response which out endpoint will return.
	// Default is 200 if no response code is passed.
	ResponseCode int

	// rendererChannel is the channel which we send a RequestPayload to when
	// receiving an incoming request to the Http protocol.
	rendererChannels []chan RequestPayload
}

// Start will start the HTTP server.
//
// Sets the channel on our struct so that incoming requests can be sent over it.
//
// In the case that we cannot start this server, we send a signal to our quit channel
// to close renderers.
func (s *Http) Start(wg *sync.WaitGroup, c []chan RequestPayload, quits []chan int) {
	addr := fmt.Sprintf("%s:%d", s.Addr, s.Port)
	errorLog := log.New(&httpErrorLog{}, "", 0)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  s.routes(),
	}

	s.rendererChannels = c

	defer wg.Done()

	err := srv.ListenAndServe()
	str := pterm.Error.WithShowLineNumber(false).Sprintf("%s\n", err)
	pterm.Printo(str) // Overwrite last line

	for _, quit := range quits {
		quit <- 1
	}
}

// routes handles the routes for our HTTP server and currently accepts any path.
func (s *Http) routes() http.Handler {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(s.defaultHandler)
	r.Use(s.logRequest)

	return r
}

// defaultHandler returns the response code which is provided as a flag.
// Defaults to 200.
func (s *Http) defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(s.ResponseCode)
}

// logRequest is the middleware that passes the request data and parameters to
// the Renderer IncomingRequest interface method.
func (s *Http) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lr := logrequest.LogRequest{Request: r, Writer: w, Handler: next}
		fields := lr.ToFields()
		params := logparams.LogParams{Request: r, HidePrefix: true}

		req := RequestPayload{
			Fields:  fields,
			Params:  params.ToString(),
			Headers: r.Header,
		}

		for _, rendererChannel := range s.rendererChannels {
			rendererChannel <- req
		}
	})
}

// httpErrorLog implements the logger interface.
type httpErrorLog struct{}

// Write let's us override the logger required for http errors and
// prints to the terminal using pterm.
func (e *httpErrorLog) Write(b []byte) (n int, err error) {
	pterm.Error.WithShowLineNumber(false).Println(string(b))
	return len(b), nil
}
