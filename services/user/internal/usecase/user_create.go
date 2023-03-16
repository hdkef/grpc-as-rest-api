package usecase

import (
	"context"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/user/domain/entity"

	"github.com/google/uuid"
)

// Create implements usecase.UserUsecase
func (u *UserArtifact) Create(d *entity.User) error {
	//generate uuid
	userid := uuid.New().String()

	//create auth w/ grpc client
	//map first
	a := authpb.CreateAuthRequest{
		Email:    d.Email,
		Password: d.Password,
		UserId:   userid,
	}

	_, err := u.authgrpcclient.CreateAuth(context.Background(), &a)
	if err != nil {
		return err
	}

	//set user uuid
	d.ID = userid

	//create user
	err = u.userrepo.Create(d)
	if err != nil {
		return err
	}
	return nil
}
