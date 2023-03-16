package usecase

import (
	"errors"
	"grpcrest/services/user/domain/entity"
	"grpcrest/services/user/domain/repository"
	grpcclient "grpcrest/services/user/internal/delivery/grpc_client"
	"grpcrest/services/user/mocks"
	"grpcrest/services/user/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUser(t *testing.T) {
	type field struct {
		authgrpcclient grpcclient.AuthGRPCClient_
		userrepo       repository.UserRepository
	}

	type args struct {
		page  int
		limit int
	}
	var fakeentities []entity.User = []entity.User{
		*testdata.NewUserEntity(), *testdata.NewUserEntity(),
	}

	mockUserRepoSuccess := new(mocks.UserRepository)
	mockUserRepoSuccess.On("GetAll", mock.Anything, mock.Anything).Return(fakeentities, nil)
	mockUserRepoFail := new(mocks.UserRepository)
	mockUserRepoFail.On("GetAll", mock.Anything, mock.Anything).Return(nil, errors.New("some error"))
	mockAuthGrpc := new(mocks.AuthGRPCClient_)

	testcases := []struct {
		Name   string
		Field  field
		Args   args
		Err    error
		Result []entity.User
	}{
		{
			Name: "success",
			Field: field{
				authgrpcclient: mockAuthGrpc,
				userrepo:       mockUserRepoSuccess,
			},
			Args: args{
				page:  1,
				limit: 2,
			},
			Err:    nil,
			Result: fakeentities,
		},
		{
			Name: "fail user repo",
			Field: field{
				authgrpcclient: mockAuthGrpc,
				userrepo:       mockUserRepoFail,
			},
			Args: args{
				page:  1,
				limit: 2,
			},
			Err:    errors.New("some error"),
			Result: nil,
		},
		{
			Name: "page invalid",
			Field: field{
				authgrpcclient: mockAuthGrpc,
				userrepo:       mockUserRepoFail,
			},
			Args: args{
				page:  0,
				limit: 2,
			},
			Err:    errors.New("some error"),
			Result: nil,
		},
		{
			Name: "limit invalid",
			Field: field{
				authgrpcclient: mockAuthGrpc,
				userrepo:       mockUserRepoFail,
			},
			Args: args{
				page:  1,
				limit: -3,
			},
			Err:    errors.New("some error"),
			Result: nil,
		},
		{
			Name: "fail page and limit repo",
			Field: field{
				authgrpcclient: mockAuthGrpc,
				userrepo:       mockUserRepoFail,
			},
			Args: args{
				page:  0,
				limit: 0,
			},
			Err:    errors.New("some error"),
			Result: nil,
		},
	}

	for _, v := range testcases {
		t.Run(v.Name, func(t *testing.T) {

			//dependency inject
			userUsecase := NewUserUsecase(v.Field.userrepo, v.Field.authgrpcclient)
			//execute
			data, err := userUsecase.GetAll(&v.Args.page, &v.Args.limit)

			//test
			if v.Err == nil {
				assert.Nil(t, err)
				assert.Equal(t, v.Result, data)
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, data)
			}
		})
	}
}
