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
