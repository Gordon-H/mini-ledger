package common

import (
	pb "github.com/miniledger/protos/peer"
	"github.com/miniledger/core/comm"
)

type CmdFactory struct {
	EndorserClient pb.EndorserClient
	//Signer          msp.SigningIdentity
	//BroadcastClient common.BroadcastClient
}

func InitCmdFactory() (*CmdFactory, error) {
	client, _ := comm.NewGRPCClient()
	address := "localhost:6601"
	conn, err := client.Connect(address)
	if err != nil {
		panic(err)
		return nil, err
	}

	endorser := pb.NewEndorserClient(conn)
	return &CmdFactory{
		EndorserClient: endorser,
	}, nil
}
