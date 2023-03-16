package grpcserver

import (
	"context"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/auth/internal/delivery/request"
)

func (s *server) DeleteAuth(ctx context.Context, auth *authpb.DeleteAuthRequest) (*authpb.DeleteAuthResponse, error) {
	//mapper
	d, err := request.MapGrpcDeleteToAuth(auth)
	if err != nil {
		return &authpb.DeleteAuthResponse{
			Message: err.Error(),
		}, err
	}
	//execute usecase
	err = s.authUC.DeleteAuth(d)
	if err != nil {
		return &authpb.DeleteAuthResponse{
			Message: err.Error(),
		}, err
	}
	return &authpb.DeleteAuthResponse{
		Message: "success delete auth",
	}, nil
}
