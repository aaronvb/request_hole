package main

import "github.com/aaronvb/request_hole/cmd"

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
	cmd.Execute(buildInfo)
}
