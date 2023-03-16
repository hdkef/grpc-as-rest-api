package grpcserver

import (
	"context"
	"errors"
	jwtS "grpcrest/pkg/auth/jwt_service"
	authpb "grpcrest/proto/_generated/auth"
	authUC "grpcrest/services/auth/domain/usecase"
	"grpcrest/services/auth/internal/config"
	"grpcrest/services/auth/mocks"
	"grpcrest/services/auth/test/testdata"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/peer"
)

func TestDeleteAuth(t *testing.T) {
	type field struct {
		authUC     authUC.AuthUseCase
		jwtService jwtS.JWTService_
		cfg        *config.AppConfig
	}

	type args struct {
		ctx  context.Context
		auth *authpb.DeleteAuthRequest
	}

	//mock auth usecase
	mockAuthUsecaseSuccess := new(mocks.AuthUseCase)
	mockAuthUsecaseSuccess.On("DeleteAuth", mock.AnythingOfType("*entity.Auth")).Return(nil)
	mockAuthUsecaseFail := new(mocks.AuthUseCase)
	mockAuthUsecaseFail.On("DeleteAuth", mock.AnythingOfType("*entity.Auth")).Return(errors.New("some error"))

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
				auth: testdata.NewDeleteAuth(),
			},
			Err: nil,
		},
		{
			Name: "fail usecase",
			Field: field{
				authUC: mockAuthUsecaseFail,
				cfg:    &cfg,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewDeleteAuth(),
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
				auth: testdata.NewDeleteAuthNoUserId(),
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
			_, err := s.DeleteAuth(v.Args.ctx, v.Args.auth)
			if v.Err == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}

		})
	}
}
