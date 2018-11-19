package utils

import (
	"github.com/golang/protobuf/proto"
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
