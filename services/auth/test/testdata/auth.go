package testdata

import (
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/auth/domain/entity"

	"github.com/bxcodec/faker/v4"
)

func NewAuthEntity() *entity.Auth {
	return &entity.Auth{
		ID:       faker.UUIDDigit(),
		UserID:   faker.UUIDDigit(),
		Password: faker.Password(),
		Email:    faker.Email(),
	}
}

func NewLoginAuth() *authpb.LoginAuthRequest {
	return &authpb.LoginAuthRequest{
		Email:    faker.Email(),
		Password: faker.Password(),
	}
}

func NewLoginAuthNoEmail() *authpb.LoginAuthRequest {
	v := NewLoginAuth()
	v.Email = ""
	return v
}

func NewLoginAuthNoPassword() *authpb.LoginAuthRequest {
	v := NewLoginAuth()
	v.Password = ""
	return v
}

func NewCreateAuthFull() *authpb.CreateAuthRequest {
	return &authpb.CreateAuthRequest{
		UserId:   faker.UUIDDigit(),
		Password: faker.Password(),
		Email:    faker.Email(),
	}
}

func NewCreateAuthNoPassword() *authpb.CreateAuthRequest {
	d := NewCreateAuthFull()
	d.Password = ""
	return d
}

func NewCreateAuthNoEmail() *authpb.CreateAuthRequest {
	d := NewCreateAuthFull()
	d.Email = ""
	return d
}

func NewDeleteAuth() *authpb.DeleteAuthRequest {
	return &authpb.DeleteAuthRequest{
		UserId: faker.UUIDDigit(),
	}
}

func NewDeleteAuthNoUserId() *authpb.DeleteAuthRequest {
	return &authpb.DeleteAuthRequest{}
}

func NewUpdateAuthFull() *authpb.UpdateAuthRequest {
	return &authpb.UpdateAuthRequest{
		UserId:   faker.UUIDDigit(),
		Password: faker.Password(),
		Email:    faker.Email(),
	}
}

func NewUpdateAuthNoPassword() *authpb.UpdateAuthRequest {
	d := NewUpdateAuthFull()
	d.Password = ""
	return d
}

func NewUpdateAuthNoEmail() *authpb.UpdateAuthRequest {
	d := NewUpdateAuthFull()
	d.Email = ""
	return d
}

func NewUpdateAuthNoUserId() *authpb.UpdateAuthRequest {
	d := NewUpdateAuthFull()
	d.UserId = ""
	return d
}
