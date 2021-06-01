# Request Hole CLI
[![go.dev Reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/aaronvb/request_hole)
[![Test](https://github.com/aaronvb/request_hole/workflows/Test/badge.svg)](https://github.com/aaronvb/request_hole/actions/workflows/tests.yml)
[![Builds](https://github.com/aaronvb/request_hole/workflows/Builds/badge.svg)](https://github.com/aaronvb/request_hole/actions/workflows/builds.yml)

`rh` is a CLI tool for creating an ephemeral endpoint for testing and inspecting requests from your application or webhook.

<img width="788" alt="Request Hole CLI" src="https://user-images.githubusercontent.com/100900/120265767-63048900-c23c-11eb-9a20-079ab9822767.png">


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

Flags:
  -a, --address string      sets the address for the endpoint (default "localhost")
      --details             shows header details in the request
  -h, --help                help for rh
  -p, --port int            sets the port for the endpoint (default 8080)
  -r, --response_code int   sets the response code (default 200)

Use "rh [command] --help" for more information about a command.
```

### Creating an HTTP endpoint
To create an http endpoint with default settings (port 8080, return status code 200):
```
$ rh http
```
<img width="785" alt="Request Hole CLI http" src="https://user-images.githubusercontent.com/100900/120266278-474db280-c23d-11eb-9e1f-4d73d18522d5.png">

### Show header details
This option shows all the header details in the incoming request.
```
$ rh http --details
```
<img width="785" alt="Request Hole CLI details" src="https://user-images.githubusercontent.com/100900/120266674-1d48c000-c23e-11eb-8107-50db997ac3cc.png">

### Exposing Request Hole to the internet
Sometimes we need to expose `rh` to the internet to test applications or webhooks from outside of our local dev env. The best way to do this is to use a tunneling service such as [ngrok](https://ngrok.com).
```
$ ngrok http 3001
$ rh http -p 3001
```
