package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/aaronvb/request_hole/cmd"
)

//go:embed web/build
var embededFiles embed.FS
var staticFS http.FileSystem

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	repo    = "github.com/aaronvb/request_hole"
	builtBy = "dev"
)

func main() {
	buildInfo := map[string]string{
		"version": version,
		"commit":  commit,
		"date":    date,
		"repo":    repo,
		"builtBy": builtBy,
	}

	// Set file system for the web UI
	useOSFS := version == "dev"
	staticFS = getFileSystem(useOSFS)

	cmd.Execute(buildInfo, staticFS)
}

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		return http.FS(os.DirFS("web/build"))
	}

	fsys, err := fs.Sub(embededFiles, "web/build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
