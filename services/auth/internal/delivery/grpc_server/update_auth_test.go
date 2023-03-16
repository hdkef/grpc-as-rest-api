package grpcserver

import (
	"context"
	"errors"
	jwtS "grpcrest/pkg/auth/jwt_service"
	authpb "grpcrest/proto/_generated/auth"
	authUC "grpcrest/services/auth/domain/usecase"
	"grpcrest/services/auth/internal/config"
	"grpcrest/services/auth/mocks"
	testdata "grpcrest/services/auth/test/testdata"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/peer"
)

func TestUpdateAuth(t *testing.T) {
	type field struct {
		authUC     authUC.AuthUseCase
		jwtService jwtS.JWTService_
		cfg        *config.AppConfig
	}

	type args struct {
		ctx  context.Context
		auth *authpb.UpdateAuthRequest
	}

	//mock auth usecase
	mockAuthUsecaseSuccess := new(mocks.AuthUseCase)
	mockAuthUsecaseSuccess.On("UpdateAuth", mock.AnythingOfType("*entity.Auth")).Return(nil)
	mockAuthUsecaseFail := new(mocks.AuthUseCase)
	mockAuthUsecaseFail.On("UpdateAuth", mock.AnythingOfType("*entity.Auth")).Return(errors.New("some error"))

	//double config
	cfg := config.AppConfig{}

	p := &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(1, 2, 3, 4), Port: 8888}}

	testcases := []struct {
		Name  string
		Field field
		Args  args
		Err   error
	}{
		{
			Name: "success",
			Field: field{
				authUC: mockAuthUsecaseSuccess,
				cfg:    &cfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewUpdateAuthFull(),
			},
			Err: nil,
		},
		{
			Name: "no user id",
			Field: field{
				authUC: mockAuthUsecaseSuccess,
				cfg:    &cfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewUpdateAuthNoUserId(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "no email",
			Field: field{
				authUC: mockAuthUsecaseSuccess,
				cfg:    &cfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewUpdateAuthNoEmail(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "no password",
			Field: field{
				authUC: mockAuthUsecaseSuccess,
				cfg:    &cfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewUpdateAuthNoPassword(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "fail usecase",
			Field: field{
				authUC: mockAuthUsecaseFail,
				cfg:    &cfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewUpdateAuthFull(),
			},
			Err: errors.New("some error"),
		},
	}

	for _, v := range testcases {
		t.Run(v.Name, func(t *testing.T) {
			//initiate delivery
			s := server{
				authUC: v.Field.authUC,
				cfg:    v.Field.cfg,
			}

			//execute
			_, err := s.UpdateAuth(v.Args.ctx, v.Args.auth)
			if v.Err == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}

		})
	}
}
