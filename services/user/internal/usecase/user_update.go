package usecase

import (
	"context"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/user/domain/entity"
)

// Update implements usecase.UserUsecase
func (u *UserArtifact) Update(d *entity.User) error {
	//update auth w/ grpc client
	//map first
	a := authpb.UpdateAuthRequest{
		UserId:   d.ID,
		Email:    d.Email,
		Password: d.Password,
	}

	_, err := u.authgrpcclient.UpdateAuth(context.Background(), &a)
	if err != nil {
		return err
	}

	//check for unique, if exist set to empty string
	exist := u.userrepo.IsExistEmail(&d.Email)
	if exist {
		d.Email = ""
	}

	//update user
	err = u.userrepo.Update(d)
	if err != nil {
		return err
	}
	return nil
}
