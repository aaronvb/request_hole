# Request Hole CLI
[![Workflow](https://img.shields.io/github/workflow/status/aaronvb/request_hole/Test?label=build%2Ftests&style=flat)](https://github.com/aaronvb/request_hole/actions/workflows/tests.yml)

`rq` is a CLI tool for creating an ephemeral endpoint for testing and inspecting requests from your application or webhook.

<img width="741" alt="rq" src="https://user-images.githubusercontent.com/100900/120058797-f9d90780-bfe8-11eb-9b1d-f65a27773600.png">

## Installation
### Clone repo and build
First make sure you have Go installed: https://golang.org/doc/install
```
git clone git@github.com:aaronvb/request_hole.git
cd request_hole
go build -o $GOBIN/rq
```

## Usage
```
rq: Request Hole
This CLI tool will let you create a temporary API endpoint for testing purposes.

Usage:
  rq [command]

Available Commands:
  help        Help about any command
  http        Creates an http endpoint
  version     Print version number of Request Hole

Flags:
  -a, --address string      sets the address for the endpoint (default "localhost")
  -h, --help                help for rq
  -p, --port int            sets the port for the endpoint (default 8080)
  -r, --response_code int   sets the response code (default 200)

Use "rq [command] --help" for more information about a command.
```


