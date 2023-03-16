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

func TestUpdateUser(t *testing.T) {
	type field struct {
		userUC usecase.UserUsecase
		cfg    *config.AppConfig
	}

	type args struct {
		ctx  context.Context
		user *userpb.UpdateUserRequest
	}

	fakeUpdateRequest := testdata.NewUpdateUserRequestFull()

	mockUserUCSuccess := new(mocks.UserUsecase)
	mockUserUCSuccess.On("Update", mock.AnythingOfType("*entity.User")).Return(nil)
	mockUserUCFail := new(mocks.UserUsecase)
	mockUserUCFail.On("Update", mock.AnythingOfType("*entity.User")).Return(errors.New("some error"))
	p := &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(1, 2, 3, 4), Port: 8888}}
	doubleCfg := &config.AppConfig{}

	sameUserIdCtx := metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("x-user-id", fakeUpdateRequest.GetUserId()))
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
				user: fakeUpdateRequest,
			},
			Err: nil,
		},
		{
			Name: "different userId than jwt userId",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(differentIdCtx, p),
				user: fakeUpdateRequest,
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
				user: fakeUpdateRequest,
			},
			Err: errors.New("some error"),
		},
		{
			Name: "empty name",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(sameUserIdCtx, p),
				user: testdata.NewUpdateUserRequestNoName(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "empty email",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(sameUserIdCtx, p),
				user: testdata.NewUpdateUserRequestNoEmail(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "empty address",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(sameUserIdCtx, p),
				user: testdata.NewUpdateUserRequestNoAddress(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "empty password",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(sameUserIdCtx, p),
				user: testdata.NewUpdateUserRequestNoPassword(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "empty user id",
			Field: field{
				userUC: mockUserUCSuccess,
				cfg:    doubleCfg,
			},
			Args: args{
				ctx:  peer.NewContext(sameUserIdCtx, p),
				user: testdata.NewUpdateUserRequestNoUserId(),
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
			_, err := s.UpdateUser(v.Args.ctx, v.Args.user)
			if v.Err == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}

		})
	}
}
