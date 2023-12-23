package cmd

import (
	"context"
	"fmt"
	"github.com/ngyewch/twirp-playground/rpc2"
	"github.com/spf13/cobra"
)

var (
	clientToUpperCmd = &cobra.Command{
		Use:   "toUpper (text)",
		Short: "To upper",
		Args:  cobra.ExactArgs(1),
		RunE:  clientToUpper,
	}
)

func clientToUpper(cmd *cobra.Command, args []string) error {
	testService2Client, err := newTestService2Client(cmd)
	if err != nil {
		return err
	}

	text := args[0]

	response, err := testService2Client.ToUpper(context.Background(), &rpc2.ToUpperRequest{
		Text: text,
	})
	if err != nil {
		return err
	}
	fmt.Printf("%s -> %s\n", text, response.Text)
	return nil
}

func init() {
	clientCmd.AddCommand(clientToUpperCmd)
}
