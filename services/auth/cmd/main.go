package main

import (
	"fmt"
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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.AppPort))
	if err != nil {
		log.Fatal(err.Error())
	}

	// var opts []grpc.ServerOption
	s := grpc.NewServer()

	//register grpc server
	grpcserver.CreateAuthHandler(s, cfg, db)
	grpcserver.DeleteAuthHandler(s, cfg, db)
	grpcserver.LoginAuthHandler(s, cfg, db)
	grpcserver.UpdateAuthHandler(s, cfg, db)

	//serve
	s.Serve(lis)
}
