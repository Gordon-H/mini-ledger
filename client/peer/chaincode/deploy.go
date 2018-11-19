package chaincode

import (
	"github.com/spf13/cobra"
	"fmt"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Connect the peer to deploy a certain chaincode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("The peer has join the channels : ch1, ch2")
	},
}

