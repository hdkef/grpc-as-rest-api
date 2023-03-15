package auth

import (
	"context"
	"errors"
	"log"

	jwtS "grpcrest/pkg/auth/jwt_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthInterceptor(methods []string, jwt jwtS.JWTService_, secret string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)
		for _, v := range methods {
			if info.FullMethod == v {
				md, ok := metadata.FromIncomingContext(ctx)
				if !ok {
					return "", errors.New("unauthorized no metadata provided")
				}
				tokenSlice := md.Get("Authorization")
				if len(tokenSlice) < 1 {
					return "", errors.New("unauthorized no bearer token provided")
				}

				token := tokenSlice[0]
				userId, err := jwt.ParseToken(token, secret)
				if err != nil {
					return "", errors.New("token failed or expired")
				}
				//set userId to grpc context
				ctx = metadata.AppendToOutgoingContext(ctx, "userId", userId)
				break
			}
		}
		return handler(ctx, req)
	}
}
