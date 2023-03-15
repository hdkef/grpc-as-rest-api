package usecase

import (
	"context"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/user/domain/entity"
	"grpcrest/services/user/domain/repository"
	"grpcrest/services/user/domain/usecase"
	grpcclient "grpcrest/services/user/internal/delivery/grpc_client"

	"github.com/google/uuid"
)

type UserArtifact struct {
	authgrpcclient grpcclient.AuthGRPCClient
	userrepo       repository.UserRepository
}

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

	_, err := u.authgrpcclient.Create.CreateAuth(context.Background(), &a)
	if err != nil {
		return err
	}

	//create user
	err = u.userrepo.Create(d)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements usecase.UserUsecase
func (u *UserArtifact) Delete(d *entity.User) error {
	return u.userrepo.Delete(d)
}

// GetAll implements usecase.UserUsecase
func (u *UserArtifact) GetAll(lastID *int, limit *int) ([]entity.User, error) {
	return u.userrepo.GetAll(lastID, limit)
}

// Update implements usecase.UserUsecase
func (u *UserArtifact) Update(d *entity.User) error {
	//update auth w/ grpc client
	//map first
	a := authpb.UpdateAuthRequest{
		UserId:   d.ID,
		Email:    d.Email,
		Password: d.Password,
	}

	_, err := u.authgrpcclient.Update.UpdateAuth(context.Background(), &a)
	if err != nil {
		return err
	}

	//update user
	err = u.userrepo.Update(d)
	if err != nil {
		return err
	}
	return nil
}

func NewUserUsecase(repo repository.UserRepository, authgrpcclient grpcclient.AuthGRPCClient) usecase.UserUsecase {
	return &UserArtifact{
		userrepo:       repo,
		authgrpcclient: authgrpcclient,
	}
}
