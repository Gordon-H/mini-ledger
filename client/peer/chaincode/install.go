package chaincode

import (
	"github.com/spf13/cobra"
	"fmt"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Connect the peer to install a certain chaincode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("The peer has join the channels : ch1, ch2")
	},
}

