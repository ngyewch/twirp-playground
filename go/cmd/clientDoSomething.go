package cmd

import (
	"context"
	"fmt"
	"github.com/ngyewch/twirp-playground/rpc2"
	"github.com/spf13/cobra"
)

var (
	clientDoSomethingCmd = &cobra.Command{
		Use:   "doSomething (text)",
		Short: "Do something",
		Args:  cobra.ExactArgs(1),
		RunE:  clientDoSomething,
	}
)

func clientDoSomething(cmd *cobra.Command, args []string) error {
	testService2Client, err := newTestService2Client(cmd)
	if err != nil {
		return err
	}

	text := args[0]

	response, err := testService2Client.DoSomething(context.Background(), &rpc2.DoSomethingRequest{
		Text: text,
	})
	if err != nil {
		return err
	}
	fmt.Printf("%s -> %s\n", text, response.Text)
	return nil
}

func init() {
	clientCmd.AddCommand(clientDoSomethingCmd)
}
