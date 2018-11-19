package comm

import (
	"net"
	"google.golang.org/grpc"
	"errors"
)

type GRPCServer struct {
	Server   *grpc.Server
	Address  string
	Listener net.Listener
}

func NewGRPCServer(address string) (*GRPCServer, error) {
	if address == "" {
		return nil, errors.New("address should not be empty")
	}
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	server := GRPCServer{grpc.NewServer(), lis.Addr().String(), lis}
	return &server, nil
}

func (t *GRPCServer) Start() error {
	return t.Server.Serve(t.Listener)
}
