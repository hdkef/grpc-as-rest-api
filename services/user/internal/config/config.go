package config

import (
	"log"
	"os"
)

type AppConfig struct {
	DBSchema   string
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	AppPort    string
	AuthPort   string
}

func NewAppConfig() *AppConfig {
	var cfg AppConfig
	AppPort, valid := os.LookupEnv("APP_PORT")
	if !valid {
		log.Fatal("app port required")
	}
	cfg.AppPort = AppPort
	DBSchema, valid := os.LookupEnv("DB_SCHEMA")
	if !valid {
		log.Fatal("db schema required")
	}
	cfg.DBSchema = DBSchema

	DBHost, valid := os.LookupEnv("DB_HOST")
	if !valid {
		log.Fatal("db host required")
	}
	cfg.DBHost = DBHost

	DBPort, valid := os.LookupEnv("DB_PORT")
	if !valid {
		log.Fatal("db port required")
	}
	cfg.DBPort = DBPort

	DBPassword, valid := os.LookupEnv("DB_PASSWORD")
	if !valid {
		log.Fatal("db password required")
	}
	cfg.DBPassword = DBPassword

	DBName, valid := os.LookupEnv("DB_NAME")
	if !valid {
		log.Fatal("db name required")
	}
	cfg.DBName = DBName

	DBUsername, valid := os.LookupEnv("DB_USERNAME")
	if !valid {
		log.Fatal("db username required")
	}
	cfg.DBUsername = DBUsername

	AuthPort, valid := os.LookupEnv("AUTH_PORT")
	if !valid {
		log.Fatal("auth port required")
	}
	cfg.AuthPort = AuthPort
	return &cfg
}
