package cmd

import (
	"context"
	"fmt"
	"github.com/ngyewch/twirp-playground/rpc"
	"github.com/spf13/cobra"
	"net/http"
	"strconv"
)

var (
	clientAddCmd = &cobra.Command{
		Use:   "add (a) (b)",
		Short: "Add",
		Args:  cobra.ExactArgs(2),
		RunE:  clientAdd,
	}
)

func clientAdd(cmd *cobra.Command, args []string) error {
	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return err
	}

	b, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return err
	}

	testServiceClient := rpc.NewTestServiceProtobufClient("http://127.0.0.1:8080", &http.Client{})

	response, err := testServiceClient.Add(context.Background(), &rpc.AddRequest{
		A: float32(a),
		B: float32(b),
	})
	if err != nil {
		return err
	}
	fmt.Printf("%f + %f = %f\n", a, b, response.Value)
	return nil
}

func init() {
	clientCmd.AddCommand(clientAddCmd)
}
