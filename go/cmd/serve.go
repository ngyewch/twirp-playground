package cmd

import (
	"github.com/ngyewch/twirp-playground/rpc"
	"github.com/ngyewch/twirp-playground/server"
	"github.com/spf13/cobra"
	"net/http"
	"strings"
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
	serverInstance := &server.Server{}
	testServiceHandler := rpc.NewTestServiceServer(serverInstance)
	testService2Handler := rpc.NewTestService2Server(serverInstance)
	handlers := []rpc.TwirpServer{
		testServiceHandler,
		testService2Handler,
	}
	return http.ListenAndServe(":8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, handler := range handlers {
				if strings.HasPrefix(r.RequestURI, handler.PathPrefix()) {
					handler.ServeHTTP(w, r)
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		}))
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
