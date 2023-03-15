package grpcserver

import (
	"context"
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/domain/entity/request"

	"github.com/google/uuid"
)

func (s *server) CreateUser(ctx context.Context, user *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	//generate uuid
	id := uuid.New().String()

	//mapper
	d, err := request.MapGRPCCreateToUser(user)
	if err != nil {
		return &userpb.CreateUserResponse{
			Message: err.Error(),
		}, err
	}

	//set id
	err = d.SetID(id)
	if err != nil {
		return &userpb.CreateUserResponse{
			Message: err.Error(),
		}, err
	}

	//execute usecase
	err = s.userUC.Create(d)
	if err != nil {
		return &userpb.CreateUserResponse{
			Message: err.Error(),
		}, err
	}
	return &userpb.CreateUserResponse{
		Message: "success create user",
	}, nil
}
