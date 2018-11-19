package main

import (
	"github.com/miniledger/common/tools/configtxgen/channel"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	RootCmd := &cobra.Command{Use: "configtxgen", Short:"config related options"}
	RootCmd.AddCommand(channel.CreateChannelTxCmd)

	if RootCmd.Execute() != nil {
		os.Exit(1)
	}


}

