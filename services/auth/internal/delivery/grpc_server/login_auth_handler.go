package grpcserver

import (
	"database/sql"
	jwtS "grpcrest/pkg/auth/jwt_service"
	authpb "grpcrest/proto/_generated/auth"
	"grpcrest/services/auth/internal/config"
	"grpcrest/services/auth/internal/repository"
	"grpcrest/services/auth/internal/usecase"

	"google.golang.org/grpc"
)

func LoginAuthHandler(s *grpc.Server, cfg *config.AppConfig, sql *sql.DB) {
	//init repo
	repo := repository.NewAuthRepo(sql, cfg)

	//init usecase
	uc := usecase.NewAuthUsecase(repo)

	//register grpc service
	authpb.RegisterLoginAuthServiceServer(s, &server{
		authUC:     uc,
		jwtService: jwtS.NewJWTService(cfg.JWTSecret),
		cfg:        cfg,
		bcrypt:     usecase.NewBcrypt(),
	})
}
