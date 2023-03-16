package usecase

import (
	"grpcrest/services/user/domain/repository"
	"grpcrest/services/user/domain/usecase"
	grpcclient "grpcrest/services/user/internal/delivery/grpc_client"
)

type UserArtifact struct {
	authgrpcclient grpcclient.AuthGRPCClient_
	userrepo       repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository, authgrpcclient grpcclient.AuthGRPCClient_) usecase.UserUsecase {
	return &UserArtifact{
		userrepo:       repo,
		authgrpcclient: authgrpcclient,
	}
}
