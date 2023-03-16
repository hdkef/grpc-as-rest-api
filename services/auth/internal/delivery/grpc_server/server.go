package grpcserver

import (
	jwtS "grpcrest/pkg/auth/jwt_service"
	authUC "grpcrest/services/auth/domain/usecase"
	"grpcrest/services/auth/internal/config"
)

type server struct {
	authUC     authUC.AuthUseCase
	jwtService jwtS.JWTService_
	cfg        *config.AppConfig
	bcrypt     authUC.Bcrypt_
}
