package testdata

import (
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/domain/entity"

	"github.com/bxcodec/faker/v4"
)

func NewUserEntity() *entity.User {
	return &entity.User{
		ID:       faker.UUIDDigit(),
		Name:     faker.Name(),
		Email:    faker.Email(),
		Address:  faker.Word(),
		Password: faker.Password(),
	}
}

func NewDeleteUserRequest() *userpb.DeleteUserRequest {
	return &userpb.DeleteUserRequest{
		UserId: faker.UUIDDigit(),
	}
}

func NewDeleteUserRequestNoUserId() *userpb.DeleteUserRequest {
	v := NewDeleteUserRequest()
	v.UserId = ""
	return v
}

func NewCreateUserRequestFull() *userpb.CreateUserRequest {
	return &userpb.CreateUserRequest{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Address:  faker.Word(),
		Password: faker.Password(),
	}
}

func NewCreateUserRequestNoName() *userpb.CreateUserRequest {
	v := NewCreateUserRequestFull()
	v.Name = ""
	return v
}

func NewCreateUserRequestNoEmail() *userpb.CreateUserRequest {
	v := NewCreateUserRequestFull()
	v.Email = ""
	return v
}

func NewCreateUserRequestNoAddress() *userpb.CreateUserRequest {
	v := NewCreateUserRequestFull()
	v.Address = ""
	return v
}

func NewCreateUserRequestNoPassword() *userpb.CreateUserRequest {
	v := NewCreateUserRequestFull()
	v.Password = ""
	return v
}

func NewUpdateUserRequestFull() *userpb.UpdateUserRequest {
	return &userpb.UpdateUserRequest{
		UserId:   faker.UUIDDigit(),
		Name:     faker.Name(),
		Email:    faker.Email(),
		Address:  faker.Word(),
		Password: faker.Password(),
	}
}

func NewUpdateUserRequestNoName() *userpb.UpdateUserRequest {
	v := NewUpdateUserRequestFull()
	v.Name = ""
	return v
}

func NewUpdateUserRequestNoEmail() *userpb.UpdateUserRequest {
	v := NewUpdateUserRequestFull()
	v.Email = ""
	return v
}

func NewUpdateUserRequestNoAddress() *userpb.UpdateUserRequest {
	v := NewUpdateUserRequestFull()
	v.Address = ""
	return v
}

func NewUpdateUserRequestNoPassword() *userpb.UpdateUserRequest {
	v := NewUpdateUserRequestFull()
	v.Password = ""
	return v
}

func NewUpdateUserRequestNoUserId() *userpb.UpdateUserRequest {
	v := NewUpdateUserRequestFull()
	v.UserId = ""
	return v
}
