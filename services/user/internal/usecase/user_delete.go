package usecase

import (
	"context"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/user/domain/entity"
)

// Delete implements usecase.UserUsecase
func (u *UserArtifact) Delete(d *entity.User) error {
	//delete auth w/ grpc client
	a := authpb.DeleteAuthRequest{
		UserId: d.ID,
	}

	// log.Println("DEBUG USECASE DELETE")

	_, err := u.authgrpcclient.Delete.DeleteAuth(context.Background(), &a)
	if err != nil {
		return err
	}

	//delete user
	err = u.userrepo.Delete(d)
	if err != nil {
		return err
	}
	return nil
}
