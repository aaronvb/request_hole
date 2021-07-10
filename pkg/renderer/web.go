package renderer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aaronvb/request_hole/graph"
	"github.com/aaronvb/request_hole/graph/generated"
	"github.com/aaronvb/request_hole/graph/model"
	"github.com/aaronvb/request_hole/pkg/protocol"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/pterm/pterm"
	"github.com/rs/cors"
)

type Web struct {
	BuildInfo     map[string]string
	Port          int
	ResponseCode  int
	RequestAddr   string
	RequestPort   int
	StaticFiles   http.FileSystem
	mu            sync.Mutex
	requests      []*protocol.RequestPayload
	subscriptions map[string]chan *protocol.RequestPayload
}

func (web *Web) Start(wg *sync.WaitGroup, rp chan protocol.RequestPayload, q chan int, e chan int) {
	// Initialize requests as an empty slice so that we can return a proper json empty array
	// if no values have been appended.
	web.requests = make([]*protocol.RequestPayload, 0)
	web.subscriptions = make(map[string]chan *protocol.RequestPayload)

	addr := fmt.Sprintf("localhost:%d", web.Port)
	errorLog := log.New(&httpErrorLog{}, "", 0)

	srv := &http.Server{
		Addr:        addr,
		ErrorLog:    errorLog,
		Handler:     web.routes(),
		IdleTimeout: 30 * time.Second,
	}

	defer wg.Done()

	go func() {
		open(fmt.Sprintf("http://%s/", addr))
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
	r.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	r.HandleFunc("/query", web.gqlHandler)
	handler := cors.Default().Handler(r)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer((web.StaticFiles))))
	return handler
}

// requestsHandler returns an array of incoming requests to our protocol server.
func (web *Web) requestsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(web.requests)
}

func (web *Web) gqlHandler(w http.ResponseWriter, r *http.Request) {
	serverInfo := model.ServerInfo{
		RequestAddress: web.RequestAddr,
		RequestPort:    web.RequestPort,
		WebPort:        web.Port,
		ResponseCode:   web.ResponseCode,
		BuildInfo:      web.BuildInfo,
	}
	// Pass pointer to requests and subscriptions
	gqlSrv := handler.New(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
			RequestPayloads:        &web.requests,
			RequestPayloadObserver: &web.subscriptions,
			Info:                   &serverInfo,
		}}))
	gqlSrv.AddTransport(transport.POST{})
	gqlSrv.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 5 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	gqlSrv.Use(extension.Introspection{})

	gqlSrv.ServeHTTP(w, r)
}

// incomingRequest is called when we receive a RequestPayload over the channel
// from the protocol server. This will add it to the request slice which our web ui
// will serve as JSON and be consumed on the front end.
func (web *Web) incomingRequest(req protocol.RequestPayload) {
	web.mu.Lock()
	web.requests = append(web.requests, &req)
	web.mu.Unlock()

	for _, subscriptions := range web.subscriptions {
		subscriptions <- &req
	}
}

// httpErrorLog implements the logger interface.
type httpErrorLog struct{}

// Write let's us override the logger required for http errors and
// prints to the terminal using pterm.
func (e *httpErrorLog) Write(b []byte) (n int, err error) {
	pterm.Error.WithShowLineNumber(false).Println(string(b))
	return len(b), nil
}

// open opens the terminal to the web UI.
// https://stackoverflow.com/a/39324149
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
