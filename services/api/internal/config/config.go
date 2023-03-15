package config

import (
	"log"
	"os"
)

type AppConfig struct {
	AppPort      string //rest api port
	GrpcPort     string //grpc api port
	AuthGrpcPort string //auth grpc port
	UserGrpcPort string //user grpc port
}

func NewAppConfig() *AppConfig {
	var cfg AppConfig
	AppPort, valid := os.LookupEnv("APP_PORT")
	if !valid {
		log.Fatal("app port required")
	}
	cfg.AppPort = AppPort

	GrpcPort, valid := os.LookupEnv("GRPC_PORT")
	if !valid {
		log.Fatal("grpc port required")
	}
	cfg.GrpcPort = GrpcPort

	UserGrpcPort, valid := os.LookupEnv("USER_GRPC_PORT")
	if !valid {
		log.Fatal("user grpc port required")
	}
	cfg.UserGrpcPort = UserGrpcPort

	AuthGrpcPort, valid := os.LookupEnv("AUTH_GRPC_PORT")
	if !valid {
		log.Fatal("auth grpc port required")
	}
	cfg.AuthGrpcPort = AuthGrpcPort
	return &cfg
}
