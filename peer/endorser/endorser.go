package endorser

import (
	"github.com/miniledger/protos/peer"
	"golang.org/x/net/context"
	"log"
	"encoding/json"
	"bytes"
)

func NewEndorserServer() *endorser {
	return &endorser{}
}

type endorser struct {
}

func (*endorser) ProcessProposal(context.Context, *protos.SignedProposal) (*protos.SignedProposal, error) {
	log.Println("============  begin to process proposal ===============")
	payload := []string{"ch1", "ch2"}
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.Encode(payload)
	proposal := protos.Proposal{Header: []byte("header"), Payload: buf.Bytes()}
	return &protos.SignedProposal{Signature: []byte("fdsf"), Proposal: &proposal}, nil
}
