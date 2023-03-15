package grpcserver

import (
	"context"
	"errors"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/auth/domain/entity/request"
	"grpcrest/services/auth/domain/entity/response"

	"golang.org/x/crypto/bcrypt"
)

const (
	errIncorrectPassword = "incorrect password"
)

func (s *server) LoginAuth(ctx context.Context, auth *authpb.LoginAuthRequest) (*authpb.LoginAuthResponse, error) {
	//mapper
	d, err := request.MapGrpcLoginToAuth(auth)
	if err != nil {
		return &authpb.LoginAuthResponse{
			Message: err.Error(),
		}, err
	}
	//find password by email
	pass, err := s.authUC.FindPasswordByEmail(d)
	if err != nil {
		return &authpb.LoginAuthResponse{
			Message: err.Error(),
		}, err
	}

	//compare hash and password
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(d.Password))
	if err != nil {
		return &authpb.LoginAuthResponse{
			Message: errIncorrectPassword,
		}, errors.New(errIncorrectPassword)
	}

	//dummy jwt TODO
	jwt := "dummyjwt"

	return response.MapToLoginResponse(jwt)
}
