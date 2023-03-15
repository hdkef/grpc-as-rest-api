package main

import (
	"context"
	"fmt"
	"grpcrest/services/api/internal/config"
	grpcserver "grpcrest/services/api/internal/delivery/grpc_server"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func customHeaderMatcher() runtime.HeaderMatcherFunc {
	return func(headerKey string) (string, bool) {
		switch headerKey {
		case "Authorization":
			//validate jwt here
			return headerKey, true
		}
		return runtime.DefaultHeaderMatcher(headerKey)
	}
}

func main() {
	//init config
	cfg := config.NewAppConfig()

	//conn grpc
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcPort))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer()

	//init runtime gateway
	gw := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(customHeaderMatcher()),
	)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(),
	}

	//register user grpc gateway
	grpcserver.NewAuthHandler(context.Background(), gw, opts, cfg)
	grpcserver.NewUserHandler(context.Background(), gw, opts, cfg)

	//serve grpc
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	//serve http
	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.AppPort),
		Handler: gw,
	}

	log.Fatalln(gwServer.ListenAndServe())
}
