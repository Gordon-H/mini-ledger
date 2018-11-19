package utils

import (
	"github.com/golang/protobuf/proto"
	"github.com/miniledger/protos/peer"
)

func MustMarshal(msg proto.Message) []byte {
	data, err := Marshal(msg)
	if err != nil {
		panic(err)
	}
	return data
}

func Marshal(msg proto.Message) ([]byte, error) {
	return proto.Marshal(msg)
}

func CreateSignedProposal(headerType protos.HeaderType, createTx proto.Message) (protos.Proposal, error) {
	proto.Marshal(createTx)
	proposal := protos.Proposal{Header: &protos.Header{Type: headerType}, Payload: []byte{}}
	return proposal, nil
}
