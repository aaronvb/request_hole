package cmd

import (
	"fmt"

	"github.com/aaronvb/request_hole/pkg/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of Request Hole",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Request Hole %s\n%s\n", version.BuildVersion, version.BuildRepo)
	},
}
