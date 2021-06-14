package renderer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pterm/pterm"
)

type Web struct {
	RendererFlagData
	requests []RequestPayload
}

func (web *Web) Start(req chan RequestPayload) {
	printFlagData(web.RendererFlagData)

	// TODO: need to spin up a WS server as well
	addr := fmt.Sprintf("localhost:%d", web.WebPort)
	errorLog := log.New(&PrinterLog{Prefix: pterm.Error}, "", 0)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  web.routes(),
	}
	r := <-req
	web.IncomingRequest(r)
	err := srv.ListenAndServe()
	web.Fatal(err)
}

// routes handles the routes for our HTTP server and currently accepts any path.
func (web *Web) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/requests", web.requestsHandler).Methods("GET")

	return r
}

func (web *Web) requestsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(web.requests)
}

// Fatal will use the Error prefix to render the error and then exit the CLI.
func (web *Web) Fatal(err error) {
	pterm.Error.WithShowLineNumber(false).Println(err)
	os.Exit(1)
}

func (web *Web) incomingRequest(req RequestPayload) {
	// TODO: this is not thread safe, use mutex
	web.requests = append(web.requests, req)
}

func (web *Web) IncomingRequestHeaders(headers map[string][]string) {

}
