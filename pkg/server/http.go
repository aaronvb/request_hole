package server

import (
	"fmt"
	"net/http"

	"github.com/aaronvb/logparams"
	"github.com/aaronvb/logrequest"
	"github.com/aaronvb/request_hole/pkg/renderer"
	"github.com/gorilla/mux"
)

type Http struct {
	// Port is the port the HTTP server will run on.
	Port int

	// Addr is the address the HTTP server will bind to.
	Addr string

	// ResponseCode is the response which out endpoint will return.
	// Default is 200 if no response code is passed.
	ResponseCode int

	// Output is the Renderer interface.
	Output renderer.Renderer
}

// Start will start the HTTP server.
func (s *Http) Start() {
	s.Output.Start()
	addr := fmt.Sprintf("%s:%d", s.Addr, s.Port)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: s.Output.ErrorLogger(),
		Handler:  s.routes(),
	}

	err := srv.ListenAndServe()
	s.Output.Fatal(err)
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
		s.Output.IncomingRequest(fields, params.ToString())
	})
}
