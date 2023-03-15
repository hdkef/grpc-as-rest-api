package grpcserver

import (
	"database/sql"
	userpb "grpcrest/proto/_generated/user"
	"grpcrest/services/user/internal/config"
	grpcclient "grpcrest/services/user/internal/delivery/grpc_client"
	"grpcrest/services/user/internal/repository"
	"grpcrest/services/user/internal/usecase"

	"google.golang.org/grpc"
)

func DeleteUserHandler(s *grpc.Server, cfg *config.AppConfig, sql *sql.DB, gc grpcclient.AuthGRPCClient) {
	//init repo
	repo := repository.NewUserRepository(sql, cfg)

	//init usecase
	uc := usecase.NewUserUsecase(repo, gc)

	//register grpc service
	userpb.RegisterDeleteUserServiceServer(s, &server{
		userUC: uc,
	})
}
