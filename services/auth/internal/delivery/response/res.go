package response

import (
	authpb "grpcrest/proto/_generated/auth"
)

func MapToLoginResponse(token string) (*authpb.LoginAuthResponse, error) {
	var data authpb.LoginAuthResponse
	data.Message = token
	return &data, nil
}

// func MapToCreateResponse(auth *entity.Auth) (*authpb.CreateAuthResponse, error) {
// 	var data authpb.CreateAuthResponse
// 	data.Message = auth.UserID
// 	return &data, nil
// }

// func MapToDeleteResponse(auth *entity.Auth) (*authpb.DeleteAuthResponse, error) {
// 	var data authpb.DeleteAuthResponse
// 	data.Message = auth.UserID
// 	return &data, nil
// }
