package auth

import (
	"context"
	"errors"

	jwtS "grpcrest/pkg/auth/jwt_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthInterceptor(methods []string, jwt jwtS.JWTService_) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
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

				tokenBearer := tokenSlice[0]
				token := tokenBearer[7:]
				userId, err := jwt.ParseToken(token)
				if err != nil {
					return "", errors.New("token failed or expired")
				}
				//set userId to grpc context
				ctx = metadata.AppendToOutgoingContext(ctx, "x-user-id", userId)
				return handler(ctx, req)
			}
		}
		return handler(ctx, req)
	}
}
