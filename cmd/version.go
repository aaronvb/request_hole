package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of Request Hole",
	Run: func(cmd *cobra.Command, args []string) {
		var date string

		buildDate, err := time.Parse(time.RFC3339, BuildInfo["date"])
		if err != nil {
			date = BuildInfo["date"]
		} else {
			date = buildDate.Format("2006-01-02 15:04:05")
		}

		fmt.Printf("Request Hole %s\n%s\n\nBuild date: %s\nCommit: %s\nBuild by: %s", BuildInfo["version"], BuildInfo["repo"], date, BuildInfo["commit"], BuildInfo["builtBy"])
	},
}
