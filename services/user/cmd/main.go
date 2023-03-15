package main

import (
	"grpcrest/services/user/internal/config"
	grpcclient "grpcrest/services/user/internal/delivery/grpc_client"
	grpcserver "grpcrest/services/user/internal/delivery/grpc_server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	//init config
	cfg := config.NewAppConfig()
	//init db psql
	db := config.InitPSQL(cfg)

	//dial grpc auth client

	authClientConn, err := grpc.Dial(cfg.AuthPort)
	if err != nil {
		log.Fatal(err.Error())
	}

	//register grpc auth client
	authClient, err := grpcclient.AuthGRPCClientHandler(authClientConn)
	if err != nil {
		log.Fatal(err.Error())
	}

	//serve grpc server
	lis, err := net.Listen("tcp", cfg.AppPort)
	if err != nil {
		log.Fatal(err.Error())
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts)

	//register grpc server
	grpcserver.CreateUserHandler(s, cfg, db, *authClient)
	grpcserver.GetAllUserHandler(s, cfg, db, *authClient)

	//serve
	s.Serve(lis)
}
