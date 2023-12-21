package cmd

import (
	"github.com/spf13/cobra"
)

var (
	clientCmd = &cobra.Command{
		Use:   "client",
		Short: "Client",
		Args:  cobra.ExactArgs(0),
		RunE:  help,
	}
)

func init() {
	rootCmd.AddCommand(clientCmd)
}
