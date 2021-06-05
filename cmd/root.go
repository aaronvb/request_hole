// Package cmd contains all of the CLI commands
package cmd

import "github.com/spf13/cobra"

var (
	Address      string
	Port         int
	ResponseCode int
	BuildInfo    map[string]string
	Details      bool
	LogFile      string
)

var rootCmd = &cobra.Command{
	Use:   "rh",
	Short: "A CLI for an ephemeral API endpoint",
	Long: `rh: Request Hole
This CLI tool will let you create a temporary API endpoint for testing purposes.`,
}

func Execute(buildInfo map[string]string) error {
	BuildInfo = buildInfo
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8080, "sets the port for the endpoint")
	rootCmd.PersistentFlags().StringVarP(&Address, "address", "a", "localhost", "sets the address for the endpoint")
	rootCmd.PersistentFlags().IntVarP(&ResponseCode, "response_code", "r", 200, "sets the response code")
	rootCmd.PersistentFlags().BoolVar(&Details, "details", false, "shows header details in the request")
	rootCmd.PersistentFlags().StringVar(&LogFile, "log", "", "writes incoming requests to the specified log file (example: --log rh.log)")
}
