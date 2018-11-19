package channel

import (
	"github.com/spf13/cobra"
	"fmt"
	"context"
	"github.com/miniledger/protos/peer"
	cm "github.com/miniledger/client/peer/common"
)

func createCmd(factory *cm.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new channel",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("the peer has joined the following channels:")
			proposal := protos.Proposal{Header: []byte("header"), Payload: []byte("payload")}
			signedProposal := protos.SignedProposal{Signature: []byte("sign"), Proposal: &proposal}
			res, _ := factory.EndorserClient.ProcessProposal(context.Background(), &signedProposal)
			fmt.Println(res.Proposal.Payload)
		},
	}
	return cmd

}
