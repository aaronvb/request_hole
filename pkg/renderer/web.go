package renderer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/aaronvb/request_hole/pkg/protocol"
	"github.com/gorilla/mux"
	"github.com/pterm/pterm"
)

type Web struct {
	Port     int
	requests []protocol.RequestPayload
}

func (web *Web) Start(wg *sync.WaitGroup, rp chan protocol.RequestPayload, q chan int, e chan int) {
	// Initialize requests as an empty slice so that we can return a proper json empty array
	// if no values have been appended.
	web.requests = make([]protocol.RequestPayload, 0)

	addr := fmt.Sprintf("localhost:%d", web.Port)
	errorLog := log.New(&httpErrorLog{}, "", 0)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  web.routes(),
	}

	defer wg.Done()

	go func() {
		err := srv.ListenAndServe()
		str := pterm.Error.WithShowLineNumber(false).Sprintf("Web: %s\n", err)
		pterm.Printo(str) // Overwrite last line
		e <- 1
	}()

	for {
		select {
		case r := <-rp:
			web.incomingRequest(r)
		case <-q:
			close(rp)
			return
		}
	}
}

// routes handles routes for our web UI.
func (web *Web) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/requests", web.requestsHandler).Methods("GET")

	return r
}

// requestsHandler returns an array of incoming requests to our protocol server.
func (web *Web) requestsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(web.requests)
}

// incomingRequest is called when we receive a RequestPayload over the channel
// from the protocol server. This will add it to the request slice which our web ui
// will serve as JSON and be consumed on the front end.
func (web *Web) incomingRequest(req protocol.RequestPayload) {
	web.requests = append(web.requests, req)
}

// httpErrorLog implements the logger interface.
type httpErrorLog struct{}

// Write let's us override the logger required for http errors and
// prints to the terminal using pterm.
func (e *httpErrorLog) Write(b []byte) (n int, err error) {
	pterm.Error.WithShowLineNumber(false).Println(string(b))
	return len(b), nil
}