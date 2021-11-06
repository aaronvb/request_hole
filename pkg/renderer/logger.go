package renderer

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/aaronvb/request_hole/pkg/protocol"
	"github.com/pterm/pterm"
)

// Logger outputs to a log file.
type Logger struct {
	// FilePathis the path to the log file which we write to.
	// Default is blank unless passed as a flag to the CLI.
	FilePath string

	// Details will log the headers with the request. Default is false unless the
	// flag is passed.
	Details bool

	// Address and Port are used for the start text
	Addr string
	Port int

	// Protocol is the protocol the web UI server will use.
	Protocol string

	// LogFile is the open log file
	logFile *os.File

	// errorChan is used to let the protocol server know that this renderer
	// has encountered an error.
	errorChan chan int
}

// Start writes the initial server start to the log file.
func (l *Logger) Start(wg *sync.WaitGroup, rp chan protocol.RequestPayload, q chan int, e chan int) {
	// Set error channel to be used in fatal
	l.errorChan = e

	f, err := os.OpenFile(l.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// If we fail to open the file, call fatal, which will print the error
		// and send a signal over the error channel.
		l.fatal(err)
	}

	defer f.Close()
	defer wg.Done()

	str := fmt.Sprintf("%s: %s\n", time.Now().Format("2006/02/01 15:04:05"), l.startText())
	f.WriteString(str)

	// Set logFile to open file
	l.logFile = f

	// Receive incoming requests on RequestPayload channel or
	// exit blocking select if quit is received from protocol
	for {
		select {
		case r := <-rp:
			l.incomingRequest(r)
		case <-q:
			close(rp)
			return
		}
	}
}

// startText returns the starting log string
func (l *Logger) startText() string {
	return fmt.Sprintf("Listening on %s://%s:%d", l.Protocol, l.Addr, l.Port)
}

// fatal will use the Error prefix to render the error and then exit the CLI.
func (l *Logger) fatal(err error) {
	pterm.Error.WithShowLineNumber(false).Println(err)
	l.errorChan <- 1
}

// incomingRequest handles the log output for incoming requests to the protocol..
func (l *Logger) incomingRequest(r protocol.RequestPayload) {
	str := fmt.Sprintf("%s: %s\n", time.Now().Format("2006/02/01 15:04:05"), l.incomingRequestText(r))
	l.logFile.WriteString(str)

	if l.Details {
		headersWithJoinedValues, keys := l.incomingRequestHeaders(r.Headers)
		for _, key := range keys {
			str := fmt.Sprintf("%s: %s: %s\n", time.Now().Format("2006/02/01 15:04:05"), key, headersWithJoinedValues[key])
			l.logFile.WriteString(str)
		}
	}
}

// incomingRequestText converts the RequestPayload into a printable string.
func (l *Logger) incomingRequestText(r protocol.RequestPayload) string {
	return fmt.Sprintf("%s %s %s", r.Fields.Method, r.Fields.Url, r.Message)
}

// incomingRequestHeaders takes the headers from the request, sorts them alphabetically,
// joins the values, and creates a new map
func (l *Logger) incomingRequestHeaders(headers map[string][]string) (map[string]string, []string) {
	headersWithJoinedValues := make(map[string]string)
	for key, val := range headers {
		value := strings.Join(val, ",")
		headersWithJoinedValues[key] = value
	}

	keys := make([]string, 0, len(headers))
	for key := range headers {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return headersWithJoinedValues, keys
}
