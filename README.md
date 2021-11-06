# Request Hole CLI
[![go.dev Reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/aaronvb/request_hole)
[![Go Tests](https://github.com/aaronvb/request_hole/workflows/Test%20Go/badge.svg)](https://github.com/aaronvb/request_hole/actions/workflows/test_go.yml)
[![JS Tests](https://github.com/aaronvb/request_hole/workflows/Test%20JS/badge.svg)](https://github.com/aaronvb/request_hole/actions/workflows/test_js.yml)
[![Builds](https://github.com/aaronvb/request_hole/workflows/Builds/badge.svg)](https://github.com/aaronvb/request_hole/actions/workflows/builds.yml)

`rh` is a CLI tool for creating an ephemeral endpoint for testing and inspecting HTTP requests, WebSocket connections, and WebSocket messages from your application or webhook.

<img width="1136" alt="Request Hole CLI web ui" src="https://user-images.githubusercontent.com/100900/125158715-9b866500-e10e-11eb-9438-36d0f8325c60.png">


## Installation
### Homebrew
```
brew install aaronvb/request_hole/rh
```

### Clone repo and build
First make sure you have Go installed: https://golang.org/doc/install
```
git clone git@github.com:aaronvb/request_hole.git
cd request_hole
go build -o $GOBIN/rh
```
### Release versions
Download the release version for your system: https://github.com/aaronvb/request_hole/releases

## Usage
```
$ rh
```

```
rh: Request Hole
This CLI tool will let you create a temporary API endpoint for testing purposes.

Usage:
  rh [command]

Available Commands:
  help        Help about any command
  http        Creates an http endpoint
  version     Print version number of Request Hole
  ws          Creates a websocket endpoint

Flags:
  -a, --address string       sets the address for the endpoint (default "localhost")
      --details              shows header details in the request
  -h, --help                 help for rh
      --log string           writes incoming requests to the specified log file (example: --log rh.log)
  -p, --port int             sets the port for the endpoint (default 8080)
  -r, --response_code int    sets the response code (default 200)
      --web                  runs the web UI to show incoming requests
      --web_address string   sets the address for the web UI (default "localhost")
      --web_port int         sets the port for the web UI (default 8081)

Use "rh [command] --help" for more information about a command.
```
## Using the Web UI
### Create an HTTP endpoint
```
$ rh http --web
```
This option will open a web UI that will display the incoming requests. Incoming requests will render live in the browser when they are received.

<img width="1136" alt="Request Hole CLI web ui" src="https://user-images.githubusercontent.com/100900/125158715-9b866500-e10e-11eb-9438-36d0f8325c60.png">

### Create a WebSocket endpoint
```
$ rh ws --web
```
This option will open a web UI that will display the WebSocket connections and incoming messages. Connections and incoming messages will render live in the browser when they are received.

<img width="1375" alt="Screen Shot 2021-11-05 at 15 03 31" src="https://user-images.githubusercontent.com/100900/140592650-fa8aa3d4-41b7-4c2f-9a2b-b81e1bbed575.png">

## Using the CLI
### Creating an HTTP endpoint
To create an http endpoint with default settings (port 8080, return status code 200):
```
$ rh http
```
<img width="785" alt="Request Hole CLI http" src="https://user-images.githubusercontent.com/100900/120266278-474db280-c23d-11eb-9e1f-4d73d18522d5.png">

### Creating a WebSocket endpoint
To create a WebSocket endpoint with default settings (port 8080):
```
$ rh ws
```
<img width="784" alt="Request Hole CLI WebSocket" src="https://user-images.githubusercontent.com/100900/140592519-a965af54-a0a3-44cd-be55-1401c8925590.png">

### Show header details
This option shows all the header details in the incoming request.
```
$ rh http --details
```
<img width="785" alt="Request Hole CLI details" src="https://user-images.githubusercontent.com/100900/120266674-1d48c000-c23e-11eb-8107-50db997ac3cc.png">

### Log to file
This option will write the CLI output to the specified log file. Works with other options such as `--details`.
```
$ rh http --log rh.log
```
<img width="787" alt="Request Hole CLI log" src="https://user-images.githubusercontent.com/100900/120877567-fac2e980-c552-11eb-8ec0-8075bc6c0cd8.png">

## Exposing Request Hole to the internet
Sometimes we need to expose `rh` to the internet to test applications or webhooks from outside of our local dev env. The best way to do this is to use a tunneling service such as [ngrok](https://ngrok.com).
```
$ ngrok http 3001
$ rh http -p 3001
```

## Running Tests and Building
It is recommended to run the JS build first so that the Go build can embed the latest web UI build.

### CLI
```
$ go test -v ./...
$ go build
```

### Web UI
```
$ cd web; yarn test
$ cd web; yarn build
```
#### For Development
```
$ cd web; yarn start
```
Visit `localhost:3000`

## Built With
- Go https://golang.org/
- logparams https://github.com/aaronvb/logparams
- logrequest https://github.com/aaronvb/logrequest
- cobra https://github.com/spf13/cobra
- gqlgen https://github.com/99designs/gqlgen (GraphQL/Websockets)
- gorilla/mux https://github.com/gorilla/mux
- Apollo https://github.com/apollographql/apollo-client (GraphQL frontend)
- create-react-app https://github.com/facebook/create-react-app
- Tailwind https://github.com/tailwindlabs/tailwindcss
- React Testing Library https://github.com/testing-library/react-testing-library
