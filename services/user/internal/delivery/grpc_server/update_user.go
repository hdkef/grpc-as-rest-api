package grpcserver

import (
	"context"
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/internal/delivery/request"
)

func (s *server) UpdateUser(ctx context.Context, user *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {

	//mapper
	d, err := request.MapGRPCUpdateToUser(user)
	if err != nil {
		return &userpb.UpdateUserResponse{
			Message: err.Error(),
		}, err
	}

	//execute usecase
	err = s.userUC.Update(d)
	if err != nil {
		return &userpb.UpdateUserResponse{
			Message: err.Error(),
		}, err
	}
	return &userpb.UpdateUserResponse{
		Message: "success update user",
	}, nil
}
