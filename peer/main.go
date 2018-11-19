package main

import (
	"github.com/miniledger/core/comm"
	"log"
	//"github.com/spf13/viper"
	pb "github.com/miniledger/protos/peer"
	"github.com/miniledger/peer/endorser"
	"github.com/spf13/viper"
)

const configName = "peer"

func main() {
	initConfig()
	address := viper.GetString("peer.listenAddress")
	log.Printf("Starting peer with address %s", address)
	//address := "localhost:6601"
	peerServer, err := comm.NewGRPCServer(address)
	if err != nil {
		log.Printf("Error creating grpc server: %s", err)
	}

	serve := make(chan error)
	serverEndorser := endorser.NewEndorserServer()
	pb.RegisterEndorserServer(peerServer.Server, serverEndorser)

	log.Println("Peer starting to serve")
	go func() {
		var grpcErr error
		if grpcErr = peerServer.Start(); grpcErr != nil {
			log.Fatalf("grpc server start with error: %s", grpcErr)
		} else {
			log.Println("Peer server exited")
		}
		serve <- grpcErr
	}()
	<-serve
}
func initConfig() {
	viper.SetConfigName(configName)
}
