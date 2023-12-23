package cmd

import (
	"fmt"
	"github.com/ngyewch/twirp-playground/rpc"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	clientCmd = &cobra.Command{
		Use:   "client",
		Short: "Client",
		Args:  cobra.ExactArgs(0),
		RunE:  help,
	}
)

func getClientEncoding(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("encoding")
}

func getClientEndpoint(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("endpoint")
}

func newTestServiceClient(cmd *cobra.Command) (rpc.TestService, error) {
	endpoint, err := getClientEndpoint(cmd)
	if err != nil {
		return nil, err
	}
	encoding, err := getClientEncoding(cmd)
	if err != nil {
		return nil, err
	}

	switch encoding {
	case "protobuf":
		return rpc.NewTestServiceProtobufClient(endpoint, &http.Client{}), nil
	case "json":
		return rpc.NewTestServiceJSONClient(endpoint, &http.Client{}), nil
	default:
		return nil, fmt.Errorf("unknown encoding")
	}
}

func newTestService2Client(cmd *cobra.Command) (rpc.TestService2, error) {
	endpoint, err := getClientEndpoint(cmd)
	if err != nil {
		return nil, err
	}
	encoding, err := getClientEncoding(cmd)
	if err != nil {
		return nil, err
	}

	switch encoding {
	case "protobuf":
		return rpc.NewTestService2ProtobufClient(endpoint, &http.Client{}), nil
	case "json":
		return rpc.NewTestService2JSONClient(endpoint, &http.Client{}), nil
	default:
		return nil, fmt.Errorf("unknown encoding")
	}
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.PersistentFlags().String("encoding", "protobuf", "Encoding (protobuf, json)")
	clientCmd.PersistentFlags().String("endpoint", "http://127.0.0.1:8080", "Endpoint")
}
