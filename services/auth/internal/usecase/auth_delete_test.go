package usecase

import (
	"errors"
	"grpcrest/services/auth/domain/entity"
	"grpcrest/services/auth/domain/repository"
	"grpcrest/services/auth/mocks"
	"grpcrest/services/auth/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteAuth(t *testing.T) {
	type field struct {
		authrepo repository.AuthRepository
	}

	type args struct {
		auth *entity.Auth
	}

	fakeEntity := testdata.NewAuthEntity()

	mockAuthRepoSuccess := new(mocks.AuthRepository)
	mockAuthRepoSuccess.On("DeleteAuth", fakeEntity).Return(nil)
	mockAuthRepoFail := new(mocks.AuthRepository)
	mockAuthRepoFail.On("DeleteAuth", fakeEntity).Return(errors.New("some error"))

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
			err := authUsecase.DeleteAuth(v.Args.auth)

			//test
			if v.Err == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
