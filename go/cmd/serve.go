package cmd

import (
	"github.com/ngyewch/twirp-playground/rpc"
	"github.com/ngyewch/twirp-playground/server"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve",
		Args:  cobra.ExactArgs(0),
		RunE:  serve,
	}
)

func serve(cmd *cobra.Command, args []string) error {
	testServiceServer := &server.Server{}
	testServiceServerHandler := rpc.NewTestServiceServer(testServiceServer)
	return http.ListenAndServe(":8080", testServiceServerHandler)
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
