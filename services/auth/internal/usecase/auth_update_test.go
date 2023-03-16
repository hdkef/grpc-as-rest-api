package usecase

import (
	"errors"
	"grpcrest/services/auth/domain/entity"
	"grpcrest/services/auth/domain/repository"
	"grpcrest/services/auth/mocks"
	"grpcrest/services/auth/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateAuth(t *testing.T) {
	type field struct {
		authrepo repository.AuthRepository
	}

	type args struct {
		auth *entity.Auth
	}

	fakeEntity := testdata.NewAuthEntity()

	mockAuthRepoSuccess := new(mocks.AuthRepository)
	mockAuthRepoSuccess.On("UpdateAuth", fakeEntity).Return(nil)
	mockAuthRepoSuccess.On("IsExistEmail", mock.AnythingOfType("*string")).Return(false)
	mockAuthRepoSuccess2 := new(mocks.AuthRepository)
	mockAuthRepoSuccess2.On("UpdateAuth", fakeEntity).Return(nil)
	mockAuthRepoSuccess2.On("IsExistEmail", mock.AnythingOfType("*string")).Return(true)
	mockAuthRepoFail := new(mocks.AuthRepository)
	mockAuthRepoFail.On("UpdateAuth", fakeEntity).Return(errors.New("some error"))
	mockAuthRepoFail.On("IsExistEmail", mock.AnythingOfType("*string")).Return(false)

	testcases := []struct {
		Name  string
		Field field
		Args  args
		Err   error
	}{
		{
			Name: "success",
			Field: field{
				authrepo: mockAuthRepoSuccess,
			},
			Args: args{
				auth: fakeEntity,
			},
			Err: nil,
		},
		{
			Name: "success",
			Field: field{
				authrepo: mockAuthRepoSuccess2,
			},
			Args: args{
				auth: fakeEntity,
			},
			Err: nil,
		},
		{
			Name: "auth repo fail",
			Field: field{
				authrepo: mockAuthRepoFail,
			},
			Args: args{
				auth: fakeEntity,
			},
			Err: errors.New("some error"),
		},
	}

	for _, v := range testcases {
		t.Run(v.Name, func(t *testing.T) {

			//dependency inject
			authUsecase := NewAuthUsecase(v.Field.authrepo)
			//execute
			err := authUsecase.UpdateAuth(v.Args.auth)

			//test
			if v.Err == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
