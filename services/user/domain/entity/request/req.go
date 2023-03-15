package request

import (
	"errors"
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/domain/entity"
)

const (
	errpagelimit = "please specify page and limit query params"
)

func GetLastIDandLimit(r *userpb.GetAllUserRequest) (page int, limit int, err error) {
	var page32 int32 = r.GetPage()
	var limit32 int32 = r.GetLimit()
	if page32 == 0 || limit32 == 0 {
		err = errors.New(errpagelimit)
		return
	}
	page = int(page32)
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
