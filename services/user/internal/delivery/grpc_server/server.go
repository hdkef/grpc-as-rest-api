package grpcserver

import (
	"grpcrest/services/user/domain/usecase"
	"grpcrest/services/user/internal/config"
)

type server struct {
	userUC usecase.UserUsecase
	cfg    *config.AppConfig
}
