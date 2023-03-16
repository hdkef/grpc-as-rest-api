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
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func TestDeleteUser(t *testing.T) {
	type field struct {
		userUC usecase.UserUsecase
		cfg    *config.AppConfig
	}

	type args struct {
		ctx  context.Context
		user *userpb.DeleteUserRequest
	}

	fakeDeleteRequest := testdata.NewDeleteUserRequest()

	mockUserUCSuccess := new(mocks.UserUsecase)
	mockUserUCSuccess.On("Delete", mock.AnythingOfType("*entity.User")).Return(nil)
	mockUserUCFail := new(mocks.UserUsecase)
	mockUserUCFail.On("Delete", mock.AnythingOfType("*entity.User")).Return(errors.New("some error"))
	p := &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(1, 2, 3, 4), Port: 8888}}
	doubleCfg := &config.AppConfig{}

	sameUserIdCtx := metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("x-user-id", fakeDeleteRequest.GetUserId()))
	differentIdCtx := metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("x-user-id", ""))

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
				ctx:  peer.NewContext(sameUserIdCtx, p),
				user: fakeDeleteRequest,
			},
			Err: nil,
		},
		{
			Name: "no user id and no metadata",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(differentIdCtx, p),
				user: testdata.NewDeleteUserRequestNoUserId(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "different userId than jwt userId",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(differentIdCtx, p),
				user: fakeDeleteRequest,
			},
			Err: errors.New("some error"),
		},
		{
			Name: "no userId metadata",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				user: fakeDeleteRequest,
			},
			Err: errors.New("some error"),
		},
		{
			Name: "fail usecase",
			Field: field{
				userUC: mockUserUCFail,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(sameUserIdCtx, p),
				user: fakeDeleteRequest,
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
			_, err := s.DeleteUser(v.Args.ctx, v.Args.user)
			if v.Err == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}

		})
	}
}
