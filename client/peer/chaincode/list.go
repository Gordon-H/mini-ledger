package chaincode

import (
	"github.com/spf13/cobra"
	"fmt"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the channel that the peer has joined.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("The peer has join the channels : ch1, ch2")
	},
}

