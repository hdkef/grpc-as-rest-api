package grpcserver

import (
	"context"
	"errors"
	"fmt"
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
	userId, pass, err := s.authUC.FindUserIdPasswordByEmail(d)
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

	tokenstr, err := s.jwtService.GenerateToken(userId)
	if err != nil {
		return &authpb.LoginAuthResponse{
			Message: errIncorrectPassword,
		}, errors.New(errIncorrectPassword)
	}

	return response.MapToLoginResponse(fmt.Sprintf("Bearer%s", tokenstr))
}
