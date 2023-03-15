package grpcserver

import (
	"context"
	"fmt"
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/api/internal/config"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func NewUserHandler(ctx context.Context, gw *runtime.ServeMux, opts []grpc.DialOption, cfg *config.AppConfig) {
	//user grpc port
	var userport string = fmt.Sprintf(":%s", cfg.UserGrpcPort)

	err := userpb.RegisterCreateUserServiceHandlerFromEndpoint(ctx, gw, userport, opts)
	if err != nil {
		log.Fatal("unable register endpoint create user grpc", err.Error())
	}

	err = userpb.RegisterUpdateUserServiceHandlerFromEndpoint(ctx, gw, userport, opts)
	if err != nil {
		log.Fatal("unable register endpoint update user grpc", err.Error())
	}

	err = userpb.RegisterDeleteUserServiceHandlerFromEndpoint(ctx, gw, userport, opts)
	if err != nil {
		log.Fatal("unable register endpoint delete user grpc", err.Error())
	}

	err = userpb.RegisterGetAllUserServiceHandlerFromEndpoint(ctx, gw, userport, opts)
	if err != nil {
		log.Fatal("unable register endpoint get all user grpc", err.Error())
	}
}
