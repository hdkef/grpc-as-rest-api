package response

import (
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/domain/entity"
)

func MapToGetAllResponse(u []entity.User) (*userpb.GetAllUserResponse, error) {
	var data []*userpb.User
	for _, v := range u {
		data = append(data, &userpb.User{
			Email:   v.Email,
			Address: v.Address,
			Name:    v.Name,
		})
	}
	return &userpb.GetAllUserResponse{
		Data: data,
	}, nil
}
