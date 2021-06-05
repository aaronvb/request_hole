package cmd

import (
	"github.com/aaronvb/request_hole/pkg/renderer"
	"github.com/aaronvb/request_hole/pkg/server"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Creates an http endpoint",
	Long: `rh: http
Create an endpoint that accepts http connections.
`,
	Run: http,
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

func http(cmd *cobra.Command, args []string) {
	renderer := &renderer.Printer{Port: Port, Addr: Address, BuildInfo: BuildInfo, Details: Details}
	httpServer := server.Http{Addr: Address, Port: Port, ResponseCode: ResponseCode, Output: renderer}
	httpServer.Start()
}
