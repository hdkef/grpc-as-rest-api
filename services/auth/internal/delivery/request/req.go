package request

import (
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/auth/domain/entity"
)

func MapGrpcUpdateToAuth(r *authpb.UpdateAuthRequest) (*entity.Auth, error) {
	return MapGrpcCreateToAuth(&authpb.CreateAuthRequest{
		UserId:   r.GetUserId(),
		Email:    r.Email,
		Password: r.Password,
	})
}

func MapGrpcCreateToAuth(r *authpb.CreateAuthRequest) (*entity.Auth, error) {
	var data entity.Auth
	err := data.SetEmail(r.GetEmail())
	if err != nil {
		return nil, err
	}
	err = data.SetPassword(r.GetPassword())
	if err != nil {
		return nil, err
	}
	err = data.SetUserID(r.GetUserId())
	if err != nil {
		return nil, err
	}
	return entity.NewAuth(&data)
}

func MapGrpcLoginToAuth(r *authpb.LoginAuthRequest) (*entity.Auth, error) {
	var data entity.Auth
	err := data.SetEmail(r.GetEmail())
	if err != nil {
		return nil, err
	}
	err = data.SetPassword(r.GetPassword())
	if err != nil {
		return nil, err
	}
	return entity.NewAuth(&data)
}

func MapGrpcDeleteToAuth(r *authpb.DeleteAuthRequest) (*entity.Auth, error) {
	var data entity.Auth
	err := data.SetUserID(r.GetUserId())
	if err != nil {
		return nil, err
	}
	return entity.NewAuth(&data)
}
