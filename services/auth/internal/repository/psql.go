package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"grpcrest/services/auth/domain/entity"
	domain "grpcrest/services/auth/domain/repository"
	cfg "grpcrest/services/auth/internal/config"

	"github.com/google/uuid"
)

const (
	errNoRowsAffected = "no rows affected"
)

type AuthRepo struct {
	sql *sql.DB
	cfg *cfg.AppConfig
}

// IsExistEmail implements repository.AuthRepository
func (u *AuthRepo) IsExistEmail(email *string) bool {
	var id string
	err := u.sql.QueryRow(fmt.Sprintf("SELECT text(id) from %s.auth WHERE email = $1", u.cfg.DBSchema), *email).Scan(&id)
	if err != nil {
		return false
	}
	return true
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
		_, err = tx.Exec(fmt.Sprintf("UPDATE %s.auth SET %s = $1 WHERE user_id = $2", u.cfg.DBSchema, v.field), v.value, userID)
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
	//convert to uuid
	userid, err := uuid.Parse(auth.UserID)
	if err != nil {
		return err
	}

	res, err := a.sql.Exec(fmt.Sprintf("DELETE FROM %s.auth WHERE text(user_id) = $1", a.cfg.DBSchema), userid)
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
	err := a.sql.QueryRow(fmt.Sprintf("SELECT password FROM %s.auth WHERE email = $1", a.cfg.DBSchema), auth.Email).Scan(&pass)
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
