package channel

import (
	"github.com/spf13/cobra"
	cm "github.com/miniledger/client/peer/common"
)

var cmd = &cobra.Command{
	Use:   "channel",
	Short: "Channel related commands.",
}

func Cmd(factory *cm.CmdFactory) *cobra.Command{
	cmd.AddCommand(listCmd(factory))
	cmd.AddCommand(createCmd(factory))
	cmd.AddCommand(joinCmd(factory))

	return cmd
}

