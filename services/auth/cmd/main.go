package main

import (
	"grpcrest/services/auth/internal/config"
	grpcserver "grpcrest/services/auth/internal/delivery/grpc_server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	//init config
	cfg := config.NewAppConfig()
	//init db psql
	db := config.InitPSQL(cfg)

	//conn grpc
	lis, err := net.Listen("tcp", cfg.AppPort)
	if err != nil {
		log.Fatal(err.Error())
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts)

	//register grpc server
	grpcserver.CreateAuthHandler(s, cfg, db)

	//serve
	s.Serve(lis)
}
