package comm

import (
	"google.golang.org/grpc"
	"context"
)

type GRPCClient struct {
	dialOpts []grpc.DialOption
}

func NewGRPCClient(option ...grpc.DialOption) (*GRPCClient, error) {
	client := GRPCClient{option}
	return &client, nil
}

func (t *GRPCClient) Connect(address string) (*grpc.ClientConn, error) {
	var dialOpts []grpc.DialOption
	dialOpts = append(dialOpts, grpc.WithInsecure())
	dialOpts = append(dialOpts, grpc.WithDefaultCallOptions())
	conn, err := grpc.DialContext(context.Background(), address, dialOpts...)
	if err != nil {
		return nil, err
	}
	return conn, nil

}
