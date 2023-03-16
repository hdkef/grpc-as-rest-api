package grpcserver

import (
	"context"
	"errors"
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/domain/usecase"
	"grpcrest/services/user/internal/config"
	"grpcrest/services/user/mocks"
	"grpcrest/services/user/test/testdata"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/peer"
)

func TestCreateUser(t *testing.T) {
	type field struct {
		userUC usecase.UserUsecase
		cfg    *config.AppConfig
	}

	type args struct {
		ctx  context.Context
		user *userpb.CreateUserRequest
	}

	mockUserUCSuccess := new(mocks.UserUsecase)
	mockUserUCSuccess.On("Create", mock.AnythingOfType("*entity.User")).Return(nil)
	mockUserUCFail := new(mocks.UserUsecase)
	mockUserUCFail.On("Create", mock.AnythingOfType("*entity.User")).Return(errors.New("some error"))
	p := &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(1, 2, 3, 4), Port: 8888}}
	doubleCfg := &config.AppConfig{}

	testcases := []struct {
		Name  string
		Args  args
		Field field
		Err   error
	}{
		{
			Name: "success",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				user: testdata.NewCreateUserRequestFull(),
			},
			Err: nil,
		},
		{
			Name: "fail usecase",
			Field: field{
				userUC: mockUserUCFail,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				user: testdata.NewCreateUserRequestFull(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "missing name",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				user: testdata.NewCreateUserRequestNoName(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "missing email",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				user: testdata.NewCreateUserRequestNoEmail(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "missing address",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				user: testdata.NewCreateUserRequestNoAddress(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "missing password",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				user: testdata.NewCreateUserRequestNoPassword(),
			},
			Err: errors.New("some error"),
		},
	}

	for _, v := range testcases {
		t.Run(v.Name, func(t *testing.T) {
			//initiate delivery
			s := server{
				userUC: v.Field.userUC,
				cfg:    v.Field.cfg,
			}

			//execute
			_, err := s.CreateUser(v.Args.ctx, v.Args.user)
			if v.Err == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}

		})
	}
}
