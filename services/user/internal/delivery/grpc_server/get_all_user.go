package grpcserver

import (
	"context"
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/internal/delivery/request"
	"grpcrest/services/user/internal/delivery/response"
)

func (s *server) GetAllUser(ctx context.Context, p *userpb.GetAllUserRequest) (*userpb.GetAllUserResponse, error) {
	//get last id and limit
	lastid, limit, err := request.GetLastIDandLimit(p)
	if err != nil {
		return nil, err
	}

	//execute usecase
	data, err := s.userUC.GetAll(&lastid, &limit)
	if err != nil {
		return nil, err
	}

	//send response
	return response.MapToGetAllResponse(data)
}
