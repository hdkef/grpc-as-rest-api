package grpcserver

import (
	"context"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/auth/domain/entity/request"

	"golang.org/x/crypto/bcrypt"
)

func (s *server) UpdateAuth(ctx context.Context, auth *authpb.UpdateAuthRequest) (*authpb.UpdateAuthResponse, error) {
	//mapper
	d, err := request.MapGrpcUpdateToAuth(auth)
	if err != nil {
		return &authpb.UpdateAuthResponse{
			Message: err.Error(),
		}, err
	}

	//hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(d.Password), 10)
	if err != nil {
		return &authpb.UpdateAuthResponse{
			Message: err.Error(),
		}, err
	}

	//set pass to hashed
	d.SetPassword(string(hashed))

	//execute usecase
	err = s.authUC.UpdateAuth(d)
	if err != nil {
		return &authpb.UpdateAuthResponse{
			Message: err.Error(),
		}, err
	}
	return &authpb.UpdateAuthResponse{
		Message: "success update auth",
	}, nil
}
