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
	Run: httpCommand,
}

var wsCmd = &cobra.Command{
	Use:   "ws",
	Short: "Creates a websocket endpoint",
	Long: `rh: ws
Create an endpoint that accepts websocket connections.
`,
	Run: wsCommand,
}

func init() {
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(wsCmd)
}

func httpCommand(cmd *cobra.Command, args []string) {
	renderers := make([]renderer.Renderer, 0)

	// Collect flag data into struct to use with renderers
	flagData := server.FlagData{
		Addr:         Address,
		BuildInfo:    BuildInfo,
		Details:      Details,
		LogFile:      LogFile,
		Port:         Port,
		Protocol:     "http",
		ResponseCode: ResponseCode,
		Web:          Web,
		WebAddress:   WebAddress,
		WebPort:      WebPort,
	}

	if Web {
		web := &renderer.Web{
			Address:      WebAddress,
			Port:         WebPort,
			StaticFiles:  StaticFS,
			RequestAddr:  Address,
			RequestPort:  Port,
			ResponseCode: ResponseCode,
			BuildInfo:    BuildInfo,
			Protocol:     "http",
		}
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
			Protocol: "http",
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

func wsCommand(cmd *cobra.Command, args []string) {
	renderers := make([]renderer.Renderer, 0)

	// Collect flag data into struct to use with renderers
	flagData := server.FlagData{
		Addr:       Address,
		BuildInfo:  BuildInfo,
		Details:    Details,
		LogFile:    LogFile,
		Port:       Port,
		Protocol:   "ws",
		Web:        Web,
		WebAddress: WebAddress,
		WebPort:    WebPort,
	}

	if Web {
		web := &renderer.Web{
			Address:      WebAddress,
			Port:         WebPort,
			StaticFiles:  StaticFS,
			RequestAddr:  Address,
			RequestPort:  Port,
			ResponseCode: ResponseCode,
			BuildInfo:    BuildInfo,
			Protocol:     "ws",
		}
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
			Protocol: "ws",
		}
		renderers = append(renderers, logger)
	}

	wsServer := &protocol.Ws{
		Addr: Address,
		Port: Port,
	}

	srv := server.Server{
		FlagData:  flagData,
		Protocol:  wsServer,
		Renderers: renderers,
	}

	srv.Start()
}
