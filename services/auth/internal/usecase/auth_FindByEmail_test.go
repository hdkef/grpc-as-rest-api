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

func TestFindUserIdPasswordByEmailAuth(t *testing.T) {
	type field struct {
		authrepo repository.AuthRepository
	}

	type args struct {
		auth *entity.Auth
	}

	fakeEntity := testdata.NewAuthEntity()

	mockAuthRepoSuccess := new(mocks.AuthRepository)
	mockAuthRepoSuccess.On("FindUserIdPasswordByEmail", fakeEntity).Return(fakeEntity.UserID, fakeEntity.Password, nil)
	mockAuthRepoFail := new(mocks.AuthRepository)
	mockAuthRepoFail.On("FindUserIdPasswordByEmail", fakeEntity).Return("", "", errors.New("some error"))

	type Result struct {
		Password string
		UserId   string
	}

	testcases := []struct {
		Name   string
		Field  field
		Args   args
		Err    error
		Result Result
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
			Result: Result{
				UserId:   fakeEntity.UserID,
				Password: fakeEntity.Password,
			},
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
			Result: Result{
				UserId:   "",
				Password: "",
			},
		},
	}

	for _, v := range testcases {
		t.Run(v.Name, func(t *testing.T) {

			//dependency inject
			authUsecase := NewAuthUsecase(v.Field.authrepo)
			//execute
			userId, pass, err := authUsecase.FindUserIdPasswordByEmail(v.Args.auth)

			//test
			if v.Err == nil {
				assert.Nil(t, err)
				assert.Equal(t, v.Result.Password, pass)
				assert.Equal(t, v.Result.UserId, userId)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, v.Result.Password, pass)
				assert.Equal(t, v.Result.UserId, userId)
			}
		})
	}
}
