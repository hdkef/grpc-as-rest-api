package grpcserver

import (
	"context"
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/domain/entity/request"
)

func (s *server) DeleteUser(ctx context.Context, user *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {

	//mapper
	d, err := request.MapGRPCDeleteToUser(user)
	if err != nil {
		return &userpb.DeleteUserResponse{
			Message: err.Error(),
		}, err
	}

	//execute usecase
	err = s.userUC.Delete(d)
	if err != nil {
		return &userpb.DeleteUserResponse{
			Message: err.Error(),
		}, err
	}
	return &userpb.DeleteUserResponse{
		Message: "success delete user",
	}, nil
}
