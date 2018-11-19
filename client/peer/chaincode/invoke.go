package chaincode

import (
	"github.com/spf13/cobra"
	"fmt"
)

var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "Connect the peer to invoke a certain chaincode.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("The peer has join the channels : ch1, ch2")
	},
}

