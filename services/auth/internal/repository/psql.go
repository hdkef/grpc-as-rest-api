package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"grpcrest/services/auth/domain/entity"
	domain "grpcrest/services/auth/domain/repository"
	cfg "grpcrest/services/auth/internal/config"
)

const (
	errNoRowsAffected = "no rows affected"
)

type AuthRepo struct {
	sql *sql.DB
	cfg *cfg.AppConfig
}

// CreateAuth implements repository.AuthRepository
func (a *AuthRepo) CreateAuth(auth *entity.Auth) error {
	res, err := a.sql.Exec(fmt.Sprintf("INSERT INTO %s.auth(user_id,email_password) VALUES($1,$2,$3)", a.cfg.DBSchema), auth.UserID, auth.Email, auth.Password)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New(errNoRowsAffected)
	}
	return nil
}

// DeleteAuth implements repository.AuthRepository
func (a *AuthRepo) DeleteAuth(auth *entity.Auth) error {
	res, err := a.sql.Exec(fmt.Sprintf("DELETE FROM %s.auth WHERE user_id = $1", a.cfg.DBSchema), auth.UserID)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New(errNoRowsAffected)
	}
	return nil
}

// FindPasswordByEmail implements repository.AuthRepository
func (a *AuthRepo) FindPasswordByEmail(auth *entity.Auth) (string, error) {
	var pass string
	err := a.sql.QueryRow(fmt.Sprintf("SELECT password from %s.auth WHERE email = ?", a.cfg.DBSchema), auth.Email).Scan(&pass)
	if err != nil {
		return "", err
	}
	return pass, nil
}

func NewAuthRepo(sql *sql.DB, cfg *cfg.AppConfig) domain.AuthRepository {
	return &AuthRepo{
		sql: sql,
		cfg: cfg,
	}
}
