package cmd

import (
	"github.com/aaronvb/request_hole/pkg/protocol"
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
	renderers := make([]renderer.Renderer, 0)

	// Collect flag data into struct to use with renderers
	flagData := server.FlagData{
		Addr:         Address,
		BuildInfo:    BuildInfo,
		Details:      Details,
		LogFile:      LogFile,
		Port:         Port,
		ResponseCode: ResponseCode,
		Web:          Web,
		WebPort:      WebPort,
	}

	if Web {
		web := &renderer.Web{Port: WebPort}
		renderers = append(renderers, web)
	} else {
		printer := &renderer.Printer{Details: Details}
		renderers = append(renderers, printer)
	}

	if LogFile != "" {
		logger := &renderer.Logger{
			FilePath: LogFile,
			Details:  Details,
			Addr:     Address,
			Port:     Port,
		}
		renderers = append(renderers, logger)
	}

	httpServer := &protocol.Http{
		Addr:         Address,
		Port:         Port,
		ResponseCode: ResponseCode,
	}

	srv := server.Server{
		FlagData:  flagData,
		Protocol:  httpServer,
		Renderers: renderers,
	}

	srv.Start()
}
