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

type fieldUpdate struct {
	field string
	value string
}

// UpdateAuth implements repository.AuthRepository
func (a *AuthRepo) UpdateAuth(auth *entity.Auth) error {
	var upField []fieldUpdate
	if auth.Email != "" {
		upField = append(upField, fieldUpdate{field: "email", value: auth.Email})
	}
	if auth.Password != "" {
		upField = append(upField, fieldUpdate{field: "password", value: auth.Password})
	}
	err := a.dynamicUpdate(auth.UserID, upField...)
	if err != nil {
		return err
	}
	return nil
}

func (u *AuthRepo) dynamicUpdate(userID string, f ...fieldUpdate) error {
	//begin transaction
	tx, err := u.sql.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	//loop
	for _, v := range f {
		_, err = tx.Exec(fmt.Sprintf("UPDATE %s.auth SET %s = $1 WHERE user_id = $3", u.cfg.DBSchema, v.field), v.value, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//commit
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// CreateAuth implements repository.AuthRepository
func (a *AuthRepo) CreateAuth(auth *entity.Auth) error {
	res, err := a.sql.Exec(fmt.Sprintf("INSERT INTO %s.auth(id,user_id,email,password) VALUES($1,$2,$3,$4)", a.cfg.DBSchema), auth.ID, auth.UserID, auth.Email, auth.Password)
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
