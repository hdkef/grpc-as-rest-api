package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitPSQL(cfg *AppConfig) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
