package grpcserver

import (
	"database/sql"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/auth/internal/config"
	"grpcrest/services/auth/internal/repository"
	"grpcrest/services/auth/internal/usecase"

	"google.golang.org/grpc"
)

func DeleteAuthHandler(s *grpc.Server, cfg *config.AppConfig, sql *sql.DB) {
	//init repo
	repo := repository.NewAuthRepo(sql, cfg)

	//init usecase
	uc := usecase.NewAuthUsecase(repo)

	//register grpc service
	authpb.RegisterDeleteAuthServiceServer(s, &server{
		authUC: uc,
	})
}
