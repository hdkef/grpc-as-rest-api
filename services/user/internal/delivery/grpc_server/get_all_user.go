package grpcserver

import (
	"context"
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/domain/entity/request"
	"grpcrest/services/user/domain/entity/response"
)

func (s *server) GetAllUser(ctx context.Context, p *userpb.GetAllUserRequest) (*userpb.GetAllUserResponse, error) {
	//get last id and limit
	lastid, limit := request.GetLastIDandLimit(p)

	//execute usecase
	data, err := s.userUC.GetAll(&lastid, &limit)
	if err != nil {
		return nil, err
	}

	//send response
	return response.MapToGetAllResponse(data)
}
