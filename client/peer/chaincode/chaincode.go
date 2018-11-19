package chaincode

import (
	"github.com/spf13/cobra"
	cm "github.com/miniledger/client/peer/common"
)

var cmd = &cobra.Command{
	Use:   "chaincode",
	Short: "Chaincode related commands.",
}

func Cmd(factory *cm.CmdFactory) *cobra.Command {
	//cmd.AddCommand(listCmd(factory))
	//cmd.AddCommand(installCmd(factory))
	//cmd.AddCommand(deployCmd(factory))
	//cmd.AddCommand(invokeCmd(factory))
	return cmd
}
