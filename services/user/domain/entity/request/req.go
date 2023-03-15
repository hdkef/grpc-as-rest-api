package request

import (
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/domain/entity"
)

func GetLastIDandLimit(r *userpb.GetAllUserRequest) (lastid int, limit int) {
	var lastid32 int32 = r.GetLastId()
	var limit32 int32 = r.GetLimit()
	lastid = int(lastid32)
	limit = int(limit32)
	return
}

func MapGRPCCreateToUser(r *userpb.CreateUserRequest) (*entity.User, error) {
	var data entity.User
	err := data.SetName(r.GetName())
	if err != nil {
		return nil, err
	}

	err = data.SetEmail(r.GetEmail())
	if err != nil {
		return nil, err
	}

	err = data.SetAddress(r.GetAddress())
	if err != nil {
		return nil, err
	}

	err = data.SetPassword(r.GetPassword())
	if err != nil {
		return nil, err
	}
	return entity.NewUser(&data)
}

func MapGRPCDeleteToUser(r *userpb.DeleteUserRequest) (*entity.User, error) {
	var data entity.User
	err := data.SetID(r.GetUserId())
	if err != nil {
		return nil, err
	}
	return entity.NewUser(&data)
}

func MapGRPCUpdateToUser(r *userpb.UpdateUserRequest) (*entity.User, error) {
	var data entity.User
	err := data.SetID(r.GetUserId())
	if err != nil {
		return nil, err
	}
	err = data.SetName(r.GetName())
	if err != nil {
		return nil, err
	}

	err = data.SetEmail(r.GetEmail())
	if err != nil {
		return nil, err
	}

	err = data.SetAddress(r.GetAddress())
	if err != nil {
		return nil, err
	}

	err = data.SetPassword(r.GetPassword())
	if err != nil {
		return nil, err
	}

	return entity.NewUser(&data)
}
