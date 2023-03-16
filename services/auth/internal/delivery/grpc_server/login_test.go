package grpcserver

import (
	"context"
	"errors"
	jwtS "grpcrest/pkg/auth/jwt_service"
	mockJWT "grpcrest/pkg/auth/mocks"
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

func TestLoginAuth(t *testing.T) {
	type field struct {
		authUC     authUC.AuthUseCase
		jwtService jwtS.JWTService_
		cfg        *config.AppConfig
		bcrypt     authUC.Bcrypt_
	}

	type args struct {
		ctx  context.Context
		auth *authpb.LoginAuthRequest
	}

	//mock auth usecase
	mockAuthUsecaseSuccess := new(mocks.AuthUseCase)
	mockAuthUsecaseSuccess.On("FindUserIdPasswordByEmail", mock.AnythingOfType("*entity.Auth")).Return("userid", "pass", nil)
	mockAuthUsecaseFail := new(mocks.AuthUseCase)
	mockAuthUsecaseFail.On("FindUserIdPasswordByEmail", mock.AnythingOfType("*entity.Auth")).Return("", "", errors.New("some error"))

	//double config
	cfg := config.AppConfig{}

	//mock jwt
	mockJWTServiceSuccess := new(mockJWT.JWTService_)
	mockJWTServiceSuccess.On("GenerateToken", mock.AnythingOfType("string")).Return("dummyjwt", nil)
	mockJWTServiceFail := new(mockJWT.JWTService_)
	mockJWTServiceFail.On("GenerateToken", mock.AnythingOfType("string")).Return("", errors.New("some error"))

	//mock bcrypt
	mockBcryptSuccess := new(mocks.Bcrypt_)
	mockBcryptSuccess.On("CompareHashAndPassword", mock.AnythingOfType("[]uint8"), mock.AnythingOfType("[]uint8")).Return(nil)
	mockBcryptFail := new(mocks.Bcrypt_)
	mockBcryptFail.On("CompareHashAndPassword", mock.AnythingOfType("[]uint8"), mock.AnythingOfType("[]uint8")).Return(errors.New("some error"))

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
				authUC:     mockAuthUsecaseSuccess,
				cfg:        &cfg,
				jwtService: mockJWTServiceSuccess,
				bcrypt:     mockBcryptSuccess,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewLoginAuth(),
			},
			Err: nil,
		},
		{
			Name: "fail usecase",
			Field: field{
				authUC:     mockAuthUsecaseFail,
				cfg:        &cfg,
				jwtService: mockJWTServiceSuccess,
				bcrypt:     mockBcryptSuccess,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewLoginAuth(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "fail jwt service",
			Field: field{
				authUC:     mockAuthUsecaseSuccess,
				cfg:        &cfg,
				jwtService: mockJWTServiceFail,
				bcrypt:     mockBcryptSuccess,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewLoginAuth(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "fail bcrypt",
			Field: field{
				authUC:     mockAuthUsecaseSuccess,
				cfg:        &cfg,
				jwtService: mockJWTServiceSuccess,
				bcrypt:     mockBcryptFail,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewLoginAuth(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "no email",
			Field: field{
				authUC:     mockAuthUsecaseSuccess,
				cfg:        &cfg,
				jwtService: mockJWTServiceSuccess,
				bcrypt:     mockBcryptSuccess,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewLoginAuthNoEmail(),
			},
			Err: errors.New("some error"),
		},
		{
			Name: "no password",
			Field: field{
				authUC:     mockAuthUsecaseSuccess,
				cfg:        &cfg,
				jwtService: mockJWTServiceSuccess,
				bcrypt:     mockBcryptSuccess,
			},
			Args: args{
				ctx:  peer.NewContext(context.TODO(), p),
				auth: testdata.NewLoginAuthNoPassword(),
			},
			Err: errors.New("some error"),
		},
	}

	for _, v := range testcases {
		t.Run(v.Name, func(t *testing.T) {
			//initiate delivery
			s := server{
				authUC:     v.Field.authUC,
				cfg:        v.Field.cfg,
				jwtService: v.Field.jwtService,
				bcrypt:     v.Field.bcrypt,
			}

			//execute
			_, err := s.LoginAuth(v.Args.ctx, v.Args.auth)
			if v.Err == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}

		})
	}
}
