package grpcserver

import (
	"context"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/auth/internal/delivery/request"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *server) CreateAuth(ctx context.Context, auth *authpb.CreateAuthRequest) (*authpb.CreateAuthResponse, error) {
	//generate id uuid
	//generate uuid
	id := uuid.New().String()

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
	//set id uuid
	d.SetID(id)

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
