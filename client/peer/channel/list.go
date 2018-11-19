package channel

import (
	"github.com/spf13/cobra"
	"fmt"
	cm "github.com/miniledger/client/peer/common"
	"context"
	"github.com/miniledger/protos/peer"
	"time"
)

func listCmd(factory *cm.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the channel that the peer has joined.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("the peer has joined the following channels:")
			proposal := protos.Proposal{Header: []byte("header"), Payload: []byte("payload")}
			signedProposal := protos.SignedProposal{Signature: []byte("sign"), Proposal: &proposal}
			res, _ := factory.EndorserClient.ProcessProposal(context.Background(), &signedProposal)
			fmt.Println("sleep...")
			time.Sleep(time.Second * 1)
			fmt.Println(res)
			fmt.Println("sleep end...")
		},
	}
	return cmd

}
