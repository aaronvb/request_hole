# Request Hole CLI
[![go.dev Reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/aaronvb/request_hole)
[![Test](https://github.com/aaronvb/request_hole/workflows/Test/badge.svg)](https://github.com/aaronvb/request_hole/actions/workflows/tests.yml)
[![Builds](https://github.com/aaronvb/request_hole/workflows/Builds/badge.svg)](https://github.com/aaronvb/request_hole/actions/workflows/builds.yml)

`rh` is a CLI tool for creating an ephemeral endpoint for testing and inspecting requests from your application or webhook.

<img width="741" alt="rh" src="https://user-images.githubusercontent.com/100900/120058797-f9d90780-bfe8-11eb-9b1d-f65a27773600.png">

## Installation
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
rh: Request Hole
This CLI tool will let you create a temporary API endpoint for testing purposes.

Usage:
  rh [command]

Available Commands:
  help        Help about any command
  http        Creates an http endpoint
  version     Print version number of Request Hole

Flags:
  -a, --address string      sets the address for the endpoint (default "localhost")
  -h, --help                help for rh
  -p, --port int            sets the port for the endpoint (default 8080)
  -r, --response_code int   sets the response code (default 200)

Use "rh [command] --help" for more information about a command.
```


