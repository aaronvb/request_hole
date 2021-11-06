package protocol

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronvb/logparams"
	"github.com/aaronvb/logrequest"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/pterm/pterm"
	"github.com/rs/cors"
)

// Ws is the protocol for accepting WS connections and messages.
type Ws struct {
	// Addr is the address the WS server will bind to.
	Addr string

	// Port is the port the WS server will run on.
	Port int

	// rendererChannel is the channel which we send a RequestPayload to when
	// receiving an incoming request to the Http protocol.
	rendererChannels     []chan RequestPayload
	rendererQuitChannels []chan int
}

// Start will start the WebSocket server.
//
// Sets the channel on our struct so that incoming message can be sent over it.
//
// In the case that we cannot start this server, we send a signal to our quit channel
// to close renderers.
func (ws *Ws) Start(c []chan RequestPayload, quits []chan int, errors []chan int) {
	addr := fmt.Sprintf("%s:%d", ws.Addr, ws.Port)
	errorLog := log.New(&wsErrorLog{}, "", 0)

	srv := &http.Server{
		Addr:        addr,
		ErrorLog:    errorLog,
		Handler:     ws.routes(),
		IdleTimeout: 30 * time.Second,
	}

	ws.rendererChannels = c
	ws.rendererQuitChannels = quits

	go func() {
		err := srv.ListenAndServe()
		str := pterm.Error.WithShowLineNumber(false).Sprintf("Websocket Protocol: %s\n", err)
		pterm.Printo(str) // Overwrite last line

		// If the server fails to start, send a quit to all renderers, which will exit
		// the main program.
		ws.quitRenderers()
	}()

	// If any of our renderers send an error signal, send a quit signal to all other
	// renderers, which will exit the main program.
	for range merge(errors) {
		ws.quitRenderers()
		return
	}
}

func (ws *Ws) quitRenderers() {
	for _, quit := range ws.rendererQuitChannels {
		quit <- 1
	}
}

// routes handles the routes for our WS server and currently accepts any path.
func (ws *Ws) routes() http.Handler {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(ws.defaultHandler)
	r.Use(ws.logRequest)

	handler := cors.AllowAll().Handler(r)

	return handler
}

// defaultHandler returns accepts the incoming connection and upgrades it to a WebSocket
// connection. Once the connection is established, it will send all incoming messages to
// the renderer channel.
func (ws *Ws) defaultHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ptermErr := pterm.Error.WithShowLineNumber(false).Sprintln(err)
		pterm.Printo(ptermErr)

	}

	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {
			ptermErr := pterm.Error.WithShowLineNumber(false).Sprintln(err)
			pterm.Printo(ptermErr)
		}
	}(c)

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			closeErrors := []int{websocket.CloseNormalClosure}
			if websocket.IsCloseError(err, closeErrors...) {
				ws.logMessage("DISCONNECTED", err.Error())
			} else {
				ws.logMessage("ERROR", err.Error())
			}
			break
		}

		// Log incoming WS message
		ws.logMessage("RECEIVE", string(message))
	}
}

// logRequest is the middleware that passes the initial WebSocket request data and parameters to
// the Renderer.
func (ws *Ws) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fields := logrequest.RequestFields{
			Method: r.Method,
			Url:    r.URL.RequestURI(),
		}

		params := logparams.LogParams{Request: r, HidePrefix: true}
		req := RequestPayload{
			ID:          uuid.New().String(),
			Fields:      fields,
			Message:     params.ToString(),
			Headers:     r.Header,
			ParamFields: params.ToFields(),
			CreatedAt:   time.Now(),
		}

		for _, rendererChannel := range ws.rendererChannels {
			rendererChannel <- req
		}

		next.ServeHTTP(w, r)
	})
}

// logMessage sends any incoming messages from the WebSocket connection to the render channel.
func (ws *Ws) logMessage(method string, msg string) {
	req := RequestPayload{
		ID:        uuid.New().String(),
		Fields:    logrequest.RequestFields{Method: method},
		CreatedAt: time.Now(),
		Message:   msg,
	}

	for _, rendererChannel := range ws.rendererChannels {
		rendererChannel <- req
	}
}

// httpErrorLog implements the logger interface.
type wsErrorLog struct{}

// Write will override the logger required for http errors and
// prints to the terminal using pterm.
func (e *wsErrorLog) Write(b []byte) (n int, err error) {
	pterm.Error.WithShowLineNumber(false).Println(string(b))
	return len(b), nil
}
