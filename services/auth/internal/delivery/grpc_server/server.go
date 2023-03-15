package grpcserver

import (
	authUC "grpcrest/services/auth/domain/usecase"
)

type server struct {
	authUC authUC.AuthUseCase
}
