package grpcserver

import (
	"context"
	"errors"
	"grpcrest/pkg/auth"
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/internal/delivery/request"
)

func (s *server) DeleteUser(ctx context.Context, user *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	//get userId, only user id match in jwt can delete
	jwtuserid, err := auth.ExtractUserID(ctx)
	if err != nil {
		return &userpb.DeleteUserResponse{
			Message: err.Error(),
		}, err
	}

	//mapper
	d, err := request.MapGRPCDeleteToUser(user)
	if err != nil {
		return &userpb.DeleteUserResponse{
			Message: err.Error(),
		}, err
	}

	if jwtuserid != d.ID {
		return &userpb.DeleteUserResponse{
			Message: "target user id did not match with your user id from token",
		}, errors.New("target user id did not match with your user id from token")
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
