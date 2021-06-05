package renderer

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/aaronvb/logrequest"
	"github.com/pterm/pterm"
)

// Logger outputs to a log file.
type Logger struct {
	// File points to the file that we write to.
	File string

	// Fields for startText
	Port int
	Addr string
}

// Start writes the initial server start to the log file.
func (l *Logger) Start() {
	if l.File == "" {
		return
	}

	f, err := os.OpenFile(l.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		l.Fatal(err)
	}

	defer f.Close()

	str := fmt.Sprintf("%s: %s\n", time.Now().Format("2006/02/01 15:04:05"), l.startText())
	f.WriteString(str)
}

func (l *Logger) startText() string {
	return fmt.Sprintf("Listening on http://%s:%d", l.Addr, l.Port)
}

// Fatal will use the Error prefix to render the error and then exit the CLI.
func (l *Logger) Fatal(err error) {
	pterm.Error.WithShowLineNumber(false).Println(err)
	os.Exit(1)
}

// IncomingRequest writes the incoming requests to the log file.
func (l *Logger) IncomingRequest(fields logrequest.RequestFields, params string) {
	if l.File == "" {
		return
	}

	f, err := os.OpenFile(l.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		l.Fatal(err)
	}

	defer f.Close()

	str := fmt.Sprintf("%s: %s\n", time.Now().Format("2006/02/01 15:04:05"), l.incomingRequestText(fields, params))
	f.WriteString(str)
}

func (l *Logger) incomingRequestText(fields logrequest.RequestFields, params string) string {
	return fmt.Sprintf("%s %s %s", fields.Method, fields.Url, params)
}

// IncomingRequestHeaders writes the incoming request headers to the log file
func (l *Logger) IncomingRequestHeaders(headers map[string][]string) {
	if l.File == "" {
		return
	}

	f, err := os.OpenFile(l.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		l.Fatal(err)
	}

	defer f.Close()

	headersWithJoinedValues, keys := l.incomingRequestHeaders(headers)

	for _, key := range keys {
		str := fmt.Sprintf("%s: %s: %s\n", time.Now().Format("2006/02/01 15:04:05"), key, headersWithJoinedValues[key])
		f.WriteString(str)
	}
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
