package usecase

import (
	"errors"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/user/domain/entity"
	"grpcrest/services/user/domain/repository"
	grpcclient "grpcrest/services/user/internal/delivery/grpc_client"
	"grpcrest/services/user/mocks"
	"grpcrest/services/user/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	type field struct {
		authgrpcclient grpcclient.AuthGRPCClient_
		userrepo       repository.UserRepository
	}

	type args struct {
		user *entity.User
	}

	fakeEntity := testdata.NewUserEntity()

	mockUserRepoSuccess := new(mocks.UserRepository)
	mockUserRepoSuccess.On("Create", fakeEntity).Return(nil)
	mockUserRepoFail := new(mocks.UserRepository)
	mockUserRepoFail.On("Create", fakeEntity).Return(errors.New("some error"))
	mockAuthGrpcClientSuccess := new(mocks.AuthGRPCClient_)
	mockAuthGrpcClientSuccess.On("CreateAuth", mock.Anything, mock.Anything).Return(&authpb.CreateAuthResponse{}, nil)
	mockAuthGrpcClientFail := new(mocks.AuthGRPCClient_)
	mockAuthGrpcClientFail.On("CreateAuth", mock.Anything, mock.Anything).Return(&authpb.CreateAuthResponse{}, errors.New("some error"))

	testcases := []struct {
		Name  string
		Field field
		Args  args
		Err   error
	}{
		{
			Name: "success",
			Field: field{
				authgrpcclient: mockAuthGrpcClientSuccess,
				userrepo:       mockUserRepoSuccess,
			},
			Args: args{
				user: fakeEntity,
			},
			Err: nil,
		},
		{
			Name: "user repo fail",
			Field: field{
				authgrpcclient: mockAuthGrpcClientSuccess,
				userrepo:       mockUserRepoFail,
			},
			Args: args{
				user: fakeEntity,
			},
			Err: errors.New("some error"),
		},
		{
			Name: "authgrpcclient fail",
			Field: field{
				authgrpcclient: mockAuthGrpcClientFail,
				userrepo:       mockUserRepoSuccess,
			},
			Args: args{
				user: fakeEntity,
			},
			Err: errors.New("some error"),
		},
		{
			Name: "authgrpcclient faila and userrepo fail",
			Field: field{
				authgrpcclient: mockAuthGrpcClientFail,
				userrepo:       mockUserRepoFail,
			},
			Args: args{
				user: fakeEntity,
			},
			Err: errors.New("some error"),
		},
	}

	for _, v := range testcases {
		t.Run(v.Name, func(t *testing.T) {

			//dependency inject
			userUsecase := NewUserUsecase(v.Field.userrepo, v.Field.authgrpcclient)
			//execute
			err := userUsecase.Create(v.Args.user)

			//test
			if v.Err == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
