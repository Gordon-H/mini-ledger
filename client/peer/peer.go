package peer

import (
	"github.com/spf13/cobra"
	"github.com/miniledger/client/peer/channel"
	"github.com/miniledger/client/peer/chaincode"
	cm "github.com/miniledger/client/peer/common"
)

var Cmd = &cobra.Command{
	Use:   "peer",
	Short: "Peer related commands",
}

func init() {
	factory, _ := cm.InitCmdFactory()
	Cmd.AddCommand(channel.Cmd(factory))
	Cmd.AddCommand(chaincode.Cmd(factory))
}

