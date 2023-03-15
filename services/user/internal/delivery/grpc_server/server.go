package grpcserver

import (
	"grpcrest/services/user/domain/usecase"
)

type server struct {
	userUC usecase.UserUsecase
}
