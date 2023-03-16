package grpcserver

import (
	"context"
	"errors"
	"grpcrest/pkg/auth"
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/internal/delivery/request"
)

func (s *server) UpdateUser(ctx context.Context, user *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	//get userId, only user id match in jwt can update
	jwtuserid, err := auth.ExtractUserID(ctx)
	if err != nil {
		return &userpb.UpdateUserResponse{
			Message: err.Error(),
		}, err
	}

	//mapper
	d, err := request.MapGRPCUpdateToUser(user)
	if err != nil {
		return &userpb.UpdateUserResponse{
			Message: err.Error(),
		}, err
	}

	if jwtuserid != d.ID {
		return &userpb.UpdateUserResponse{
			Message: "target user id did not match with your user id from token",
		}, errors.New("target user id did not match with your user id from token")
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
