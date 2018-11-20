package utils

import (
	"github.com/golang/protobuf/proto"
	"github.com/miniledger/protos/peer"
	"log"
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

func CreateSignedProposal(headerType protos.HeaderType, createTx proto.Message) (protos.SignedProposal, error) {
	payload, err := proto.Marshal(createTx)
	if err != nil {
		log.Fatal("creating signed proposal failed!")
	}
	proposal := protos.Proposal{Header: &protos.Header{Type: headerType}, Payload: payload}
	return protos.SignedProposal{Signature: nil, Proposal: &proposal}, nil
}
