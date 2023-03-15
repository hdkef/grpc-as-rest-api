package grpcserver

import (
	"context"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/auth/domain/entity/request"

	"golang.org/x/crypto/bcrypt"
)

func (s *server) CreateAuth(ctx context.Context, auth *authpb.CreateAuthRequest) (*authpb.CreateAuthResponse, error) {
	//mapper
	d, err := request.MapGrpcCreateToAuth(auth)
	if err != nil {
		return &authpb.CreateAuthResponse{
			Message: err.Error(),
		}, err
	}
	//hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(d.Password), 10)
	if err != nil {
		return &authpb.CreateAuthResponse{
			Message: err.Error(),
		}, err
	}

	//set pass to hashed
	d.SetPassword(string(hashed))

	//execute usecase
	err = s.authUC.CreateAuth(d)
	if err != nil {
		return &authpb.CreateAuthResponse{
			Message: err.Error(),
		}, err
	}
	return &authpb.CreateAuthResponse{
		Message: "success create auth",
	}, nil
}
