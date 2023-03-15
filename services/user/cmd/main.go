package main

import (
	"fmt"
	"grpcrest/services/user/internal/config"
	grpcclient "grpcrest/services/user/internal/delivery/grpc_client"
	grpcserver "grpcrest/services/user/internal/delivery/grpc_server"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//init config
	cfg := config.NewAppConfig()
	//init db psql
	db := config.InitPSQL(cfg)

	//dial grpc auth client

	authClientConn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.AuthHost, cfg.AuthPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}

	//register grpc auth client
	authClient, err := grpcclient.AuthGRPCClientHandler(authClientConn)
	if err != nil {
		log.Fatal(err.Error())
	}

	//serve grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.AppPort))
	if err != nil {
		log.Fatal(err.Error())
	}

	// var opts []grpc.ServerOption
	s := grpc.NewServer()

	//register grpc server
	grpcserver.CreateUserHandler(s, cfg, db, *authClient)
	grpcserver.GetAllUserHandler(s, cfg, db, *authClient)

	//serve
	s.Serve(lis)
}
