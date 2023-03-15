package grpcserver

import (
	"context"
	"fmt"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/api/internal/config"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func NewAuthHandler(ctx context.Context, gw *runtime.ServeMux, opts []grpc.DialOption, cfg *config.AppConfig) {
	//auth grpc port
	var authport string = fmt.Sprintf(":%s", cfg.AuthGrpcPort)

	err := authpb.RegisterLoginAuthServiceHandlerFromEndpoint(ctx, gw, authport, opts)
	if err != nil {
		log.Fatal("unable register endpoint login auth grpc", err.Error())
	}
}
