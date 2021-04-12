package main

import (
	"fmt"

	"github.com/leophys/fart/server"
	"github.com/spf13/cobra"
	ffmt "gopkg.in/ffmt.v1"
)

var params server.ServeParams

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long: `This subcommand starts an instance of the intercept proxy server
			and listens on the chosen local address.

			The intercepted calls (both requests and responses) are served via
			websocket and a control socket is also bound, to allow a client to
			change the internal state of the server.`,
	// The values of the flags are stored in the global variables in params.
	// `args` instead contains the positional arguments.
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		if debug {
			ffmt.Pjson(params)
		}
		cobra.CheckErr(server.Server(params))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&params.BindAddr, "proxy-addr", "p", ":8080", "Address to bind the proxy server to")
	serveCmd.Flags().StringVarP(&params.CtrlAddr, "ctrl-addr", "c", ":51324", "Address to bind the control socket to")
	serveCmd.Flags().StringVarP(&params.WebsocketAddr, "websocket", "w", ":51325", "Address to bind the websocket to")

	// Whitelist has priority to blacklist. If defined blacklist is ignored
	serveCmd.Flags().StringSliceVarP(&params.WhitelistTarget, "whitelist", "W", []string{}, "List of address you want to proxy")
	serveCmd.Flags().StringSliceVarP(&params.BlacklistTarget, "blacklist", "B", []string{}, "List of address you don't want to proxy. Note that if you have defined whitelist flag this flag will be ignored")
}
